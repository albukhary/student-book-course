package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	bookpb "github.com/albukhary/student-book-course/book_service/bookpb"
	studentpb "github.com/albukhary/student-book-course/student_service/studentpb"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
)

// Declare pointer variable to database
var db *sqlx.DB
var err error

var studentServiceClient studentpb.StudentServiceClient

type bookServer struct {
	bookpb.UnimplementedBookServiceServer
}

const (
	insertBook      = `INSERT INTO book (title, author) VALUES ($1, $2)`
	selectBookTitle = `SELECT book_id, title, author FROM book WHERE author = $1 AND title = $2`
	selectBookId    = `SELECT book_id, title, author FROM book WHERE book_id = $1`
)

func (*bookServer) CreateBook(ctx context.Context, req *bookpb.CreateBookRequest) (*bookpb.CreateBookResponse, error) {

	var book bookpb.Book

	book.Author = req.Book.Author
	book.Title = req.Book.Title

	db.MustExec(insertBook, book.Title, book.Author)

	// fetch the book details from database
	err := db.QueryRow(selectBookTitle, book.Author, book.Title).Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		log.Fatal(err)
	}

	res := &bookpb.CreateBookResponse{
		Book: &book,
	}

	return res, nil
}

func (*bookServer) GetBook(ctx context.Context, req *bookpb.GetBookRequest) (*bookpb.GetBookResponse, error) {

	var book bookpb.Book

	book.Id = req.Id

	// fetch the book details from database
	rows, err := db.Query(selectBookId, book.Id)
	if err != nil {
		log.Fatal(err)
	}

	// iterate over each row
	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			log.Fatal(err)
		}
	}

	res := &bookpb.GetBookResponse{
		Book: &book,
	}

	return res, nil
}

func (*bookServer) UpdateBook(ctx context.Context, req *bookpb.UpdateBookRequest) (*bookpb.UpdateBookResponse, error) {
	var book bookpb.Book

	book.Id = req.Book.Id
	book.Author = req.Book.Author
	book.Title = req.Book.Title

	db.MustExec("UPDATE book SET title = $1, author = $2 WHERE book_id = $3", book.Title, book.Author, book.Id)

	err := db.QueryRow("SELECT title, author FROM book WHERE book_id = $1", book.Id).Scan(&book.Title, &book.Author)
	if err != nil {
		log.Fatal(err)
	}

	res := &bookpb.UpdateBookResponse{
		Book: &book,
	}

	return res, nil
}

func (*bookServer) DeleteBook(ctx context.Context, req *bookpb.DeleteBookRequest) (*bookpb.DeleteBookResponse, error) {
	var book bookpb.Book

	book.Id = req.Id

	err := db.QueryRow("SELECT book_id, title, author FROM book WHERE book_id = $1", book.Id).Scan(&book.Id, &book.Title, &book.Author)
	if err != nil {
		log.Fatalf("There is no book with such id. %v\n", err)
	}

	db.MustExec("DELETE FROM book WHERE book_id = $1", book.Id)

	res := &bookpb.DeleteBookResponse{
		Book: &book,
	}

	return res, nil
}

func (*bookServer) GetAllBooks(ctx context.Context, req *empty.Empty) (*bookpb.GetAllBooksResponse, error) {

	rows, err := db.Query("SELECT book_id, title, author FROM book")
	if err != nil {
		log.Fatal(err)
	}

	var books []*bookpb.Book
	var book bookpb.Book

	for rows.Next() {
		err = rows.Scan(&book.Id, &book.Title, &book.Author)
		if err != nil {
			log.Fatal(err)
		}

		books = append(books, &bookpb.Book{
			Id:     book.Id,
			Title:  book.Title,
			Author: book.Author,
		})
	}

	res := &bookpb.GetAllBooksResponse{
		Books: books,
	}

	return res, nil
}

func (*bookServer) GetBorrowingStudents(ctx context.Context, req *bookpb.GetBorrowingStudentsRequest) (*bookpb.GetBorrowingStudentsResponse, error) {
	bookId := req.BookId

	var studentIds []int
	fmt.Println("Book Id: ", req.BookId)
	err := db.Select(&studentIds, "SELECT student_id FROM student_book WHERE book_id=$1", bookId)
	if err != nil {
		log.Fatalf("There are no students who borrowed this book\n. %v\n", err)
	}

	var students []*bookpb.Student

	for _, studentId := range studentIds {

		studentReq := &studentpb.GetStudentRequest{
			Id: int32(studentId),
		}
		studentRes, err := studentServiceClient.GetStudent(context.Background(), studentReq)
		if err != nil {
			log.Fatal(err)
		}

		students = append(students, &bookpb.Student{
			Id:        studentRes.Student.Id,
			FirstName: studentRes.Student.FirstName,
			LastName:  studentRes.Student.LastName,
			Email:     studentRes.Student.Email,
		})
	}

	res := &bookpb.GetBorrowingStudentsResponse{
		Students: students,
	}

	return res, nil
}

func main() {
	setUpDatabase()

	/*---------------Establish connection with student service -------------------------*/
	// Open an INSECURE client connection(cc) with Student Service Server
	studentServiceConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect to Student Service %v\n", err)
	}
	defer studentServiceConnection.Close()

	// Register our client to that Dialing connection
	studentServiceClient = studentpb.NewStudentServiceClient(studentServiceConnection)

	// Create a network listener, bind it to port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50052")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register a service to the gRPC server
	bookpb.RegisterBookServiceServer(s, &bookServer{})

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
