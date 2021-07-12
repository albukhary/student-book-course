package main

import (
	"context"
	"log"
	"strconv"

	studentpb "github.com/albukhary/student-book-course/student_service/studentpb"

	swagger "github.com/arsmn/fiber-swagger/v2"
	"github.com/gofiber/fiber/v2"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/types/known/emptypb"

	// docs are generated by Swag CLI, you have to import them.
	_ "github.com/albukhary/student-book-course/api_gateway/docs"
)

type Student struct {
	ID        int    `db:"student_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

type Course struct {
	CourseId   int    `db:"course_id"`
	Instructor string `db:"instructor"`
	Title      string `db:"title"`
}

type Book struct {
	BookId int    `db:"book_id"`
	Title  string `db:"title"`
	Author string `db:"author"`
}

type Student_Book_Ids struct {
	Book_id    int32
	Student_id int32
}

type Student_Course_Ids struct {
	Course_id  int32
	Student_id int32
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
	app.Get("/courses/student/:id", getEnrolledCourses)
	app.Get("/books/student/:id", getBorrowedBooks)
	app.Post("/create/student", createStudent)
	app.Post("/student/borrow/book", borrowBook)
	app.Post("/student/enroll/course", enrollCourse)
	app.Delete("/delete/student/:id", deleteStudent)
	app.Delete("/student/handin/book", handInBook)
	app.Delete("/student/drop/course", dropCourse)
	app.Put("update/student/:id", updateStudent)
}

// createStudent godoc
// @tags Student Related
// @Summary Creates a student record with user input details and writes into database
// @Accept json
// @Produce json
// @Param details body Student true "Student details"
// @Success 200 {object} Student
// @Router /create/student [post]
func createStudent(c *fiber.Ctx) error {
	var student Student

	c.BodyParser(&student)

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
		log.Fatalf("Error creating a student from gRPC Student Service Server: %v", err)
	}

	// return it as a JSON
	return c.JSON(res.Student)
}

// getStudents godoc
// @tags Student Related
// @Summary Retrieves the list of all students
// @Produce json
// @Success 200 {object} []Student
// @Router /students [get]
func getStudents(c *fiber.Ctx) error {

	var students []Student
	var student Student

	res, err := studentServiceClient.GetAllStudents(context.Background(), &emptypb.Empty{})
	if err != nil {
		log.Fatalf("Error getting all students via gRPC. Error : %v\n", err)
	}

	messageStudents := res.Students

	for _, messageStudent := range messageStudents {

		// traverse through the message Student list
		// and write to our own Student struct
		student.ID = int(messageStudent.Id)
		student.FirstName = messageStudent.FirstName
		student.LastName = messageStudent.LastName
		student.Email = messageStudent.Email

		// append to the list of all students
		students = append(students, student)
	}

	//return the slice of students to http
	return c.JSON(students)
}

// getStudent godoc
// @tags Student Related
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

	// Create gRPC request
	req := &studentpb.GetStudentRequest{
		Id: int32(id),
	}

	// Make RPC and get RPC response
	res, err := studentServiceClient.GetStudent(context.Background(), req)
	if err != nil {
		log.Fatalf("Error getting getStudent RPC. Error : %v\n", err)
	}
	// Read the details from gRPC response to our own Student struct
	student.ID = int(res.Student.Id)
	student.FirstName = res.Student.FirstName
	student.LastName = res.Student.LastName
	student.Email = res.Student.Email

	// Return the student to REST as JSON
	return c.JSON(student)
}

// deleteStudent godoc
// @tags Student Related
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

	req := &studentpb.DeleteStudentRequest{
		Id: int32(id),
	}

	res, err := studentServiceClient.DeleteStudent(context.Background(), req)
	if err != nil {
		log.Fatalf("Error getting deleteStudent RPC. Error : %v\n", err)
	}

	// Read the details from gRPC response to our own Student struct
	student.ID = int(res.Student.Id)
	student.FirstName = res.Student.FirstName
	student.LastName = res.Student.LastName
	student.Email = res.Student.Email

	return c.JSON(student)
}

// updateStudent godoc
// @tags Student Related
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

	req := &studentpb.UpdateStudentRequest{
		Student: &studentpb.Student{
			Id:        int32(student.ID),
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Email:     student.Email,
		},
	}

	res, err := studentServiceClient.UpdateStudent(context.Background(), req)
	if err != nil {
		log.Fatalf("Error getting updateStudent RPC. Error : %v\n", err)
	}

	// Read the details from gRPC response to our own Student struct
	student.ID = int(res.Student.Id)
	student.FirstName = res.Student.FirstName
	student.LastName = res.Student.LastName
	student.Email = res.Student.Email

	// print the newly updated student details
	return c.JSON(student)
}

// getEnrolledCourses godoc
// @tags Course Related
// @Summary Retrieves the enrolled courses of the student based on given ID
// @Produce json
// @Param id path integer true "Student ID"
// @Success 200 {object} []Course
// @Router /courses/student/{id} [get]
func getEnrolledCourses(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// ID is initially a string when we get it from JSON
	// convert into int to use in a query
	id, err1 := strconv.Atoi(idParam)
	if err1 != nil {
		log.Fatal(err1)
	}

	// Create RPC request
	req := &studentpb.GetEnrolledCoursesRequest{
		StudentId: int32(id),
	}

	// Send RPc request and get RPC response
	res, err := studentServiceClient.GetEnrolledCourses(context.Background(), req)
	if err != nil {
		log.Fatalf("Error getting enroller courses RPC. Error : %v\n", err)
	}

	var courses []Course
	var course Course
	// iterate through courses and write them to courses array struct
	for _, messageCourse := range res.Courses {

		course.CourseId = int(messageCourse.CourseId)
		course.Instructor = messageCourse.Instructor
		course.Title = messageCourse.Title

		courses = append(courses, course)
	}

	// reutrn coursesto the browser as JSON
	return c.JSON(courses)
}

// getBorrowedBooks godoc
// @tags Book Related
// @Summary Retrieves the borrowed books of the student based on given ID
// @Produce json
// @Param id path integer true "Student ID"
// @Success 200 {object} []Book
// @Router /books/student/{id} [get]
func getBorrowedBooks(c *fiber.Ctx) error {
	idParam := c.Params("id")

	// ID is initially a string when we get it from JSON
	// convert into int to use in a query
	id, err1 := strconv.Atoi(idParam)
	if err1 != nil {
		log.Fatal(err1)
	}

	// Create RPC request
	req := &studentpb.GetBorrowedBooksRequest{
		StudentId: int32(id),
	}

	// Send RPc request and get RPC response
	res, err := studentServiceClient.GetBorrowedBooks(context.Background(), req)
	if err != nil {
		log.Fatalf("Error getting borrowed books RPC. Error : %v\n", err)
	}

	var books []Book
	var book Book
	// iterate through books and write them to books array struct
	for _, messageBook := range res.Books {

		book.BookId = int(messageBook.Id)
		book.Title = messageBook.Title
		book.Author = messageBook.Author

		books = append(books, book)
	}

	// reutrn books to the browser as JSON
	return c.JSON(books)
}

// borrowBook godoc
// @tags Book Related
// @Summary Borrows the book of a given Id for a given student ID
// @Accept  json
// @Produce json
// @Param student_book_ids body Student_Book_Ids true "Book and Student IDs"
// @Success 200 {object} Book
// @Router /student/borrow/book [POST]
func borrowBook(c *fiber.Ctx) error {

	var ids Student_Book_Ids
	c.BodyParser(&ids)

	// Prepare a gRPC request
	req := &studentpb.BorrowBookRequest{
		StudentId: ids.Student_id,
		BookId:    ids.Book_id,
	}

	// Make a CAll
	res, err := studentServiceClient.BorrowBook(context.Background(), req)
	if err != nil {
		log.Fatalf("Error from gRPC borrow book call.\n %v\n", err)
	}

	var book Book

	book.BookId = int(res.Book.Id)
	book.Title = res.Book.Title
	book.Author = res.Book.Author

	return c.JSON(book)
}

// handInBook godoc
// @tags Book Related
// @Summary Hands in the book of a given Id for a given student ID
// @Accept  json
// @Produce json
// @Param student_book_ids body Student_Book_Ids true "Book and Student IDs"
// @Success 200 {object} Book
// @Router /student/handin/book [DELETE]
func handInBook(c *fiber.Ctx) error {
	var student_book_id Student_Book_Ids
	var book Book

	c.BodyParser(&student_book_id)

	req := &studentpb.HandInBookRequest{
		StudentId: student_book_id.Student_id,
		BookId:    student_book_id.Book_id,
	}

	res, err := studentServiceClient.HandInBook(context.Background(), req)
	if err != nil {
		log.Fatalf("Error HandInBook RPC. ERROR : %v\n", err)
	}

	book.BookId = int(res.Book.Id)
	book.Title = res.Book.Title
	book.Author = res.Book.Author

	return c.JSON(book)
}

// enrollCourse godoc
// @tags Course Related
// @Summary Enrolls the student of the given Id to the given course ID
// @Accept  json
// @Produce json
// @Param student_course_ids body Student_Course_Ids true "Student and Course IDs"
// @Success 200 {object} Course
// @Router /student/enroll/course [POST]
func enrollCourse(c *fiber.Ctx) error {
	var ids Student_Course_Ids
	c.BodyParser(&ids)

	// Prepare a gRPC request
	req := &studentpb.EnrollCourseRequest{
		StudentId: ids.Student_id,
		CourseId:  ids.Course_id,
	}

	// Make a CAll
	res, err := studentServiceClient.EnrollCourse(context.Background(), req)
	if err != nil {
		log.Fatalf("Error from gRPC enroll course call.\n %v\n", err)
	}

	var course Course

	course.CourseId = int(res.Course.CourseId)
	course.Title = res.Course.Title
	course.Instructor = res.Course.Instructor

	return c.JSON(course)
}

// dropCourse godoc
// @tags Course Related
// @Summary Drop the student of the given Id from the given course(ID)
// @Accept  json
// @Produce json
// @Param student_course_ids body Student_Course_Ids true "Student and Course IDs"
// @Success 200 {object} Course
// @Router /student/drop/course [DELETE]
func dropCourse(c *fiber.Ctx) error {
	var student_course_ids Student_Course_Ids
	var course Course

	c.BodyParser(&student_course_ids)

	req := &studentpb.DropCourseRequest{
		StudentId: student_course_ids.Student_id,
		CourseId:  student_course_ids.Course_id,
	}

	res, err := studentServiceClient.DropCourse(context.Background(), req)
	if err != nil {
		log.Fatalf("Error from gRPC DROP course call.\n %v\n", err)
	}

	course.CourseId = int(res.Course.CourseId)
	course.Title = res.Course.Title
	course.Instructor = res.Course.Instructor

	return c.JSON(course)
}
