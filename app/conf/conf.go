package conf

import (
	"strconv"
	"strings"
)

var (
	Configs map[string]map[string]string
)

func Init() {
	Configs = make(map[string]map[string]string)
}

func Add(key string, config map[string]string) {
	Configs[pars(key)] = parsMap(config)
}

func Get(key string) string {
	keys := strings.Split(pars(key), ".")
	if val, ok := Configs[pars(keys[0])][pars(keys[1])]; ok {
		return val
	}
	return ""
}

func pars(key string) string {
	key = strings.ToUpper(key)
	return strings.ReplaceAll(key, "-", "_")
}

func parsMap(keys map[string]string) map[string]string {
	var mapKey = make(map[string]string)
	for k, v := range keys {
		mapKey[pars(k)] = v
	}
	return mapKey
}

func GetInt(key string) int {
	keys := strings.Split(key, ".")
	value, err := strconv.Atoi(Configs[keys[0]][keys[1]])
	if err != nil {
		return 0
	}
	return value
}

func GetBool(key string) bool {
	keys := strings.Split(key, ".")
	value, err := strconv.ParseBool(Configs[pars(keys[0])][pars(keys[1])])
	if err != nil {
		return false
	}
	return value
}
