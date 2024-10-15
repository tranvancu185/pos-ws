package utils

// Where function filter elements based on conditions.
// Where return a new slice.
//
//	 Example:
//		evens := Filter(list, func(i int) bool {return i % 2 == 0})
func Where[T any](list []T, f func(T) bool) []T {
	var newList []T
	for _, v := range list {
		if f(v) {
			newList = append(newList, v)
		}
	}
	return newList
}

// Find function find first element based on conditions.
// Find return a new value T.
//
//	 Example:
//		item := Find(list, func(n name) bool {return n == "ABC"})
func Find[T any](list []T, f func(T) bool) T {
	var newValue T
	for _, v := range list {
		if f(v) {
			newValue = v
			break
		}
	}
	return newValue
}

// Distinct function remove duplicates from slice.
// Distinct return a new slice.
//
//	 Example:
//		newList := Distinct(oldList)
func Distinct[T string | int32 | int64 | float32 | float64](list []T) []T {
	allKeys := make(map[T]bool)
	newList := []T{}
	for _, item := range list {
		if _, value := allKeys[item]; !value {
			allKeys[item] = true
			newList = append(newList, item)
		}
	}
	return newList
}

func ArrayIn[T int | int64 | string](val T, arr []T) bool {
	for _, item := range arr {
		if item == val {
			return true
		}
	}

	return false
}
