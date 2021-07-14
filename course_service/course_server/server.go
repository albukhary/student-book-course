package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jmoiron/sqlx"

	coursepb "github.com/albukhary/student-book-course/course_service/coursepb"

	"google.golang.org/grpc"
)

var db *sqlx.DB
var err error

type server struct {
	coursepb.UnimplementedCourseServiceServer
}

func (*server) CreateCourse(ctx context.Context, req *coursepb.CreateCourseRequest) (*coursepb.CreateCourseResponse, error) {
	var course coursepb.Course

	course.Instructor = req.Course.Instructor
	course.Title = req.Course.Title

	db.MustExec("INSERT INTO course (instructor, title) VALUES $1, $2", course.Instructor, course.Title)

	err := db.QueryRow("SELECT course_id, instructor, title FROM course WHERE instructor = $1 AND title = $2", course.Instructor, course.Title).Scan(&course.CourseId, &course.Instructor, &course.Title)
	if err != nil {
		log.Fatal(err)
	}

	res := &coursepb.CreateCourseResponse{
		Course: &course,
	}

	return res, nil
}

func main() {
	setUpDatabase()

	// Create a network listener, bind it to port 50051
	lis, err := net.Listen("tcp", "0.0.0.0:50053")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// Create a gRPC server
	s := grpc.NewServer()

	// Register a service to the gRPC server
	coursepb.RegisterCourseServiceServer(s, &server{})

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
