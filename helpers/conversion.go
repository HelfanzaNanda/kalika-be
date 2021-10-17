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

func FormatRupiah(amount float64) string {
	humanizeValue := FormatPrice(int64(amount))
	stringValue := strings.Replace(humanizeValue, ",", ".", -1)
	return "Rp. " + stringValue
}

func FormatPrice(n int64) string {
    in := strconv.FormatInt(n, 10)
    numOfDigits := len(in)
    if n < 0 {
        numOfDigits-- // First character is the - sign (not a digit)
    }
    numOfCommas := (numOfDigits - 1) / 3

    out := make([]byte, len(in)+numOfCommas)
    if n < 0 {
        in, out[0] = in[1:], '-'
    }

    for i, j, k := len(in)-1, len(out)-1, 0; ; i, j = i-1, j-1 {
        out[j] = in[i]
        if i == 0 {
            return string(out)
        }
        if k++; k == 3 {
            j, k = j-1, 0
            out[j] = ','
        }
    }
}