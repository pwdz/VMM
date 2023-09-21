package server

// import (
// 	"net/http"

// 	"github.com/labstack/echo/v4"
// 	"github.com/pwdz/VMM/code/backend/models"
// )

// func AdminMiddleware(next echo.HandlerFunc) echo.HandlerFunc {
//     return func(c echo.Context) error {
//         user := getCurrentUser(c) // Implement a function to get the current user (e.g., from JWT token)

//         if user.Role != "admin" {
//             return c.JSON(http.StatusForbidden, models.VMResponse{Error: "Access denied"})
//         }

//         return next(c)
//     }
// }
