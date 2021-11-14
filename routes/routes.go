package routes

import (
	"github.com/RohitKuwar/go_api_gin/controllers"
	"github.com/gin-gonic/gin"
)

func SetupRouter() *gin.Engine {
	router := gin.Default()

	router.GET("/", (func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "root route"})
	}))

	grp := router.Group("/api")
	{
		grp.GET("/", (func(c *gin.Context) {
			c.JSON(200, gin.H{"message": "api route"})
		}))
		grp.GET("/goals", controllers.GetGoals)
		grp.GET("/goals/:id", controllers.GetGoal)
		grp.POST("/goals", controllers.CreateGoal)
		grp.PATCH("/goals/:id", controllers.UpdateGoal)
		grp.DELETE("/goals/:id", controllers.DeleteGoal)
	}
	return router
}
