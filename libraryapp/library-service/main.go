package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "libraryapp/library-service/book"
	"net"
)
const (
	port ="localhost:50051"

)
func main() {
	e:=echo.New()
	e.GET("/books/get",getBook)//getBook
	e.PUT("/books/put",giveBook)//giveBook
	e.GET("/books/put",giveBook)//giveBook
	e.POST("/books/create",createBook)//createBook
	e.GET("/books/create",createBook)
	e.DELETE("/books/delete",deleteBook)
	e.GET("/books/delete",deleteBook)
	fmt.Println("Введите 1,чтобы запустить grpc сервер\nВведите 2, чтобы запустить самостоятельный rest-сервер")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		s := grpc.NewServer()
		h := &handler{pb.UnimplementedLibraryServiceServer{}}
		pb.RegisterLibraryServiceServer(s, h)
		lis, err := net.Listen("tcp", port)

		if err != nil {
			fmt.Printf("failed to listen: %v", err)
		}
		defer db.Close()

		reflection.Register(s)
		if err := s.Serve(lis); err != nil {
			fmt.Printf("failed to serve: %v", err)
		}
	case 2:
		e.Logger.Fatal(e.Start(port))
	}
}