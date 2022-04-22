package types

import (
	"fmt"
)

const (
	AuthDefault = Auth(0)
)

type Auth int32

func FromRoles(rs []Role) Auth {
	auth := Auth(0)
	for _, r := range rs {
		auth |= r.Auth()
	}
	return auth
}

func (a Auth) Roles() (rs []Role) {
	for _, r := range Role_value {
		if a.Access(Role(r).Auth()) {
			rs = append(rs, Role(r))
		}
	}
	return rs
}

func (a Auth) Access(auth Auth) bool {
	return (a & auth) > 0
}

func (a Auth) IsRootAdmin() bool {
	return (a & RoleRootAdmin.Auth()) > 0
}

func (a Auth) IsPermAdmin() bool {
	return (a & RolePermAdmin.Auth()) > 0
}

// Auth return the auth of the role
func (r Role) Auth() Auth {
	return 1 << r
}

func GetRolesFromStr(strRoles ...string) (roles []Role, err error) {
	for _, strRole := range strRoles {
		role, err := RoleFromstring(strRole)
		if err != nil {
			return roles, err
		}
		roles = append(roles, role)
	}
	return roles, nil
}

// RoleFromstring turns a string into a Auth
func RoleFromstring(str string) (Role, error) {
	option, ok := Role_value[str]
	if !ok {
		return Role(0xff), fmt.Errorf("'%s' is not a valid vote option", str)
	}
	return Role(option), nil
}

// ValidRole returns true if the role is valid and false otherwise.
func ValidRole(role Role) bool {
	for _, r := range Role_value {
		if role == Role(r) {
			return true
		}
	}
	return false
}

// Marshal needed for protobuf compatibility
func (r Role) Marshal() ([]byte, error) {
	return []byte{byte(r)}, nil
}

// Unmarshal needed for protobuf compatibility
func (r *Role) Unmarshal(data []byte) error {
	*r = Role(data[0])
	return nil
}

// Format implements the fmt.Formatter interface.
// nolint: errcheck
func (r Role) Format(s fmt.State, verb rune) {
	switch verb {
	case 's':
		s.Write([]byte(r.String()))
	default:
		s.Write([]byte(fmt.Sprintf("%v", byte(r))))
	}
}
