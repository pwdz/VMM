package jwt

import (
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
	jwt.StandardClaims
}

// Middleware function to validate JWT tokens.
func JwtMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		tokenString := c.Request().Header.Get("Authorization")

		if tokenString == "" {
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
				return c.JSON(http.StatusUnauthorized, map[string]string{
					"error": "Invalid token signature",
				})
			}
			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Unauthorized",
			})
		}

		// Check if the token is valid
		if claims, ok := token.Claims.(*CustomClaims); ok && token.Valid {
			c.Set("username", claims.Username)
			return next(c)
		}

		return c.JSON(http.StatusUnauthorized, map[string]string{
			"error": "Unauthorized",
		})
	}
}

























// import (
// 	"log"
// 	"net/http"

// 	"github.com/dgrijalva/jwt-go"
// 	"github.com/labstack/echo/v4"
// 	"github.com/pwdz/VMM/backend/internal/server/constants"
// 	"github.com/pwdz/VMM/backend/internal/server/user"
// )








// func Authorize(next echo.HandlerFunc) echo.HandlerFunc {
// 	return func(c echo.Context) error {
// 		log.Println("Middleware")
// 		// Get the JWT string from the header
// 		tknStr := c.Request().Header.Get("token")

// 		// Initialize a new instance of `Claims`
// 		claims := &user.User{}

// 		// Parse the JWT string and store the result in `claims`.
// 		// Note that we are passing the key in this method as well. This method will return an error
// 		// if the token is invalid (if it has expired according to the expiry time we set on sign in),
// 		// or if the signature does not match
// 		tkn, err := jwt.ParseWithClaims(tknStr, claims, func(token *jwt.Token) (interface{}, error) {
// 			return constants.JwtKey, nil
// 		})
// 		if err != nil {
// 			if err == jwt.ErrSignatureInvalid {
// 				return c.String(http.StatusUnauthorized, "Unauthorized")
// 			}
// 			return c.String(http.StatusBadRequest, "Bad request")
// 		}
// 		if !tkn.Valid {
// 			return c.String(http.StatusUnauthorized, "Unauthorized")
// 		}

// 		// Finally, return the welcome message to the user, along with their
// 		// username given in the token
// 		c.Set("username", claims.Username)
// 		c.Set("role", claims.Role)
// 		return next(c)
// 	}
// }
