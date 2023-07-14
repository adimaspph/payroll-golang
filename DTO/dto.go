package DTO

import "github.com/adimaspph/payroll-golang/models"

type PayrollResponse struct {
	DetailEmployee *models.Employee `json:"detail-employee"`
	TotalPaycut    float64          `json:"total-paycut"`
	TotalAllowance float64          `json:"total-allowance"`
	HeadOfFamily   float64          `json:"head-of-family"`
	BasicSalary    float64          `json:"basic-salary"`
	TotalSalary    float64          `json:"total-salary"`
	Date           *models.Date     `json:"date"`
}

type PayrollRequest struct {
	Id             string       `json:"id"`
	HariMasuk      int          `json:"hari-masuk"`
	HariTidakMasuk int          `json:"hari-tidak-masuk"`
	Date           *models.Date `json:"date"`
}
