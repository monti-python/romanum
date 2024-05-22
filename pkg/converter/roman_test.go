package converter

import (
	"testing"
)

func TestToRoman(t *testing.T) {
	tests := []struct {
		input    int
		expected string
		hasError bool
	}{
		{1, "I", false},
		{2, "II", false},
		{3, "III", false},
		{4, "IV", false},
		{5, "V", false},
		{6, "VI", false},
		{7, "VII", false},
		{8, "VIII", false},
		{9, "IX", false},
		{10, "X", false},
		{37, "XXXVII", false},
		{42, "XLII", false},
		{58, "LVIII", false},
		{949, "CMXLIX", false},
		{1994, "MCMXCIV", false},
		{0, "", true},
		{4000, "", true},
	}

	for _, test := range tests {
		t.Run("ConvertToRoman", func(t *testing.T) {
			output, err := ToRoman(test.input)
			if test.hasError {
				if err == nil {
					t.Errorf("Expected error for input %d, but got none", test.input)
				}
			} else {
				if err != nil {
					t.Errorf("Did not expect error for input %d, but got: %v", test.input, err)
				}
				if output != test.expected {
					t.Errorf("For input %d, expected %s but got %s", test.input, test.expected, output)
				}
			}
		})
	}
}
