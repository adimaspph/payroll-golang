package models

import "fmt"

type Employee struct {
	IdEmployee   string `json:"id"`
	NameEmployee string `json:"name"`
	Gender       string `json:"gender"`
	Grade        int    `json:"grade"`
	IsMarried    bool   `json:"is-married"`
}

func (e Employee) String() string {
	return fmt.Sprintf("ID: %s, Name: %s, Gender: %s, Grade: %d, isMarried: %t", e.IdEmployee, e.NameEmployee, e.Gender, e.Grade, e.IsMarried)
}
