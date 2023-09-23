package jwt

import (
	"fmt"
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	// "github.com/labstack/echo/v4/middleware"
)

// Define a secret key for signing and validating JWT tokens.
var JwtSecret = []byte("your-secret-key")

// CustomClaims represents custom claims to be included in the JWT token.
type CustomClaims struct {
	Username string `json:"username"`
	Role   	 string `json:"role"`
	jwt.StandardClaims
}

// Middleware function to validate JWT tokens.
func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
			// return c.Redirect(http.StatusFound, "/login")
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
		}

		// Parse the token
		token, err := jwt.ParseWithClaims(tokenString, &CustomClaims{}, func(token *jwt.Token) (interface{}, error) {
			return JwtSecret, nil
		})

		if err != nil {
			if err == jwt.ErrSignatureInvalid {
				// return c.Redirect(http.StatusFound, "/login")
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token signature",
				})
			}
			// return c.Redirect(http.StatusFound, "/login")
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			fmt.Println("natals")
			fmt.Println(token.Claims)
			fmt.Println(token.Claims.(*CustomClaims))
			fmt.Println(claims)
			c.Set("username", claims.Username)
			c.Set("role", claims.Role)

			fmt.Println("natals2")
			return next(c)
		}

		// return c.Redirect(http.StatusFound, "/login")
		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}
}

