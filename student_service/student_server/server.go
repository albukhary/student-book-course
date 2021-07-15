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

	coursepb "github.com/albukhary/student-book-course/course_service/coursepb"

	bookpb "github.com/albukhary/student-book-course/book_service/bookpb"

	"google.golang.org/grpc"
)

// Declare pointer variable to database
var db *sqlx.DB
var err error

var bookServiceClient bookpb.BookServiceClient
var courseServiceClient coursepb.CourseServiceClient

const (
	insertStudent  = `INSERT INTO student (first_name, last_name, email) VALUES ($1, $2, $3);`
	updateStudent  = `UPDATE student SET first_name=$1, last_name=$2, email=$3 WHERE student_id=$4;`
	deleteQuery    = `DELETE FROM student WHERE student_id=$1`
	courseIDsQuery = `SELECT course_id FROM student_course WHERE student_id = $1`
	getByEmail     = `SELECT student_id, first_name, last_name, email FROM student WHERE email = $1`
	getById        = `SELECT student_id, first_name, last_name, email FROM student WHERE student_id = $1`
)

type server struct {
	pb.UnimplementedStudentServiceServer
}

func (*server) CreateStudent(stx context.Context, req *pb.CreateStudentRequest) (*pb.CreateStudentResponse, error) {

	student := req.Student

	// Execute the query
	db.MustExec(insertStudent, student.FirstName, student.LastName, student.Email)

	// fetch the student details from database
	err := db.QueryRow(getByEmail, student.Email).Scan(&student.Id, &student.FirstName, &student.LastName, &student.Email)
	if err != nil {
		log.Fatal(err)
	}

	res := &pb.CreateStudentResponse{
		Student: student,
	}

	return res, nil
}

func (*server) GetStudent(ctx context.Context, req *pb.GetStudentRequest) (*pb.GetStudentResponse, error) {

	var student pb.Student

	// fetch the student details from database
	err := db.QueryRow(getById, req.Id).Scan(&student.Id, &student.FirstName, &student.LastName, &student.Email)
	if err != nil {
		log.Fatal(err)
	}

	res := &pb.GetStudentResponse{
		Student: &student,
	}

	return res, nil
}

func (*server) GetAllStudents(ctx context.Context, req *empty.Empty) (*pb.GetAllStudentsResponse, error) {

	// fetch the students rows from database
	rows, err := db.Query("SELECT student_id, first_name, last_name, email FROM student ORDER BY student_id ASC")
	if err != nil {
		log.Fatal(err)
	}

	var student pb.Student
	var students []*pb.Student

	for rows.Next() {
		err = rows.Scan(&student.Id, &student.FirstName, &student.LastName, &student.Email)
		if err != nil {
			log.Fatal(err)
		}
		students = append(students, &pb.Student{
			Id:        student.Id,
			FirstName: student.FirstName,
			LastName:  student.LastName,
			Email:     student.Email,
		})
	}

	res := &pb.GetAllStudentsResponse{
		Students: students,
	}

	return res, nil
}

func (*server) UpdateStudent(ctx context.Context, req *pb.UpdateStudentRequest) (*pb.UpdateStudentResponse, error) {
	var student pb.Student

	// Write details from gRPC request to a strcut
	student.Id = req.Student.GetId()
	student.FirstName = req.Student.GetFirstName()
	student.LastName = req.Student.GetLastName()
	student.Email = req.Student.GetEmail()

	// Insert the student into the database
	db.MustExec(updateStudent, student.FirstName, student.LastName, student.Email, student.Id)

	// fetch the student details from database
	err := db.QueryRow(getById, student.Id).Scan(&student.Id, &student.FirstName, &student.LastName, &student.Email)
	if err != nil {
		log.Fatal(err)
	}
	res := &pb.UpdateStudentResponse{
		Student: req.Student,
	}

	return res, nil
}

func (*server) DeleteStudent(ctx context.Context, req *pb.DeleteStudentRequest) (*pb.DeleteStudentResponse, error) {

	var student pb.Student

	student.Id = req.Id

	// find the requested student from database
	err := db.QueryRow(getById, student.Id).Scan(&student.Id, &student.FirstName, &student.LastName, &student.Email)
	if err != nil {
		fmt.Println("There is no student with such ID")
		log.Fatal(err)
	}

	//execute deletion
	db.MustExec(deleteQuery, student.Id)

	res := &pb.DeleteStudentResponse{
		Student: &student,
	}

	return res, nil
}

func (*server) GetEnrolledCourses(ctx context.Context, req *pb.GetEnrolledCoursesRequest) (*pb.GetEnrolledCoursesResponse, error) {

	var course_ids []int

	studentId := req.StudentId

	// Use db.Select() to write all the rows in a slice
	err := db.Select(&course_ids, "SELECT course_id FROM student_course WHERE student_id = $1 ORDER BY course_id ASC", studentId)
	if err != nil {
		log.Fatal(err)
	}

	var messageCourses []*pb.Course

	for _, course_id := range course_ids {

		courseReq := &coursepb.GetCourseRequest{
			Id: int32(course_id),
		}

		courseRes, err := courseServiceClient.GetCourse(context.Background(), courseReq)
		if err != nil {
			log.Fatal(err)
		}

		messageCourses = append(messageCourses, &pb.Course{
			CourseId:   courseRes.Course.CourseId,
			Instructor: courseRes.Course.Instructor,
			Title:      courseRes.Course.Title,
		})

	}

	// generate gRPC response
	res := &pb.GetEnrolledCoursesResponse{
		Courses: messageCourses,
	}

	return res, nil
}

func (*server) GetBorrowedBooks(ctx context.Context, req *pb.GetBorrowedBooksRequest) (*pb.GetBorrowedBooksResponse, error) {

	var book_ids []int

	studentId := req.StudentId

	// Use db.Select() to write all the rows in a slice
	query := fmt.Sprintf("SELECT book_id FROM student_book WHERE student_id=%v ORDER BY book_id ASC", studentId)
	err := db.Select(&book_ids, query)
	if err != nil {
		log.Fatal(err)
	}

	var messageBooks []*pb.Book

	for _, book_id := range book_ids {

		bookReq := &bookpb.GetBookRequest{
			Id: int32(book_id),
		}

		bookRes, err := bookServiceClient.GetBook(context.Background(), bookReq)
		if err != nil {
			log.Fatal(err)
		}

		messageBooks = append(messageBooks, &pb.Book{
			Id:     bookRes.Book.Id,
			Title:  bookRes.Book.Title,
			Author: bookRes.Book.Author,
		})
	}

	// generate gRPC response
	res := &pb.GetBorrowedBooksResponse{
		Books: messageBooks,
	}

	return res, nil
}

func (*server) BorrowBook(ctx context.Context, req *pb.BorrowBookRequest) (*pb.BorrowBookResponse, error) {
	var book pb.Book

	studentId := req.StudentId

	book.Id = req.BookId

	_, err := db.Exec("INSERT INTO student_book (student_id, book_id) VALUES ($1, $2)", studentId, book.Id)
	if err != nil {
		log.Printf("There is no student or book with such Id.\n Error : %v\n", err)
		res := &pb.BorrowBookResponse{}
		return res, nil
	}

	//  Retrieve the book details and make RPC response
	// fetch the book details from database
	err = db.QueryRow("SELECT book_id, title, author FROM book WHERE book_id = $1", book.Id).Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare a RPC response which sends newly Borrowed book details
	res := &pb.BorrowBookResponse{
		Book: &book,
	}

	return res, nil
}

func (*server) HandInBook(ctx context.Context, req *pb.HandInBookRequest) (*pb.HandInBooksResponse, error) {
	student_id := req.StudentId

	var book pb.Book
	book.Id = req.BookId

	db.MustExec("DELETE FROM student_book WHERE student_id = $1 AND book_id = $2", student_id, book.Id)

	//  Retrieve the book details and make RPC response
	// fetch the student details from database
	err = db.QueryRow("SELECT book_id, title, author FROM book WHERE book_id = $1", book.Id).Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		log.Fatal(err)
	}

	res := &pb.HandInBooksResponse{
		Book: &book,
	}

	return res, nil
}

func (*server) EnrollCourse(ctx context.Context, req *pb.EnrollCourseRequest) (*pb.EnrollCourseResponse, error) {
	var course pb.Course

	studentId := req.StudentId

	course.CourseId = req.CourseId

	_, err := db.Exec("INSERT INTO student_course (student_id, course_id) VALUES ($1, $2)", studentId, course.CourseId)
	if err != nil {
		log.Printf("There is no student or course with such Id.\n Error : %v\n", err)
		res := &pb.EnrollCourseResponse{}
		return res, nil
	}

	//  Retrieve the book details and make RPC response
	// fetch the course details from database
	err = db.QueryRow("SELECT course_id, instructor, title FROM course WHERE course_id = $1", course.CourseId).Scan(&course.CourseId, &course.Instructor, &course.Title)
	if err != nil {
		log.Fatal(err)
	}
	// Prepare a RPC response which sends newly Borrowed book details
	res := &pb.EnrollCourseResponse{
		Course: &course,
	}

	return res, nil
}

func (*server) DropCourse(ctx context.Context, req *pb.DropCourseRequest) (*pb.DropCourseResponse, error) {
	student_id := req.StudentId

	var course pb.Course
	course.CourseId = req.CourseId

	db.MustExec("DELETE FROM student_course WHERE student_id = $1 AND course_id = $2", student_id, course.CourseId)

	//  Retrieve the book details and make RPC response
	// fetch the course details from database
	err = db.QueryRow("SELECT course_id, instructor, title FROM course WHERE course_id = $1", course.CourseId).Scan(&course.CourseId, &course.Instructor, &course.Title)
	if err != nil {
		log.Fatal(err)
	}

	// Prepare a RPC response which sends newly Borrowed book details
	res := &pb.DropCourseResponse{
		Course: &course,
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

	/*-----------------------------------------------------------------------------------------------*/
	/*-------------------------------Connection for Book Services------------------------------------*/
	/*-----------------------------------------------------------------------------------------------*/
	// Open an INSECURE client connection(cc) with Book Service Server
	bookServiceConnection, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect %v\n", err)
	}
	defer bookServiceConnection.Close()

	// Register our client to that Dialing connection
	bookServiceClient = bookpb.NewBookServiceClient(bookServiceConnection)

	/*-----------------------------------------------------------------------------------------------*/
	/*-------------------------------Connection for Course Services------------------------------------*/
	/*-----------------------------------------------------------------------------------------------*/
	// Open an INSECURE client connection(cc) with Book Service Server
	courseServiceConnection, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect %v\n", err)
	}
	defer courseServiceConnection.Close()

	courseServiceClient = coursepb.NewCourseServiceClient(courseServiceConnection)

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
