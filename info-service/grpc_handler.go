package main
import (
	"context"
	pb "info-service/info"
)
type handler struct {
	pb.UnimplementedInfoServiceServer
}
func (s *handler) Create(ctx context.Context,e *pb.Empty) (*pb.Response,error) {
	_,mes:=serve("create")
	return &pb.Response{Response: mes},nil
}
func (s *handler) Get(ctx context.Context,e *pb.Empty) (*pb.Response,error) {
	_,mes:=serve("get")
	return &pb.Response{Response: mes},nil
}
func (s *handler) Give(ctx context.Context,e *pb.Empty) (*pb.Response,error) {
	_,mes:=serve("give")
	return &pb.Response{Response: mes},nil
}
func (s *handler) Home(ctx context.Context,e *pb.Empty) (*pb.Response,error) {
	_,mes:=serve("home")
	return &pb.Response{Response: mes},nil
}


