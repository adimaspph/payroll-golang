package main

import (
	"fmt"
	"github.com/adimaspph/payroll-golang/handler"
	"net/http"
)

func main() {
	var employee http.HandlerFunc = handler.Employee
	var salary http.HandlerFunc = handler.SalaryMatrix
	var payroll http.HandlerFunc = handler.Payroll
	var download http.HandlerFunc = handler.Download

	mux := http.NewServeMux()
	mux.HandleFunc("/employee", employee)
	mux.HandleFunc("/salary", salary)
	mux.HandleFunc("/payroll", payroll)
	mux.HandleFunc("/download", download)

	server := http.Server{
		Addr:    "localhost:8080",
		Handler: mux,
	}
	fmt.Printf("Server running on %s", server.Addr)
	err := server.ListenAndServe()
	if err != nil {
		panic(err.Error())
	}
}
