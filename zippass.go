package zippass

import (
	_ "fmt"
	"net/http"
	_ "html/template"
	
	"github.com/gin-gonic/gin"
)

// initialize
func init() {
	// http.HandleFunc("/", handler)
	http.Handle("/", GetMainEngine())
}

func GetMainEngine() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/test/:id", TestGet)
	router.POST("/zippass/", ZipPass)
	return router
}

func TestGet(g *gin.Context) {
	ids := g.Param("id")
	g.HTML(http.StatusOK, "t.html", gin.H{
		"KeyIDs": ids,
	})
}

func ZipPass(g *gin.Context) {
	g.String(200, "Request is received.")
}
