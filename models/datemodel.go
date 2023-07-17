package models

import "fmt"

type Date struct {
	Month int `json:"month" validate:"required"`
	Year  int `json:"year" validate:"required"`
}

func (d Date) String() string {
	return fmt.Sprint(d.Month, "-", d.Year)
}
