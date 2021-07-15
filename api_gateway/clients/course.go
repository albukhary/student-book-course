package clients

import (
	"log"

	coursepb "github.com/albukhary/student-book-course/course_service/coursepb"

	"google.golang.org/grpc"
)

var CourseServiceClient coursepb.CourseServiceClient
var CourseServiceConnection *grpc.ClientConn

func SetCourseClient() {
	/*-----------------------------------------------------------------------------------------------*/
	/*-------------------------------Connection for Course Services------------------------------------*/
	/*-----------------------------------------------------------------------------------------------*/
	// Open an INSECURE client connection(cc) with Book Service Server
	CourseServiceConnection, err := grpc.Dial("localhost:50053", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect %v\n", err)
	}

	CourseServiceClient = coursepb.NewCourseServiceClient(CourseServiceConnection)
}
