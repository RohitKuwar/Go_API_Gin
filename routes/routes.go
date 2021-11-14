package routes

import (
	"github.com/RohitKuwar/go_api_gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()
	grp := router.Group("/api")
	{
		grp.GET("/goals", controllers.GetGoals)
		grp.GET("/goals/:id", controllers.GetGoal)
		grp.POST("/goals", controllers.CreateGoal)
		grp.PATCH("/goals/:id", controllers.UpdateGoal)
		grp.DELETE("/goals/:id", controllers.DeleteGoal)
	}
	return router
}
