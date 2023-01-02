package utils

import (
	"errors"
	"roli/config"
)

const (
	Everyone = iota
	Moderator
	Admin
	Owner
	Developer
)

func HavePermission(userRoles []string, whitelist config.Roles, requiredLevel int) (bool, error) {

	switch {

	case requiredLevel <= Everyone:
		return true, nil

	case requiredLevel <= Moderator:
		return HaveIntersect(userRoles, whitelist.ModeratorRoles), nil

	case requiredLevel <= Admin:
		return HaveIntersect(userRoles, whitelist.AdminRoles), nil

	case requiredLevel <= Developer:
		return HaveIntersect(userRoles, whitelist.DevRoles), nil

	default:
		return false, errors.New("invalid permission level")
	}
}
