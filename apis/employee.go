package apis

import (
	"employee-management-system/models"
	"fmt"
	"strconv"

	"github.com/gin-gonic/gin"
)

//Login API used for login
func Login(c *gin.Context) {
	email, password := c.PostForm("email"), c.PostForm("password")
	employee, err := models.Login(email, password)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Invalid credentials",
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Login successful",
		"employee": employee,
	})
}

// GetEmployeeById GetById API use when user after login or AD, DL want to see user information
func GetEmployeeById(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	employee, err := models.GetEmployeeById(id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Employee not found",
		})
		return
	}
	c.IndentedJSON(200, employee)
}

// CheckIn API use when user check in
func CheckIn(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	message := models.HandleCheckIn(id)

	if message == "error" {
		c.JSON(404, gin.H{"message": "check in fail"})
		return
	}

	c.JSON(200, gin.H{"check in": message})
}

// CheckOut API use when user check out
func CheckOut(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	message := models.HandleCheckOut(id)

	if message == "error" {
		c.JSON(404, gin.H{"message": "check out fail"})
		return
	} else if message == "forgot" {
		c.JSON(404, gin.H{"message": "you forgot check out, now you check in today"})
		return
	} else if message == "didn't check in" {
		c.JSON(404, gin.H{"message": "you must check in"})
		return
	}

	c.JSON(200, gin.H{"check out": message})
}

// GetTimesheet API use for get timesheet of employee
func GetTimesheet(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	timesheet, err := models.GetTimesheetById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err})
	}

	c.JSON(200, gin.H{"amount": len(timesheet), "timesheet": timesheet, "message": "success"})
}

// GetTimesheetByMonth API use for get employee's timesheet by specific month
func GetTimesheetByMonth(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	month := c.Param("month")
	year := c.Param("year")
	timesheet, err := models.GetTimesheetByMonth(id, year, month)
	if err != nil {
		c.JSON(404, gin.H{"error": err})
	}

	c.JSON(200, gin.H{"amount": len(timesheet), "timesheet": timesheet, "message": "success"})
}

func GetRequest(c *gin.Context) {
	id, _ := strconv.Atoi(c.Param("id"))
	requests, err := models.GetRequestByEmployeeId(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err})
		return
	}

	c.JSON(200, gin.H{"amount": len(requests), "requests": requests, "message": "success"})
}

func CreateRequest(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return
	}
	var request models.Request
	err = c.BindJSON(&request)
	fmt.Printf("request: %+v\n", request)
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error(), "message": "invalid request"})
		return
	}
	request.EmployeeID = id
	requestResponse, err := models.CreateRequest(request)
	if err != nil {
		c.JSON(400, gin.H{"error": err.Error(), "message": "create request fail"})
		return
	}

	c.JSON(201, gin.H{"message": "create request success", "request": requestResponse})
}

//CheckPermission func use for check permission of employee
func CheckPermission(c *gin.Context) (int, error) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.IndentedJSON(400, gin.H{"error": err.Error()})
		return 0, err
	}
	employee, err := models.GetEmployeeById(id)
	if err != nil {
		c.JSON(404, gin.H{"error": err.Error(), "message": "employee not found"})
		return 0, err
	}
	return employee.RoleId, nil
}
