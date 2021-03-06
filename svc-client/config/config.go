package config

import (
	"os"
)

type Config struct {
	Server Server
}

func New() Config {
	return Config{
		Server: NewServer(),
	}
}

type Server struct {
	Port         string
	IdentityHost string
	ProfileHost  string
}

func NewServer() Server {
	return Server{
		Port:         os.Getenv("MESHER_CLIENT_PORT"),
		IdentityHost: os.Getenv("MESHER_IDENTITY_HOST"),
		ProfileHost:  os.Getenv("MESHER_PROFILE_HOST"),
	}
}
