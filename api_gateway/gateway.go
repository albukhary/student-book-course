package main

import (
	"context"
	"fmt"
	"log"

	studentpb "github.com/albukhary/student-book-course/student_service/studentpb"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"google.golang.org/grpc"
)

type Student struct {
	ID        int    `db:"student_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

var studentServiceClient studentpb.StudentServiceClient

func main() {
	app := fiber.New()

	app.Get("/swagger/*", swagger.Handler)

	setupRoutes(app)

	// Open an INSECURE client connection(cc) with Student Service Server
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect %v\n", err)
	}
	defer cc.Close()

	// Register our client to that Dialing connection
	studentServiceClient = studentpb.NewStudentServiceClient(cc)

	app.Listen(":8084")
}

func setupRoutes(app *fiber.App) {

	app.Get("/", func(c *fiber.Ctx) error {
		return c.SendString("Assalamu alaykum 👋!")
	})
	app.Get("/students", getStudents)
	app.Get("/student/:id", getStudent)
	app.Post("/create/student", createStudent)
	app.Delete("/delete/student/:id", deleteStudent)
	app.Put("update/student/:id", updateStudent)
}

// createStudent godoc
// @Summary Creates a student record with user input details and writes into database
// @Accept json
// @Produce json
// @Param details body Student true "Student details"
// @Success 200 {object} Student
// @Router /create/student [post]
func createStudent(c *fiber.Ctx) error {
	var student Student

	// parse JSON to a student struct
	c.BodyParser(&student)

	fmt.Println(student)

	// create a request
	req := &studentpb.CreateStudentRequest{
		Student: &studentpb.Student{
			Id:        int32(student.ID),
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Email:     student.Email,
		},
	}

	// Invoke gRPC call to the StudentService Server
	res, err := studentServiceClient.CreateStudent(context.Background(), req)
	if err != nil {
		log.Fatal("Error creating a student from gRPC Student Service Server: %v", err)
	}

	// writethe details from gRPC response to the struct
	student.ID = int(res.GetStudent().GetId())
	student.FirstName = res.GetStudent().GetFirstName()
	student.LastName = res.GetStudent().GetLastName()
	student.Email = res.Student.GetEmail()

	// return it as a JSON
	return c.JSON(student)
}

// getStudents godoc
// @Summary Retrieves the list of all students
// @Produce json
// @Success 200 {object} []Student
// @Router /students [get]
func getStudents(c *fiber.Ctx) error {
	var students []Student

	// Use db.Select() to write all the rows in a slice
	err := db.Select(&students, "SELECT * FROM student")
	if err != nil {
		log.Fatal(err)
	}

	//return the slice of students to http
	return c.JSON(students)
}

// getStudent godoc
// @Summary Retrieves user based on given ID
// @Produce json
// @Param id path integer true "Student ID"
// @Success 200 {object} Student
// @Router /student/{id} [get]
func getStudent(c *fiber.Ctx) error {
	idParam := c.Params("id")

	var student Student

	// ID is initially a string when we get it from JSON
	// convert into int to use in a query
	id, err := strconv.Atoi(idParam)
	if err != nil {
		log.Fatal(err)
	}

	// query the database using db.Get()
	err = db.Get(&student, "SELECT id, name, email, age FROM student WHERE id=$1", id)
	if err != nil {
		log.Fatal(err)
	}

	return c.JSON(student)
}
// deleteStudent godoc
// @Summary Deletes a student with the specified ID
// @Produce json
// @Param id path integer true "Student ID"
// @Success 200 {object} Student
// @Router /delete/student/{id} [delete]
func deleteStudent(c *fiber.Ctx) error {
	idParam := c.Params("id")

	var student Student

	// ID is initially a string when we get it from JSON
	// convert into int to use in a query
	id, err1 := strconv.Atoi(idParam)
	if err1 != nil {
		log.Fatal(err1)
	}

	// find the requested student from database
	err = db.Get(&student, "SELECT id, name, email, age FROM student WHERE id=$1", id)
	if err != nil {
		fmt.Println("There is no student with such ID")
		log.Fatal(err)
	}

	deleteQuery := `DELETE FROM student WHERE id=$1`

	//execute deletion
	db.MustExec(deleteQuery, student.ID)

	return c.JSON(student)
}

// updateStudent godoc
// @Summary Updates a student record with user input details and writes into database
// @Accept json
// @Produce json
// @Param details body Student true "Updated Student Details"
// @Success 200 {object} Student
// @Router /update/student/{id} [put]
func updateStudent(c *fiber.Ctx) error {
	var student Student

	// parses JSON to struct
	c.BodyParser(&student)

	updateStudent := `UPDATE student SET name=$1, email=$2, age=$3 WHERE id=$4;`

	// Insert the student into the database
	db.MustExec(updateStudent, student.Name, student.Email, student.Age, student.ID)

	// print the newly updated student details
	return c.JSON(student)
}