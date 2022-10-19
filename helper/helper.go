package helper

func Contains(array []string, c string) bool {
	for _, a := range array {
		if a == c {
			return true
		}
	}
	return false
}
