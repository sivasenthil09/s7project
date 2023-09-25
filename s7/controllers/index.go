package controllers

import (
	"net/http"
	"s7/interfaces"
	"s7/models"

	"github.com/gin-gonic/gin"
)



type UserController struct {
	Userservices interfaces.UserService
}
func New (userservice interfaces.UserService)UserController{
	return UserController{
		Userservices: userservice,
	}
}
func (uc*UserController)CreateUser(ctx *gin.Context){
	var user models.User
	if err:=ctx.ShouldBindJSON(&user);err !=nil {
		ctx.JSON(http.StatusBadRequest,gin.H{"message":err.Error()})
		return
	}
	err :=uc.Userservices.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})

}
func (uc*UserController)GetUser(ctx *gin.Context){
	username :=ctx.Param("name")
	user,err :=uc.Userservices.GetUser(&username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}
// func (uc *UserController)UpdateUser(ctx *gin.Context){
// 	var user models.User
// 	if err := ctx.ShouldBindJSON(&user); err != nil {
// 		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
// 		return
// 	}
// 	err := uc.Userservices.UpdateUser(&user)
// 	if err != nil {
// 		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
// 		return
// 	}
// 	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
// }
func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username :=ctx.Param("name")
	err:=uc.Userservices.DeleteUser(&username)
	if err != nil {
	   ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
	   return
   }
   ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) GetAll(ctx *gin.Context) {
	users,err:=uc.Userservices.GetAll()
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, users)
}