package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// Config holds the application configuration
type Config struct {
	Port        int
	Env         string
	DatabaseURL string
	LogLevel    string

	// Mempool monitoring configuration
	EthereumRPCURL  string  // HTTP RPC endpoint
	EthereumWSURL   string  // WebSocket endpoint
	UseWebSocket    bool    // Use WebSocket (true) or HTTP (false)
	MinGasPriceGwei float64 // Minimum gas price to monitor (in Gwei)
	MaxTransactions int     // Max transactions to process per second
}

// Validate validates the configuration values
func (c *Config) Validate() error {
	if c.Port < 1 || c.Port > 65535 {
		return fmt.Errorf("invalid port: %d (must be between 1 and 65535)", c.Port)
	}

	validEnvs := map[string]bool{
		"development": true,
		"staging":     true,
		"production":  true,
	}
	if !validEnvs[c.Env] {
		return fmt.Errorf("invalid environment: %s (must be development, staging, or production)", c.Env)
	}

	validLogLevels := map[string]bool{
		"debug": true,
		"info":  true,
		"warn":  true,
		"error": true,
	}
	if !validLogLevels[c.LogLevel] {
		return fmt.Errorf("invalid log level: %s (must be debug, info, warn, or error)", c.LogLevel)
	}

	if c.DatabaseURL == "" {
		return errors.New("database URL cannot be empty")
	}

	// Validate Ethereum node configuration
	if c.UseWebSocket {
		if c.EthereumWSURL == "" {
			return errors.New("Ethereum WebSocket URL is required when USE_WEBSOCKET=true")
		}
	} else {
		if c.EthereumRPCURL == "" {
			return errors.New("Ethereum RPC URL is required when USE_WEBSOCKET=false")
		}
	}

	// Validate MaxTransactions
	if c.MaxTransactions < 1 {
		return fmt.Errorf("invalid max transactions per second: %d (must be at least 1)", c.MaxTransactions)
	}

	// Validate MinGasPriceGwei (if set, must be positive)
	if c.MinGasPriceGwei < 0 {
		return fmt.Errorf("invalid minimum gas price: %.2f Gwei (must be non-negative)", c.MinGasPriceGwei)
	}

	return nil
}

// LoadConfig loads configuration from environment variables
func LoadConfig() (*Config, error) {
	cfg := &Config{
		Port:        getEnvAsInt("PORT", 8080),
		Env:         getEnv("ENV", "development"),
		DatabaseURL: getEnv("DATABASE_URL", "postgres://localhost/frontrun_monitor"),
		LogLevel:    getEnv("LOG_LEVEL", "info"),

		// Mempool monitoring config
		EthereumRPCURL:  getEnv("ETHEREUM_RPC_URL", ""),
		EthereumWSURL:   getEnv("ETHEREUM_WS_URL", ""),
		UseWebSocket:    getEnv("USE_WEBSOCKET", "true") == "true",
		MinGasPriceGwei: getEnvAsFloat("MIN_GAS_PRICE_GWEI", 0.0),
		MaxTransactions: getEnvAsInt("MAX_TRANSACTIONS_PER_SEC", 1000),
	}

	return cfg, nil
}

// getEnv retrieves an environment variable with a default value
func getEnv(key, defaultVal string) string {
	if value, exists := os.LookupEnv(key); exists {
		return value
	}
	return defaultVal
}

// getEnvAsInt retrieves an environment variable as an integer with a default value
func getEnvAsInt(key string, defaultVal int) int {
	valStr := getEnv(key, "")
	if val, err := strconv.Atoi(valStr); err == nil {
		return val
	}
	return defaultVal
}

// getEnvAsFloat retrieves an environment variable as a float with a default value
func getEnvAsFloat(key string, defaultVal float64) float64 {
	valStr := getEnv(key, "")
	if val, err := strconv.ParseFloat(valStr, 64); err == nil {
		return val
	}
	return defaultVal
}
