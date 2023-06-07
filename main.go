package main

import (
	"context"
	"fmt"
	"go-grpc-intro/bookpb/bookpb"
	"io"
	"log"
	"net"
	"os"
	"os/signal"

	"google.golang.org/grpc"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/reflection"
	"google.golang.org/grpc/status"
)

// Book represents book data in local storage
type Book struct {
	ID          string
	Title       string
	Description string
	Author      string
}

var database []Book = []Book{}

type Server struct {
	bookpb.UnimplementedBookServiceServer
}

// implement service server
func (srv *Server) CreateBook(ctx context.Context, request *bookpb.CreateBookRequest) (*bookpb.CreateBookResponse, error) {
	book := Book{
		ID:          request.GetId(),
		Title:       request.GetTitle(),
		Description: request.GetDescription(),
		Author:      request.GetAuthor(),
	}

	database = append(database, book)

	return &bookpb.CreateBookResponse{
		Status: true,
		Book: &bookpb.Book{
			Id:          book.ID,
			Title:       book.Title,
			Description: book.Description,
			Author:      book.Author,
		},
	}, nil
}

func (srv *Server) GetBooks(request *bookpb.GetBooksRequest, stream bookpb.BookService_GetBooksServer) error {
	for _, book := range database {
		stream.Send(&bookpb.GetBooksResponse{
			Status: true,
			Book: &bookpb.Book{
				Id:          book.ID,
				Title:       book.Title,
				Description: book.Description,
				Author:      book.Author,
			},
		})
	}

	return nil
}

func (srv *Server) CreateManyBooks(stream bookpb.BookService_CreateManyBooksServer) error {
	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return stream.SendAndClose(&bookpb.CreateManyBooksResponse{
				Status:  true,
				Message: "All book data inserted successfully!",
			})
		}
		if err != nil {
			return status.Errorf(
				codes.Internal,
				fmt.Sprintf("Internal error, insert batch failed: %v", err),
			)
		}

		var book Book = Book{
			ID:          req.GetId(),
			Title:       req.GetTitle(),
			Description: req.GetDescription(),
			Author:      req.GetAuthor(),
		}

		database = append(database, book)
	}
}

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	fmt.Println("Book service started")

	lis, err := net.Listen("tcp", "0.0.0.0:8080")
	if err != nil {
		log.Fatalf("Failed to listen: %v\n", err)
	}

	s := grpc.NewServer()

	bookpb.RegisterBookServiceServer(s, &Server{})
	reflection.Register(s)

	go func() {
		fmt.Println("Starting server...")
		if err := s.Serve(lis); err != nil {
			log.Fatalf("Failed to serve: %v", err)
		}
	}()

	// other codes...
	ch := make(chan os.Signal, 1)
	signal.Notify(ch, os.Interrupt)

	<-ch
	fmt.Println("Stopping the server..")
	s.Stop()
	fmt.Println("Stopping listener...")
	lis.Close()
	fmt.Println("End of Program")
}
