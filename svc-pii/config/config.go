package config

import (
	"os"
)

type Config struct {
	Server   Server
}

func New() Config {
	return Config{
		Server:   NewServer(),
	}
}

type Server struct {
	Host         string
}

func NewServer() Server {
	return Server{
		Host: os.Getenv("MESHER_PII_HOST"),
	}
}
