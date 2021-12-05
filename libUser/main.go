package main

import (
	"database/sql"
	"fmt"
	"github.com/labstack/echo/v4"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	pb "libUser/user"
	"net"
)

const (
	port=":50053"
)
func main() {
	fmt.Println("Введите 1,чтобы запустить grpc-сервер.\nВведите 2, чтобы запустить rest-сервер")
	var choice int
	fmt.Scan(&choice)
	switch choice {
	case 1:
		s := grpc.NewServer()
		db,err:=sql.Open("mysql",config.FormatDSN())
		if err!=nil {
			fmt.Println("DB not found")
			return
		}
		h := &GRPCHandler{authService{NewRepository(db)},pb.UnimplementedUserServiceServer{}}
		pb.RegisterUserServiceServer(s,h)
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
		e:=echo.New()
		err:=SetEchoHandlers(e)
		if err!=nil {
			fmt.Println(err)
		}
		e.Logger.Error(e.Start(port))
	}
}
