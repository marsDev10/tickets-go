package enums

import (
	"database/sql/driver"
	"fmt"
)

// UserRole representa los diferentes roles que puede tener un usuario
type UserRole string

const (
	// Roles a nivel organizacional (GlobalRole)
	Admin UserRole = "admin" // Acceso total a la organización
	User  UserRole = "user"  // Usuario regular de la organización

	// Roles a nivel de equipo (TeamMember.Role)
	Manager    UserRole = "manager"    // Gestiona el equipo, asigna tickets, ve métricas del equipo
	Supervisor UserRole = "supervisor" // Revisa y aprueba tickets, acceso a reportes
	Agent      UserRole = "agent"      // Trabaja en tickets asignados, puede actualizar y resolver
	Member     UserRole = "member"     // Miembro básico del equipo, trabaja tickets
	Viewer     UserRole = "viewer"     // Solo visualiza tickets (stakeholders, clientes internos)
)

// Scan implementa sql.Scanner para leer desde la base de datos
func (r *UserRole) Scan(value interface{}) error {
	if value == nil {
		*r = Member // valor por defecto
		return nil
	}

	switch v := value.(type) {
	case string:
		*r = UserRole(v)
		return nil
	case []byte:
		*r = UserRole(string(v))
		return nil
	default:
		return fmt.Errorf("cannot scan type %T into UserRole", value)
	}
}

// Value implementa driver.Valuer para escribir a la base de datos
func (r UserRole) Value() (driver.Value, error) {
	return string(r), nil
}

// String retorna la representación en string del rol
func (r UserRole) String() string {
	return string(r)
}

// IsValid verifica si el rol es válido
func (r UserRole) IsValid() bool {
	switch r {
	case Admin, User, Manager, Supervisor, Agent, Member, Viewer:
		return true
	default:
		return false
	}
}

// IsTeamRole verifica si el rol es válido para equipos
func (r UserRole) IsTeamRole() bool {
	switch r {
	case Manager, Supervisor, Agent, Member, Viewer:
		return true
	default:
		return false
	}
}

// IsGlobalRole verifica si el rol es válido a nivel organizacional
func (r UserRole) IsGlobalRole() bool {
	switch r {
	case Admin, User:
		return true
	default:
		return false
	}
}

// GetPermissionLevel retorna un nivel numérico de permisos (útil para comparaciones)
func (r UserRole) GetPermissionLevel() int {
	switch r {
	case Admin:
		return 100
	case Manager:
		return 80
	case Supervisor:
		return 60
	case Agent:
		return 40
	case Member:
		return 30
	case User:
		return 20
	case Viewer:
		return 10
	default:
		return 0
	}
}

// HasHigherOrEqualPermission compara si este rol tiene permisos mayores o iguales que otro
func (r UserRole) HasHigherOrEqualPermission(other UserRole) bool {
	return r.GetPermissionLevel() >= other.GetPermissionLevel()
}

// AllTeamRoles retorna todos los roles válidos para equipos
func AllTeamRoles() []UserRole {
	return []UserRole{Manager, Supervisor, Agent, Member, Viewer}
}

// AllGlobalRoles retorna todos los roles válidos a nivel organizacional
func AllGlobalRoles() []UserRole {
	return []UserRole{Admin, User}
}
