package helpers

import (
	"fmt"
	"strconv"
)

func IntToString(n int) string {
	return strconv.Itoa(n)
}

func StringToInt(s string) int {
	i, err := strconv.Atoi(s)
	if err != nil {
		fmt.Println(err)
	}
	return i
}
