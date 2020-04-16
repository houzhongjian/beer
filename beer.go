package beer

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Engine interface {
	Run(addr string) error
	GET(path string, handler hFunc)
	POST(path string, handler hFunc)
	DELETE(path string, handler hFunc)
	PUT(path string, handler hFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
	Static(path string, dir string)
}

type hFunc func(*Context)

type beerHandler struct {
	Method string
	Path string
}

type Handler struct {
	router map[beerHandler]hFunc
	fsRouter map[string]string
}

func New() Engine {
	e := new(Handler)
	e.router = make(map[beerHandler]hFunc)
	e.fsRouter = make(map[string]string)
	return e
}

func (srv *Handler) Run(addr string) error {
	return http.ListenAndServe(addr, srv)
}

func (srv *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paramsUrl := strings.Split(r.RequestURI,"?")
	path := paramsUrl[0]

	//判断是否为资源文件.
	fPath, ok := srv.fsRouter[path]
	if ok {
		//返回文件.
		http.ServeFile(w, r, fPath)
		return
	}

	paramsMp := map[string]string{}
	for router, handler := range srv.router {
		//判断当前的路由.
	 	//判断是否有:.
	 	index := strings.Index(router.Path,":")
	 	if index > -1 {
	 		//正则路由.
	 		//1.先找出参数字段:[\w]+
			reg, err := regexp.Compile(`:[\w]+`)
			if err != nil {
				log.Printf("err:%+v\n",err)
				return
			}
			sArr := reg.FindAllString(router.Path, -1)
			route := "^"+router.Path
			for _, param := range sArr {
				//将router中的:xxx部分替换成[\w]+.
				route = strings.Replace(route, param, `[\w]+`,-1)
			}
			route += "$"

			//匹配map中的路由.
			reg, err = regexp.Compile(route)
			if err != nil {
				log.Printf("err:%+v\n",err)
				return
			}
			if reg.MatchString(path) {
				sArr := strings.Split(path,"/")
				rArr := strings.Split(router.Path,"/")

				for i := 0; i < len(sArr); i++ {
					if sArr[i] != rArr[i] {
						key := strings.Replace(rArr[i],":","",-1)
						val := sArr[i]
						paramsMp[key] = val
					}
				}

				if router.Method != r.Method {
					_, _ = w.Write([]byte("not found"))
					return
				}
				ctx := &Context{
					Method: r.Method,
					Request:  r,
					Response: w,
					Params: paramsMp,
				}
				handler(ctx)
				return
			}
		}
	 }

	 h := beerHandler{
		 Method: r.Method,
		 Path:   r.RequestURI,
	 }
	handler, ok := srv.router[h]
	if !ok {
		_, _ = w.Write([]byte("not found"))
		return
	}

	ctx := &Context{
		Method: r.Method,
		Request:  r,
		Response: w,
	}
	handler(ctx)
}
