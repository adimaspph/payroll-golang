package usecase

import (
	"errors"
	"github.com/adimaspph/payroll-golang/database"
	"github.com/adimaspph/payroll-golang/models"
	"strconv"
)

func AddSalaryMatrix(matrix *models.SalaryMatrix) (*models.SalaryMatrix, error) {

	c, err := database.ReadCompany()

	if err != nil {
		return nil, err
	}

	matrix.IdSalary = "SALARY-" + strconv.Itoa(len(c.SalaryMatrixList))

	// Check if grade exist
	for _, salaryMatrix := range c.SalaryMatrixList {
		if salaryMatrix.Grade == matrix.Grade {
			return matrix, errors.New("grade already exist")
		}
	}

	// Check if minus
	if matrix.BasicSalary < 0 || matrix.PayCut < 0 || matrix.Allowance < 0 || matrix.HeadOfFamily < 0 {
		return matrix, errors.New("cant assign negative value")
	}

	c.SalaryMatrixList = append(c.SalaryMatrixList, *matrix)

	err = database.WriteCompany(c)
	if err != nil {
		return nil, err
	}

	return matrix, nil
}

func GetAllSalaryMatrix() (*[]models.SalaryMatrix, error) {
	c, err := database.ReadCompany()

	if err != nil {
		return nil, err
	}

	return &c.SalaryMatrixList, nil
}

func GetSalaryByGrade(grade int) (*models.SalaryMatrix, error) {
	c, err := database.ReadCompany()

	if err != nil {
		return nil, err
	}

	for _, matrix := range c.SalaryMatrixList {
		if matrix.Grade == grade {
			return &matrix, nil
		}
	}

	return nil, errors.New("salary grade not found")
}
