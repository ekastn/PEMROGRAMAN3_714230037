package config

var allowedOrigins = []string{
	"http://localhost:8080/",
	"http://localhost:5173/",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
