package main

import (
	_ "awesomeProject/docs"
	"awesomeProject/internal/config"
	"awesomeProject/internal/handlers"
	"awesomeProject/internal/repositories"
	"awesomeProject/internal/services"
	"database/sql"
	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
	"go.uber.org/zap"
)

func main() {
	router := gin.Default()

	logger := config.InitLogger()
	defer logger.Sync()

	database := config.InitDB()
	defer database.Close()

	router = setupSwagger(router)
	router = setupUserRoutes(database, logger, router)

	router.Run("localhost:8080")
}

func setupSwagger(router *gin.Engine) *gin.Engine {
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	return router
}

func setupUserRoutes(database *sql.DB, sugar *zap.SugaredLogger, router *gin.Engine) *gin.Engine {
	var userRepository = repositories.NewUserRepository(database, sugar)
	userService := services.NewUserService(userRepository, sugar)

	userHandler := handlers.NewUserHandler(userService, sugar)

	router.GET("/user/all", userHandler.GetAllUsers)
	router.GET("/user/:id", userHandler.GetUserById)
	router.POST("/user", userHandler.AddUser)
	router.DELETE("/user/:id", userHandler.DeleteUser)
	router.PUT("/user/:id", userHandler.UpdateUser)
	router.GET("/user/all/books", userHandler.GetAll)

	return router
}
