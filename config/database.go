package config

import (
	"errors"
	"fmt"
	"os"
)

func GetPGUser() (string, error) {
	if v, found := os.LookupEnv("POSTGRES_USER"); found {
		return v, nil
	} else {
		return "", errors.New("POSTGES_USER is not set")
	}
}

func GetPGPassword() (string, error) {
	if v, found := os.LookupEnv("POSTGRES_PASSWORD"); found {
		return v, nil
	} else {
		return "", errors.New("POSTGRES_PASSWORD is not set")
	}
}

func GetPGDB() (string, error) {
	if v, found := os.LookupEnv("POSTGRES_DB"); found {
		return v, nil
	} else {
		return "", errors.New("POSTGRES_DB is not set")
	}
}

func GetPGAddress() (string, error) {
	if v, found := os.LookupEnv("POSTGRES_ADDRESS"); found {
		return v, nil
	} else {
		return "", errors.New("POSTGRES_DB is not set")
	}
}

func GetPostgresDNS() (string, error) {
	user, err := GetPGUser()
	if err != nil {
		return "", err
	}

	password, err := GetPGPassword()
	if err != nil {
		return "", err
	}

	db, err := GetPGDB()
	if err != nil {
		return "", err
	}

	address, err := GetPGAddress()
	if err != nil {
		return "", err
	}

	return fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True", user, password, address, 3306, db), nil
}
