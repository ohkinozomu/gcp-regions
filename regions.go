package regions

func IsValid(name string) bool {
	if _, ok := REGION_DATA[name]; ok {
		return true
	}
	return false
}
