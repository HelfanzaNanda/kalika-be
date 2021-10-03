package middlewares

import (
	"github.com/labstack/echo"
	"kalika-be/helpers"
	"strings"
)

func Auth(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// DISABLED ONPRODUCTION
		//return next(c)

		headerAuthorization := c.Request().Header.Get(echo.HeaderAuthorization)
		if !strings.Contains(headerAuthorization, "Bearer") {
			res := map[string]interface{}{}
			res["code"] = 400
			res["message"] = "Invalid Token Header Format"
			res["data"] = nil
			return c.JSON(400, res)
		}

		tokenString := strings.Replace(headerAuthorization, "Bearer ", "", -1)
		claimToken, err := helpers.JwtParse(tokenString)
		if err != nil {
			return c.JSON(400, claimToken)
		}
		claimToken["token"] = tokenString
		c.Set("userInfo", claimToken)

		return next(c)
	}
}