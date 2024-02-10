package models

import (
	"context"
	"errors"
	"log"
	"time"

	"github.com/esleirter/e-learning-api/utils"
)

// Course represents the structure of a course
type Course struct {
	CourseId    int    `json:"course_id"`
	Name        string `json:"name"`
	Description string `json:"description"`
	CreatedAt   string `json:"created_at"`
	UpdatedAt   string `json:"updated_at"`
}

var ctx = context.Background()
var ErrCourseNotFound = errors.New("course not found")

// GetAllCourses retrieves all courses from the database
func GetAllCourses() ([]Course, error) {
	db := utils.DB
	// Prepare statement for query optimization and security
	stmt, err := db.PrepareContext(ctx, "SELECT course_id, name, description, created_at, updated_at FROM courses")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return nil, err
	}
	defer stmt.Close()

	// Execute query with context
	rows, err := stmt.QueryContext(ctx)
	if err != nil {
		log.Println("Error querying courses:", err)
		return nil, err
	}
	defer rows.Close()

	var courses []Course
	for rows.Next() {
		var course Course
		if err := rows.Scan(&course.CourseId, &course.Name, &course.Description, &course.CreatedAt, &course.UpdatedAt); err != nil {
			log.Println("Error scanning course row:", err)
			return nil, err
		}
		courses = append(courses, course)
	}

	return courses, nil
}

func GetCourse(id string) (Course, error) {
	db := utils.DB

	var course Course

	// Prepare statement for query optimization and security
	stmt, err := db.PrepareContext(ctx, "SELECT course_id, name, description, created_at, updated_at FROM courses WHERE course_id = ?")
	if err != nil {
		log.Println("Error preparing statement:", err)
		return course, err
	}
	defer stmt.Close()

	// Execute query with context
	rows, err := stmt.QueryContext(ctx, id)
	if err != nil {
		log.Println("Error querying courses:", err)
		return course, err
	}
	defer rows.Close()

	// Scan the result
	if rows.Next() {
		if err := rows.Scan(&course.CourseId, &course.Name, &course.Description, &course.CreatedAt, &course.UpdatedAt); err != nil {
			log.Println("Error scanning course row:", err)
			return course, err
		}
	} else {
		// No rows found with the given courseID
		return course, ErrCourseNotFound
	}

	return course, nil
}

// CreateCourse creates a new course in the database
func CreateCourse(course *Course) error {
	//ctx := context.Background()
	db := utils.DB

	query := "INSERT INTO courses (name, description, created_at, updated_at) VALUES (?, ?, ?, ?)"

	_, err := db.ExecContext(ctx, query, course.Name, course.Description, time.Now(), time.Now())
	if err != nil {
		log.Println("Error creating course:", err)
		return err
	}

	return nil
}

// UpdateCourse updates an existing course in the database
func UpdateCourse(id string, course *Course) error {
	db := utils.DB
	query := "UPDATE courses SET name = ?, description = ?, updated_at = ? WHERE course_id = ?"

	_, err := db.ExecContext(ctx, query, course.Name, course.Description, time.Now(), id)
	if err != nil {
		log.Println("Error updating course:", err)
		return err
	}

	return nil
}

// DeleteCourse deletes a course from the database
func DeleteCourse(id string) error {
	db := utils.DB
	query := "DELETE FROM courses WHERE course_id = ?"

	_, err := db.ExecContext(ctx, query, id)
	if err != nil {
		log.Println("Error deleting course:", err)
		return err
	}

	return nil
}
