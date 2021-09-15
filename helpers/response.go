package helpers

import (
	"kalika-be/models/web"
	"strings"
)

func Response(status string, message string, data interface{}) (res web.Response) {
	var code int
	splitStatus := strings.Split(status, "|")
	status = splitStatus[0]
	if len(splitStatus) > 1 {
		if message == "" {
			message = splitStatus[1]
		}
	}

	switch status {
		case "OK":
			code = 200
		case "CREATED":
			code = 201
		case "NOT_FOUND":
			code = 404
		case "INTERNAL_ERROR":
			code = 500
		case "UNAUTHORIZED":
			code = 401
		default:
			code = 400
	}

	res.Code = code
	res.Status = status
	res.Data = data
	res.Message = message

	return
}
