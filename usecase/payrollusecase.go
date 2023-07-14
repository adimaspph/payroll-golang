package usecase

import (
	"errors"
	"github.com/adimaspph/payroll-golang/DTO"
	"github.com/adimaspph/payroll-golang/database"
	"github.com/adimaspph/payroll-golang/models"
	"strconv"
)

func AddPayroll(id string, hariMasuk int, hariTidakMasuk int, period models.Date) (*models.Payroll, error) {
	c, err := database.ReadCompany()
	if err != nil {
		return nil, err
	}

	employee, err := GetEmployeeByID(id)

	if err != nil {
		return nil, err
	}

	if employee == nil {
		return nil, errors.New("employee not exist")
	}
	result, err := CheckDuplicatePayroll(employee, period)

	if err != nil {
		return nil, err
	}

	if result {
		return nil, errors.New("duplicate payroll")
	}

	salary, err := GetSalaryByGrade(employee.Grade)

	if err != nil {

	}

	if salary == nil {
		return nil, errors.New("Salary doesnt exist")
	}

	var headOfFamilyAllowance float64 = 0

	if employee.Gender == "Male" && employee.IsMarried {
		headOfFamilyAllowance += salary.HeadOfFamily
	}

	newPayroll := models.Payroll{
		IdPayroll:        "PAYROLL-" + strconv.Itoa(len(c.PayrollList)),
		BasicSalary:      salary.BasicSalary,
		PayCut:           salary.PayCut * float64(hariTidakMasuk),
		AdditionalSalary: (salary.Allowance * float64(hariMasuk)) + headOfFamilyAllowance,
		Employee:         employee,
		Period:           &period,
	}

	c.PayrollList = append(c.PayrollList, newPayroll)

	err = database.WriteCompany(c)
	if err != nil {
		return nil, err
	}

	return &newPayroll, nil
}

func CheckDuplicatePayroll(employee *models.Employee, period models.Date) (bool, error) {
	c, err := database.ReadCompany()

	if err != nil {
		return false, err
	}

	for _, payroll := range c.PayrollList {
		if payroll.Employee.IdEmployee == employee.IdEmployee && payroll.Period.Month == period.Month && payroll.Period.Year == period.Year {
			return true, nil
		}
	}
	return false, nil
}

func GetPayrollByIdEmployee(idEmployee string) ([]DTO.PayrollResponse, error) {
	c, err := database.ReadCompany()
	if err != nil {
		return nil, err
	}

	employee, err := GetEmployeeByID(idEmployee)

	if err != nil {
		return nil, err
	}

	salary, err := GetSalaryByGrade(employee.Grade)

	if err != nil {
		return nil, err
	}

	var result []DTO.PayrollResponse = make([]DTO.PayrollResponse, 0)

	var headOfFamilyAllowance float64 = 0

	if employee.Gender == "male" && employee.IsMarried {
		headOfFamilyAllowance += salary.HeadOfFamily
	}

	for _, payroll := range c.PayrollList {
		if payroll.Employee.IdEmployee == employee.IdEmployee {

			result = append(result, DTO.PayrollResponse{
				DetailEmployee: employee,
				TotalPaycut:    payroll.PayCut,
				TotalAllowance: payroll.AdditionalSalary + headOfFamilyAllowance,
				HeadOfFamily:   headOfFamilyAllowance,
				BasicSalary:    payroll.BasicSalary,
				TotalSalary:    payroll.BasicSalary + payroll.AdditionalSalary + headOfFamilyAllowance - payroll.PayCut,
				Date:           payroll.Period,
			})
		}
	}

	return result, nil
}
