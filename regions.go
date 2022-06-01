package regions

func IsValid(name string) bool {
	for _, r := range REGIONS {
		if r == name {
			return true
		}
	}
	return false
}
