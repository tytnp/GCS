package resources

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed all:admin-dist
var AdminFs embed.FS

//go:embed admin-dist/index.html
var AdminIndexFile []byte

// EmbedAdminVuePage 嵌入后台管理页面
func EmbedAdminVuePage(r *gin.RouterGroup) {
	assets, _ := fs.Sub(AdminFs, "admin-dist/admin-assets")
	r.StaticFS("/admin-assets", http.FS(assets))
	r.GET("/admin", func(c *gin.Context) {
		// 直接返回 HTML 内容
		//c.Header("Content-Type", "text/html; charset=utf-8")
		//c.Status(http.StatusOK)
		//c.Writer.Write(AdminIndexFile)
		c.Data(http.StatusOK, "text/html", AdminIndexFile)
	})
}
