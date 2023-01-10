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

	// ЧЕРЕЗ МНОЖЕСТВО
	// получить часть вайтлиста под требуемый уровень
	// составить множество ролей, которых достаточно для выполнения команды
	// составить множество ролей пользователя
	// найти пересечение

	// ЧЕРЕЗ MAP
	//

	for _, role := range userRoles {
		if ArrayContains(whitelist.ModeratorRoles, role) {
			if requiredLevel <= Moderator {
				return true
			}
		}
		if ArrayContains(whitelist.AdminRoles, role) {
			if requiredLevel <= Admin {
				return true
			}
		}
		if ArrayContains(whitelist.OwnerRoles, role) {
			if requiredLevel <= Owner {
				return true
			}
		}
		if ArrayContains(whitelist.DevRoles, role) {
			if requiredLevel <= Developer {
				return true
			}
		}
	}

	return false
}
