package router

import "github.com/gin-gonic/gin"

// 添加请求数据处理中间件
func RequestDataMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		switch c.Request.Method {
		case "GET":
			// 处理 GET 请求的 query 参数
			query := c.Request.URL.Query()
			for key, values := range query {
				if len(values) > 0 {
					c.Set("geq_"+key, values[0]) // 如果有多个相同参数，只取第一个
				}
			}

		case "POST", "PUT":
			// 处理 POST/PUT 请求的 JSON body
			var jsonData map[string]interface{}
			if err := c.ShouldBindJSON(&jsonData); err == nil {
				for key, value := range jsonData {
					c.Set("geq_"+key, value)
				}
			}

			// 处理表单数据
			if err := c.Request.ParseForm(); err == nil {
				for key, values := range c.Request.PostForm {
					if len(values) > 0 {
						c.Set("geq_"+key, values[0])
					}
				}
			}
		}

		c.Next()
	}
}
