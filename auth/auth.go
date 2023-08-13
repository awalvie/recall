package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"strings"

	"github.com/awalvie/recall/config"
	"github.com/labstack/echo/v4"
)

// IsAuthenticated checks if the user is authenticated by checking the session cookie
func IsAuthenticated(c echo.Context) bool {
	// Get the config from the context
	a := c.Get("app").(*config.App)

	// Get the session cookie
	cookie, err := c.Cookie("auth")

	// Check if the cookie is nil or if there was an error
	if cookie == nil || err != nil {
		return false
	}

	// Get the username and secret from the cookie value
	parts := strings.Split(cookie.Value, ":")

	// Check if the cookie value is in the correct format
	if len(parts) != 2 {
		return false
	}

	username := parts[0]
	secret := parts[1]

	// Check if username and secret are correct
	if username == a.Config.Auth.Username && secret == Secret(username, a.Config.Auth.Password) {
		return true
	}

	return false
}

// Secret returns a hex encoded HMAC-SHA256 hash of the message using the key
func Secret(msg, key string) string {
	// Create a new HMAC-SHA256 hash
	mac := hmac.New(sha256.New, []byte(key))

	// Write the message to the hash
	mac.Write([]byte(msg))

	// Get the hash sum
	src := mac.Sum(nil)

	// Return the hex encoded hash sum
	return hex.EncodeToString(src)
}

// IsPublic checks if the path is publically accessible
func IsPublic(path string) bool {
	// List of public paths
	publicPaths := []string{
		"/login",
		"/static",
	}

	// Check if the path is in the list of public paths
	for _, publicPath := range publicPaths {
		if strings.HasPrefix(path, publicPath) {
			return true
		}
	}

	return false
}
