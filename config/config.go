package config

import (
	"github.com/joho/godotenv"
)

// LoadEnv carrega variáveis de ambiente do arquivo .env
func LoadEnv() error {
	return godotenv.Load()
}
