package converter

import (
	"testing"
)

func TestConvertNumber(t *testing.T) {
	tests := []struct {
		input    int
		system   string
		expected string
		hasError bool
	}{
		{1, "roman", "I", false},
		{2, "roman", "II", false},
		{3, "roman", "III", false},
		{4, "roman", "IV", false},
		{5, "roman", "V", false},
		{6, "roman", "VI", false},
		{7, "roman", "VII", false},
		{8, "roman", "VIII", false},
		{9, "roman", "IX", false},
		{10, "roman", "X", false},
		{37, "roman", "XXXVII", false},
		{42, "roman", "XLII", false},
		{58, "roman", "LVIII", false},
		{949, "roman", "CMXLIX", false},
		{1994, "roman", "MCMXCIV", false},
		{0, "roman", "", true},
		{4000, "roman", "", true},
		{10, "binary", "1010", false},
		{255, "binary", "11111111", false},
		{16, "hexadecimal", "10", false},
		{255, "hexadecimal", "FF", false},
	}

	for _, test := range tests {
		t.Run("ConvertNumber", func(t *testing.T) {
			output, err := ConvertNumber(test.input, test.system)
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
