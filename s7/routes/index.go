package routes

import (
	"net/http"
	"s7/controllers"

	"github.com/gin-gonic/gin"
)

func RegisterUserRoutes(router *gin.Engine,controllers controllers.UserController){
	router.POST("/create",controllers.CreateUser)
	router.GET("/get/:name",controllers.GetUser)
    router.GET("/getall",controllers.GetAll)
	//router.PATCH("/update",controllers.UpdateUser)
	router.DELETE("/delete/:name",controllers.DeleteUser)
}

func Default(router *gin.Engine){
	router.GET("/api",func(ctx*gin.Context){
		ctx.JSON(http.StatusOK,gin.H{"status":"success","message":"server is healthy"})
	})
}

