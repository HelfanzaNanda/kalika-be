package config

import (
	"github.com/joho/godotenv"
	"os"
	"strconv"
	"strings"
	"time"
)

var Default = map[string]Config{
	"APP_URL": "localhost",
	"APP_ENV": "production",
	"APP_PORT": "4000",
	"DB_DRIVER": "mysql",
	"DB_HOST": "127.0.0.1",
	"DB_NAME": "kalika_pos",
	"DB_PORT": "3306",
	"DB_USER": "root",
	"DB_PASSWORD": "",
	"DB_MAX_OPEN_CONNS": "100",
	"DB_MAX_IDLE_CONNS": "2",
	"DB_CONN_MAX_LIFETIME": "0ms",
	"JWT_KEY": "74217b8dadbb365783580f9e92487d701cd38fe6",
}

type Config string

func init() {
	godotenv.Load()
}

func Get(key string) Config {
	value := Config(os.Getenv(key))
	if value == "" {
		value = Default[key]
	}

	return value
}

func (c Config) String() string {
	return string(c)
}

func (c Config) Int() int {
	v, err := strconv.Atoi(c.String())
	if err != nil {
		return 0
	}
	return v
}

func (c Config) Bool() bool {
	if strings.ToLower(c.String()) == "true" {
		return true
	}
	return false
}

func (c Config) Duration() time.Duration {
	v, err := time.ParseDuration(c.String())
	if err != nil {
		return 0
	}
	return v
}