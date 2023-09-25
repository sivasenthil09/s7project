package main

import (
	"context"
	"fmt"
	"log"
	"s7/config"
	"s7/constants"
	"s7/controllers"
	"s7/routes"
	"s7/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	mongoclient *mongo.Client
	ctx         context.Context
	server      *gin.Engine
)

func initRoutes() {
	routes.Default(server)
}
func initApp(mongoClient *mongo.Client) {
	ctx = context.TODO()
	UserCollection := mongoClient.Database(constants.DatabaseName).Collection("project")
	UserService := services.NewUseService(UserCollection, ctx)
	UserController := controllers.New(UserService)
	routes.RegisterUserRoutes(server, UserController)
}
func main() {
	server = gin.Default()
	mongoclient, err := config.ConnectDatabase()
	defer mongoclient.Disconnect(ctx)
	if err != nil {
		panic(err)

	}
	initRoutes()
	initApp(mongoclient)
	fmt.Println("server running on port", constants.Port)
	log.Fatal(server.Run(constants.Port))
}
