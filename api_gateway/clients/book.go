package clients

import (
	"log"

	bookpb "github.com/albukhary/student-book-course/book_service/bookpb"

	"google.golang.org/grpc"
)

var BookServiceClient bookpb.BookServiceClient
var BookServiceConnection *grpc.ClientConn

func SetBookClient() {
	/*-----------------------------------------------------------------------------------------------*/
	/*-------------------------------Connection for Book Services------------------------------------*/
	/*-----------------------------------------------------------------------------------------------*/
	// Open an INSECURE client connection(cc) with Book Service Server
	BookServiceConnection, err := grpc.Dial("localhost:50052", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Client Could not connect %v\n", err)
	}

	// Register our client to that Dialing connection
	BookServiceClient = bookpb.NewBookServiceClient(BookServiceConnection)
}
