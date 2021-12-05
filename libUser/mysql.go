package main

import (
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"net/http"
)

type Repository interface {
	Create(context.Context,*User) error
	GetAll(ctx context.Context) ([]*User,error)
	GetById(ctx context.Context,id string) (*User,error)
	GetByName(ctx context.Context,name string) (*User,error)
	Delete(ctx context.Context,id string) error
}
type db struct {
	storage *sql.DB
}
func NewRepository(storage *sql.DB) Repository {
	return &db{storage: storage}
}
func (s *db) Create(ctx context.Context,u *User) error {

	_, err := s.storage.Exec("insert into libtest.users (name,email,password) values (?,?,?)",u.Name,u.Email,u.Password)
	if err != nil {

		return err
	}
	return nil
}
func (s *db) GetById(ctx context.Context,id string) (*User,error) {
	rows:=s.storage.QueryRow("SELECT * from libtest.users where id=?",id)
	u:=User{}
	err:=rows.Scan(&u.Id,&u.Name,&u.Email,&u.Password)
	if err!=nil {
		return nil,err
	}
	return &u,nil
}
func (s *db) GetAll(ctx context.Context) ([]*User,error) {
	rows,err:=s.storage.Query("select * from libtest.users")
	if err!=nil {
		return nil,err
	}
	defer rows.Close()
	users:=make([]*User,0)
	for rows.Next() {
		u:=User{}
		err:=rows.Scan(&u.Id,&u.Name,&u.Email,&u.Password)
		if err!=nil {
			return users,err
		}
		users= append(users, &u)
	}
	return users,nil
}
func (s *db) GetByName(ctx context.Context,name string) (*User,error) {
	rows:=s.storage.QueryRow("SELECT * from libtest.users where name=?",name)
	u:=User{}
	err:=rows.Scan(&u.Id,&u.Name,&u.Email,&u.Password)
	if err!=nil {
		return nil,err
	}
	return &u,nil
}
func (s *db) Delete(ctx context.Context,id string) error {
	res,err:=s.storage.Exec("delete from libtest.users where id=?",id)
	if err!=nil {
		return err
	}
	rows,err:= res.RowsAffected()
	if err!=nil {
		return err
	}
	if rows==0 {

		return ValidationError{Message: "library doesn't have this user",Status: http.StatusBadRequest}
	}else {
		return ValidationError{Message: "User was deleted!",Status: http.StatusOK}
	}

}