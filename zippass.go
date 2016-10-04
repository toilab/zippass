package zippass

import (
	"fmt"
	"net/http"
	_ "html/template"
	"log"
	
	"github.com/gin-gonic/gin"
	// "github.com/comail/colog"
	
	"entity"
	
	
)

// initialize
func init() {
	// set logger
	
	// SetLogger()
	
	// http.HandleFunc("/", handler)
	http.Handle("/", GetMainEngine())
	
	// colog.Register()
	// log.Printf("info: colog registered.")
}

/*
func SetLogger() {
	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors:	true,
		Flag:	log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()
	
	log.Printf("info: colog registered.")
}
*/
func GetMainEngine() *gin.Engine {
	log.Printf("debug: start %s", "GetMainEngine")
	router := gin.Default()
	router.LoadHTMLGlob("templates/*")
	router.GET("/", RootGet)
	router.GET("/test/:id", TestGet)
	router.POST("/zippass/", ZipPass)
	router.POST("/test/", TestPost)
	log.Printf("debug: end %s", "GetMainEngine")
	return router
}

func RootGet(g *gin.Context) {
	log.Printf("debug: start %s", "RootGet")
	g.HTML(http.StatusOK, "index.html", nil)
}

func TestGet(g *gin.Context) {
	log.Printf("debug: start %s", "TestGet")
	ids := g.Param("id")
	g.HTML(http.StatusOK, "t.html", gin.H{
		"KeyIDs": ids,
	})
}

func TestPost(g *gin.Context) {
	log.Printf("debug: start %s", "TestPost")
	var passInfo entity.PassInfo
	j := g.BindJSON(&passInfo)
	log.Printf(fmt.Sprintf("PassTypeID is %s.", passInfo.PassTypeID))
	if j != nil {
		
		g.String(200, "kitakita")
	}
}

func ZipPass(g *gin.Context) {
	log.Printf("debug: start %s", "ZipPass")
	
	g.String(200, "Request is received.")
}
