net/http提供了基础的Web功能，即监听端口，映射静态路由，解析HTTP报文。一些Web开发中简单的需求并不支持，需要手工实现。

动态路由：例如hello/:name，hello/*这类的规则。
鉴权：没有分组/统一鉴权的能力，需要在每个路由映射的handler中实现。
模板：没有统一简化的HTML机制。
…

Go语言内置了 net/http库，封装了HTTP网络编程的基础的接口，我们实现的Gee Web 框架便是基于net/http的


ListenAndServe的第二个参数 Handler是实现框架的关键入口
