package models

import (
	"employee-management-system/database"
	"log"
)

type Request struct {
	ID            int    `json:"id"`
	EmployeeID    int    `json:"employee_id"`
	TypeRequest   string `json:"type_request"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Reason        string `json:"reason"`
	StatusRequest string `json:"status_request"`
}

type RequestResponse struct {
	ID            int    `json:"id"`
	EmployeeName  string `json:"employee_name"`
	TypeRequest   string `json:"type_request"`
	StartDate     string `json:"start_date"`
	EndDate       string `json:"end_date"`
	Reason        string `json:"reason"`
	StatusRequest string `json:"status_request"`
}

// GetRequestByEmployeeId function get request of employee by id
func GetRequestByEmployeeId(id int) ([]RequestResponse, error) {
	db = database.DB
	var requests []Request
	rows, err := db.Query("SELECT * FROM requests WHERE employeeId = ?", id)
	if err != nil {
		log.Fatal(err)
		return nil, err
	}
	for rows.Next() {
		var request Request
		err = rows.Scan(&request.ID, &request.TypeRequest, &request.EmployeeID, &request.StartDate, &request.EndDate, &request.Reason, &request.StatusRequest)
		if err != nil {
			log.Fatal(err)
			return nil, err
		}
		requests = append(requests, request)
	}
	var requestResponses = make([]RequestResponse, len(requests))
	for i, request := range requests {
		requestResponses[i], err = GetRequestResponse(request)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}

	return requestResponses, nil
}

// CreateRequest create request of employee
func CreateRequest(request Request) (RequestResponse, error) {
	db = database.DB
	var requestResponse RequestResponse
	stmt, err := db.Prepare("INSERT INTO requests (employeeId, typeRequest, startDate, endDate, reason, statusRequest) VALUES (?, ?, ?, ?, ?, ?)")
	if err != nil {
		log.Fatal(err)
		return requestResponse, err
	}
	res, err := stmt.Exec(request.EmployeeID, request.TypeRequest, request.StartDate, request.EndDate, request.Reason, request.StatusRequest)
	if err != nil {
		log.Fatal(err)
		return requestResponse, err
	}
	id, err := res.LastInsertId()
	if err != nil {
		log.Fatal(err)
		return requestResponse, err
	}
	request.ID = int(id)
	requestResponse, err = GetRequestResponse(request)
	if err != nil {
		log.Println(err)
		return requestResponse, err
	}
	return requestResponse, nil
}

// GetAllRequests func get all request
func GetAllRequests() ([]RequestResponse, error) {
	db = database.DB
	var requests []Request
	rows, err := db.Query("SELECT * FROM requests")
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var request Request
		err = rows.Scan(&request.ID, &request.TypeRequest, &request.EmployeeID, &request.StartDate, &request.EndDate, &request.Reason, &request.StatusRequest)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		requests = append(requests, request)
	}
	var requestResponses = make([]RequestResponse, len(requests))
	for i, request := range requests {
		requestResponses[i], err = GetRequestResponse(request)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return requestResponses, nil
}

// GetRequestResponse func get request response
func GetRequestResponse(request Request) (RequestResponse, error) {
	db = database.DB
	var requestResponse RequestResponse
	var employeeName string
	err := db.QueryRow("SELECT name FROM employees WHERE id = ?", request.EmployeeID).Scan(&employeeName)
	if err != nil {
		log.Println(err)
		return requestResponse, err
	}

	requestResponse.ID = request.ID
	requestResponse.EmployeeName = employeeName
	requestResponse.TypeRequest = request.TypeRequest
	requestResponse.StartDate = request.StartDate
	requestResponse.EndDate = request.EndDate
	requestResponse.Reason = request.Reason
	requestResponse.StatusRequest = request.StatusRequest

	return requestResponse, nil
}

func UpdateRequest(id int, status string) error {
	db = database.DB
	stmt, err := db.Prepare("UPDATE requests SET statusRequest = ? WHERE id = ?")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(status, id)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}

func GetDivisionRequests(id int) ([]RequestResponse, error) {
	db = database.DB
	lead, err := GetEmployeeById(id)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	rows, err := db.Query("SELECT * FROM employees where departmentId = ?", lead.DepartmentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var employees []EmployeeRequest
	for rows.Next() {
		var employee EmployeeRequest
		err = rows.Scan(&employee.Id, &employee.Name, &employee.Email, &employee.Password, &employee.Birthday, &employee.RoleId, &employee.DepartmentId)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		employees = append(employees, employee)
	}
	var requestResponses []RequestResponse
	for _, employee := range employees {
		requests, err := GetRequestByEmployeeId(employee.Id)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		for _, request := range requests {
			requestResponses = append(requestResponses, request)
		}
	}
	return requestResponses, nil
}

func UpdateRequestByLeader(request Request) error {
	db = database.DB
	stmt, err := db.Prepare("UPDATE requests SET statusRequest = ? WHERE id = ?")
	if err != nil {
		log.Println(err)
		return err
	}
	_, err = stmt.Exec(request.StatusRequest, request.ID)
	if err != nil {
		log.Println(err)
		return err
	}
	return nil
}
