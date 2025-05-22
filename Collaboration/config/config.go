package config

import "os"

type Config struct {
	MONGO_URI string
	GRPC_PORT string
	DB string
}

func Load() *Config {
	return &Config{
		MONGO_URI: os.Getenv("MONGO_URI"),
		GRPC_PORT: getConfig("GRPC_PORT", "50051"),
		DB: getConfig("DB", "teams"),
	}
}

func getConfig(key, alt string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return alt
}