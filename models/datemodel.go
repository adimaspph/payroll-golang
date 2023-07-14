package models

import "fmt"

type Date struct {
	Month int `json:"month"`
	Year  int `json:"year"`
}

func (d Date) String() string {
	return fmt.Sprint(d.Month, "-", d.Year)
}
