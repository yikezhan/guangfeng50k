package main

//
import (
	"entryTask/configs"
	"entryTask/internal/controller"
	router "entryTask/route"
	"flag"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"log"
	"net/http"
	_ "net/http/pprof"
)

func main() {
	go func() {
		log.Println(http.ListenAndServe(":6060", nil))
	}()
	flag.Parse()
	//2、初始化日志配置
	configs.InitLog(zap.InfoLevel)
	defer configs.Log.Sync()
	//3、初始化服务
	controller.Init()
	//4、初始化路由
	r := gin.Default()
	r = router.Load(r)
	r.Run()
}
