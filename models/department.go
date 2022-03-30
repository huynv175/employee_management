package models

import (
	"employee-management-system/database"
	"log"
)

type DepartmentRequest struct {
	Id        int    `json:"id"`
	Name      string `json:"name"`
	ManagerId int    `json:"manager_id"`
}

type DepartmentResponse struct {
	Id      int    `json:"id"`
	Name    string `json:"name"`
	Manager string `json:"manager"`
}

// GetAllDepartment Get all departments
func GetAllDepartment() ([]DepartmentResponse, error) {
	var departments []DepartmentRequest
	db = database.DB
	rows, err := db.Query("SELECT * FROM departments")
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var department DepartmentRequest
		err = rows.Scan(&department.Id, &department.Name, &department.ManagerId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		departments = append(departments, department)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	departmentsResponse := make([]DepartmentResponse, len(departments))
	for i, department := range departments {
		departmentsResponse[i], err = GetDepartmentResponse(department)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return departmentsResponse, nil
}

// GetEmployeesOfDepartment func Get employee of department
func GetEmployeesOfDepartment(departmentId int) ([]EmployeeResponse, error) {
	var employees []EmployeeRequest
	var employeesResponse []EmployeeResponse
	db = database.DB
	rows, err := db.Query("SELECT * FROM employees WHERE departmentId = ?", departmentId)
	if err != nil {
		log.Println(err)
		return employeesResponse, err
	}

	for rows.Next() {
		var employee EmployeeRequest
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.Birthday, &employee.RoleId, &employee.DepartmentId)
		if err != nil {
			log.Println(err)
			return employeesResponse, err
		}
		employees = append(employees, employee)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return employeesResponse, err
	}
	employeesResponse = make([]EmployeeResponse, len(employees))
	for i, employee := range employees {
		employeesResponse[i], err = GetEmployeeResponse(employee)
		if err != nil {
			log.Println(err)
			return employeesResponse, err
		}
	}
	return employeesResponse, nil
}

func GetDepartmentById(id int) (DepartmentResponse, error) {
	db = database.DB
	var departmentResponse DepartmentResponse
	var department DepartmentRequest
	err := db.QueryRow("SELECT * FROM departments WHERE id = ?", id).Scan(&department.Id, &department.Name, &department.ManagerId)
	if err != nil {
		log.Println(err)
		return departmentResponse, err
	}
	departmentResponse, err = GetDepartmentResponse(department)
	if err != nil {
		log.Println(err)
		return departmentResponse, err
	}
	return departmentResponse, nil
}

// CreateDepartment Create new department
func CreateDepartment(departmentResponse DepartmentResponse) (DepartmentResponse, error) {
	db = database.DB
	var department DepartmentRequest
	stmt, err := db.Prepare("INSERT INTO departments (name, managerId) VALUES (?, ?)")
	if err != nil {
		log.Println(err)
		return departmentResponse, err
	}

	department, err = GetDepartmentRequest(departmentResponse)
	if err != nil {
		log.Println(err)
		return departmentResponse, err
	}

	res, err := stmt.Exec(departmentResponse.Name, department.ManagerId)
	if err != nil {
		log.Println(err)
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Println(err)
		return departmentResponse, err
	}
	departmentResponse.Id = int(id)
	return departmentResponse, nil
}

// GetDepartmentResponse func  transform DepartmentRequest to DepartmentResponse
func GetDepartmentResponse(request DepartmentRequest) (DepartmentResponse, error) {
	var response DepartmentResponse
	response.Id = request.Id
	response.Name = request.Name
	department, err := GetEmployeeById(request.ManagerId)
	if err != nil {
		log.Println(err)
		return response, err
	}
	response.Manager = department.Name
	return response, nil
}

// GetDepartmentRequest  func transform DepartmentResponse to DepartmentRequest
func GetDepartmentRequest(response DepartmentResponse) (DepartmentRequest, error) {
	var request DepartmentRequest
	request.Id = response.Id
	request.Name = response.Name
	employee, err := GetEmployeeByEmail(response.Manager)
	if err != nil {
		log.Println(err)
		return request, err
	}
	request.ManagerId = employee.Id
	return request, nil
}

func UpdateDepartment(id int, response DepartmentResponse) (DepartmentResponse, error) {
	db = database.DB
	var department DepartmentRequest
	stmt, err := db.Prepare("UPDATE departments SET name = ?, managerId = ? WHERE id = ?")
	if err != nil {
		log.Println(err)
		return response, err
	}
	department, err = GetDepartmentRequest(response)
	if err != nil {
		log.Println(err)
		return response, err
	}
	_, err = stmt.Exec(response.Name, department.ManagerId, id)
	if err != nil {
		log.Println(err)
		return response, err
	}
	return response, nil
}

func DeleteDepartment(id int) error {
	db = database.DB
	stmt, err := db.Prepare("DELETE FROM departments WHERE id = ?")
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
