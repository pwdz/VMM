package handler

import (
	"net/http"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"

	// "strings"
	// "errors"
	"github.com/pwdz/VMM/backend/internal/server/constants"
	"github.com/pwdz/VMM/backend/internal/server/middleware"
	"github.com/pwdz/VMM/backend/internal/server/user"

	// "os"
	// "io/ioutil"
	"encoding/json"
	// "fmt"
)

// Create a struct to read the username and password from the request body
type Credentials struct {
	Password string `json:"password"`
	Username string `json:"username"`
}

type VMRequest struct {
	VMName    string `json:"vmName"`
	OSType    string `json:"osType"`
	AmountInMB string `json:"amountInMB"`
	NumCPUs   string `json:"numCPUs"`
	VDIPath   string `json:"vdiPath"`
	ISOPath   string `json:"isoPath"`
}

type VMResponse struct {
    Message string `json:"message"`
    Error   string `json:"error,omitempty"`
    Data    string `json:"data,omitempty"`
}

// Handler for CreateVM
func CreateVMHandler(c echo.Context) error {
    req := new(VMRequest)
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to create a VM here using req

    return c.JSON(http.StatusOK, VMResponse{Message: "VM created successfully"})
}

// Handler for DeleteVM
func DeleteVMHandler(c echo.Context) error {
    req := new(struct {
        VMName string `json:"vmName"`
    })
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to delete a VM here using req.VMName

    return c.JSON(http.StatusOK, VMResponse{Message: "VM deleted successfully"})
}

func CloneVMHandler(c echo.Context) error {
    req := new(struct {
        SourceVMName string `json:"sourceVMName"`
        NewVMName    string `json:"newVMName"`
    })
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to clone a VM here using req.SourceVMName and req.NewVMName

    return c.JSON(http.StatusOK, VMResponse{Message: "VM cloned successfully"})
}
func ChangeVMSettingsHandler(c echo.Context) error {
    req := new(struct {
        VMName       string `json:"vmName"`
        SettingName  string `json:"settingName"`
        SettingValue string `json:"settingValue"`
    })
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to change VM settings here using req.VMName, req.SettingName, and req.SettingValue

    return c.JSON(http.StatusOK, VMResponse{Message: "VM settings changed successfully"})
}
func PowerOffVMHandler(c echo.Context) error {
    req := new(struct {
        VMName string `json:"vmName"`
    })
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to power off a VM here using req.VMName

    return c.JSON(http.StatusOK, VMResponse{Message: "VM powered off successfully"})
}

func PowerOnVMHandler(c echo.Context) error {
    req := new(struct {
        VMName string `json:"vmName"`
    })
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to power on a VM here using req.VMName

    return c.JSON(http.StatusOK, VMResponse{Message: "VM powered on successfully"})
}
func GetVMStatusHandler(c echo.Context) error {
    req := new(struct {
        VMName string `json:"vmName"`
    })
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to get VM status here using req.VMName

    return c.JSON(http.StatusOK, VMResponse{Message: "VM status retrieved successfully"})
}
func GetAvailableVMsHandler(c echo.Context) error {
    // Implement the logic to get a list of available VMs

    return c.JSON(http.StatusOK, VMResponse{Message: "List of available VMs retrieved successfully"})
}
func UploadFileToVMHandler(c echo.Context) error {
    req := new(struct {
        VMName    string `json:"vmName"`
        LocalFile string `json:"localFile"`
        GuestFile string `json:"guestFile"`
    })
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to upload a file to a VM using req.VMName, req.LocalFile, and req.GuestFile

    return c.JSON(http.StatusOK, VMResponse{Message: "File uploaded to VM successfully"})
}
func TransferFileBetweenVMsHandler(c echo.Context) error {
    req := new(struct {
        SourceVMName string `json:"sourceVMName"`
        SourceFile   string `json:"sourceFile"`
        DestVMName   string `json:"destVMName"`
        DestFile     string `json:"destFile"`
    })
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to transfer a file between VMs using req.SourceVMName, req.SourceFile, req.DestVMName, and req.DestFile

    return c.JSON(http.StatusOK, VMResponse{Message: "File transferred between VMs successfully"})
}
func ExecuteCommandOnVMHandler(c echo.Context) error {
    req := new(struct {
        VMName    string `json:"vmName"`
        PathToExe string `json:"pathToExe"`
        Arguments string `json:"arguments"`
    })
    if err := c.Bind(req); err != nil {
        return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
    }

    // Implement the logic to execute a command on a VM using req.VMName, req.PathToExe, and req.Arguments

    return c.JSON(http.StatusOK, VMResponse{Message: "Command executed on VM successfully"})
}


// Handler for user registration (Sign-up)
func SignupHandler(c echo.Context) error {
	user := new(user.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
	}

	// TODO
	/*
	// Check if the user already exists
	if _, exists := usersDB[user.Username]; exists {
		return c.JSON(http.StatusConflict, VMResponse{Error: "User already exists"})
	}

	// Store the user in the database (you should hash the password in a real application)
	usersDB[user.Username] = *user
*/


	// Generate a JWT token for the new user
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	tokenString, _ := token.SignedString(middleware.JwtSecret) // Use JwtSecret from jwt_middleware.go

	return c.JSON(http.StatusOK, VMResponse{Message: "User registered successfully", Error: "", Data: tokenString})
}

// Handler for user login
func LoginHandler(c echo.Context) error {
	user := new(user.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, VMResponse{Error: "Invalid request"})
	}

	// TODO
	/*
	
	// Check if the user exists and the password is correct (you should hash the password in a real application)
	storedUser, exists := usersDB[user.Username]
	if !exists || storedUser.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, VMResponse{Error: "Invalid credentials"})
	}
	*/

	// Generate a JWT token for the authenticated user
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	tokenString, _ := token.SignedString(middleware.JwtSecret) // Use JwtSecret from jwt_middleware.go

	return c.JSON(http.StatusOK, VMResponse{Message: "Login successful", Error: "", Data: tokenString})
}
































// Create the Signin handler
func Login(c echo.Context) error{
	var creds Credentials
	// Get the JSON body and decode into credentials
	err := json.NewDecoder(c.Request().Body).Decode(&creds)
	if err != nil {
		// If the structure of the body is wrong, return an HTTP error
		return c.String(http.StatusBadRequest, "")
	}

	// Get the expected password from our in memory map
	ok := false
	var expectedPassword, role string
	for _, user := range constants.Users{ 
		if user.Username == creds.Username{
			expectedPassword = user.Password
			ok = true
			role = user.Role
		}
	}

	// If a password exists for the given user
	// AND, if it is the same as the password we received, the we can move ahead
	// if NOT, then we return an "Unauthorized" status
	if !ok || expectedPassword != creds.Password {
		return c.String(http.StatusUnauthorized, "")
	}

	// Declare the expiration time of the token
	// here, we have kept it as 5 minutes
	// expirationTime := time.Now().Add(5 * time.Minute)
	// Create the JWT claims, which includes the username and expiry time
	claims := &user.User{
		Username: creds.Username,
		Role: role,
	}

	// Declare the token with the algorithm used for signing, and the claims
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	// Create the JWT string
	tokenString, err := token.SignedString(constants.JwtKey)
	if err != nil {
		// If there is an error in creating the JWT return an internal server error
		return c.String(http.StatusInternalServerError, "")
	}

	return c.String(http.StatusOK, tokenString)
}
// func EndPointHandler(c echo.Context) error{
	
// 	headerContentType := c.Request().Header.Get("Content-Type")
// 	var cmd command

// 	if headerContentType == "application/json" {		
// 		var unmarshalErr *json.UnmarshalTypeError

// 		decoder := json.NewDecoder(c.Request().Body)
// 		decoder.DisallowUnknownFields()

// 		err := decoder.Decode(&cmd)
// 		if err != nil {
// 			if errors.As(err, &unmarshalErr) {
// 				return c.String(http.StatusBadRequest, "Bad Request. Wrong Type provided for field "+unmarshalErr.Field)
// 			} else {
// 				return c.String(http.StatusBadRequest, "Bad Request "+err.Error())
// 			}
// 		}

// 	}else if strings.Contains(headerContentType, "multipart/form-data"){
// 		c.Request().ParseMultipartForm(10 << 20)
// 		file, handler, err := c.Request().FormFile("file")
// 		if err != nil{
// 			return err
// 		}
// 		defer file.Close()
			
// 		emptyFile, err := os.Create(handler.Filename)
// 		if err != nil {
// 			return err
// 		}
// 		fileBytes, err := ioutil.ReadAll(file)
// 		if err != nil {
// 			return err
// 		}
// 		emptyFile.Write(fileBytes)
// 		emptyFile.Close()

// 		vmName := c.Request().FormValue("vmName")
// 		dstPath := c.Request().FormValue("destPath") 

// 		cmd = command{
// 			Type: constants.CMDUpload,
// 			DestVM: vmName,
// 			DestPath: dstPath,
// 			OriginPath: handler.Filename,
// 		}
// 	}
// 	role := fmt.Sprintf("%v", c.Get("role"))
// 	jsonResponse := cmd.handleCommand(role)

// 	return c.JSONBlob(http.StatusOK, jsonResponse)
// }
