package config

import (
	"errors"
	"log"
	"os"
	"path/filepath"
	"strings"

	"github.com/spf13/viper"
)

type (
	AppConfig struct {
		SecretKey string
		DBConfig
	}

	DBConfig struct {
		FirebaseCred string
	}
)

func NewAppConfig() (*AppConfig, error) {
	viper.AutomaticEnv()
	viper.SetEnvKeyReplacer(strings.NewReplacer(".", "_"))

	// Get path to executable
	ext, err := os.Executable()
	if err != nil {
		return nil, err
	}

	// Set config file paths
	viper.SetConfigName("config")
	viper.AddConfigPath(filepath.Join(filepath.Dir(ext), "conf"))
	viper.AddConfigPath(filepath.Join("cmd", "conf"))
	viper.SetConfigType("yml")

	// Read config file if exists
	if err := viper.ReadInConfig(); err != nil {
		log.Printf("Failed to read config file: %v", err)
	}

	// Fetch configuration values
	firebaseCred := getConfigValue("FIREBASE_CREDENTIALS", "database.firebase.credentials")
	secretKey := getConfigValue("SECRET_KEY", "auth.secret")

	// Validate required configuration values
	if firebaseCred == "" {
		return nil, errors.New("FIREBASE_CREDENTIALS is required")
	}

	if secretKey == "" {
		return nil, errors.New("SECRET_KEY is required")
	}

	return &AppConfig{
		SecretKey: secretKey,
		DBConfig: DBConfig{
			FirebaseCred: firebaseCred,
		},
	}, nil
}

// getConfigValue fetches the value from environment variable first, if not found then from config file
func getConfigValue(envKey, configKey string) string {
	value := viper.GetString(envKey)
	if value == "" {
		value = viper.GetString(configKey)
	}
	return value
}
