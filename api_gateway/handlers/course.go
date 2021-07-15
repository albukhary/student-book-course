package handlers

import (
	"context"
	"log"
	"strconv"

	_ "github.com/albukhary/student-book-course/api_gateway/model"

	coursepb "github.com/albukhary/student-book-course/course_service/coursepb"

	"github.com/albukhary/student-book-course/api_gateway/clients"

	"github.com/gofiber/fiber/v2"

	"google.golang.org/protobuf/types/known/emptypb"

	// docs are generated by Swag CLI, you have to import them.
	_ "github.com/albukhary/student-book-course/api_gateway/docs"
)

/*-------------------------------Course Service Handlers------------------------------------*/

// CreateCourse godoc
// @tags Course Related
// @Summary Creates a Course record with user input details and writes into database
// @Accept json
// @Produce json
// @Param details body model.CreateCourseModel true "Course details"
// @Success 200 {object} model.Course
// @Router /create/course [post]
func CreateCourse(c *fiber.Ctx) error {
	var course coursepb.Course

	c.BodyParser(&course)

	req := &coursepb.CreateCourseRequest{
		Course: &course,
	}

	res, err := clients.CourseServiceClient.CreateCourse(context.Background(), req)
	if err != nil {
		log.Fatalf("Error creating a course from gRPC Book Service Server: %v", err)
	}

	return c.JSON(res.Course)
}

// GetCourse godoc
// @tags Course Related
// @Summary Gets details of the Course from User input book ID
// @Produce json
// @Param id path integer true "Course ID"
// @Success 200 {object} model.Course
// @Router /course/{id} [get]
func GetCourse(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// ID is initially a string when we get it from JSON
	// convert into int to use in a query
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	// Create gRPC request
	req := &coursepb.GetCourseRequest{
		Id: int32(id),
	}

	// Make a call and get gRPC response
	res, err := clients.CourseServiceClient.GetCourse(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling getCourse RPC. Error : %v\n", err)
	}

	return c.JSON(res.Course)
}

// GetAllCourses godoc
// @tags Course Related
// @Summary Gets the list of all the courses
// @Produce json
// @Success 200 {object} []model.Course
// @Router /courses [get]
func GetAllCourses(c *fiber.Ctx) error {

	// Make a call and get gRPC response
	res, err := clients.CourseServiceClient.GetAllCourses(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error calling getAllCourses RPC. Error : %v\n", err)
	}

	return c.JSON(res.Courses)
}

// UpdateCourse godoc
// @tags Course Related
// @Summary Updates the course record with user input details
// @Accept json
// @Produce json
// @Param details body model.Course true "Course details"
// @Success 200 {object} model.Course
// @Router /update/course [put]
func UpdateCourse(c *fiber.Ctx) error {

	var course coursepb.Course

	c.BodyParser(&course)

	req := &coursepb.UpdateCourseRequest{
		Course: &course,
	}

	// Make a call and get gRPC response
	res, err := clients.CourseServiceClient.UpdateCourse(context.Background(), req)
	if err != nil {
		log.Fatalf("Error updateCourse RPC. Error : %v\n", err)
	}

	return c.JSON(res.Course)
}

// DeleteCourse godoc
// @tags Course Related
// @Summary Deletes the course with the given ID
// @Produce json
// @Param id path integer true "Course ID"
// @Success 200 {object} model.Course
// @Router /delete/course/{id} [delete]
func DeleteCourse(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// ID is initially a string when we get it from JSON
	// convert into int to use in a query
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	// Create gRPC request
	req := &coursepb.DeleteCourseRequest{
		Id: int32(id),
	}

	// Make a call and get gRPC response
	res, err := clients.CourseServiceClient.DeleteCourse(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling deleteCourse RPC. Error : %v\n", err)
	}

	return c.JSON(res.Course)
}

// GetEnrolledStudents godoc
// @tags Course Related
// @Summary Gets the list of all students who are enrolled to a particular course
// @Produce json
// @Param id path integer true "Course ID"
// @Success 200 {object} []model.Student
// @Router /students/course/{id} [get]
func GetEnrolledStudents(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// ID is initially a string when we get it from JSON
	// convert into int to use in a query
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	// Create gRPC request
	req := &coursepb.GetEnrolledStudentsRequest{
		CourseId: int32(id),
	}

	// Make a call and get gRPC response
	res, err := clients.CourseServiceClient.GetEnrolledStudents(context.Background(), req)
	if err != nil {
		log.Fatalf("Error calling getEnrolledStudents RPC.\n Error : %v\n", err)
	}

	return c.JSON(res.Students)
}