package routers

import (
	"employee-management-system/apis"
	"github.com/gin-gonic/gin"
	"log"
)

func Router() {
	router := gin.Default()

	// employee screen
	router.POST("/login", apis.Login)
	router.GET("/:id/profile", apis.GetEmployeeById)
	router.POST("/:id/checkin", apis.CheckIn)
	router.POST("/:id/checkout", apis.CheckOut)
	router.GET("/:id/timesheet", apis.GetTimesheet)
	router.GET("/:id/timesheet/:year/:month", apis.GetTimesheetByMonth)
	router.GET("/:id/request", apis.GetRequest)
	router.POST("/:id/request", apis.CreateRequest)

	// admin screen (only admin can access)
	//employee
	router.GET("/:id/admin/employee", apis.GetAllEmployees)
	router.POST("/:id/admin/employee", apis.CreateEmployee)
	router.GET("/:id/admin/employee/:employee_id", apis.GetEmployeeByIdAdmin)
	router.PUT("/:id/admin/employee/:employee_id/edit", apis.UpdateEmployee)
	router.DELETE("/:id/admin/employee/:employee_id", apis.DeleteEmployee)
	router.POST("/:id/admin/employee/:employee_id/password", apis.UpdatePassword)
	//department
	router.GET("/:id/admin/department", apis.GetAllDepartments)
	router.GET("/:id/admin/department/:department_id", apis.GetDepartmentById)
	router.GET("/:id/admin/department/:department_id/employee", apis.GetEmployeeOfDepartment)
	router.POST("/:id/admin/department", apis.CreateDepartment)
	router.PUT("/:id/admin/department/:department_id/edit", apis.UpdateDepartment)
	router.DELETE("/:id/admin/department/:department_id", apis.DeleteDepartment)
	//request
	router.GET("/:id/admin/request", apis.GetAllRequests)
	router.GET("/:id/admin/request/:employee_id", apis.GetRequestByEmployeeId)
	router.POST("/:id/admin/request/:request_id/edit", apis.UpdateRequest)

	//division leader screen
	router.GET("/:id/division/employee", apis.GetDivisionEmployees)
	router.PUT("/:id/profile/edit", apis.UpdateDivisionLead)
	router.GET("/:id/division/request", apis.GetDivisionRequests)
	router.PUT("/:id/division/request/:request_id/edit", apis.UpdateRequestByLeader)
	router.GET("/:id/division/timesheet", apis.GetDivisionTimesheet)

	router.GET("/:id/division/timesheet/:year/:month", apis.GetDivisionTimesheetByMonth)
	router.GET("/:id/division/timesheet/:year/:month/export", apis.ExportTimesheet)

	router.GET("/:id/division/employee/export", apis.ExportEmployee)

	log.Fatal(router.Run(":8080"))
}
