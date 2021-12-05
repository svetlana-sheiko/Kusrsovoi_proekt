package main

import (
	"context"
	"fmt"
	"github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"time"
)

const (
	takenTTL   = 12 * time.Hour
	signingKey = "web324o8wr981r32nw3efweifjowief32423r"
)

type tokenClaims struct {
	jwt.StandardClaims
	UserId string `json:"user_id"`
}
type Authorization interface {
	CreateUser(user User) (int, error)
	Encode(username, password string) (string, error)
	Decode(token string) (*tokenClaims, error)
	SignIn(username, password string, token *string) (int, error)
	Validate(user User) error
}
type authService struct {
	repository Repository
}

func (a *authService) CreateUser(user *User) (int, error) {
	if err := a.Validate(user); err == nil {
		hashedPass, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
		if err != nil {
			return http.StatusInternalServerError, err
		}

		user.Password = string(hashedPass)

		err = a.repository.Create(context.Background(), user)
		if err != nil {
			return http.StatusInternalServerError, err
		}
		return http.StatusCreated, nil
	} else {
		return http.StatusBadRequest, err
	}
}
func (a *authService) Validate(u *User) error {
	if u.Name == "" || u.Password == "" || u.Email == "" {
		return &ValidationError{Message: "Missing fields\n"}
	}
	if _, err := a.repository.GetByName(context.Background(), u.Name); err == nil {
		fmt.Println("User with this name already exists.Please change your name")
		return &ValidationError{Message: "User with this name already exists.Please change your name"}
	}

	return nil
}
func (a *authService) Encode(username, password string) (string, error) {
	user, err := a.repository.GetByName(context.Background(), username)
	if err != nil {
		return "", err
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		return "", err
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, &tokenClaims{
		jwt.StandardClaims{
			ExpiresAt: time.Now().Add(takenTTL).Unix(),
			IssuedAt:  time.Now().Unix(),
		}, user.Id,
	})

	return token.SignedString([]byte(signingKey))
}
func (a *authService) SignIn(username, password string, token *string) (int, error) {

	var err error
	*token, err = a.Encode(username, password)
	if err != nil {
		return http.StatusInternalServerError, err
	}

	return http.StatusAccepted, nil
}
