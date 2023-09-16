package server

import (
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/pwdz/VMM/backend/configs"
	"github.com/pwdz/VMM/backend/db"
	jwtMiddleware "github.com/pwdz/VMM/backend/jwt"
)

var e* echo.Echo
var Cfg configs.ServerConfig
var DB *db.Database

func InitCfg(){
	err := cleanenv.ReadEnv(&Cfg)
	log.Printf("%+v", Cfg)
	if err != nil{
		e.Logger.Fatal("Unable to load configs")
	}

	fmt.Println("adasdad")
	DB, err = db.NewDatabase(configs.GetDBConfig())
	if err != nil{
		e.Logger.Fatal("Unable to load configs")
	}

}
func InitServer(){
	e = echo.New()

   // Middleware to set content-type to JSON for all routes
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			c.Response().Header().Set("Content-Type", "application/json; charset=utf-8")
			return next(c)
		}
	})

	// Middleware to add JWT token validation to protected routes.
	e.Use(jwtMiddleware.JwtMiddleware)

	// Define routes for each endpoint
	e.POST("/signup", SignupHandler) 
	e.POST("/login", LoginHandler)   
	e.POST("/create-vm", CreateVMHandler)
	e.POST("/clone-vm", CloneVMHandler)
	e.POST("/change-vm-settings", ChangeVMSettingsHandler)
	e.POST("/power-off-vm", PowerOffVMHandler)
	e.POST("/power-on-vm", PowerOnVMHandler)
	e.POST("/get-vm-status", GetVMStatusHandler)
	e.GET("/get-available-vms", GetAvailableVMsHandler)
	e.POST("/upload-file-to-vm", UploadFileToVMHandler)
	e.POST("/transfer-file-between-vms", TransferFileBetweenVMsHandler)
	e.POST("/execute-command-on-vm", ExecuteCommandOnVMHandler)
}