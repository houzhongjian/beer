package beer

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Engine interface {
	Run(addr string) error
	GET(path string, handler beerFunc)
	POST(path string, handler beerFunc)
	DELETE(path string, handler beerFunc)
	PUT(path string, handler beerFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Static(path string, dir string)
	SetTemplateDir(dir string)
	Use(middleware ...beerFunc)
}

type beerFunc func(*Context)

type beerHandler struct {
	Method string
	Path   string
}

type Handler struct {
	router           map[beerHandler]beerFunc
	fsRouter         map[string]string
	templateDir      string
	templateData     map[string]string
	middleware       []beerFunc
	middlewareRouter map[string][]beerFunc
}

func New() Engine {
	e := new(Handler)
	e.router = make(map[beerHandler]beerFunc)
	e.fsRouter = make(map[string]string)
	e.middlewareRouter = make(map[string][]beerFunc)
	e.templateData = make(map[string]string)
	return e
}


func (srv *Handler) Run(addr string) error {
	if len(addr) < 1 {
		panic("addr 不能为空")
	}
	return http.ListenAndServe(addr, srv)
}

func (srv *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paramsUrl := strings.Split(r.RequestURI, "?")
	path := paramsUrl[0]

	//判断是否为资源文件.
	fPath, ok := srv.fsRouter[path]
	if ok {
		//返回文件.
		http.ServeFile(w, r, fPath)
		return
	}

	//固定路由优先.
	h := beerHandler{
		Method: r.Method,
		Path:   path,
	}
	handler, ok := srv.router[h]
	if ok {
		middleware := srv.middlewareRouter[path]
		srv.beerFunc(w, r, nil, handler, middleware)
		return
	}

	paramsMp := map[string]string{}
	for router, handler := range srv.router {
		//判断当前的路由.
		//判断是否有:.
		index := strings.Index(router.Path, ":")
		if index > -1 {
			//正则路由.
			//1.先找出参数字段:[\w]+
			reg, err := regexp.Compile(`:[\w]+`)
			if err != nil {
				log.Printf("err:%+v\n", err)
				return
			}
			sArr := reg.FindAllString(router.Path, -1)
			route := "^" + router.Path
			for _, param := range sArr {
				//将router中的:xxx部分替换成[\w]+.
				route = strings.Replace(route, param, `[\w]+`, -1)
			}
			route += "$"

			//匹配map中的路由.
			reg, err = regexp.Compile(route)
			if err != nil {
				log.Printf("err:%+v\n", err)
				return
			}
			if reg.MatchString(path) {
				sArr := strings.Split(path, "/")
				rArr := strings.Split(router.Path, "/")

				for i := 0; i < len(sArr); i++ {
					if sArr[i] != rArr[i] {
						key := strings.Replace(rArr[i], ":", "", -1)
						val := sArr[i]
						paramsMp[key] = val
					}
				}

				if router.Method != r.Method {
					_, _ = w.Write([]byte("not found"))
					return
				}
				middleware := srv.middlewareRouter[router.Path]
				srv.beerFunc(w, r, paramsMp, handler, middleware)
				return
			}
		}
	}

	w.WriteHeader(http.StatusNotFound)
	_, _ = w.Write([]byte("not found"))
}

//parseParams 解析http请求中的参数.
//记录到Params当中.
func (srv *Handler) parseParams(r *http.Request) map[string]string {
	paramsMp := map[string]string{}
	_ = r.ParseForm()
	for key, val := range r.Form {
		paramsMp[key] = val[0]
	}
	return paramsMp
}

func (srv *Handler) beerFunc(w http.ResponseWriter, r *http.Request, params map[string]string, handler beerFunc, middleware []beerFunc) {
	remoteAddr := strings.Split(r.RemoteAddr, ":")
	ctx := &Context{
		Method:    r.Method,
		Request:   r,
		Response:  w,
		params:    params,
		UserAgent: r.UserAgent(),
		Url:       r.URL.String(),
		Header:    r.Header,
		IP:        remoteAddr[0],
		step:      1, //步长默认为1.
		Data:      make(map[string]interface{}),
	}

	//执行中间件.
	for _, midd := range middleware {
		midd(ctx)
		log.Println(midd)
		if ctx.step < 1 {
			return
		}
	}
	handler(ctx)
}
