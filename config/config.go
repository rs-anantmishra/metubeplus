package config

import (
	"fmt"
	"os"
	"path/filepath"

	"github.com/joho/godotenv"
)

// Config func to get env value

func Config(key string) string {
	relPath := "..\\"
	err := godotenv.Load(filepath.Join(relPath, ".env"))
	if err != nil {
		fmt.Println("Error loading .env file")
	}

	return os.Getenv(key)
}
