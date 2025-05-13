package config

import (
	"log"
	"os"
	"strconv"

	"github.com/joho/godotenv"
)

var (
	MiddlewareUrl    string
	MiddlewareSecret string
	Port             string
	TopazChannel     string
	ApplicationId    string
	DeviceId         string
	DeviceIp         string
	RedisHost        string
	RedisTTL         int
	DatabaseConexion string
)

func LoadConfig() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	MiddlewareUrl = os.Getenv("MIDDLEWARE_URL")
	MiddlewareSecret = os.Getenv("MIDDLEWARE_SECRET")
	Port = os.Getenv("PORT")
	TopazChannel = os.Getenv("TOPAZ_CHANNEL")
	ApplicationId = os.Getenv("APPLICATION_ID")
	DeviceId = os.Getenv("DEVICE_ID")
	DeviceIp = os.Getenv("DEVICE_IP")
	RedisHost = os.Getenv("REDIS_HOST")
	RedisTTL, _ = strconv.Atoi(os.Getenv("REDIS_TTL"))
	DatabaseConexion = os.Getenv("DATABASE_CONEXION")
}
