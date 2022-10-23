package service

import "gin_sample/db"

func GetUser(name string) (string, bool) {
	value, ok := db.DB[name]
	return value, ok
}

func UpdateUser(name string, value string) {
	db.DB[name] = value
}
