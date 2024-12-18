package converter

import (
	"fmt"
)

func ToBinary(num int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("number %d out of range", num)
	}
	return fmt.Sprintf("%b", num), nil
}

func ToHexadecimal(num int) (string, error) {
	if num < 0 {
		return "", fmt.Errorf("number %d out of range", num)
	}
	return fmt.Sprintf("%X", num), nil
}
