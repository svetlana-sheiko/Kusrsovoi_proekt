package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	pb1 "libraryClientApi/client-library/book"
	pb2 "libraryClientApi/client-library/info"
	pb3 "libraryClientApi/client-library/user"
)

const (
	port=":50050"
	address1 ="localhost:50051"
	address2 ="localhost:50052"
	address3 ="localhost:50053"
)

func main() {
		conn1, err := grpc.Dial(address1, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Не могу подключиться: %v", err)
		}
		defer conn1.Close()

		client1 := pb1.NewLibraryServiceClient(conn1)

		conn2, err := grpc.Dial(address2, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Не могу подключиться: %v", err)
		}
		defer conn2.Close()

		client2 := pb2.NewInfoServiceClient(conn2)
		conn3, err := grpc.Dial(address3, grpc.WithInsecure())
		if err != nil {
			fmt.Printf("Не могу подключиться: %v", err)
		}
		defer conn2.Close()
		client3 := pb3.NewUserServiceClient(conn3)
		server := &EchoServer{echo.New(), client1, client2, client3}
		err=server.Register()
		if err!=nil {
			fmt.Printf("Register echo error")
			return
		}
		err = server.Echo.Start(port)
		if err != nil {
			fmt.Printf(err.Error())
			return
		}

}

