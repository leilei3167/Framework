package gee

import (
	"net/http"
)

type router struct {
	//储存处理器
	handlers map[string]HandlerFunc
	//储存每种请求方式的Trie树根节点
	roots map[string]*node
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc),
		roots: make(map[string]*node)}
}




func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	r.handlers[key] = handler
}

func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
