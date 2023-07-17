package models

type Company struct {
	EmployeeList     []Employee     `json:"employee-list,omitempty" validate:"required"`
	PayrollList      []Payroll      `json:"payroll-list,omitempty" validate:"required"`
	SalaryMatrixList []SalaryMatrix `json:"salary-matrix-list,omitempty" validate:"required"`
}

// Soal 1
//func (c *Company) AddEmployee(employee *Employee) {
//	//newEmployee := Employee{
//	//	IdEmployee:   "ID" + strconv.Itoa(len(c.EmployeeList)),
//	//	NameEmployee: name,
//	//	Gender:       gender,
//	//	Grade:        grade,
//	//	IsMarried:    isMarried,
//	//}
//
//	c.EmployeeList = append(c.EmployeeList, newEmployee)
//}

// Soal 3
//func (c *Company) AddPayroll(id string, hariMasuk int, hariTidakMasuk int, period Date) (*Payroll, error) {
//	employee := c.getEmployeeByID(id)
//
//	if employee == nil {
//		return nil, errors.New("Employee not exist")
//	}
//
//	if c.checkDuplicatePayroll(employee, period) {
//		return nil, errors.New("Duplicate payroll")
//	}
//
//	salary := c.getSalaryByGrade(employee.Grade)
//
//	if salary == nil {
//		return nil, errors.New("Salary doesnt exist")
//	}
//
//	var headOfFamilyAllowance float64 = 0
//
//	if employee.Gender == "Male" && employee.IsMarried {
//		headOfFamilyAllowance += salary.HeadOfFamily
//	}
//
//	newPayroll := Payroll{
//		IdPayroll:        "ID" + strconv.Itoa(len(c.PayrollList)),
//		BasicSalary:      salary.BasicSalary,
//		PayCut:           salary.PayCut * float64(hariTidakMasuk),
//		AdditionalSalary: (salary.Allowance * float64(hariMasuk)) + headOfFamilyAllowance,
//		Employee:         employee,
//		Period:           &period,
//	}
//
//	c.PayrollList = append(c.PayrollList, newPayroll)
//
//	return &newPayroll, nil
//}

//func (c *Company) getEmployeeByID(id string) *Employee {
//	for _, e := range c.EmployeeList {
//		if e.IdEmployee == id {
//			return &e
//		}
//	}
//	return nil
//}

//func (c *Company) getSalaryByGrade(grade int) *SalaryMatrix {
//	for _, matrix := range c.SalaryMatrixList {
//		if matrix.Grade == grade {
//			return &matrix
//		}
//	}
//	return nil
//}

//func (c *Company) checkDuplicatePayroll(employee *Employee, period Date) bool {
//	for _, payroll := range c.PayrollList {
//		if payroll.Employee.IdEmployee == employee.IdEmployee && payroll.Period.Month == period.Month && payroll.Period.Year == period.Year {
//			return true
//		}
//	}
//	return false
//}

// Soal 2
//func (c *Company) ShowAllEmployee() {
//	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
//
//	fmt.Fprintln(w, "No\tID\tName\tGender\tGrade\tIs Married\t")
//	for i, e := range c.EmployeeList {
//		fmt.Fprintf(w, "%d\t%s\t%s\t%s\t%d\t%t\t\n", i, e.IdEmployee, e.NameEmployee, e.Gender, e.Grade, e.IsMarried)
//	}
//	err := w.Flush()
//	if err != nil {
//		return
//	}
//}

//func (c *Company) ShowPayrollByIdEmployee(idEmployee string) {
//	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
//
//	employee := c.getEmployeeByID(idEmployee)
//
//	if employee == nil {
//		fmt.Println("Employee not exist")
//		return
//	}
//
//	//fmt.Fprintln(w, "Keteranga\tData\t")
//	fmt.Fprintf(w, "Detail Employee\t:(%s)\t\n", employee)
//
//	salary := c.getSalaryByGrade(employee.Grade)
//
//	if salary == nil {
//		fmt.Println("Employee salary not exist")
//		return
//	}
//
//	var headOfFamilyAllowance float64 = 0
//
//	if employee.Gender == "Male" && employee.IsMarried {
//		headOfFamilyAllowance += salary.HeadOfFamily
//	}
//	fmt.Fprintf(w, "Head Of Family\t: %.2f\t\n", headOfFamilyAllowance)
//	fmt.Fprintf(w, "Basic Salary\t: %.2f\t\n", salary.BasicSalary)
//
//	var totalPayCut float64 = 0
//	var totalAllowance float64 = 0
//	var totalSalary float64 = 0
//
//	for _, payroll := range c.PayrollList {
//		if payroll.Employee.IdEmployee == employee.IdEmployee {
//			totalPayCut += payroll.PayCut
//			totalAllowance += payroll.AdditionalSalary
//			totalSalary += payroll.BasicSalary
//		}
//	}
//
//	totalSalary += totalAllowance
//	totalSalary -= totalPayCut
//
//	fmt.Fprintf(w, "Total Pay Cut\t: %.2f\t\n", totalPayCut)
//	fmt.Fprintf(w, "Total Allowance\t: %.2f\t\n", totalAllowance)
//	fmt.Fprintf(w, "Total Salary\t: %.2f\t\n", totalSalary)
//
//	err := w.Flush()
//	if err != nil {
//		return
//	}
//}

//func (c *Company) AddSalaryMatrix(matrix SalaryMatrix) (*SalaryMatrix, error) {
//
//	matrix.IdSalary = strconv.Itoa(len(c.SalaryMatrixList))
//
//	// Check if grade exist
//	for _, salaryMatrix := range c.SalaryMatrixList {
//		if salaryMatrix.Grade == matrix.Grade {
//			return &matrix, errors.New("Grade already exist")
//		}
//	}
//
//	// Check if minus
//	if matrix.BasicSalary < 0 || matrix.PayCut < 0 || matrix.Allowance < 0 || matrix.HeadOfFamily < 0 {
//		return &matrix, errors.New("Cant assign negative value")
//	}
//
//	c.SalaryMatrixList = append(c.SalaryMatrixList, matrix)
//
//	return &matrix, nil
//}

//func (c *Company) ShowSalaryMatrix() {
//	w := tabwriter.NewWriter(os.Stdout, 1, 1, 1, ' ', 0)
//
//	fmt.Fprintln(w, "Grade\tBasic Salary\tPay Cut\tAllowance\tHead of Family\t")
//
//	for _, matrix := range c.SalaryMatrixList {
//		fmt.Fprintf(w, "%d\t%.2f\t%.2f\t%.2f\t%.2f\t\n", matrix.Grade, matrix.BasicSalary, matrix.PayCut, matrix.Allowance, matrix.HeadOfFamily)
//	}
//
//	err := w.Flush()
//	if err != nil {
//		return
//	}
//}

//func (c *Company) GetGrade() []int {
//	result := []int{}
//
//	for _, salary := range c.SalaryMatrixList {
//		result = append(result, salary.Grade)
//	}
//
//	return result
//}

//
