package main

import "os"

var redisConfigs = RedisConfigs{
	network: getEnv("REDIS_NETWORK", "tcp"),                   // Default to "tcp"
	address: getEnv("REDIS_ADDRESS", "127.0.0.1:6379"),        // Default to "127.0.0.1:6379"
}

var serverConfigs = ServerConfigs{
	addr:   getEnv("SERVER_ADDRESS", "127.0.0.1:8080"),        // Default to "127.0.0.1:8080"
	static: getEnv("STATIC_PATH", "/app/static"),              // Default to "/app/static" (adjust as needed)
}

// Helper function to get environment variables or default values
func getEnv(key, defaultValue string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultValue
}
