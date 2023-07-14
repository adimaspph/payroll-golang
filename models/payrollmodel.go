package models

type Payroll struct {
	IdPayroll        string    `json:"id-payroll"`
	BasicSalary      float64   `json:"basic-salary"`
	PayCut           float64   `json:"pay-cut"`
	AdditionalSalary float64   `json:"additional-salary"`
	Employee         *Employee `json:"employee"`
	Period           *Date     `json:"period"`
}
