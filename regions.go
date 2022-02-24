//go:generate go run tools/generator/generate.go
//go:generate gofmt -w data.go

package regions

func IsValid(name string) bool {
	if _, ok := REGION_DATA[name]; ok {
		return true
	}
	return false
}
