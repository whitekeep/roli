package utils

import (
	"roli/config"
)

const (
	Everyone = iota
	Moderator
	Admin
	Owner
	Developer
)

// HavePermission TODO: Переделать - сделать алгоритм более эффективным
func HavePermission(userRoles []string, whitelist config.Roles, requiredLevel int) bool {

	if requiredLevel == Everyone {
		return true
	}

	for _, role := range userRoles {
		if requiredLevel <= Moderator {
			if ArrayContains(whitelist.ModeratorRoles, role) {
				return true
			}
		}
		if requiredLevel <= Admin {
			if ArrayContains(whitelist.AdminRoles, role) {
				return true
			}
		}
		if requiredLevel <= Owner {
			if ArrayContains(whitelist.OwnerRoles, role) {
				return true
			}
		}
		if requiredLevel <= Developer {
			if ArrayContains(whitelist.DevRoles, role) {
				return true
			}
		}
	}

	return false
}
