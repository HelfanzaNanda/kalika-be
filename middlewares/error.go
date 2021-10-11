package middlewares

import (
	"encoding/json"
	"fmt"
	"github.com/labstack/echo"
	"kalika-be/config"
	"runtime"
	"sort"
)

func ErrorHandler(err error, c echo.Context) {
	he, ok := err.(*echo.HTTPError)
	if ok {
		if he.Internal != nil {
			if herr, ok := he.Internal.(*echo.HTTPError); ok {
				he = herr
			}
		}
	} else {
		he = &echo.HTTPError{
			Code:    500,
			Message: "Failed to connect to the server, please try again later.",
		}
	}

	code := he.Code

	env := config.Get("APP_ENV").String()
	if code >= 500 && (env == "production" || env == "development") {
		temp := map[int]string{}
		trace := []string{}
		for i := 0; i <= 15; i++ {
			fun, file, no, _ := runtime.Caller(i)
			if file != "" {
				temp[i] = fmt.Sprintf("%s:%d on %s", file, no, runtime.FuncForPC(fun).Name())
			}
		}
		index := make([]int, 0)
		for i := range temp {
			index = append(index, i)
		}
		sort.Ints(index)
		for _, i := range index {
			trace = append(trace, fmt.Sprintf("#%d ", i)+temp[i])
		}

		log := map[string]interface{}{}
		log["env"] = env
		log["error"] = err.Error()
		log["request"] = c.Request().Method + " " + c.Path()
		log["body"] = BindBodyRequest(c)
		log["trace"] = trace

		dataJson, _ := json.MarshalIndent(log, "", "  ")
		if env == "production" || env == "development" {
			fmt.Println(string(dataJson))
		} else {
			fmt.Println(string(dataJson))
		}
	}
}

func BindBodyRequest(c echo.Context) echo.Map {
	body := echo.Map{}
	c.Bind(&body)
	return body
}
