package config

import "os"

type Config struct {
	MongoURI     string
	GRPCPort     string
	RESTPort     string
	DatabaseName string
}

func Load() *Config {
	return &Config{
		MongoURI:     os.Getenv("MONGO_URI"),
		GRPCPort:     getEnv("GRPC_PORT", "50051"),
		RESTPort:     getEnv("REST_PORT", "8080"),
		DatabaseName: getEnv("MONGO_DB", "analytics"),
	}
}

func getEnv(key, fallback string) string {
	if val, ok := os.LookupEnv(key); ok {
		return val
	}
	return fallback
}
