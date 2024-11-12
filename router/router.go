package router

import (
	"github.com/gin-gonic/gin"
)

// 路由信息结构
type RouteInfo struct {
	Method      string
	URL         string
	HandlerFunc interface{}
}

func InitRouter() *gin.Engine {
	r := gin.Default()

	// 使用请求数据处理中间件
	r.Use(RequestDataMiddleware())

	// 解析handler目录下的路由信息
	routes := ParseHandlerFiles("./handler")

	// 扩展路由注册
	for _, route := range routes {
		switch route.Method {
		case "GET":
			r.GET(route.URL, route.HandlerFunc.(gin.HandlerFunc))
		case "POST":
			r.POST(route.URL, route.HandlerFunc.(gin.HandlerFunc))
		case "PUT":
			r.PUT(route.URL, route.HandlerFunc.(gin.HandlerFunc))
		case "DELETE":
			r.DELETE(route.URL, route.HandlerFunc.(gin.HandlerFunc))
		}
	}

	return r
}
