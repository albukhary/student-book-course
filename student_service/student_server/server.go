package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/golang/protobuf/ptypes/empty"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	pb "github.com/albukhary/student-book-course/student_service/studentpb"

	"google.golang.org/grpc"
)

type Student struct {
	ID        int    `db:"student_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

type Course struct {
	course_id  int32  `db:"course_id"`
	instructor string `db:"instructor"`
	title      string `db:"title"`
}

type Book struct {
	book_id int32  `db:"book_id"`
	title   string `db:"title"`
	author  string `db:"author"`
}

// Declare pointer variable to database
var db *sqlx.DB
var err error

const (
	insertStudent  = `INSERT INTO student (id, first_name, last_name, email) VALUES ($1, $2, $3, $4);`
	updateStudent  = `UPDATE student SET name=$1, email=$2, age=$3 WHERE id=$4;`
	deleteQuery    = `DELETE FROM student WHERE id=$1`
	courseIDsQuery = `SELECT * FROM student_course WHERE student_id = $1`
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (*server) CreateStudent(stx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {

	// Get new student details from Client
	student := &Student{
		ID:        int(req.GetStudent().GetId()),
		FirstName: req.GetStudent().FirstName,
		LastName:  req.GetStudent().LastName,
		Email:     req.Student.GetEmail(),
	}

	// Execute the query
	db.MustExec(insertStudent, student.ID, student.FirstName, student.LastName, student.Email)

	// create a response
	res := &pb.CreateStudentResponse{
		Student: &pb.Student{
			Id:        int32(student.ID),
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Email:     student.Email,
		},
	}

	return res, nil
}

func (*server) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {

	var student Student

	studentID := req.GetId()

	err = db.Get(&student, "SELECT id, first_name, last_name, email FROM student WHERE id=$1", studentID)
	if err != nil {
		log.Fatal(err)
	}

	res := &pb.GetStudentResponse{
		Student: &pb.Student{
			Id:        int32(student.ID),
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Email:     student.Email,
		},
	}

	return res, nil
}

func (*server) GetAllStudents(ctx context.Context, req *empty.Empty) (*pb.GetAllStudentsResponse, error) {
	var students []Student

	// Use db.Select() to write all the rows in a slice
	err := db.Select(&students, "SELECT * FROM student")
	if err != nil {
		log.Fatal(err)
	}

	var messageStudents []*pb.Student
	var messageStudent pb.Student

	//traversy the query results to write all tehir content to the proto message
	for _, student := range students {

		messageStudent.Id = int32(student.ID)
		messageStudent.FirstName = student.FirstName
		messageStudent.LastName = student.LastName
		messageStudent.Email = student.Email

		messageStudents = append(messageStudents, &messageStudent)
	}

	// generate gRPC response
	res := &pb.GetAllStudentsResponse{
		Students: messageStudents,
	}

	return res, nil
}

func (*server) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentResponse, error) {
	var student Student

	// Write details from gRPC request to a strcut
	student.ID = int(req.Student.GetId())
	student.FirstName = req.Student.GetFirstName()
	student.LastName = req.Student.GetLastName()
	student.Email = req.Student.GetEmail()

	// Insert the student into the database
	db.MustExec(updateStudent, student.ID, student.FirstName, student.LastName, student.Email)

	res := &pb.UpdateStudentResponse{
		Student: &pb.Student{
			Id:        int32(student.ID),
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Email:     student.Email,
		},
	}

	return res, nil
}

func (*server) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {
	var student Student

	student.ID = int(req.Id)

	// find the requested student from database
	err = db.Get(&student, "SELECT id, name, email, age FROM student WHERE id=$1", student.ID)
	if err != nil {
		fmt.Println("There is no student with such ID")
		log.Fatal(err)
	}

	//execute deletion
	db.MustExec(deleteQuery, student.ID)

	res := &pb.DeleteStudentResponse{
		Student: &pb.Student{
			Id:        int32(student.ID),
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Email:     student.Email,
		},
	}

	return res, nil
}

func (*server) GetEnrolledCourses(ctx context.Context, req *pb.GetEnrolledCoursesRequest) (*pb.GetEnrolledCoursesResponse, error) {
	var student Student
	type student_course struct {
		student_id int32
		course_id  int32
	}
	var student_courses []student_course

	var course Course

	student.ID = int(req.StudentId)

	// Use db.Select() to write all the rows in a slice
	err := db.Select(&student_courses, "SELECT * FROM student_course")
	if err != nil {
		log.Fatal(err)
	}

	var messageCourses []*pb.Course
	var messageCourse pb.Course

	for _, student_course := range student_courses {

		// Use db.Select() to write all the rows in a slice
		err := db.Get(&course, "SELECT * FROM course WHERE course_id = $1", student_course.course_id)
		if err != nil {
			log.Fatal(err)
		}

		messageCourse.CourseId = course.course_id
		messageCourse.Instructor = course.instructor
		messageCourse.Title = course.title

		messageCourses = append(messageCourses, &messageCourse)

	}

	// generate gRPC response
	res := &pb.GetEnrolledCoursesResponse{
		Courses: messageCourses,
	}

	return res, nil
}

func (*server) GetBorrowedBooks(ctx context.Context, req *pb.GetBorrowedBooksRequest) (*pb.GetBorrowedBooksResponse, error) {
	var student Student
	type student_book struct {
		student_id int32
		book_id    int32
	}
	var student_books []student_book

	var book Book

	student.ID = int(req.StudentId)

	// Use db.Select() to write all the rows in a slice
	err := db.Select(&student_books, "SELECT * FROM student_book")
	if err != nil {
		log.Fatal(err)
	}

	var messageBooks []*pb.Book
	var messageBook pb.Book

	for _, student_book := range student_books {

		// Use db.Select() to write all the rows in a slice
		err := db.Get(&book, "SELECT * FROM book WHERE book_id = $1", student_book.book_id)
		if err != nil {
			log.Fatal(err)
		}

		messageBook.Id = book.book_id
		messageBook.Title = book.title
		messageBook.Author = book.author

		messageBooks = append(messageBooks, &messageBook)

	}

	// generate gRPC response
	res := &pb.GetBorrowedBooksResponse{
		Books: messageBooks,
	}

	return res, nil
}

func main() {

	setUpDatabase()

	// Create a network listener, bind it to port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register a service to the gRPC server
	pb.RegisterStudentServiceServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve %v", err)
	}
}

func setUpDatabase() {
	//Loading environment variables for DATABASE connection
	dialect := os.Getenv("DIALECT")
	host := os.Getenv("HOST")
	dbPort := os.Getenv("DBPORT")
	user := os.Getenv("USER")
	dbName := os.Getenv("NAME")
	password := os.Getenv("PASSWORD")

	// Database connection string
	dbURI := fmt.Sprintf("host=%s user=%s dbname=%s sslmode=disable password=%s port=%s", host, user, dbName, password, dbPort)

	//open and connect to the database at the same time
	db, err = sqlx.Connect(dialect, dbURI)
	if err != nil {
		log.Fatalf("Database connection error : %v\n", err)
	}
}
