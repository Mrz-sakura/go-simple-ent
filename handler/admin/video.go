package admin

import (
	"fmt"

	"github.com/gin-gonic/gin"
)

type Handler struct{}

// @url /admin/video/list
func (h *Handler) GetVideoList(c *gin.Context) {
	// c.Get("geq_page")
	fmt.Println(c.Get("geq_page"))
}
