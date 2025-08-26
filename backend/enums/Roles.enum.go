package enums

import (
	"database/sql/driver"
	"fmt"
)

type UserRole string

const (
	// Roles de usuario
	RoleAdmin      UserRole = "admin"
	RoleAgent      UserRole = "agent"
	RoleCustomer   UserRole = "customer"
	RoleManager    UserRole = "manager"
	RoleSupervisor UserRole = "supervisor"
)

func (ur *UserRole) Scan(value interface{}) error {
	if value == nil {
		*ur = RoleCustomer
		return nil
	}
	if str, ok := value.(string); ok {
		*ur = UserRole(str)
		return nil
	}
	return fmt.Errorf("cannot scan %T into UserRole", value)
}

func (ur UserRole) Value() (driver.Value, error) {
	return string(ur), nil
}
