package env

import "os"

type environment string

const (
	Prod environment = "prod"
	Dev  environment = "dev"
)

func IsProd() bool {
	return os.Getenv("ENV") == string(Prod)
}

func IsDev() bool {
	return os.Getenv("ENV") == string(Dev)
}
