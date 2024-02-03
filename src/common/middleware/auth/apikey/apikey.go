package apikey

import (
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/keyauth"
)

var (
	apiKey    = "pass1"
	adminPath = regexp.MustCompile("/admin")
)

func ValidateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	if key == apiKey {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func AdminAuthFilter(c *fiber.Ctx) bool {
	originalURL := strings.ToLower(c.OriginalURL())

	if adminPath.MatchString(originalURL) {
		return false
	}
	return true
}
