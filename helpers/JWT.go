package helpers

import (
	"fmt"
	"github.com/golang-jwt/jwt"
	"kalika-be/config"
)

func JwtGenerator(name, username, roleId, storeId, key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"username":  username,
		"name":  name,
		"roleId": roleId,
		"store_id": storeId,
	})

	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		return err.Error()
	}
	return tokenString
}

func JwtParse(accessToken string) (map[string]interface{}, error) {
	res := map[string]interface{}{}
	res["code"] = 400
	res["message"] = "Signing Method Invalid"
	res["data"] = nil

	token, err := jwt.Parse(accessToken, func(token *jwt.Token) (interface{}, error) {
		if method, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Signing Method Invalid")
		} else if method != jwt.SigningMethodHS256 {
			return nil, fmt.Errorf("Signing method invalid")
		}

		return []byte(config.Get("JWT_KEY").String()), nil
	})

	if err != nil {
		fmt.Println(err.Error())
		return res, err
	}

	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok || !token.Valid {
		res["message"] = "Invalid Token"
		return res, err
	}

	return claims, nil
}
