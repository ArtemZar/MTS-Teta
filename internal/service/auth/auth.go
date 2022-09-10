package auth

import (
	"github.com/ArtemZar/MTS-Teta/internal/config"
)

// TODO checkink from DB with hashing pass
func CheckCredentials(cfg *config.Config, username, password string) bool {
	if pass, ok := cfg.Credentials[username]; ok && password == pass {
		return true
	}
	return false
}
