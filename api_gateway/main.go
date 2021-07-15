package main

import (
	"github.com/albukhary/student-book-course/api_gateway/clients"

	"github.com/albukhary/student-book-course/api_gateway/routes"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	clients.SetBookClient()
	clients.SetStudentClient()
	clients.SetCourseClient()

	routes.SetUpRoutes(app)

	defer clients.BookServiceConnection.Close()
	defer clients.CourseServiceConnection.Close()
	defer clients.StudentServiceConnection.Close()

	app.Listen(":8084")
}
