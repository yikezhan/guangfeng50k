package router

import (
	"github.com/gin-gonic/gin"
	"guangfeng/internal/controller"
	"net/http"
)

func Load(r *gin.Engine) *gin.Engine {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})
	r.GET("/ping", controller.Ping)
	v1 := r.Group("/room")
	{
		v1.POST("/create", controller.CreateRoom)
		v1.POST("/update", controller.UpdateRoom)
		v1.POST("/enter_room", controller.EnterRoom)
	}
	v2 := r.Group("/game")
	{
		v2.POST("/submit_data", controller.SubmitGameData) //提交对局结果
		v2.POST("/calculate", controller.Calculate)        //只有房主能触发
		v2.POST("/result", controller.GetGameResult)       //查看本轮结果
		v2.POST("/confirm", controller.ConfirmGameResult)  //确认对局结果
		v2.POST("/next_game", controller.NextGame)         //下一局,需要房主先触发才能开启
	}
	v3 := r.Group("/user")
	{
		v3.POST("/profile", controller.UserProfile) //当前游戏输赢情况概览
		v3.POST("/detail", controller.GameDetail)   //游戏总体详情
	}
	return r
}
