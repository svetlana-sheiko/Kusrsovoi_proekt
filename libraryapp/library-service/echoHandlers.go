package main

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"libraryapp/CreateDBConnect"
	pb "libraryapp/library-service/book"
	"net/http"
)

var db, _ = CreateDBConnect.CreateClient(0)

func getBook(c echo.Context) error {
	return c.String(getBookInner(c.QueryParam("name"), c.QueryParam("author"), c.QueryParam("year")))
}
func deleteBook(c echo.Context) error {

	return c.String(deleteBookInner(c.QueryParam("id")))
}
func giveBook(c echo.Context) error {
	return c.String(giveBookInner(c.QueryParam("name"), c.QueryParam("author"), c.QueryParam("year")))
}

func createBook(c echo.Context) error {
	_, err := db.Exec("insert into libtest.books (name, year, available, author,section) values (?,?,?,?,?)",
		c.QueryParam("name"), c.QueryParam("year"), true, c.QueryParam("author"), "")

	if err != nil {
		fmt.Println("EXEC create...")
		return err
	}
	return c.String(http.StatusCreated, "book successfully added to library")
}

func getBookInner(name, author, year string) (int, string) {

	var Name, Author, empty string
	var Year, Id int
	var Available bool
	row := db.QueryRow("SELECT * FROM books WHERE name=? && author=? && year=? && available=?", name,
		author, year, true)
	err := row.Scan(&Id, &Name, &Year, &Available, &Author, &empty)
	if err != nil {
		return http.StatusBadRequest, "book is not available"
	}
	book := &pb.Book{
		Id:        int32(Id),
		Name:      Name,
		Year:      int32(Year),
		Available: Available,
		Author:    Author,
	}
	_, err = db.Exec("update books set available=? WHERE name=? && author=? && year=? && available=?",
		false, name, author, year, true)
	if err != nil {
		return http.StatusExpectationFailed, "Inner database error"
	}
	return http.StatusOK, book.String()
}
func giveBookInner(name, author, year string) (int, string) {
	res, err := db.Exec("update libtest.books set available=? WHERE name=? && author=? && year=? && available=?",
		true, name, author, year, false)
	if err != nil {
		return http.StatusExpectationFailed, "Inner database error"
	}
	rowsAffected, err := res.RowsAffected()
	if err != nil {
		return http.StatusExpectationFailed, "Inner database error"
	}
	if rowsAffected == 0 {
		return http.StatusBadRequest, "your book doesn't belong to this library"
	}
	return http.StatusAccepted, "book is returned"
}
func createBookInner(name, author, year string) (int, string) {
	_, err := db.Exec("insert into libtest.books (name, year, available, author,section) values (?,?,?,?,?)",
		name, year, true, author, "")

	if err != nil {
		return http.StatusExpectationFailed, "Inner database error"
	}
	return http.StatusCreated, "book successfully added to library"
}
func deleteBookInner(id string) (int, string) {
	res, err := db.Exec("delete from libtest.books where id=?", id)
	if err != nil {
		return http.StatusBadRequest, "Inner database error"
	}
	rows, err := res.RowsAffected()
	if err != nil {
		return http.StatusBadRequest, "Inner database error"
	}
	if rows == 0 {
		return http.StatusNoContent, "library doesn't belong this book"
	} else {
		return http.StatusOK, "book was successfully deleted"
	}
}
