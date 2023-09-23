package server

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/pwdz/VMM/code/backend/models"
)

func CheckJsonFormatMiddlware(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
		return next(c)
	}
}

// Middleware function to check user role.
func CheckUserRoleMiddleware(expectedRole models.Role) echo.MiddlewareFunc {
	return func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {

			fmt.Println("ASDASDA0")
			// Get the user's role from the context
			role := models.Role(c.Get("role").(string))

			fmt.Println("ASDASDA")
			// Check if the user's role matches the expected role
			if role == expectedRole {
				return next(c)
			}

			return c.JSON(http.StatusUnauthorized, map[string]string{
				"error": "Access Denied",
			})
		}
	}
}

// Create a middleware function to extract user_id from JWT claims
func ExtractUserIDMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
    return func(c echo.Context) error {
        // Retrieve the username from JWT claims
        username := c.Get("username").(string)
        
        // Query your database to find the user by username and retrieve user_id
		if storedUser := DB.FindUserByUsername(username); storedUser != nil {
            return c.JSON(http.StatusUnauthorized, map[string]string{
                "error": "Unauthorized",
            })
		}else{
			// Set user_id in the context
			c.Set("user_id", storedUser.ID)
        	return next(c)
		}        
    }
}
