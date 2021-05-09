package phonenumber

import (
	"errors"
	"strings"
	"unicode"
)

var phoneNumberLen = 10
var areaEnd = 3
var midDigitsEnd = 3
var midDigitsStart = 3
var lastDigitsStart = 6

func Number(num string) (string, error) {
	numStr := strings.ReplaceAll(num, "+", "")
	numStr = strings.ReplaceAll(numStr, "(", "")
	numStr = strings.ReplaceAll(numStr, ")", "")
	numStr = strings.ReplaceAll(numStr, "-", "")
	numStr = strings.ReplaceAll(numStr, " ", "")
	numStr = strings.ReplaceAll(numStr, ".", "")
	if len(numStr) > phoneNumberLen {
		numStr = strings.TrimPrefix(numStr, "1")
	}
	for _, r := range numStr {
		if !unicode.IsDigit(r) {
			return "", errors.New("incorrect input")
		}
	}
	numStrSlice := strings.Split(numStr, "")
	if numStrSlice[0] == "0" || numStrSlice[0] == "1" {
		return numStr, errors.New("incorrect number of digits")
	}
	if numStrSlice[3] == "0" || numStrSlice[3] == "1" {
		return numStr, errors.New("incorrect number of digits")
	}
	if len(numStr) != phoneNumberLen {
		return numStr, errors.New("incorrect number of digits")
	}
	return numStr, nil

}
func Format(num string) (string, error) {
	numStr, expectErr := Number(num)
	if expectErr == nil {
		formatted := []string{"("}
		phone := []string{}
		phone = strings.Split(numStr[midDigitsStart:], "")
		lastDigits := numStr[lastDigitsStart:]
		midDigits := phone[:midDigitsEnd]
		midDigits = append(midDigits, "-")
		midStr := strings.Join(midDigits, "")
		area, _ := AreaCode(numStr)
		formatted = append(formatted, area, ")", " ", midStr, lastDigits)
		return strings.Join(formatted, ""), nil
	}
	return "", expectErr

}
func AreaCode(num string) (string, error) {
	numStr, expectErr := Number(num)
	if expectErr == nil {
		area := numStr[:areaEnd]
		return area, nil
	}
	return numStr, expectErr
}
