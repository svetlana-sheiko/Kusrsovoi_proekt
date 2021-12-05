package main

import (
	"context"
	"database/sql"
	"github.com/go-sql-driver/mysql"
	"github.com/labstack/echo/v4"
	"net/http"
)
var config = mysql.Config{
	User: "hekapoo",
	Passwd: "1234",
	Net: "tcp",
	Addr: "localhost:3306",
	DBName: "libtest",
	Collation: "",
}

type echoHandler struct {
	a authService
}
func(eHandler *echoHandler) getById(c echo.Context) error {
	u,err:=eHandler.a.repository.GetById(context.Background(),c.QueryParam("Id"))
	if err!=nil {
		return err
	}
	return c.String(http.StatusOK,u.String())
}
func (eHandler *echoHandler) create(c echo.Context) error {
	Name:=c.QueryParam("name")
	Email:=c.QueryParam("email")
	Pwd:=c.QueryParam("password")
	u:=User{Name: Name,Email: Email,Password: Pwd}
	i,err:=eHandler.a.CreateUser(&u)
	if err!=nil {
		er,ok:=err.(*ValidationError)
		if ok {
			return c.String(http.StatusBadRequest,er.Error())
		}else {
			return err
		}
	}
	return c.String(i,"User created!")
}
func (eHandler *echoHandler) getAll(c echo.Context) error {
	usersInfo:="UsersInfo:\n"
	users,err:=eHandler.a.repository.GetAll(context.Background())
	if err!=nil {
		return err
	}
	for _,user:= range users {
		usersInfo+=user.String()
	}
	return c.String(http.StatusOK,usersInfo)
}
func (eHandler *echoHandler) getByName(c echo.Context) error {
	u,err:=eHandler.a.repository.GetByName(context.Background(),c.QueryParam("name"))
	if err!=nil {
		return err
	}
	return c.String(http.StatusOK,u.String())
}
func (eHandler *echoHandler) delete(c echo.Context) error {
	err:=eHandler.a.repository.Delete(context.Background(),c.QueryParam("id"))
	if err!=nil {
		if er, ok := err.(ValidationError); ok {
			return c.String(er.Status, er.Message)
		}else {
			return err
		}
	}
	return nil
}
func (eHandler *echoHandler) signIn(c echo.Context) error {
	name:=c.QueryParam("name")
	pwd:=c.QueryParam("password")
	var token string=""
	res,err:=eHandler.a.SignIn(name, pwd,&token)
	if err!=nil {
		return err
	}
	return c.String(res,token)
}
func SetEchoHandlers(e *echo.Echo) error {
	db,err:=sql.Open("mysql",config.FormatDSN())
	if err!=nil {
		return err
	}
	eHandlers:=&echoHandler{a:authService{repository: NewRepository(db)}}
	e.GET("/getById",eHandlers.getById)
	e.GET("/getByName",eHandlers.getByName)
	e.POST("/create",eHandlers.create)
	e.GET("/create",eHandlers.create)
	e.GET("/getAll",eHandlers.getAll)
	e.DELETE("/delete",eHandlers.delete)
	e.GET("/delete",eHandlers.delete)
	e.GET("/signIn",eHandlers.signIn)
	return nil
}
