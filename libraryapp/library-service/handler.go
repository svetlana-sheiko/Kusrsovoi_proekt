package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	pb "libraryapp/library-service/book"
	"net/http"
	"strconv"
)

const (
	signingKey = "web324o8wr981r32nw3efweifjowief32423r"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}
type handler struct {
	pb.UnimplementedLibraryServiceServer
}

func (s *handler) CreateBook(ctx context.Context, in *pb.Request) (*pb.Empty, error) {
	token := in.Token
	if tokenCheck(token) {
		_, mes := createBookInner(in.Book.Name, in.Book.Author, strconv.Itoa(int(in.Book.Year)))
		return &pb.Empty{Resp: mes}, nil
	} else {
		return &pb.Empty{}, jwt.ValidationError{Errors: http.StatusNotAcceptable}
	}
}
func (s *handler) GetBook(ctx context.Context, in *pb.Request) (*pb.Empty, error) {
	token := in.Token
	if tokenCheck(token) {

		_, mes := getBookInner(in.Book.Name, in.Book.Author, strconv.Itoa(int(in.Book.Year)))
		return &pb.Empty{Resp: mes}, nil
	} else {
		return &pb.Empty{}, jwt.ValidationError{Errors: http.StatusNotAcceptable}
	}
}
func (s *handler) GiveBook(ctx context.Context, in *pb.Request) (*pb.Empty, error) {
	token := in.Token
	if tokenCheck(token) {
		_, mes := giveBookInner(in.Book.Name, in.Book.Author, strconv.Itoa(int(in.Book.Year)))
		return &pb.Empty{Resp: mes}, nil
	} else {
		return &pb.Empty{}, jwt.ValidationError{Errors: http.StatusNotAcceptable}
	}
}

func Decode(token string) (*tokenClaims, error) {
	tokenType, err := jwt.ParseWithClaims(token, &tokenClaims{}, func(token *jwt.Token) (interface{}, error) {
		return []byte(signingKey), nil
	})
	if err != nil {
		return nil, err
	}
	if claims, ok := tokenType.Claims.(*tokenClaims); ok && tokenType.Valid {
		fmt.Println(claims)
		return claims, nil
	} else {

		return nil, err
	}
}

func tokenCheck(token string) bool {
	tClaims, err := Decode(token)
	if err != nil {
		return false
	}
	fmt.Println(tClaims.UserId + " ID")
	jwtToken := jwt.NewWithClaims(jwt.SigningMethodHS256, tClaims)
	str, err := jwtToken.SignedString([]byte(signingKey))
	if token == str {
		return true
	} else {

		return false
	}
}
