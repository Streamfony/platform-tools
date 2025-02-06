package env

import "os"

func IsProd() bool {
	return os.Getenv("ENV") == "prod"
}

func IsDev() bool {
	return !IsProd()
}
