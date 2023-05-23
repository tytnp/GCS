package resources

import (
	"embed"
	"github.com/gin-gonic/gin"
	"io/fs"
	"net/http"
)

//go:embed all:gpt-dist
var GptFs embed.FS

//go:embed gpt-dist/index.html
var GptIndexFile []byte

// EmbedGptVuePage 嵌入后台管理页面
func EmbedGptVuePage(r *gin.RouterGroup) {
	assets, _ := fs.Sub(GptFs, "gpt-dist/gpt-assets")
	r.StaticFS("/gpt-assets", http.FS(assets))
	r.GET("/gpt", func(c *gin.Context) {
		// 直接返回 HTML 内容
		//c.Header("Content-Type", "text/html; charset=utf-8")
		//c.Status(http.StatusOK)
		//c.Writer.Write(AdminIndexFile)
		c.Data(http.StatusOK, "text/html", GptIndexFile)
	})
}
