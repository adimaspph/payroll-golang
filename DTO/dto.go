package DTO

import "github.com/adimaspph/payroll-golang/models"

type PayrollResponse struct {
	DetailEmployee *models.Employee `json:"detail-employee" validate:"required"`
	TotalPaycut    float64          `json:"total-paycut" validate:"required"`
	TotalAllowance float64          `json:"total-allowance" validate:"required"`
	HeadOfFamily   float64          `json:"head-of-family" validate:"required"`
	BasicSalary    float64          `json:"basic-salary" validate:"required"`
	TotalSalary    float64          `json:"total-salary" validate:"required"`
	Date           *models.Date     `json:"date" validate:"required"`
}

type PayrollRequest struct {
	Id             string       `json:"id" validate:"required"`
	HariMasuk      int          `json:"hari-masuk" validate:"required"`
	HariTidakMasuk int          `json:"hari-tidak-masuk" validate:"required"`
	Date           *models.Date `json:"date" validate:"required"`
}
