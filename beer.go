package beer

import (
	"log"
	"net/http"
	"regexp"
	"strings"
)

type Engine interface {
	Run(addr string) error
	GET(path string, handler beerHandlerFunc)
	POST(path string, handler beerHandlerFunc)
	DELETE(path string, handler beerHandlerFunc)
	PUT(path string, handler beerHandlerFunc)
	ServeHTTP(w http.ResponseWriter, r *http.Request)
}

type beerHandlerFunc func(*Context)

type Handler struct {
	router map[string]beerHandlerFunc
}

func New() Engine {
	e := new(Handler)
	e.router = make(map[string]beerHandlerFunc)
	return e
}

func (srv *Handler) Run(addr string) error {
	return http.ListenAndServe(addr, srv)
}

func (srv *Handler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	paramsUrl := strings.Split(r.RequestURI,"?")
	path := paramsUrl[0]
	paramsMp := map[string]string{}
	 for router, funcName := range srv.router {
		//判断当前的路由.
	 	//判断是否有:.
	 	index := strings.Index(router,":")
	 	if index > -1 {
	 		//正则路由.
	 		//1.先找出参数字段:[\w]+
			reg, err := regexp.Compile(`:[\w]+`)
			if err != nil {
				log.Printf("err:%+v\n",err)
				return
			}
			sArr := reg.FindAllString(router, -1)
			route := "^"+router
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
				rArr := strings.Split(router,"/")

				for i := 0; i < len(sArr); i++ {
					if sArr[i] != rArr[i] {
						key := strings.Replace(rArr[i],":","",-1)
						val := sArr[i]
						paramsMp[key] = val
					}
				}

				ctx := &Context{
					Request:  r,
					Response: w,
					Params: paramsMp,
				}
				funcName(ctx)
				return
			}
		}
	 }

	funcName, ok := srv.router[r.RequestURI]
	if !ok {
		_, _ = w.Write([]byte("not found"))
		return
	}
	ctx := &Context{
		Request:  r,
		Response: w,
	}
	funcName(ctx)
}
