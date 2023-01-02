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

	userLevel := Everyone

	for _, role := range userRoles {
		if ArrayContains(whitelist.ModeratorRoles, role) {
			userLevel = Moderator
		}
		if ArrayContains(whitelist.AdminRoles, role) {
			userLevel = Admin
		}
		if ArrayContains(whitelist.OwnerRoles, role) {
			userLevel = Owner
		}
		if ArrayContains(whitelist.DevRoles, role) {
			userLevel = Developer
		}
	}

	return userLevel >= requiredLevel
}
