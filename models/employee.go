package models

import (
	"database/sql"
	"employee-management-system/database"
	"encoding/csv"
	"fmt"
	"log"
	"os"
)

type EmployeeRequest struct {
	Id           int    `json:"id"`
	Name         string `json:"name"`
	Email        string `json:"email"`
	Password     string `json:"password"`
	Birthday     string `json:"birthday"`
	RoleId       int    `json:"role_id"`
	DepartmentId int    `json:"department_id"`
}

type EmployeeResponse struct {
	Id         int    `json:"id"`
	Name       string `json:"name"`
	Email      string `json:"email"`
	Birthday   string `json:"birthday"`
	Role       string `json:"role"`
	Department string `json:"department"`
}

var db *sql.DB

// GetAllEmployees GetEmployeeList get all employee
func GetAllEmployees() ([]EmployeeResponse, error) {
	db = database.DB
	var employees []EmployeeRequest
	rows, err := db.Query("SELECT * FROM employees")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var employee EmployeeRequest
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.Birthday, &employee.RoleId, &employee.DepartmentId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}

	var employeesResponse = make([]EmployeeResponse, len(employees))
	for i, employee := range employees {
		employeesResponse[i], err = GetEmployeeResponse(employee)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return employeesResponse, nil
}

// GetEmployeeById function
func GetEmployeeById(id int) (EmployeeRequest, error) {
	db = database.DB
	var employee EmployeeRequest
	err := db.QueryRow("SELECT * FROM employees WHERE id=?", id).Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.Birthday, &employee.RoleId, &employee.DepartmentId)
	if err != nil {
		log.Println(err)
		return employee, err
	}

	return employee, nil
}

// CreateEmployee GetEmployeeByEmail function
func CreateEmployee(employee EmployeeRequest) (EmployeeRequest, error) {
	db = database.DB
	stmt, err := db.Prepare("INSERT INTO employees(name, email, password, birthday, roleId, departmentId) VALUES(?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Println(err)
		return employee, err
	}
	res, err := stmt.Exec(employee.Name, employee.Email, employee.Password, employee.Birthday, employee.RoleId, employee.DepartmentId)
	if err != nil {
		log.Println(err)
		return employee, err
	}
	fmt.Println(res)
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return employee, err
	}
	employee.Id = int(id)
	return employee, nil
}

// UpdateEmployee function
func UpdateEmployee(employee EmployeeRequest) error {
	db = database.DB
	stmt, err := db.Prepare("UPDATE employees SET name=?, email=?, password=?, birthday=?, roleId=?, departmentId=? WHERE id=?")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(employee.Name, employee.Email, employee.Password, employee.Birthday, employee.RoleId, employee.DepartmentId, employee.Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// DeleteEmployee function
func DeleteEmployee(id int) error {
	db = database.DB
	stmt, err := db.Prepare("DELETE FROM employees WHERE id=?")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

// GetEmployeeResponse function, transform EmployeeRequest to EmployeeResponse
func GetEmployeeResponse(employee EmployeeRequest) (EmployeeResponse, error) {
	db = database.DB
	var employeeResponse EmployeeResponse
	role, err := GetRoleById(employee.RoleId)
	if err != nil {
		log.Println(err)
		return employeeResponse, err
	}
	department, err := GetDepartmentById(employee.DepartmentId)
	if err != nil {
		log.Println(err)
		return employeeResponse, err
	}

	employeeResponse = EmployeeResponse{
		employee.Id,
		employee.Name,
		employee.Email,
		employee.Birthday,
		role.Name,
		department.Name,
	}
	return employeeResponse, nil
}

//Login function
func Login(email string, password string) (EmployeeResponse, error) {
	db = database.DB
	var employee EmployeeRequest
	var employeeResponse EmployeeResponse

	err := db.QueryRow("SELECT * FROM employees WHERE email=? AND password=?", email, password).Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.Birthday, &employee.RoleId, &employee.DepartmentId)
	if err != nil {
		log.Println(err)
		return employeeResponse, err
	}
	employeeResponse, err = GetEmployeeResponse(employee)
	if err != nil {
		log.Println(err)
		return employeeResponse, err
	}
	return employeeResponse, nil
}

// GetEmployeeByEmail function to get employee by email
func GetEmployeeByEmail(email string) (EmployeeRequest, error) {
	db = database.DB
	var employee EmployeeRequest
	err := db.QueryRow("SELECT * FROM employees WHERE email=?", email).Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.Birthday, &employee.RoleId, &employee.DepartmentId)
	if err != nil {
		log.Println(err)
		return employee, err
	}
	return employee, nil
}

//UpdatePassword function
func UpdatePassword(password string, id int) error {
	db = database.DB
	stmt, err := db.Prepare("UPDATE employees SET password=? WHERE id=?")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(password, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//GetDivisionEmployees function get employees by division
func GetDivisionEmployees(id int) ([]EmployeeResponse, error) {
	db = database.DB
	var employees []EmployeeResponse
	var departmentId int
	err := db.QueryRow("SELECT departmentId FROM employees WHERE id=?", id).Scan(&departmentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM employees WHERE departmentId=?", departmentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	var employeeRequests []EmployeeRequest
	for rows.Next() {
		var employee EmployeeRequest
		err := rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.Birthday, &employee.RoleId, &employee.DepartmentId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		fmt.Println(employee)
		employeeRequests = append(employeeRequests, employee)
	}
	employees = make([]EmployeeResponse, len(employeeRequests))
	for i, employeeRequest := range employeeRequests {
		employees[i], err = GetEmployeeResponse(employeeRequest)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return employees, nil
}

func UpdateDivisionLead(response EmployeeResponse) error {
	db = database.DB
	employee, err := GetEmployeeRequest(response)
	fmt.Printf("employee: %v\n", employee)
	if err != nil {
		log.Println(err)
		return err
	}
	stmt, err := db.Prepare("UPDATE employees SET name=?, birthday=?, roleId=?, departmentId=? WHERE id=?")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(employee.Name, employee.Birthday, employee.RoleId, employee.DepartmentId, employee.Id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

//GetEmployeeRequest function transform EmployeeResponse to EmployeeRequest
func GetEmployeeRequest(response EmployeeResponse) (EmployeeRequest, error) {
	db = database.DB
	var employee EmployeeRequest
	err := db.QueryRow("SELECT * FROM employees WHERE id=?", response.Id).Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.Birthday, &employee.RoleId, &employee.DepartmentId)
	if err != nil {
		log.Println(err)
		return employee, err
	}
	if response.Name != "" {
		employee.Name = response.Name
	}
	if response.Birthday != "" {
		employee.Birthday = response.Birthday
	}
	if response.Role != "" {
		role, err := GetRoleByName(response.Role)
		if err != nil {
			log.Println(err)
			return employee, err
		}
		employee.RoleId = role.Id
	}
	return employee, nil
}

func ExportEmployee(leadId int) ([]EmployeeResponse, error) {
	employee, err := GetDivisionEmployees(leadId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	file, err := os.Create("/data/employees.csv")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	defer file.Close()
	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := []string{"Name", "Email", "Birthday", "Role", "Department"}
	err = writer.Write(header)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, employee := range employee {
		row := []string{employee.Name, employee.Email, employee.Birthday, employee.Role, employee.Department}
		err = writer.Write(row)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return employee, nil
}
