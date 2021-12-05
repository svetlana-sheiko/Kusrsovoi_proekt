package main

import (
	"context"
	pb "libUser/user"
	"strconv"
)
type GRPCHandler struct {
	a authService
	pb.UnimplementedUserServiceServer
}
func (h *GRPCHandler) CreateUser(ctx context.Context,user *pb.User) (*pb.Response,error) {
	_,err:=h.a.CreateUser(Marshal(user))
	if err!=nil {
		return &pb.Response{},err
	}
	return &pb.Response{Response: "User was successfully created\n"},nil
}
func (h *GRPCHandler) DeleteUser(ctx context.Context,id *pb.Id) (*pb.Response,error) {
	ID:=strconv.Itoa(int(id.Id))
	err:=h.a.repository.Delete(context.Background(),ID)
	if err!=nil {
		return &pb.Response{},err
	}
	return &pb.Response{Response: "User was deleted"},nil
}
func (h *GRPCHandler) GetById(ctx context.Context,id *pb.Id) (*pb.Response,error) {
	ID:=strconv.Itoa(int(id.Id))
	u,err:=h.a.repository.GetById(context.Background(),ID)
	if err!=nil {
		return &pb.Response{},err
	}
	return &pb.Response{Response: u.String()},nil
}
func (h *GRPCHandler) GetAll(ctx context.Context,e *pb.Empty) (*pb.Response,error) {
	usersInfo:="UsersInfo:\n"
	users,err:=h.a.repository.GetAll(context.Background())
	if err!=nil {
		return &pb.Response{},err
	}
	for _,user:= range users {
		usersInfo+=user.String()
	}
	return &pb.Response{Response: usersInfo},nil
}
func (h *GRPCHandler) SignIn(ctx context.Context,info *pb.SignInInfo) (*pb.Response,error) {
	var token string
	_,err:=h.a.SignIn(info.Name,info.Password,&token)
	if err!=nil {
		return &pb.Response{},err
	}
	return &pb.Response{Response: token},nil
}

func Marshal(user *pb.User) *User {
	u:=&User{Name: user.Name,Password: user.Password,Email: user.Email}
	return u
}
