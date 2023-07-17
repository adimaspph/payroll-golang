package handler

import (
	"encoding/json"
	"github.com/adimaspph/payroll-golang/DTO"
	"github.com/adimaspph/payroll-golang/models"
	"github.com/adimaspph/payroll-golang/usecase"
	"github.com/go-playground/validator/v10"
	"io"
	"net/http"
	"os"
	"strconv"
)

func Employee(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method == http.MethodGet {
		employeeList, err := usecase.GetAllEmployee()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(employeeList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(result)

	} else if r.Method == http.MethodPost {
		var employee models.Employee

		err := json.NewDecoder(r.Body).Decode(&employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		validate := validator.New()
		err = validate.Struct(employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		empl, err := usecase.AddEmployee(employee)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(empl)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func SalaryMatrix(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method == http.MethodGet {
		salaryList, err := usecase.GetAllSalaryMatrix()

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(salaryList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(result)

	} else if r.Method == http.MethodPost {
		var salary models.SalaryMatrix

		err := json.NewDecoder(r.Body).Decode(&salary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validation
		validate := validator.New()
		err = validate.Struct(salary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		sal, err := usecase.AddSalaryMatrix(&salary)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(sal)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func Payroll(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json; charset=utf-8")

	if r.Method == http.MethodGet {
		query := r.URL.Query()
		idEmployee := query.Get("id")

		if idEmployee == "" {
			http.Error(w, "employee not found", http.StatusBadRequest)
			return
		}

		payrollList, err := usecase.GetPayrollByIdEmployee(idEmployee)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(payrollList)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		w.Write(result)

	} else if r.Method == http.MethodPost {
		var payrollRequest DTO.PayrollRequest

		err := json.NewDecoder(r.Body).Decode(&payrollRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		// Validation
		validate := validator.New()
		err = validate.Struct(payrollRequest)
		if err != nil {
			http.Error(w, err.Error(), http.StatusBadRequest)
			return
		}

		pyr, err := usecase.AddPayroll(payrollRequest.Id, payrollRequest.HariMasuk, payrollRequest.HariTidakMasuk, *payrollRequest.Date)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		result, err := json.Marshal(pyr)
		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Write(result)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}

func Download(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {

		filePath := "database/database.json"
		file, err := os.Open(filePath)
		if err != nil {
			http.Error(w, "File not found.", http.StatusNotFound)
			return
		}
		defer file.Close()

		// Mendapatkan informasi file
		fileInfo, _ := file.Stat()

		// Menetapkan header HTTP
		w.Header().Set("Content-Disposition", "attachment; filename="+fileInfo.Name())
		w.Header().Set("Content-Type", "application/octet-stream")
		w.Header().Set("Content-Length", strconv.FormatInt(fileInfo.Size(), 10))

		// Mengirim isi file kepada client
		io.Copy(w, file)

	} else {
		w.WriteHeader(http.StatusMethodNotAllowed)
		return
	}
}
