package utils

import "testing"

func TestGenerateString(t *testing.T) {
	test := GenerateString(5)
	if len(test) != 5 {
		t.Error("String's length is wrong")
	}
}
