package helpers

import (
	"fmt"
	"net/http"
	"regexp"
	"strconv"
	"strings"
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

func ParseFormCollection(r *http.Request, typeName string) (result []interface{}) {
	var temp []map[string]string
	r.ParseForm()
	for key, values := range r.Form {
		re := regexp.MustCompile(typeName + "\\[([0-9]+)\\]\\[([a-zA-Z]+)\\]")
		matches := re.FindStringSubmatch(key)

		if len(matches) >= 3 {

			index, _ := strconv.Atoi(matches[1])

			for ; index >= len(temp); {
				temp = append(temp, map[string]string{})
			}

			temp[index][matches[2]] = strings.TrimSpace(values[0])
		}
	}
	result = append(result, temp)
	return result
}