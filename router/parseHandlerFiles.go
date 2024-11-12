package router

import (
	"fmt"
	"go/ast"
	"go/parser"
	"go/token"
	"log"
	"os"
	"path/filepath"
	"reflect"
	"strings"

	"dwbackend/handler"

	"github.com/gin-gonic/gin"
)

// 解析处理器文件获取路由信息
func ParseHandlerFiles(dir string) []RouteInfo {
	var routes []RouteInfo

	// 递归遍历handler目录及其子目录
	var files []string
	err := filepath.Walk(dir, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}
		if !info.IsDir() && strings.HasSuffix(path, ".go") {
			files = append(files, path)
		}
		return nil
	})
	if err != nil {
		log.Printf("遍历目录错误: %v\n", err)
		return routes
	}

	for _, file := range files {
		fset := token.NewFileSet()
		node, _ := parser.ParseFile(fset, file, nil, parser.ParseComments)

		// 遍历所有函数声明
		ast.Inspect(node, func(n ast.Node) bool {
			if fn, ok := n.(*ast.FuncDecl); ok {
				if fn.Doc != nil {
					for _, comment := range fn.Doc.List {
						if strings.Contains(comment.Text, "@url") {
							// 解析URL
							url := strings.TrimSpace(strings.Split(comment.Text, "@url")[1])

							// 扩展HTTP方法判断
							method := "POST" // 默认POST
							fnName := fn.Name.Name
							switch {
							case strings.HasPrefix(fnName, "Get") || strings.HasPrefix(fnName, "GET"):
								method = "GET"
							case strings.HasPrefix(fnName, "Put") || strings.HasPrefix(fnName, "PUT"):
								method = "PUT"
							case strings.HasPrefix(fnName, "Delete") || strings.HasPrefix(fnName, "DELETE"):
								method = "DELETE"
							}

							fmt.Println(fn.Name.Name)

							routes = append(routes, RouteInfo{
								Method:      method,
								URL:         url,
								HandlerFunc: getHandlerFuncByName(fn.Name.Name),
							})
						}
					}
				}
			}
			return true
		})
	}
	return routes
}

func getHandlerFuncByName(funcName string) gin.HandlerFunc {
	handlerPkg := reflect.ValueOf(handler.GlobalHandler)
	handlerFunc := handlerPkg.MethodByName(funcName)

	if !handlerFunc.IsValid() {
		log.Printf("警告: 未找到处理器函数 %s\n", funcName)
		return nil
	}

	// 返回一个包装函数，该函数符合 gin.HandlerFunc 的签名
	return func(c *gin.Context) {
		handlerFunc.Call([]reflect.Value{reflect.ValueOf(c)})
	}
}
