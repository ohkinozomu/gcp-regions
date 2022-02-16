package regions

import (
	"testing"
)

func TestIsValid(t *testing.T) {
	validRegionName := "us-central1"
	if IsValid("us-central1") != true {
		t.Fatalf("failed test: %v\n", validRegionName)
	}
	notValidRegionName := "us-central1000"
	if IsValid(notValidRegionName) != false {
		t.Fatalf("failed test: %v\n", notValidRegionName)
	}
}
