package config

import "os"

func GetPort() string {
	port := os.Getenv("PORT")

	if port == "" {
		port = ":80"
	} else {
		port = ":" + port
	}

	return port
}
