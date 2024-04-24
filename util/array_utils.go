package util

func InArray[T comparable](target T, array []T) bool {
	for _, element := range array {
		if element == target {
			return true
		}
	}
	return false
}

func SliceTypedToSliceInterface[T comparable](arr []T) []interface{} {
	var interfaceSlice []interface{}
	for _, v := range arr {
		interfaceSlice = append(interfaceSlice, v)
	}

	return interfaceSlice
}
