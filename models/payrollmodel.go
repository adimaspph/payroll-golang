package models

type Payroll struct {
	IdPayroll        string    `json:"id-payroll" `
	BasicSalary      float64   `json:"basic-salary" validate:"required"`
	PayCut           float64   `json:"pay-cut" validate:"required"`
	AdditionalSalary float64   `json:"additional-salary" validate:"required"`
	Employee         *Employee `json:"employee" validate:"required"`
	Period           *Date     `json:"period" validate:"required"`
}
