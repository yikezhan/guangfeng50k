package router

import (
	"entryTask/internal/controller"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(r *gin.Engine) *gin.Engine {
	r.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "404 not found")
	})

	v1 := r.Group("/room")
	{
		v1.POST("/create", controller.CreateRoom)
		v1.POST("/update", controller.UpdateRoom)
		v1.POST("/enter_room", controller.EnterRoom)
	}
	v2 := r.Group("/game")
	{
		v2.POST("/next_number", controller.CreateGameResult) //下一局,只有房主能触发
		v2.POST("/calculate", controller.Calculate)          //只有房主能触发
		v2.POST("/submit_data", controller.UpdateGameResult) //提交对局结果
		v2.POST("/confirm", controller.ConfirmGameResult)    //确认对局结果
		v2.POST("/", controller.ConfirmGameResult)           //确认对局结果

	}
	v3 := r.Group("/record")
	{
		v3.POST("/profile", controller.CreateRoom)
		v3.POST("/detail", controller.CreateRoom)
	}
	return r
}
