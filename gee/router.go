package gee

import (
	"net/http"
	"strings"
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

func parsePattern(parttern string) []string {
	//以/分割
	vs := strings.Split(parttern, "/")

	parts := make([]string, 0)
	//排除掉空字符
	for _, item := range vs {
		if item != "" {

			parts = append(parts, item)
			if item[0] == '*' {

				break
			}

		}

	}
	return parts

}

func (r *router) addRoute(method string, pattern string, handler HandlerFunc) {
	parts := parsePattern(pattern)
	key := method + "-" + pattern
	//查是否有对应的路由
	_, ok := r.roots[method]
	if !ok {
		r.roots[method] = &node{}
	}
	r.roots[method].insert(pattern, parts, 0)
	r.handlers[key] = handler

}

func (r *router) getRoute(method string, path string) (*node, map[string]string) {
	//解析成字符切片
	searchParts := parsePattern(path)
	//查找是否有已注册的方法
	params := make(map[string]string)
	root, ok := r.roots[method]
	if !ok {
		return nil, nil

	}

	//查找对应的节点
	n := root.search(searchParts, 0)

	if n != nil {
		parts := parsePattern(n.pattern)
		for index, part := range parts {
			if part[0] == ':' {

				params[part[1:]] = searchParts[index]

			}
			if part[0] == '*' && len(part) > 1 {
				//将字符切片用/链接成一个整体
				params[part[1:]] = strings.Join(searchParts[index:], "/")
				break
			}

		}
		return n, params
	}
	return nil, nil
}

func (r *router) getRoutes(method string) []*node {
	root, ok := r.roots[method]
	if !ok {
		return nil
	}
	nodes := make([]*node, 0)
	root.travel(&nodes)
	return nodes
}

//在调用匹配到的handler前，将解析出来的路由参数赋值给了c.Params。
//这样就能够在handler中，通过Context对象访问到具体的值了。

func (r *router) handle(c *Context) {
	n, params := r.getRoute(c.Method, c.Path)
	if n != nil {
		c.Params = params
		key := c.Method + "-" + n.pattern
		c.handlers = append(c.handlers, r.handlers[key])

	} else {
		c.handlers = append(c.handlers, func(ctx *Context) {

			c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)

		})

	}
	c.Next()
}
