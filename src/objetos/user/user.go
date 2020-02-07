package user

import (
	"time"
)

type Usuario struct {
	Id     int
	Name   string
	UpDate time.Time
	Status bool
}

func (user *Usuario) CreateUser(id int, name string, upDate time.Time, status bool) {
	user.Id = id
	user.Name = name
	user.UpDate = upDate
	user.Status = status
}
