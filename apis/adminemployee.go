package apis

import (
	"employee-management-system/constants"
	"employee-management-system/models"
	"fmt"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
)

// GetAllEmployees get information of all employees
func GetAllEmployees(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.ADMIN {
		fmt.Println(permission)
		c.JSON(401, gin.H{
			"message": "Denied Permission",
		})
		return
	}
	employees, err := models.GetAllEmployees()
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "success",
		"data":    employees,
	})
}
func GetEmployeeByIdAdmin(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.ADMIN {
		fmt.Println(permission)
		c.JSON(401, gin.H{
			"message": "Denied Permission",
		})
		return
	}
	id, _ := strconv.Atoi(c.Param("employee_id"))
	employee, err := models.GetEmployeeById(id)
	if err != nil {
		c.JSON(404, gin.H{
			"message": "Employee not found",
			"error":   err.Error(),
		})
		return
	}
	c.IndentedJSON(200, gin.H{"message": "success", "employee": employee})
}
func CreateEmployee(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.ADMIN {
		fmt.Println(permission)
		c.JSON(401, gin.H{
			"message": "Denied Permission",
		})
		return
	}
	var employee models.EmployeeRequest
	err = c.BindJSON(&employee)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{"message": "Bad request", "error": err.Error()})
		return
	}
	employee, err = models.CreateEmployee(employee)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success", "data": employee,
	})
}

func UpdateEmployee(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.ADMIN {
		fmt.Println(permission)
		c.JSON(401, gin.H{
			"message": "Denied Permission",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}
	var employee models.EmployeeRequest
	err = c.BindJSON(&employee)
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}
	employee.Id = id
	err = models.UpdateEmployee(employee)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "success", "employee": employee,
	})
}

func DeleteEmployee(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.ADMIN {
		fmt.Println(permission)
		c.JSON(401, gin.H{
			"message": "Denied Permission",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	err = models.DeleteEmployee(id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "Delete Employee successfully",
	})
}

func UpdatePassword(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		log.Println(err)
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	} else if permission != constants.ADMIN {
		fmt.Println(permission)
		c.JSON(401, gin.H{
			"message": "Denied Permission",
		})
		return
	}
	id, err := strconv.Atoi(c.Param("employee_id"))
	if err != nil {
		log.Println(err)
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}
	password := c.PostForm("password")
	err = models.UpdatePassword(password, id)
	if err != nil {
		c.JSON(400, gin.H{
			"message": "Bad Request",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message": "reset password successfully",
	})
}
