package apis

import (
	"employee-management-system/constants"
	"employee-management-system/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"strconv"
)

// GetAllRequests  Get all admin requests
func GetAllRequests(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.ADMIN {
		c.JSON(401, gin.H{
			"message": "Denied Permission",
		})
		return
	}
	requests, err := models.GetAllRequests()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to get all requests",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":  "Get requests successfully",
		"requests": requests,
	})
}

// GetRequestByEmployeeId func Get request by id
func GetRequestByEmployeeId(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.ADMIN {
		c.JSON(401, gin.H{
			"message": "Denied Permission",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to get request",
			"error":   err.Error(),
		})
		return
	}
	request, err := models.GetRequestByEmployeeId(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to get request",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Get request successfully",
		"request": request,
	})
}

func UpdateRequest(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.ADMIN {
		c.JSON(401, gin.H{
			"message": "Denied Permission",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("request_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to get request",
			"error":   err.Error(),
		})
		return
	}

	var request models.Request
	err = c.BindJSON(&request)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to get request",
			"error":   err.Error(),
		})
		return
	}
	fmt.Println("request: ", request)
	err = models.UpdateRequest(id, request.StatusRequest)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "failed to update request",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Update request successfully",
	})
}
