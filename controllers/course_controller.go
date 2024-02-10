package controllers

import (
	"errors"
	"net/http"

	"github.com/esleirter/e-learning-api/models"

	"github.com/gin-gonic/gin"
)

// GetCourses retrieves all courses from the database
func GetAllCourses(c *gin.Context) {
	courses, err := models.GetAllCourses()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch courses"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

func GetCourse(c *gin.Context) {
	id := c.Param("id")
	courses, err := models.GetCourse(id)
	if err != nil {
		if errors.Is(err, models.ErrCourseNotFound) {
			// Devolver un código de estado 404 si el curso no existe
			c.JSON(http.StatusNotFound, gin.H{"error": "Course not found"})
			return
		}
		// Devolver un código de estado 500 para otros errores
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to fetch course"})
		return
	}
	c.JSON(http.StatusOK, courses)
}

// CreateCourse creates a new course
func CreateCourse(c *gin.Context) {
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := models.CreateCourse(&course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to create course"})
		return
	}

	c.JSON(http.StatusCreated, course)
}

// UpdateCourse updates an existing course
func UpdateCourse(c *gin.Context) {
	id := c.Param("id")
	var course models.Course
	if err := c.ShouldBindJSON(&course); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}

	if err := models.UpdateCourse(id, &course); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to update course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course updated successfully"})
}

// DeleteCourse deletes a course by ID
func DeleteCourse(c *gin.Context) {
	id := c.Param("id")

	if err := models.DeleteCourse(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to delete course"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"message": "Course deleted successfully"})
}
