package gee

import (
	"fmt"
	"net/http"
)

//定义处理器函数(实现了ServeHTTP接口的才能作为处理器)
type HandlerFunc func(http.ResponseWriter, *http.Request)

/* 添加了一张路由映射表router，key 由请求方法和静态路由地址构成，
例如GET-/、GET-/hello、POST-/hello，这样针对相同的路由，
如果请求方法不同,可以映射不同的处理方法(Handler)，
value 是用户映射的处理方法 */
type Engine struct {
	roter map[string]HandlerFunc
}

//给调用者创建Engin的方法
func New() *Engine {
	return &Engine{
		roter: make(map[string]HandlerFunc),
	}
}

//注册处理器,内部调用
func (e *Engine) addRoute(method string, pattern string, handler HandlerFunc) {
	//Get-/  Get-/hello 实现同方法不同路径对应不同的处理器
	key := method + "-" + pattern
	e.roter[key] = handler

}

//供外部调用,用户根据需求注册服务
func (e *Engine) GET(pattern string, handler HandlerFunc) {
	e.addRoute("GET", pattern, handler)

}

func (engine *Engine) POST(pattern string, handler HandlerFunc) {
	engine.addRoute("POST", pattern, handler)
}

//让Engin实现ServeHTTP方法
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	//根据获取到的key去router注册表中查询
	if handler, ok := e.roter[key]; ok {
		//查询到了则执行
		handler(w, req)

	} else {
		fmt.Fprintf(w, "404 page not found: %s\n", req.URL)
	}

}

//开启服务
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)

}
