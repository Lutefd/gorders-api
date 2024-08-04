package application

import (
	"os"
	"strconv"
)

type Config struct {
	RedisAddr  string
	RedisPass  string
	ServerPort uint16
}

func LoadConfig() Config {
	redisPassword := os.Getenv("REDIS_PASSWORD")
	redisAddr := os.Getenv("REDIS_ADDR")

	serverPort := os.Getenv("SERVER_PORT")
	parsedServerPort, err := strconv.ParseUint(serverPort, 10, 16)
	if err != nil {
		panic(err)
	}

	return Config{
		RedisAddr:  redisAddr,
		RedisPass:  redisPassword,
		ServerPort: uint16(parsedServerPort),
	}
}
