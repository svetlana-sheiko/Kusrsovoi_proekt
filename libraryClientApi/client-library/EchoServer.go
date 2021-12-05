package main

import (
	"context"
	"github.com/labstack/echo/v4"
	pb1 "libraryClientApi/client-library/book"
	pb2 "libraryClientApi/client-library/info"
	pb3 "libraryClientApi/client-library/user"
	"net/http"
	"strconv"
	"time"
)

type EchoServer struct {
	*echo.Echo
	pb1.LibraryServiceClient
	pb2.InfoServiceClient
	pb3.UserServiceClient
}

func (e *EchoServer) Register() error {
	e.GET("/user/getAll", e.getAll)
	e.GET("/user/getById", e.getUserById)
	e.GET("/signIn", e.signIn)
	e.GET("/library/create", e.createBook)
	e.GET("/library/get", e.getBook)
	e.GET("/library/give", e.giveBook)
	e.GET("/info/create", e.createInfo)
	e.GET("/info/give", e.giveInfo)
	e.GET("/info/get", e.getInfo)
	e.GET("/info", e.homeInfo)
	return nil
}
func (e *EchoServer) getUserById(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return err
	}
	res, err := e.UserServiceClient.GetById(context.Background(), &pb3.Id{Id: int32(ID)})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) getAll(c echo.Context) error {
	res, err := e.UserServiceClient.GetAll(context.Background(), &pb3.Empty{})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) createUser(c echo.Context) error {
	u := &pb3.User{Name: c.QueryParam("name"),
		Email:    c.QueryParam("email"),
		Password: c.QueryParam("password")}
	res, err := e.UserServiceClient.CreateUser(context.Background(), u)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) deleteUser(c echo.Context) error {
	ID, err := strconv.Atoi(c.QueryParam("id"))
	if err != nil {
		return err
	}
	res, err := e.UserServiceClient.DeleteUser(context.Background(), &pb3.Id{Id: int32(ID)})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) signIn(c echo.Context) error {
	name := c.QueryParam("name")
	pwd := c.QueryParam("password")
	res, err := e.UserServiceClient.SignIn(context.Background(), &pb3.SignInInfo{Name: name, Password: pwd})
	if err != nil {
		c.SetCookie(&http.Cookie{Name: "token", Value: "", HttpOnly: true, Expires: time.Now().Add(10 * time.Minute)})
		return c.String(http.StatusBadRequest, "There is no such user")
	}
	c.SetCookie(&http.Cookie{Name: "token", Value: res.Response, HttpOnly: true, Expires: time.Now().Add(10 * time.Minute)})
	return c.String(http.StatusAccepted, "You entered successfully")
}
func (e *EchoServer) createBook(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use library")
	}
	token := cookies.Value
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		return err
	}

	res, err := e.LibraryServiceClient.CreateBook(context.Background(), &pb1.Request{Book: &pb1.Book{
		Name: c.QueryParam("name"), Author: c.QueryParam("author"),
		Year: int32(year),
	}, Token: token})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Resp)
}
func (e *EchoServer) getBook(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use library")
	}
	token := cookies.Value
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		return err
	}
	res, err := e.LibraryServiceClient.GetBook(context.Background(), &pb1.Request{Book: &pb1.Book{
		Name: c.QueryParam("name"), Year: int32(year), Author: c.QueryParam("author"),
	}, Token: token})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Resp)
}
func (e *EchoServer) giveBook(c echo.Context) error {
	cookies, err := c.Request().Cookie("token")
	if err != nil {
		return c.String(http.StatusMethodNotAllowed, "Please sign in to use library")
	}
	token := cookies.Value
	year, err := strconv.Atoi(c.QueryParam("year"))
	if err != nil {
		return err
	}
	res, err := e.LibraryServiceClient.GiveBook(context.Background(), &pb1.Request{Book: &pb1.Book{
		Name: c.QueryParam("name"), Year: int32(year), Author: c.QueryParam("author"),
	}, Token: token})
	if err != nil {
		return err
	}
	return c.String(http.StatusAccepted, res.Resp)
}
func (e *EchoServer) createInfo(c echo.Context) error {
	res, err := e.InfoServiceClient.Create(context.Background(), &pb2.Empty{})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) giveInfo(c echo.Context) error {
	res, err := e.InfoServiceClient.Give(context.Background(), &pb2.Empty{})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) getInfo(c echo.Context) error {
	res, err := e.InfoServiceClient.Get(context.Background(), &pb2.Empty{})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
func (e *EchoServer) homeInfo(c echo.Context) error {
	res, err := e.InfoServiceClient.Home(context.Background(), &pb2.Empty{})
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, res.Response)
}
