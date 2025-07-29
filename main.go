package main

import (
	"crypto/sha256"
	"crypto/subtle"
	"fmt"
	"os"
	"regexp"
	"strings"

	"github.com/gofiber/fiber/v2/middleware/keyauth"

	"github.com/gofiber/fiber/v2"

	"github.com/kahnwong/todotxt-api/todo"
)

var (
	apiKey        = os.Getenv("TODOTXT_API_KEY")
	protectedURLs = []*regexp.Regexp{
		regexp.MustCompile("^/today$"),
	}
)

func validateAPIKey(c *fiber.Ctx, key string) (bool, error) {
	hashedAPIKey := sha256.Sum256([]byte(apiKey))
	hashedKey := sha256.Sum256([]byte(key))

	if subtle.ConstantTimeCompare(hashedAPIKey[:], hashedKey[:]) == 1 {
		return true, nil
	}
	return false, keyauth.ErrMissingOrMalformedAPIKey
}

func authFilter(c *fiber.Ctx) bool {
	originalURL := strings.ToLower(c.OriginalURL())

	for _, pattern := range protectedURLs {
		if pattern.MatchString(originalURL) {
			return false
		}
	}
	return true
}

func main() {
	// init
	app := fiber.New()

	// auth
	app.Use(keyauth.New(keyauth.Config{
		Next:      authFilter,
		KeyLookup: "header:X-API-Key",
		Validator: validateAPIKey,
	}))

	// routes
	app.Get("/today", todo.TodoTodayController)

	// start server
	err := app.Listen(os.Getenv("LISTEN_ADDR"))
	if err != nil {
		fmt.Println("Error starting server", err)
	}
}
