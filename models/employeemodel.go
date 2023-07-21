package models

import "fmt"

type Employee struct {
	IdEmployee   string `json:"id"`
	NameEmployee string `json:"name" validate:"required"`
	Gender       string `json:"gender" validate:"required"`
	Grade        int    `json:"grade" validate:"required"`
	IsMarried    bool   `json:"is-married" validate:"boolean"`
}

func (e Employee) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Gender: %s, Grade: %d, isMarried: %t", e.IdEmployee, e.NameEmployee, e.Gender, e.Grade, e.IsMarried)
}
