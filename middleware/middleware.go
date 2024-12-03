package middleware

import (
	util "Alya-Ecommerce-Go/utils"
	"net/http"
	"strings"

	"github.com/gofiber/fiber/v2"
	"github.com/supabase-community/supabase-go"
)

func ValidateTokenMiddleware(c *supabase.Client) fiber.Handler {
	return func(ctx *fiber.Ctx) error {
		authHeader := ctx.Get("Authorization")
		if authHeader == "" {
			return util.GenerateResponse(ctx, http.StatusUnauthorized, "Missing Authorization header", "")
		}

		tokenParts := strings.Split(authHeader, " ")
		if len(tokenParts) != 2 || tokenParts[0] != "Bearer" {
			return util.GenerateResponse(ctx, http.StatusUnauthorized, "Invalid Authorization format", "")
		}

		token := tokenParts[1]

		isValid := util.CheckToken(c, token)

		if !isValid {
			return util.GenerateResponse(ctx, http.StatusUnauthorized, "Invalid or Expired Token", "")
		}
		return ctx.Next()
	}
}
