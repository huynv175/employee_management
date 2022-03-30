package models

import (
	"employee-management-system/database"
	"log"
)

type Role struct {
	Id   int    `json:"id"`
	Name string `json:"name"`
}

// GetRoleById returns a role by id
func GetRoleById(id int) (Role, error) {
	db = database.DB
	var role Role
	err := db.QueryRow("SELECT id, name FROM roles WHERE id = ?", id).Scan(&role.Id, &role.Name)
	if err != nil {
		log.Println(err)
		return role, err
	}
	return role, nil
}

// GetRoleByName returns a role by name
func GetRoleByName(name string) (Role, error) {
	db = database.DB
	var role Role
	err := db.QueryRow("SELECT id, name FROM roles WHERE name = ?", name).Scan(&role.Id, &role.Name)
	if err != nil {
		log.Println(err)
		return role, err
	}
	return role, nil
}
