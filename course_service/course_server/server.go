package main

import (
	"context"
	"fmt"
	"log"
	"net"
	"os"

	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"

	coursepb "github.com/albukhary/student-book-course/course_service/coursepb"

	studentpb "github.com/albukhary/student-book-course/student_service/studentpb"

	"google.golang.org/grpc"

	"github.com/golang/protobuf/ptypes/empty"
)

var db *sqlx.DB
var err error

var studentServiceClient studentpb.StudentServiceClient

type server struct {
	coursepb.UnimplementedCourseServiceServer
}

func (*server) CreateCourse(ctx context.Context, req *coursepb.CreateCourseRequest) (*coursepb.CreateCourseResponse, error) {
	var course coursepb.Course

	course.Instructor = req.Course.Instructor
	course.Title = req.Course.Title

	db.MustExec("INSERT INTO course (instructor, title) VALUES ($1, $2)", course.Instructor, course.Title)

	err := db.QueryRow("SELECT course_id, instructor, title FROM course WHERE instructor = $1 AND title = $2", course.Instructor, course.Title).Scan(&course.CourseId, &course.Instructor, &course.Title)
	if err != nil {
		log.Fatal(err)
	}

	res := &coursepb.CreateCourseResponse{
		Course: &course,
	}

	return res, nil
}

func (*server) GetCourse(ctx context.Context, req *coursepb.GetCourseRequest) (*coursepb.GetCourseResponse, error) {
	var course coursepb.Course

	course.CourseId = req.Id

	err := db.QueryRow("SELECT instructor, title FROM course WHERE course_id = $1", course.CourseId).Scan(&course.Instructor, &course.Title)
	if err != nil {
		log.Fatal(err)
	}

	res := &coursepb.GetCourseResponse{
		Course: &course,
	}

	return res, nil
}

func (*server) UpdateCourse(ctx context.Context, req *coursepb.UpdateCourseRequest) (*coursepb.UpdateCourseResponse, error) {
	var course coursepb.Course

	course.CourseId = req.Course.CourseId
	course.Instructor = req.Course.Instructor
	course.Title = req.Course.Title

	db.MustExec("UPDATE course SET instructor = $1, title = $2 WHERE course_id = $3", course.Instructor, course.Title, course.CourseId)

	err := db.QueryRow("SELECT instructor, title FROM course WHERE course_id = $1", course.CourseId).Scan(&course.Instructor, &course.Title)
	if err != nil {
		log.Fatal(err)
	}

	res := &coursepb.UpdateCourseResponse{
		Course: &course,
	}

	return res, nil
}

func (*server) DeleteCourse(ctx context.Context, req *coursepb.DeleteCourseRequest) (*coursepb.DeleteCourseResponse, error) {
	var course coursepb.Course

	course.CourseId = req.Id

	err := db.QueryRow("SELECT instructor, title FROM course WHERE course_id = $1", course.CourseId).Scan(&course.Instructor, &course.Title)
	if err != nil {
		log.Fatal("There is no course such course", err)
	}

	db.MustExec("DELETE FROM course WHERE course_id = $1", course.CourseId)

	res := &coursepb.DeleteCourseResponse{
		Course: &course,
	}

	return res, nil
}

func (*server) GetAllCourses(ctx context.Context, req *empty.Empty) (*coursepb.GetAllCoursesResponse, error) {
	var courses []*coursepb.Course

	rows, err := db.Query("SELECT course_id, instructor, title FROM course")
	if err != nil {
		log.Fatal(err)
	}

	for rows.Next() {

		var course coursepb.Course

		err = rows.Scan(&course.CourseId, &course.Instructor, &course.Title)
		if err != nil {
			log.Fatal(err)
		}

		courses = append(courses, &coursepb.Course{
			CourseId:   course.CourseId,
			Instructor: course.Instructor,
			Title:      course.Title,
		})
	}

	res := &coursepb.GetAllCoursesResponse{
		Courses: courses,
	}

	return res, nil
}

func (*server) GetEnrolledStudents(ctx context.Context, req *coursepb.GetEnrolledStudentsRequest) (*coursepb.GetEnrolledStudentsResponse, error) {

	courseId := req.CourseId

	var studentIds []int

	db.Select(&studentIds, "SELECT student_id from student_course WHERE course_id = $1", courseId)

	var students []*coursepb.Student
	for _, studentId := range studentIds {

		studentReq := &studentpb.GetStudentRequest{
			Id: int32(studentId),
		}

		studentRes, err := studentServiceClient.GetStudent(context.Background(), studentReq)
		if err != nil {
			log.Fatal(err)
		}

		students = append(students, &coursepb.Student{
			Id:        studentRes.Student.Id,
			FirstName: studentRes.Student.FirstName,
			LastName:  studentRes.Student.LastName,
			Email:     studentRes.Student.Email,
		})
	}

	res := &coursepb.GetEnrolledStudentsResponse{
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
