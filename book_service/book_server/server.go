package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	bookpb "github.com/albukhary/student-book-course/book_service/bookpb"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	"google.golang.org/grpc"
)

type Student struct {
	ID        int    `db:"student_id"`
	FirstName string `db:"first_name"`
	LastName  string `db:"last_name"`
	Email     string `db:"email"`
}

type Book struct {
	BookId int32  `db:"book_id"`
	Title  string `db:"title"`
	Author string `db:"author"`
}

// Declare pointer variable to database
var db *sqlx.DB
var err error

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

	db.MustExec(insertBook, book.Author, book.Title)

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

func main() {
	setUpDatabase()

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
