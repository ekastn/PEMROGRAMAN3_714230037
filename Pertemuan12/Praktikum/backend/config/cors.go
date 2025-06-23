package config

var allowedOrigins = []string{
	"http://localhost:8080/",
	"http://localhost:5173/",
	"http://localhost:8088/",
	"http://127.0.0.1:8088/",
}

func GetAllowedOrigins() []string {
	return allowedOrigins
}
