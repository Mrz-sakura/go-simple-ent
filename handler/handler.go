package handler

import (
	"dwbackend/handler/admin"
)

// Handler 结构体用于组织所有处理器函数
type Handler struct {
	*admin.Handler
}

// 创建一个全局的 Handler 实例
var GlobalHandler = &Handler{}
