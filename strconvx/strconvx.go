package strconvx

import (
	"log"
	"strconv"
)

// Convert the string to int and return the value.
func StrToInt(s string) int {
	num, err := strconv.Atoi(s)
	if err != nil {
		log.Panicf("error when converting %v in StrToInt: %v", s, err)
	}
	return num
}

// Convert the int to string and return the value.
func IntToStr(n int) string {
	return strconv.Itoa(n)
}
