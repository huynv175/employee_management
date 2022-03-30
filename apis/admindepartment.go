package apis

import (
	"employee-management-system/constants"
	"employee-management-system/models"
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"strconv"
)

// GetAllDepartments API to get all departments
func GetAllDepartments(c *gin.Context) {
	permission, err := CheckPermission(c)
	if err != nil {
		log.Println(err)
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
	var departments []models.DepartmentResponse
	departments, err = models.GetAllDepartment()
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":     "Success",
		"departments": departments,
	})
	return
}

// GetEmployeeOfDepartment  Get a department
func GetEmployeeOfDepartment(c *gin.Context) {
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
	id, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	employees, err := models.GetEmployeesOfDepartment(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":                 "Success",
		"employees of department": employees,
	})
}

// CreateDepartment API create a new department
func CreateDepartment(c *gin.Context) {
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
	var departmentResponse models.DepartmentResponse
	err = c.BindJSON(&departmentResponse)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	department, err := models.CreateDepartment(departmentResponse)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}

	c.JSON(200, gin.H{
		"message":    "Department created successfully",
		"department": department,
	})
}

// UpdateDepartment API update a department
func UpdateDepartment(c *gin.Context) {
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
	id, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	var departmentResponse models.DepartmentResponse
	err = c.BindJSON(&departmentResponse)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	departmentResponse, err = models.UpdateDepartment(id, departmentResponse)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message":    "Department updated successfully",
		"department": departmentResponse,
	})
}

// GetDepartmentById API get a department by id
func GetDepartmentById(c *gin.Context) {
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
	id, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	department, err := models.GetDepartmentById(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"department": department,
		"message":    "Department found successfully",
	})
}

// DeleteDepartment API delete a department
func DeleteDepartment(c *gin.Context) {
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
	id, err := strconv.Atoi(c.Param("department_id"))
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	err = models.DeleteDepartment(id)
	if err != nil {
		c.JSON(500, gin.H{
			"message": "Internal server error",
			"error":   err.Error(),
		})
		return
	}
	c.JSON(200, gin.H{
		"message": "Department deleted successfully",
	})

}

//func CheckPermissionAdmin(c *gin.Context) (bool, error) {
//	permission, err := CheckPermission(c)
//	if err != nil {
//		log.Println(err)
//		return false, err
//	} else if permission != ADMIN {
//		return false, nil
//	}
//	return true, nil
//}
