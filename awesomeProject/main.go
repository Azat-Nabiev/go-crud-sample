package main

import (
	_ "awesomeProject/docs"
	"awesomeProject/internal/config"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/repositories"
	"awesomeProject/internal/services"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func main() {
	router := gin.Default()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}
	defer logger.Sync()

	database := config.InitDB()
	defer database.Close()

	sugar := logger.Sugar()

	var userRepository = repositories.NewUserRepository(database, sugar)
	userService := services.NewUserService(userRepository, sugar)
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	userHandler := handlers.NewUserHandler(userService, sugar)

	router.GET("/user/all", userHandler.GetAllUsers)
	router.GET("/user/:id", userHandler.GetUserById)
	router.POST("/user", userHandler.AddUser)
	router.DELETE("/user/:id", userHandler.DeleteUser)
	router.PUT("/user/:id", userHandler.UpdateUser)
	router.GET("/user/all/books", userHandler.GetAll)

	router.Run("localhost:8080")
}
