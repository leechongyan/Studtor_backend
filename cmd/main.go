package main

import (
	"net/http"

	gin "github.com/gin-gonic/gin"

	sqlite_controller "github.com/leechongyan/Studtor_backend/internal/controllers/sqlitecontroller"
)

func main() {
	// Set the router as the default one shipped with Gin
	router := gin.Default()

	// Serve 404 page when page not found
	router.NoRoute(func(c *gin.Context) {
		c.JSON(404, gin.H{"code": "PAGE_NOT_FOUND", "message": "Page not found"})
	})

	// Set up test endpoint "/ping"
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// This handler will match /user/1 but will not match /user/ or /user
	router.GET("/students/:id", func(c *gin.Context) {
		id := c.Param("id")
		name, err := sqlite_controller.SelectStudentByID(id)
		if err != nil {
			c.String(http.StatusBadRequest, "StatusBadRequest")
		}
		c.String(http.StatusOK, "Hello %s", name)
	})

	// Set up test endpoint to create random user -> to be converted to POST in future
	router.GET("/students/create/:name", func(c *gin.Context) {
		name := c.Param("name")
		err := sqlite_controller.InsertStudent(name)
		if err != nil {
			c.String(http.StatusBadRequest, "StatusBadRequest")
		}
		c.String(http.StatusOK, "Hello %s", name)
	})

	// Start and run the server
	router.Run(":8080")
}
