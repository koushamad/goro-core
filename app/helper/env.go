package helper

import (
	"github.com/joho/godotenv"
	"github.com/koushamad/goro/app/exception/throw"
	"os"
	"strconv"
	"strings"
)

func Env(env string, def string) string {
	err := godotenv.Load()
	throw.Fatal("Cannot Load Env File", 107, err)
	if os.Getenv(strings.ToUpper(env)) == "" {
		return def
	}
	return os.Getenv(strings.ToUpper(env))
}

func EnvToInt(env string, def int) int {
	err := godotenv.Load()
	throw.Fatal("Cannot Load Env File", 108, err)
	value := Env(env, "")
	if value == "" {
		return def
	}
	config, err := strconv.Atoi(value)
	throw.Fatal("Cannot Get Env Worker Number", 109, err)
	return config
}

func EnvToBool(env string, def bool) bool {
	err := godotenv.Load()
	throw.Fatal("Cannot Load Env File", 110, err)
	value := Env(env, "")
	if value == "" {
		return def
	}
	config, err := strconv.ParseBool(value)
	throw.Fatal("Cannot Get Env Worker Number", 111, err)
	return config
}
