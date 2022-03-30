package models

import (
	"employee-management-system/constants"
	"employee-management-system/database"
	"encoding/csv"
	"fmt"
	"log"
	"os"
	"strconv"
	"time"
)

type Timesheet struct {
	Id         int    `json:"id"`
	EmployeeId int    `json:"employee_id"`
	Year       string `json:"year"`
	Month      string `json:"month"`
	Day        string `json:"day"`
	CheckIn    string `json:"check_in"`
	CheckOut   string `json:"check_out"`
}

type TimesheetResponse struct {
	Id       int    `json:"id"`
	Employee string `json:"employee"`
	Year     string `json:"year"`
	Month    string `json:"month"`
	Day      string `json:"day"`
	CheckIn  string `json:"check_in"`
	CheckOut string `json:"check_out"`
}

type WorkingHours struct {
	Name string `json:"name"`
	Day  int    `json:"day"`
	Hour int    `json:"hour"`
}

// HandleCheckIn function hande checkin
func HandleCheckIn(id int) string {
	checked := checked(id)
	// checked = 0: checked in today, checked = 1: checked out today, checked = 2: forgot to check out, checked = 3: checked out yesterday, checked = 4: error
	if checked == 0 {
		return "you checked in"
	} else if checked == 1 {
		return "you checked out"
	} else if checked == 2 {
		return "you forgot to check out"
	} else if checked == 4 {
		return "error"
	}
	db = database.DB
	currentYear, currentMonth, currentDay, currentTime := getTime()
	_, err := db.Exec("INSERT INTO timesheet (employeeId, year, month, day, checkIn, checkOut) VALUES (?, ?, ?, ?, ?, ?)", id, currentYear, currentMonth, currentDay, currentTime, constants.DID_NOT_CHECK_OUT)
	if err != nil {
		log.Println(err)
		return "error"
	}
	return currentTime
}

// HandleCheckOut function hande checkout
func HandleCheckOut(id int) string {
	checked := checked(id)
	//checked = 0: checked in today, checked = 1: checked out today, checked = 2: forgot to check out, checked = 3: checked out yesterday, checked = 4: error

	if checked == 1 {
		return "you checked out"
	} else if checked == 2 {
		return "you forgot to check out"
	} else if checked == 3 {
		return "you forgot to check in"
	} else if checked == 4 {
		return "error"
	}

	db = database.DB
	currentYear, currentMonth, currentDay, currentTime := getTime()
	_, err := db.Exec("UPDATE timesheet SET checkout = ? where employeeId = ? AND year = ? AND month = ? AND day = ?", currentTime, id, currentYear, currentMonth, currentDay)
	if err != nil {
		log.Println(err)
		return "error"
	}
	return currentTime
}

// function use to get time at location
func getTime() (string, string, string, string) {
	location, _ := time.LoadLocation("Asia/Ho_Chi_Minh")
	nowTime := time.Now().In(location)
	currentYear := strconv.Itoa(nowTime.Year())
	currentMonth := strconv.Itoa(int(nowTime.Month()))
	currentDay := strconv.Itoa(nowTime.Day())
	currentTime := nowTime.Format("2006-01-02 15:04:05")

	return currentYear, currentMonth, currentDay, currentTime
}

// GetTimesheetById Get timesheet by user ID
func GetTimesheetById(id int) ([]TimesheetResponse, error) {
	db = database.DB
	var timesheet []Timesheet
	rows, err := db.Query("SELECT * FROM timesheet where employeeId = ?", id)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var newTimesheet Timesheet
		err := rows.Scan(&newTimesheet.Id, &newTimesheet.EmployeeId, &newTimesheet.Year, &newTimesheet.Month, &newTimesheet.Day, &newTimesheet.CheckIn, &newTimesheet.CheckOut)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		timesheet = append(timesheet, newTimesheet)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	timesheetResponse := make([]TimesheetResponse, len(timesheet))
	for i, timesheetEmployee := range timesheet {
		timesheetResponse[i], err = GetTimesheetResponse(timesheetEmployee)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return timesheetResponse, nil
}

// Kiem tra xem da check in chua
func checked(id int) int {
	db = database.DB
	currentYear, currentMonth, currentDay, _ := getTime()
	fmt.Printf("currentDate: %s %s %s ", currentYear, currentMonth, currentDay)
	timesheet, err := GetTimesheetByMonth(id, currentYear, currentMonth)
	if err != nil {
		log.Println(err)
		return 4
	}
	if len(timesheet) == 0 {
		return 3
	}
	lastCheck := timesheet[len(timesheet)-1]
	if lastCheck.Day == currentDay {
		if lastCheck.CheckOut == "did not check out" {
			return 0 // checked in today
		} else {
			return 1 // checked out today
		}
	} else {
		if lastCheck.CheckOut == "did not check out" {
			return 2 // forgot to check out yesterday
		} else {
			return 3 // checked out yesterday
		}
	}
}

// GetTimesheetByMonth get timesheet response by month
func GetTimesheetByMonth(id int, year string, month string) ([]TimesheetResponse, error) {
	db = database.DB
	var timesheet []Timesheet
	rows, err := db.Query("SELECT * FROM timesheet where employeeId = ? AND year = ? AND month = ?", id, year, month)
	if err != nil {
		log.Println(err)
		return nil, err
	}

	for rows.Next() {
		var newTimesheet Timesheet
		err := rows.Scan(&newTimesheet.Id, &newTimesheet.EmployeeId, &newTimesheet.Year, &newTimesheet.Month, &newTimesheet.Day, &newTimesheet.CheckIn, &newTimesheet.CheckOut)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		timesheet = append(timesheet, newTimesheet)
	}
	if err := rows.Err(); err != nil {
		log.Println(err)
		return nil, err
	}
	timesheetResponse := make([]TimesheetResponse, len(timesheet))
	for i, timesheetEmployee := range timesheet {
		timesheetResponse[i], err = GetTimesheetResponse(timesheetEmployee)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return timesheetResponse, nil
}

//GetDivisionTimesheet get timesheet response by division
func GetDivisionTimesheet(leadId int) ([]TimesheetResponse, error) {
	db = database.DB
	var departmentId string
	err := db.QueryRow("SELECT departmentId FROM employees where id = ?", leadId).Scan(&departmentId)
	fmt.Println(departmentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var timesheet []Timesheet
	rows, err := db.Query("SELECT * FROM timesheet where employeeId IN (SELECT id FROM employees WHERE departmentId = ?)", departmentId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for rows.Next() {
		var newTimesheet Timesheet
		err := rows.Scan(&newTimesheet.Id, &newTimesheet.EmployeeId, &newTimesheet.Year, &newTimesheet.Month, &newTimesheet.Day, &newTimesheet.CheckIn, &newTimesheet.CheckOut)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		timesheet = append(timesheet, newTimesheet)
	}
	fmt.Printf("timesheets: %+v\n", timesheet)
	timesheetResponse := make([]TimesheetResponse, len(timesheet))
	for i, timesheetEmployee := range timesheet {
		timesheetResponse[i], err = GetTimesheetResponse(timesheetEmployee)
		if err != nil {
			log.Println(err)
			return nil, err
		}
	}
	return timesheetResponse, nil

}

//GetDivisionTimesheetByMonth get timesheet response of division by month
func GetDivisionTimesheetByMonth(leadId int, year string, month string) ([]TimesheetResponse, error) {
	timesheets, err := GetDivisionTimesheet(leadId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	var timesheetByMonth []TimesheetResponse
	for _, timesheet := range timesheets {
		if timesheet.Year == year && timesheet.Month == month {
			timesheetByMonth = append(timesheetByMonth, timesheet)
		}
	}

	return timesheetByMonth, nil
}

func ExportTimesheetToCSV(workingHours []WorkingHours) error {
	file, err := os.Create("/data/timesheet.csv")
	defer file.Close()
	if err != nil {
		log.Println(err)
		return err
	}
	writer := csv.NewWriter(file)
	defer writer.Flush()
	header := []string{"Name", "Working Days", "Working Hours"}
	if err := writer.Write(header); err != nil {
		log.Println(err)
		return err
	}
	for _, workingHour := range workingHours {
		row := []string{workingHour.Name, strconv.Itoa(workingHour.Day), strconv.Itoa(workingHour.Hour)}
		err := writer.Write(row)
		if err != nil {
			log.Println(err)
			return err
		}
	}
	return nil
}

// CalculateWorkingHours calculate working hours of employee

func CalculateWorkingHours(timesheet []Timesheet, leadId int) ([]WorkingHours, error) {
	var workingHours []WorkingHours
	employees, err := GetDivisionEmployees(leadId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	workingHours = make([]WorkingHours, len(employees))
	for i, employee := range employees {
		workingHours[i].Name = employee.Name
		workingHours[i].Day = 0
		workingHours[i].Hour = 0
		for _, timesheetEmployee := range timesheet {
			if employee.Id == timesheetEmployee.EmployeeId {
				if timesheetEmployee.CheckOut == constants.DID_NOT_CHECK_OUT {
					workingHours[i].Hour += 0
				} else {
					hours, _, _ := CalculateTime(timesheetEmployee.CheckIn, timesheetEmployee.CheckOut)
					workingHours[i].Hour += hours
				}
				workingHours[i].Day++
			}
		}
	}
	fmt.Printf("workingHours: %+v\n", workingHours)
	return workingHours, nil

}

// CalculateTime  function to calculate time between two time
func CalculateTime(checkIn string, checkOut string) (int, int, int) {
	var checkInTime time.Time
	var checkOutTime time.Time
	var duration time.Duration
	var hours int
	var minutes int
	var seconds int
	checkInTime, _ = time.Parse("2006-01-02 15:04:05", checkIn)
	checkOutTime, _ = time.Parse("2006-01-02 15:04:05", checkOut)
	duration = checkOutTime.Sub(checkInTime)
	hours = int(duration.Hours())
	minutes = int(duration.Minutes())
	seconds = int(duration.Seconds())
	return hours, minutes, seconds
}

//ExportTimesheet handle export timesheet
func ExportTimesheet(leadId int, year string, month string) ([]WorkingHours, error) {
	timesheets, err := GetDivisionNormalTimesheetByMonth(leadId, year, month)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	workingHours, err := CalculateWorkingHours(timesheets, leadId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	err = ExportTimesheetToCSV(workingHours)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return workingHours, nil
}

//GetTimesheetResponse function to transform timesheet to timesheet response
func GetTimesheetResponse(timesheet Timesheet) (TimesheetResponse, error) {
	var response TimesheetResponse
	response.Id = timesheet.Id
	response.Year = timesheet.Year
	response.Month = timesheet.Month
	response.Day = timesheet.Day
	response.CheckIn = timesheet.CheckIn
	response.CheckOut = timesheet.CheckOut
	employee, err := GetEmployeeById(timesheet.EmployeeId)
	if err != nil {
		log.Println(err)
		return response, err
	}
	response.Employee = employee.Name
	return response, nil
}

//GetDivisionNormalTimesheetByMonth return timesheet of division by month
func GetDivisionNormalTimesheetByMonth(leadId int, year string, month string) ([]Timesheet, error) {
	db = database.DB
	var timesheets []Timesheet
	employees, err := GetDivisionEmployees(leadId)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	for _, employee := range employees {
		rows, err := db.Query("SELECT * FROM timesheet WHERE employeeId = ? AND year = ? AND month = ?", employee.Id, year, month)
		if err != nil {
			log.Println(err)
			return nil, err
		}
		for rows.Next() {
			var newTimesheet Timesheet
			err := rows.Scan(&newTimesheet.Id, &newTimesheet.EmployeeId, &newTimesheet.Year, &newTimesheet.Month, &newTimesheet.Day, &newTimesheet.CheckIn, &newTimesheet.CheckOut)
			if err != nil {
				log.Println(err)
				return nil, err
			}
			timesheets = append(timesheets, newTimesheet)
		}
	}
	return timesheets, nil
}
