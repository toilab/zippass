package zippass

import (
	"fmt"
	"net/http"
	_ "html/template"
	"log"
	
	"github.com/gin-gonic/gin"
	
	"google.golang.org/appengine"
    "google.golang.org/appengine/datastore"
	
	// "github.com/vira0223/zippass/entity"
	"entity"

	"golang.org/x/net/context"
	"time"
	"strconv"
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
	log.Printf("debug: end %s", "GetMainEngine")
	return router
}

func RootGet(g *gin.Context) {
	log.Printf("debug: start %s", "RootGet")
	g.HTML(http.StatusOK, "index.html", nil)
}

func TestGet(g *gin.Context) {
	log.Printf("debug: start %s", "TestGet")
	ids, _ := strconv.ParseInt(g.Param("id"), 10, 64)

	log.Printf("debug: ids-> %s", ids)

	// Datastoreからデータ取得。ids=KeyNameとして検索
	ctx := appengine.NewContext(g.Request)

	var passInfo entity.PassInfo
	//k := datastore.NewKey(ctx, "PassInfo", "", ids, nil)
	headKey := datastore.NewKey(ctx, "PassHeader", "", 5629499534213120, nil)
	log.Printf("debug: headkey-> %s", fmt.Sprint(headKey))
	// _ = datastore.Get(ctx, k, &passInfo)
	q := datastore.NewQuery("PassInfo").Ancestor(headKey)
	if _, err := q.GetAll(ctx, &passInfo); err != nil {
		log.Fatalf("とってもやばす：%v", err)
		return
	}

	log.Printf("debug: PassInfoのGET結果 %s", fmt.Sprint(passInfo.Beacons))

	g.HTML(http.StatusOK, "t.html", gin.H{
		"description": passInfo.Description,
		"formatversion": passInfo.FormatVersion,
		"organizationname": passInfo.OrganizationName,
		"passtypeid": passInfo.PassTypeId,
		"serialnumber": passInfo.SerialNumber,
		"teamid": passInfo.TeamId,
		"applaunchurl": passInfo.AppLaunchURL,
		"associatedstoreid": passInfo.AssociatedStoreId,
		"userinfo": passInfo.Userinfo,
		"expirationdate": passInfo.ExpirationDate,
		"voided": passInfo.Voided,

//		"relevance": []entity.Beacon{ {passInfo.Beacons[0].Major, passInfo.Beacons[0].Minor, passInfo.Beacons[0].ProximityUUID, passInfo.Beacons[0].RelevantText},
//									  {passInfo.Beacons[1].Major, passInfo.Beacons[1].Minor, passInfo.Beacons[1].ProximityUUID, passInfo.Beacons[1].RelevantText} },
	})
}

// 開発環境では、http://localhost:8000/datastore でローカルに保存してあるデータが確認できる
func ZipPass(g *gin.Context) {
	log.Printf("debug: start %s", "ZipPass")
	ctx := appengine.NewContext(g.Request)

	/*
	ids, _ := AllocateID(ctx, "PassInfo")
	log.Printf("debug: Allocated %s", ids)
	*/

	header := entity.PassHeader{}
	header.PassId = 22222
	header.PassDescription = "いい感じのpass"
	header.PassUrl = "https://passurl.toilab.jp/xxx/yyy/paaaas.pkpass"
	header.PassName = "開店祝い用クーポン"
	header.OwnerUserId = 1029384756
	header.CreatePassCnt = 0
	header.CreateDt = time.Now()
	header.CreateUserId = 1029384756
	header.UpdateDt = time.Now()
	header.UpdateUserId = 1029384756
	header.DelFlg = false

	k := datastore.NewIncompleteKey(ctx, "PassHeader", nil)
	log.Printf("debug: passheader key %s", fmt.Sprint(k))

	if _, err := datastore.Put(ctx, k, &header); err != nil {
		log.Fatalf("やばす：%v", err)
		return
	}

	pass := entity.PassInfo{}
	pass.Description = "オープン記念!!90%引!!!"
	pass.FormatVersion = 1
	pass.OrganizationName = "ToI研"
	pass.PassTypeId = "pass.jp.toilab.coupon.dev"
	pass.SerialNumber = "gabigabino7"
	pass.TeamId = "W64S6BJKKZ"

	pass.ExpirationDate = time.Now().AddDate(0, 3, 0)
	pass.Voided = false

	b1 := entity.NewBeacon(12345, 65412, "shirangana", "ビーコンに反応しているぞ!!")
	b2 := entity.NewBeacon(3456, 4123, "bibibiii", "むしろビーコンが反応しているのか!?")
	pass.Beacons = []entity.Beacon{b1, b2}

	headKey := datastore.NewKey(ctx, "PassInfo", "", 5629499534213120, nil)
	k2 := datastore.NewIncompleteKey(ctx, "PassInfo", headKey)
	log.Printf("debug: passInfo key %s", fmt.Sprint(k2))

	if _, err := datastore.Put(ctx, k2, &pass); err != nil {
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
