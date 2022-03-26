package gee

import "strings"

/* 为了实现动态路由,需采用前缀树的形式
 */

type node struct {
	pattern  string  //待匹配的路由 如/p/:lang
	part     string  //路由中的一部分 如 :lang
	children []*node //子节点
	isWild   bool    //是否精确匹配，part 含有 : 或 * 时为true

}

//第一个匹配成功的节点,用于查找
func (n *node) matchChild(part string) *node {
	for _, child := range n.children {
		if child.part == part || child.isWild {
			return child
		}

	}
	return nil

}

//所有匹配成功的节点,用于查找
func (n *node) matchChildren(part string) []*node {
	nodes := make([]*node, 0)

	for _, child := range n.children {
		if child.part == part || child.isWild {
			nodes = append(nodes, child)
		}
	}

	return nodes

}

//插入和查找逻辑,对应开发中的注册处理器和调用处理器
func (n *node) insert(partten string, parts []string, height int) {

	if len(parts) == 0 {
		n.pattern = partten
		return
	}
	part := parts[height]
	child := n.matchChild(part)
	if child == nil {
		child = &node{part: part, isWild: part[0] == ':' || part[0] == '*'}
		n.children = append(n.children, child)
	}
	child.insert(partten, parts, height+1)

}

func (n *node) search(parts []string, height int) *node {
	if len(parts) == height || strings.HasPrefix(n.part, "*") {
		if n.pattern == "" {
			return nil
		}
		return n
	}

	part := parts[height]
	children := n.matchChildren(part)

	for _, child := range children {
		result := child.search(parts, height)
		if result != nil {
			return result
		}

	}
	return nil

}
