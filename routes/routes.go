package routes

import (
	"github.com/esleirter/e-learning-api/controllers"
	"github.com/gin-gonic/gin"
)

// InitializeRoutes initializes API routes
func InitializeRoutes(router *gin.Engine) {
	router.GET("/course/:id", controllers.GetCourse)
	router.GET("/courses", controllers.GetAllCourses)
	router.POST("/courses/create", controllers.CreateCourse)
	router.PUT("/courses/update/:id", controllers.UpdateCourse)
	router.DELETE("/courses/delete/:id", controllers.DeleteCourse)
}
