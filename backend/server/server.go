package server

import (
	"fmt"
	"log"
	"strings"

	"github.com/ilyakaznacheev/cleanenv"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/pwdz/VMM/code/backend/configs"
	"github.com/pwdz/VMM/code/backend/db"
	jwtMiddleware "github.com/pwdz/VMM/code/backend/jwt"
	"github.com/pwdz/VMM/code/backend/models"
	"github.com/pwdz/VMM/code/backend/pricing"
	vboxWrapper "github.com/pwdz/VMM/code/backend/vbox"
	"github.com/swaggo/echo-swagger"
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

	if vmsWithUsers, err := DB.GetAllVMs(); err == nil{
		for _, vmWithUser := range vmsWithUsers{
			if currStatus, err2 := vboxWrapper.GetVMStatus(vmWithUser.VM.Name); err2 == nil{
				fmt.Println(currStatus)
				if strings.Contains(currStatus, "off"){
					currStatus = "off"
				}else if strings.Contains(currStatus, "running"){
					currStatus = "on"
				}
				DB.UpdateVMStatus(vmWithUser.VM.ID, currStatus)
			}
		}
	}
	priceConfigs, err := DB.GetPriceConfigs()
    if err != nil {
        return 
    }

	var c,r,h int
    // Iterate through the priceConfigs and update the variables
    for _, config := range priceConfigs {
        switch config.Type {
        case "cpu":
            c = config.CostPerUnit
        case "ram":
            r = config.CostPerUnit
        case "hdd":
            h = config.CostPerUnit
        // Add more cases if you have other types
        }
    }
	pricing.UpdatePriceConfig(c, r, h)
}
func InitServer() {
	e = echo.New()

	// Serve Swagger UI
	e.GET("/swagger/*", echoSwagger.WrapHandler)

	// Middleware to set content-type to JSON for all routes
	e.Use(CheckJsonFormatMiddlware)
	e.Use(RequestDumper)

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
	
	// api Group
	apiGroup := e.Group("/api")
	apiGroup.Use(jwtMiddleware.JwtMiddleware)
	apiGroup.GET("/get-role", GetRoleHandler)

	// Admin Group
	adminGroup := e.Group("/admin")
	adminGroup.Use(jwtMiddleware.JwtMiddleware)
	adminGroup.Use(CheckUserRoleMiddleware(models.AdminRole))
	adminGroup.GET("/users", GetUsersHandler)
	adminGroup.GET("/export-users", ExportUsersHandler)
	adminGroup.GET("/vms", GetAllVMsHandler)
	adminGroup.GET("/export-vms", ExportAllVMsHandler)
	adminGroup.GET("/pricing", GetPriceSettingHandler)
	adminGroup.POST("/pricing-update", ChangePriceSettingHandler)

	// Create a route group with the authorization middleware
	userGroup := e.Group("/user")
	userGroup.Use(jwtMiddleware.JwtMiddleware) // Apply the authorization middleware to this group
	userGroup.Use(ExtractUserIDMiddleware)
	userGroup.Use(CheckUserRoleMiddleware(models.UserRole))

	userGroup.GET("/profile", GetProfileDataHandler)
	userGroup.POST("/create-vm", CreateVMHandler)
	userGroup.POST("/delete-vm", DeleteVMHandler)
	userGroup.POST("/clone-vm", CloneVMHandler)
	userGroup.POST("/change-vm-settings", ChangeVMSettingsHandler)
	userGroup.POST("/power-off-vm", PowerOffVMHandler)
	userGroup.POST("/power-on-vm", PowerOnVMHandler)
	userGroup.POST("/get-vm-status", GetVMStatusHandler)
	userGroup.GET("/get-vms", GetVMsHandler)
	userGroup.POST("/upload-file-to-vm", UploadFileToVMHandler)
	userGroup.POST("/transfer-file-between-vms", TransferFileBetweenVMsHandler)
	userGroup.POST("/execute-command-on-vm", ExecuteCommandOnVMHandler)

	e.Start(fmt.Sprintf("%s:%s", Cfg.Host, Cfg.Port))
}
