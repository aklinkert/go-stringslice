package stringslice

// Contains returns whether the given slice contains the given item
func Contains(list []string, item string) bool {
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

// Remove the given item from the given list if exists
func Remove(list []string, item string) []string {
	for i, v := range list {
		if v == item {
			list = append(list[:i], list[i+1:]...)
		}
	}
	return list
}

// Equals checks whether two string slices are equal, i.e. have the same length and contain the same elements
func Equals(list1, list2 []string) bool {
	if len(list1) != len(list2) {
		return false
	}

	if list1 == nil || list2 == nil {
		return false
	}

	for _, s := range list1 {
		if !Contains(list2, s) {
			return false
		}
	}

	for _, s := range list2 {
		if !Contains(list1, s) {
			return false
		}
	}

	return true
}
