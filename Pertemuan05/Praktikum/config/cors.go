package config

var allowedOrigins = []string{
	"http://localhost:8080/",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
