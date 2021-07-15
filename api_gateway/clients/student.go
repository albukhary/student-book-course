package clients

import (
	"log"

	studentpb "github.com/albukhary/student-book-course/student_service/studentpb"

	"google.golang.org/grpc"
)

var StudentServiceClient studentpb.StudentServiceClient
var StudentServiceConnection *grpc.ClientConn

func SetStudentClient() {
	/*-----------------------------------------------------------------------------------------------*/
	/*-------------------------------Connection for Student Services------------------------------------*/
	/*-----------------------------------------------------------------------------------------------*/
	// Open an INSECURE client connection(cc) with Student Service Server
	StudentServiceConnection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect %v\n", err)
	}

	// Register our client to that Dialing connection
	StudentServiceClient = studentpb.NewStudentServiceClient(StudentServiceConnection)
}
