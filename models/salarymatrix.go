package models

type SalaryMatrix struct {
	IdSalary     string  `json:"id" validate:"required"`
	Grade        int     `json:"grade" validate:"required"`
	BasicSalary  float64 `json:"basic-salary" validate:"required"`
	PayCut       float64 `json:"pay-cut" validate:"required"`
	Allowance    float64 `json:"allowance" validate:"required"`
	HeadOfFamily float64 `json:"head-of-family" validate:"required"`
}
