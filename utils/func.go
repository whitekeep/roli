package utils

// ArrayContains Checks if a string is in an array
func ArrayContains(s []string, str string) bool {
	for _, v := range s {
		if v == str {
			return true
		}
	}

	return false
}

// HaveIntersect Checks if two arrays have an intersection
func HaveIntersect(a []string, b []string) bool {
	for _, val := range a {
		for _, val2 := range b {
			if val == val2 {
				return true
			}
		}
	}
	return false
}
