package users

import (
	"fmt"
	"time"
)

type User struct {
	Id        int64     `json:"id"`
	Email     string    `json:"email"`
	Password  string    `json:"password"`
	FirstName string    `json:"firstName"`
	LastName  string    `json:"lastName"`
	Phone     string    `json:"phone"`
	CreateAt  time.Time `json:"createAt"`
}

func (u *User) GetId() int64 {
	return u.Id
}

func NewUser(id int64, email string, password string, firstName string, lastName string, phone string) *User {
	return &User{
		Id:        id,
		Email:     email,
		Password:  password,
		FirstName: firstName,
		LastName:  lastName,
		Phone:     phone,
		CreateAt:  time.Now(),
	}
}

func (u *User) ToString() string {
	return fmt.Sprintf("User: %d, %s, %s, %s, %s, %s, %s", u.Id, u.Email, u.Password, u.FirstName, u.LastName, u.Phone, u.CreateAt.Format("2006-01-02 15:04:05"))
}
