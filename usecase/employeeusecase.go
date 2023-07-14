package usecase

import (
	"errors"
	"github.com/adimaspph/payroll-golang/database"
	"github.com/adimaspph/payroll-golang/models"
	"strconv"
)

func AddEmployee(employee models.Employee) (*models.Employee, error) {
	company, err := database.ReadCompany()

	if err != nil {
		return nil, err
	}

	employee.IdEmployee = "ID-" + strconv.Itoa(len(company.EmployeeList))

	company.EmployeeList = append(company.EmployeeList, employee)

	err = database.WriteCompany(company)
	if err != nil {
		return nil, err
	}

	return &employee, nil
}

func GetEmployeeByID(id string) (*models.Employee, error) {
	company, err := database.ReadCompany()

	if err != nil {
		return nil, err
	}

	for _, e := range company.EmployeeList {
		if e.IdEmployee == id {
			return &e, nil
		}
	}
	return nil, errors.New("employee is not found")
}

func GetAllEmployee() ([]models.Employee, error) {
	company, err := database.ReadCompany()

	if err != nil {
		return nil, err
	}

	return company.EmployeeList, nil
}
