package main

import (
	"schedule-api-specification/controller"
	"schedule-api-specification/docs"
	"schedule-api-specification/middleware"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

func main() {
	// Register the controller
	//http.HandleFunc("/schedule", scheduleController)
	// Start the server
	//http.ListenAndServe(":8080", nil)

	router := gin.Default()

	// Initialize Swagger doc info
	docs.SwaggerInfo.BasePath = "/"

	// Swagger endpoint
	router.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Schedule CRUD routes
	schedule := router.Group("/schedule")
	{
		schedule.GET("/:date", middleware.CheckScope("volunteer"), controller.GetSchedule)

		schedule.GET("/task/:id", middleware.CheckScope("volunteer"), controller.GetTask)
		schedule.POST("/task", middleware.CheckScope("team-lead"), controller.CreateTask)
		schedule.PUT("/task/:id", middleware.CheckScope("team-lead"), controller.UpdateTask)
		schedule.DELETE("/task/:id", middleware.CheckScope("team-lead"), controller.DeleteTask)
		schedule.POST("/task/:id", middleware.CheckScope("volunteer"), controller.CheckIn)
	}

	router.Run(":8080")
}
