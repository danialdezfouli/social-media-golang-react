package common

func Contains[T comparable](arr []T, elm T) bool {
	for _, x := range arr {
		if x == elm {
			return true
		}
	}
	return false

}
