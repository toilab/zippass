package zippass

import (
	"fmt"
	"net/http"
	_ "html/template"
	"log"
	
	"github.com/gin-gonic/gin"
	
	"google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
	
	"github.com/vira0223/zippass/entity"

	"golang.org/x/net/context"
)

// initialize
func init() {

	http.Handle("/", GetMainEngine())

}

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
	ctx := appengine.NewContext(g.Request)

	ids, _ := AllocateID(ctx, "Pass")
	log.Printf("debug: Allocated %s", ids)

	k := datastore.NewKey(ctx, "Pass", ids, 0, nil)
	log.Printf("debug: getDatastore %s", fmt.Sprint(k))

	pass := entity.PassInfo{}
	pass.PassTypeID = "kakakakaka"
	pass.SerialNumber = "ser_99"
	
	if _, err := datastore.Put(ctx, k, &pass); err != nil {
		log.Fatalf("やばす：%v", err)
		return
	}
	
	g.String(200, "Put success.")
}

// UniqueなIDを文字列で返してくれるいいやつ
func AllocateID(c context.Context, kind string) (string, error) {
    id, _, err := datastore.AllocateIDs(c, kind, nil, 1)
    return fmt.Sprint(id), err
}
