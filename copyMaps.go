package functions

//MakeMapCopy will received a map
func CopyMapStringString(m map[string]string) map[string]string {
	copy := make(map[string]string)
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

func CopyMapStringInt(m map[string]int64) map[string]int64 {
	copy := make(map[string]int64)
	for key, value := range m {
		copy[key] = value
	}
	return copy
}

