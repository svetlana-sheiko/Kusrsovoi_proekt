package main

type ValidationError struct {
	Message string
	Status int
}
func (e ValidationError) Error() string {
	return e.Message
}
type User struct {
	Id string
	Name string
	Email string
	Password string
}
func (u *User) String() string {
	return "Name: "+u.Name+" Email: "+u.Email+" Id: "+u.Id+"\n"
}