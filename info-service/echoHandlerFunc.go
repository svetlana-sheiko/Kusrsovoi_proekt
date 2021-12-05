package main

import (
	"github.com/labstack/echo/v4"
	"info-service/CreateDBConnect"
	"net/http"
)
var db,_=CreateDBConnect.CreateClient(0)

func get(c echo.Context) error {
	return c.String(serve("get"))
}
func give(c echo.Context) error {
	return c.String(serve("give"))
}
func create(c echo.Context) error {
	return c.String(serve("create"))
}
func home(c echo.Context) error {
	return c.String(serve("home"))
}
func serve(param string) (int,string) {
	var resp1,resp2 string

	row:=db.QueryRow("SELECT * FROM info WHERE name=?",param)
	err:=row.Scan(&resp1,&resp2)
	if err!=nil {
		return http.StatusBadRequest,"Inner database error"+param
	}
	return http.StatusOK,resp2
}


