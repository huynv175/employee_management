package apis

import (
	"employee-management-system/constants"
	"employee-management-system/models"
	"github.com/gin-gonic/gin"
	"strconv"
)

func GetDivisionEmployees(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error checking permission",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.DIVISION_LEAD && permission != constants.ADMIN {
		c.JSON(403, gin.H{
			"error":   "Forbidden",
			"message": "You do not have permission to access this resource",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	employees, err := models.GetDivisionEmployees(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting division employees",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":   "Successfully retrieved division employees",
		"employees": employees,
	})
}

func UpdateDivisionLead(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error checking permission",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.DIVISION_LEAD && permission != constants.ADMIN {
		c.JSON(403, gin.H{
			"error":   "Forbidden",
			"message": "You do not have permission to access this resource",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error converting id to int",
			"error":   err.Error(),
		})
		return
	}
	var employee models.EmployeeResponse
	err = c.BindJSON(&employee)
	employee.Id = id
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error binding JSON",
			"error":   err.Error(),
		})
		return
	}
	err = models.UpdateDivisionLead(employee)
	c.JSON(200, gin.H{
		"message": "Successfully updated division lead",
	})
}

// GetDivisionRequests gets requests for a division
func GetDivisionRequests(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error checking permission",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.DIVISION_LEAD && permission != constants.ADMIN {
		c.JSON(403, gin.H{
			"error":   "Forbidden",
			"message": "You do not have permission to access this resource",
		})
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	requests, err := models.GetDivisionRequests(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting division requests",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Successfully retrieved division requests",
		"requests": requests,
	})
}

// UpdateRequestByLeader updates a request by a division leader
func UpdateRequestByLeader(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error checking permission",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.DIVISION_LEAD && permission != constants.ADMIN {
		c.JSON(403, gin.H{
			"error":   "Forbidden",
			"message": "You do not have permission to access this resource",
		})
		return
	}

	requestId, err := strconv.Atoi(c.Param("request_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error converting id to int",
			"error":   err.Error(),
		})
		return
	}
	var request models.Request
	err = c.BindJSON(&request)
	request.ID = requestId
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error binding JSON",
			"error":   err.Error(),
		})
		return
	}
	err = models.UpdateRequestByLeader(request)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error updating request",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Successfully updated request",
	})
}

func GetDivisionTimesheet(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error checking permission",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.DIVISION_LEAD && permission != constants.ADMIN {
		c.JSON(403, gin.H{
			"error":   "Forbidden",
			"message": "You do not have permission to access this resource",
		})
		return
	}
	leadId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error converting id to int",
			"error":   err.Error(),
		})
		return
	}

	timesheet, err := models.GetDivisionTimesheet(leadId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting division timesheet",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":   "Successfully retrieved division timesheet",
		"timesheet": timesheet,
	})

}

func GetDivisionTimesheetByMonth(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error checking permission",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.DIVISION_LEAD && permission != constants.ADMIN {
		c.JSON(403, gin.H{
			"error":   "Forbidden",
			"message": "You do not have permission to access this resource",
		})
		return
	}
	year := c.Param("year")
	month := c.Param("month")
	leadId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error converting id to int",
			"error":   err.Error(),
		})
		return
	}
	timesheet, err := models.GetDivisionTimesheetByMonth(leadId, year, month)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting division timesheet",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":            "Successfully retrieved division timesheet by month",
		"timesheet by month": timesheet,
	})
}

func ExportTimesheet(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error checking permission",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.DIVISION_LEAD && permission != constants.ADMIN {
		c.JSON(403, gin.H{
			"error":   "Forbidden",
			"message": "You do not have permission to access this resource",
		})
		return
	}
	year := c.Param("year")
	month := c.Param("month")
	leadId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error converting id to int",
			"error":   err.Error(),
		})
		return
	}
	workingHours, err := models.ExportTimesheet(leadId, year, month)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting division timesheet",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":       "Successfully retrieved division timesheet by month",
		"working hours": workingHours,
	})
}

func ExportEmployee(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error checking permission",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.DIVISION_LEAD && permission != constants.ADMIN {
		c.JSON(403, gin.H{
			"error":   "Forbidden",
			"message": "You do not have permission to access this resource",
		})
		return
	}

	leadId, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error converting id to int",
			"error":   err.Error(),
		})
		return
	}
	employees, err := models.ExportEmployee(leadId)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Error getting division timesheet",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":   "Successfully exported division timesheet by month",
		"employees": employees,
	})
}
