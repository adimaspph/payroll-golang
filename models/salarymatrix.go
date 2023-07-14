package models

type SalaryMatrix struct {
	IdSalary     string  `json:"id"`
	Grade        int     `json:"grade"`
	BasicSalary  float64 `json:"basic-salary"`
	PayCut       float64 `json:"pay-cut"`
	Allowance    float64 `json:"allowance"`
	HeadOfFamily float64 `json:"head-of-family"`
}
