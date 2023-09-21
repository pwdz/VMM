package server

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pwdz/VMM/code/backend/configs"
	"github.com/pwdz/VMM/code/backend/db"
	jwtMiddleware "github.com/pwdz/VMM/code/backend/jwt"
)

var e *echo.Echo
var Cfg configs.ServerConfig
var DB *db.Database

func InitCfg() {
	err := cleanenv.ReadEnv(&Cfg)
	log.Printf("%+v", Cfg)
	if err != nil {
		e.Logger.Fatal("Unable to load configs")
	}

	fmt.Println("adasdad")
	DB, err = db.NewDatabase(configs.GetDBConfig())
	if err != nil {
		e.Logger.Fatal("Unable to load configs")
	}

}
func InitServer() {
	e = echo.New()

	// Middleware to set content-type to JSON for all routes
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
			return next(c)
		}
	})

	// CORS middleware configuration
	corsConfig := middleware.CORSConfig{
		// Define the allowed origins (replace with your frontend domain)
		AllowOrigins: []string{"*"},
		AllowMethods: []string{echo.GET, echo.PUT, echo.POST, echo.DELETE},
	}
	e.Use(middleware.CORSWithConfig(corsConfig))

	// Define routes for each endpoint
	e.POST("/signup", SignupHandler)
	e.POST("/login", LoginHandler)

	// Create a route group with the authorization middleware
	apiGroup := e.Group("/api")
	// Middleware to add JWT token validation to protected routes.
	apiGroup.Use(jwtMiddleware.JwtMiddleware) // Apply the authorization middleware to this group

	apiGroup.POST("/create-vm", CreateVMHandler)
	apiGroup.POST("/clone-vm", CloneVMHandler)
	apiGroup.POST("/change-vm-settings", ChangeVMSettingsHandler)
	apiGroup.POST("/power-off-vm", PowerOffVMHandler)
	apiGroup.POST("/power-on-vm", PowerOnVMHandler)
	apiGroup.POST("/get-vm-status", GetVMStatusHandler)
	apiGroup.GET("/get-available-vms", GetAvailableVMsHandler)
	apiGroup.POST("/upload-file-to-vm", UploadFileToVMHandler)
	apiGroup.POST("/transfer-file-between-vms", TransferFileBetweenVMsHandler)
	apiGroup.POST("/execute-command-on-vm", ExecuteCommandOnVMHandler)

	e.Start(fmt.Sprintf("%s:%s", Cfg.Host, Cfg.Port))
}
