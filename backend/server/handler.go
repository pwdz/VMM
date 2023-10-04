package server

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"strconv"

	"github.com/dgrijalva/jwt-go"
	"github.com/labstack/echo/v4"
	"github.com/tealeg/xlsx"

	jwtMiddleware "github.com/pwdz/VMM/code/backend/jwt"
	"github.com/pwdz/VMM/code/backend/models"
	vbox "github.com/pwdz/VMM/code/backend/vbox"
)

// Handler for CreateVM
func CreateVMHandler(c echo.Context) error {
	    // Parse the request
		var req models.VMRequest
		if err := c.Bind(&req); err != nil {
			return c.JSON(http.StatusBadRequest, models.VMResponse{Error: "Invalid request"})
		}
	
		// Call CreateVM function to create the VM
		vmName := req.VMName
		osType := req.OSType
		ramInMB := req.RamInMB
		numCPUs := req.NumCPUs
		isoPath := req.ISOPath
		
		// Convert RAM and CPU to integers
		ramInMBInt, err := strconv.Atoi(ramInMB)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.VMResponse{Error: "Invalid RAM format"})
		}

		numCPUsInt, err := strconv.Atoi(numCPUs)
		if err != nil {
			return c.JSON(http.StatusBadRequest, models.VMResponse{Error: "Invalid CPU format"})
		}

		if err := vbox.CreateVM(vmName, osType, ramInMB, numCPUs, isoPath); err != nil {
			return c.JSON(http.StatusInternalServerError, models.VMResponse{Error: "Failed to create VM"})
		}
	
		// Retrieve the user_id from the context
		userID := c.Get("user_id").(uint)
	
		// Now, add a record to the database
		vm := &models.VM{
			UserID: userID,
			Name:   vmName,
			OSType: osType,
			RAM:    ramInMBInt,
			CPU:    numCPUsInt,
			Status: "off",
			IsDeleted: false,
		}
	
		if err := DB.CreateVM(vm); err != nil {
			return c.JSON(http.StatusInternalServerError, models.VMResponse{Error: "Failed to create VM record in the database"})
		}
	
		return c.JSON(http.StatusOK, models.VMResponse{Message: "VM created successfully"})
}

// Handler for DeleteVM
func DeleteVMHandler(c echo.Context) error {
	// Extract the VM ID from the request JSON
	var req models.VMRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Call the DeleteVM function to delete the virtual machine in vboxWrapper
	// if err := vbox.DeleteVM(req.VMName); err != nil {
	// 	return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to delete VM"})
	// }
	
	// Call the DeleteVM function to set IsDeleted to true in the database
	if err := DB.DeleteVM(req.VMID); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to mark VM as deleted in the database"})
	}

	return c.JSON(http.StatusOK, map[string]string{"message": "VM deleted successfully"})
}

func CloneVMHandler(c echo.Context) error {
	// Extract the source VM name and new VM name from the request JSON
	var req models.VMRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	// Assuming you have a naming convention for the new VM based on the user's input
	sourceVMName := req.VMName
	newVMName := req.NewVMName

	// Query the database to find the source VM's details
	sourceVM := DB.FindVMByName(sourceVMName)
	if sourceVM == nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Source VM not found"})
	}

	// Clone the VM using vboxWrapper
	if err := vbox.CloneVM(sourceVMName, newVMName); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to clone VM"})
	}

	// Retrieve the user_id from the context
	userID := c.Get("user_id").(uint)
	
	// Create a new VM record in the database based on the source VM's details
	newVM := &models.VM{
		UserID:    userID,        // Use the appropriate user ID
		Name:      newVMName,     // Set the new VM name
		OSType:    sourceVM.OSType,
		RAM:       sourceVM.RAM,
		CPU:       sourceVM.CPU,
		Status:    "off",         // Assuming the new VM is initially off
		IsDeleted: false,         // Default to false
	}

	if err := DB.CreateVM(newVM); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to create VM in the database"})
	}

	return c.JSON(http.StatusOK, newVM)
}
func ChangeVMSettingsHandler(c echo.Context) error {
	// Extract the VM ID, setting name, and setting value from the request JSON
	var req models.VMRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}


	// Check if the VM is turned off
	vm := DB.FindVMByID(req.VMID)
	if vm == nil {
		fmt.Println(")))))))))))))))))))", vm, req.VMID)
		return c.JSON(http.StatusNotFound, map[string]string{"error": "VM not found"})
	}

	if vm.Status == "on" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "VM is currently running. Turn it off to change settings."})
	}

	cpuReq, _ := strconv.Atoi(req.NumCPUs)
	// Validate and apply the CPU setting
	if req.NumCPUs != "" && cpuReq != vm.CPU{
		cpu, err := strconv.Atoi(req.NumCPUs)
		if err != nil || cpu < 1 || cpu > 8 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid CPU value. CPU must be an integer between 1 and 8."})
		}

		// Apply the CPU setting using vboxWrapper
		if err := vbox.ChangeVMSettings(vm.Name, "cpus", req.NumCPUs); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to change CPU settings"})
		}

		// Update the CPU setting in the database
		if err := DB.UpdateVMSetting(vm.ID, "cpu", cpuReq); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update CPU setting in the database"})
		}
	}

	ramReq, _ := strconv.Atoi(req.RamInMB)
	// Validate and apply the RAM setting
	if req.RamInMB != "" && ramReq != vm.RAM {
		ram, err := strconv.Atoi(req.RamInMB)
		if err != nil || ram < 1 || ram > 4096 {
			return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid RAM value. RAM must be an integer between 1 and 4096."})
		}

		// Apply the RAM setting using vboxWrapper
		if err := vbox.ChangeVMSettings(vm.Name, "memory", req.RamInMB); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to change RAM settings"})
		}

		// Update the RAM setting in the database
		if err := DB.UpdateVMSetting(vm.ID, "ram", ramReq); err != nil {
			return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update RAM setting in the database"})
		}
	}

	return c.JSON(http.StatusOK, models.VMResponse {Message:  fmt.Sprintf("VM %d settings updated successfully", req.VMID)})
}

func PowerOffVMHandler(c echo.Context) error {
	// Extract the VM ID and VM Name from the request JSON
	var req models.VMRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}

	var vm *models.VM

	// Check if VM ID is provided, if not, try to find the VM by name
	if req.VMID > 0 {
		vm = DB.FindVMByID(req.VMID)
		if vm == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "VM not found"})
		}
	} else if req.VMName != "" {
		vm = DB.FindVMByName(req.VMName)
		if vm == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "VM not found"})
		}
	} else {
		fmt.Println("NOT GGGGGGGG", req)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "VM ID or VM Name is required"})
	}

	// Get the current status of the VM from the database
	// Assuming that the status in the database is either "on" or "off"
	currentStatus := vm.Status

	// Check if the VM is already powered off
	if currentStatus == "off" {
		fmt.Println("NOT GGGGGGGG2", req)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "VM is already powered off"})
	}

	// Perform the power off operation using vboxWrapper
	if err := vbox.PowerOffVM(vm.Name); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to power off VM"})
	}

	// Update the VM status in the database
	if err := DB.UpdateVMStatus(vm.ID, "off"); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update VM status in the database"})
	}

	response := models.VMResponse{
		Message: fmt.Sprintf("VM %d powered off successfully", vm.ID),
	}

	return c.JSON(http.StatusOK, response)
}

func PowerOnVMHandler(c echo.Context) error {
	// Extract the VM ID and VM Name from the request JSON

	var req models.VMRequest
	if err := c.Bind(&req); err != nil {
		fmt.Println(":|", err)
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	}


	// var req models.VMRequest
	// if err := c.Bind(&req); err != nil {
	// 	fmt.Println(":|", err)
	// 	return c.JSON(http.StatusBadRequest, map[string]string{"error": "Invalid request"})
	// }

	var vm *models.VM

	// Check if VM ID is provided, if not, try to find the VM by name
	if req.VMID > 0 {
		vm = DB.FindVMByID(req.VMID)
		if vm == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "VM not found"})
		}
	} else if req.VMName != "" {
		vm = DB.FindVMByName(req.VMName)
		if vm == nil {
			return c.JSON(http.StatusNotFound, map[string]string{"error": "VM not found"})
		}
	} else {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "VM ID or VM Name is required"})
	}

	// Get the current status of the VM from the database
	// Assuming that the status in the database is either "on" or "off"
	currentStatus := vm.Status

	// Check if the VM is already powered on
	if currentStatus == "on" {
		return c.JSON(http.StatusBadRequest, map[string]string{"error": "VM is already powered on"})
	}

	// Perform the power on operation using vboxWrapper
	if err := vbox.PowerOnVM(vm.Name); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to power on VM"})
	}

	// Update the VM status in the database
	if err := DB.UpdateVMStatus(vm.ID, "on"); err != nil {
		return c.JSON(http.StatusInternalServerError, map[string]string{"error": "Failed to update VM status in the database"})
	}

	response := models.VMResponse{
		Message: fmt.Sprintf("VM %d powered on successfully", vm.ID),
	}

	return c.JSON(http.StatusOK, response)
}

// GetVMStatusHandler retrieves the status of a VM based on the provided VM ID.
func GetVMStatusHandler(c echo.Context) error {
	// Parse the request using the VMRequest struct from models
	var req models.VMRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Get the vm_id from the request
	vmID := req.VMID

	// Query the database for the VM's status
	vmStatus, err := DB.GetVMStatus(vmID)
	if err != nil {
		// Handle database errors
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}

	// Return the VM's status as JSON
	return c.JSON(http.StatusOK, map[string]string{
		"status": vmStatus,
	})
}

func GetVMsHandler(c echo.Context) error {
	fmt.Println(":))))))))))))")
	// Retrieve the user_id from the context
	userID := c.Get("user_id").(uint)

	// Query the database for VMs associated with the user_id
	vms, err := DB.GetVMsByUserID(userID)
	if err != nil {
		// Handle database errors
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Internal Server Error",
		})
	}
	
	// Return the VMs as JSON
	return c.JSON(http.StatusOK, vms)
}

func UploadFileToVMHandler(c echo.Context) error {
	// Parse the request using the VMRequest struct from models
	var req models.VMRequest
	if err := c.Bind(&req); err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Invalid request body",
		})
	}

	// Check if the VM with the given VM ID belongs to the user
	vm := DB.FindVMByID(req.VMID)
	if vm == nil {
		return c.JSON(http.StatusNotFound, map[string]string{"error": "VM not found"})
	}

	userID := c.Get("user_id").(uint)
	// Ensure the VM belongs to the user
	if vm.UserID != userID {
		return c.JSON(http.StatusForbidden, map[string]string{
			"error": "Access Denied",
		})
	}

	// Decode the base64-encoded file content from the request
	fileContent, err := base64.StdEncoding.DecodeString(req.FileContent)
	if err != nil {
		return c.JSON(http.StatusBadRequest, map[string]string{
			"error": "Failed to decode file content",
		})
	}

	fmt.Println("chill bro")
	// Call the function to upload the file content to the VM
	if err := vbox.UploadFileToVM(vm.Name, req.GuestFilePath, fileContent); err != nil {
		// Handle the error from the vbox.UploadFileContentToVM function
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to upload file content to VM",
		})
	}

	// File upload was successful
	return c.JSON(http.StatusOK, map[string]string{
		"message": "File uploaded successfully",
	})
}


func TransferFileBetweenVMsHandler(c echo.Context) error {
	type TransferFileRequest struct {
		SourceVMID      uint   `json:"source_vm_id"`
		SourceFilePath  string `json:"source_file_path"`
		DestinationVMID uint   `json:"destination_vm_id"`
		DestinationPath string `json:"destination_path"`
	}
	// Parse the request body
    var req TransferFileRequest
    if err := c.Bind(&req); err != nil {
        return c.JSON(http.StatusBadRequest, map[string]string{
            "error": "Invalid request body",
        })
    }

    // Check if the source and destination VMs belong to the same user
    userID := c.Get("user_id").(uint) // Assuming you store the user ID in the context

    // Check if the source VM belongs to the user
    sourceVM := DB.FindVMByID(req.SourceVMID)
    if sourceVM == nil || sourceVM.UserID != userID {
        return c.JSON(http.StatusForbidden, map[string]string{
            "error": "Source VM does not exist or does not belong to the user",
        })
    }

    // Check if the destination VM belongs to the user
    destinationVM := DB.FindVMByID(req.DestinationVMID)
    if destinationVM == nil || destinationVM.UserID != userID {
        return c.JSON(http.StatusForbidden, map[string]string{
            "error": "Destination VM does not exist or does not belong to the user",
        })
    }

    // Perform the file transfer between VMs
    err := vbox.TransferFileBetweenVMs(
        sourceVM.Name,
        req.SourceFilePath,
        destinationVM.Name,
        req.DestinationPath,
    )

    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to transfer the file between VMs",
        })
    }

    // Return a success response
    return c.JSON(http.StatusOK, map[string]string{
        "message": "File transfer successful",
    })
}

// TODO
func ExecuteCommandOnVMHandler(c echo.Context) error {
	req := new(struct {
		VMName    string `json:"vm_name"`
		Command string `json:"command"`
	})
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, models.VMResponse{Error: "Invalid request"})
	}

	// Implement the logic to execute a command on a VM using req.VMName, req.PathToExe, and req.Arguments
	if err := vbox.ExecuteCommandOnVM(req.VMName, req.Command); err != nil {
		// Handle the error from the vbox.UploadFileContentToVM function
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to upload file content to VM",
		})
	}




	return c.JSON(http.StatusOK, models.VMResponse{Message: "Command executed on VM successfully"})
}

// Handler for user registration (Sign-up)
func SignupHandler(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, models.VMResponse{Error: "Invalid request"})
	}

	// TODO
	if storedUser := DB.FindUserByUsername(user.Username); storedUser != nil {
		return c.JSON(http.StatusConflict, models.VMResponse{Error: "User already exists"})
	}

	DB.CreateUser(user)

	// Generate a JWT token for the new user
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["role"] = models.UserRole
	tokenString, _ := token.SignedString(jwtMiddleware.JwtSecret) // Use JwtSecret from jwt_middleware.go

	return c.JSON(http.StatusOK, models.VMResponse{Message: "User registered successfully", Error: "", Data: tokenString})
}

// Handler for user login
func LoginHandler(c echo.Context) error {
	user := new(models.User)
	if err := c.Bind(user); err != nil {
		return c.JSON(http.StatusBadRequest, models.VMResponse{Error: "Invalid request"})
	}

	storedUser := DB.FindUserByUsername(user.Username)
	if storedUser == nil || storedUser.Password != user.Password {
		return c.JSON(http.StatusUnauthorized, models.VMResponse{Error: "Invalid credentials"})
	}


	// Generate a JWT token for the authenticated user
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["username"] = user.Username
	claims["role"] = storedUser.Role
	tokenString, _ := token.SignedString(jwtMiddleware.JwtSecret) // Use JwtSecret from jwt_middleware.go
	fmt.Println(user.Password, user.Username)
	fmt.Println(claims)
	fmt.Println(tokenString)
	return c.JSON(http.StatusOK, models.VMResponse{Message: "Login successful", Error: "", Data: tokenString})
}

func GetRoleHandler(c echo.Context) error {
	fmt.Println("WTF")
	// Get the user's role from the context set by the JWT middleware
	userRole := c.Get("role").(string)

	// You can now use the userRole as needed, for example, return it in the response
	return c.JSON(http.StatusOK, map[string]string{
		"role": userRole,
	})
}

// Define the GetUsersHandler to retrieve non-admin users
func GetUsersHandler(c echo.Context) error {
	fmt.Println("ridi")
    // Assuming you have a method in your database package to retrieve non-admin users
    nonAdminUsers, err := DB.GetNonAdminUsersWithVMCounts()
    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to retrieve non-admin users",
        })
    }
	fmt.Println(nonAdminUsers)
    // Return the list of non-admin users in JSON format
    return c.JSON(http.StatusOK, nonAdminUsers)
}

func ExportUsersHandler(c echo.Context) error {
    // Retrieve the list of users
    users, err := DB.GetNonAdminUsersWithVMCounts()
    if err != nil {
        // Handle the error, e.g., return an error response
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to retrieve users",
        })
    }

    // Create a new Excel file
    file := xlsx.NewFile()
    sheet, err := file.AddSheet("Users")
    if err != nil {
        // Handle the error, e.g., return an error response
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to create Excel sheet",
        })
    }

    // Add headers to the Excel sheet
    headerRow := sheet.AddRow()
    headerRow.AddCell().SetString("ID")
    headerRow.AddCell().SetString("Username")
    headerRow.AddCell().SetString("Email")
    headerRow.AddCell().SetString("Active VMs")
    headerRow.AddCell().SetString("Inactive VMs")

    // Add user data to the Excel sheet
    for _, user := range users {
        userRow := sheet.AddRow()
        userRow.AddCell().SetInt(int(user.ID))
        userRow.AddCell().SetString(user.Username)
        userRow.AddCell().SetString(user.Email)
        userRow.AddCell().SetInt(user.ActiveVMCount)
        userRow.AddCell().SetInt(user.InactiveVMCount)
    }

    // Set the content type for the response
    contentType := "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet"
    c.Response().Header().Set("Content-Type", contentType)

    // Set the content disposition to trigger a download
    fileName := "users.xlsx"
    c.Response().Header().Set("Content-Disposition", "attachment; filename="+fileName)

    // Write the Excel file to the response body
    err = file.Write(c.Response().Writer)
    if err != nil {
        // Handle the error, e.g., return an error response
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to write Excel file",
        })
    }

    // Return a success response
    return c.NoContent(http.StatusOK)
}

// GetAllVMsHandler retrieves a list of all VMs in the database.
func GetAllVMsHandler(c echo.Context) error {
    // Query the database for all VMs
    VMWithUsers, err := DB.GetAllVMs()
    if err != nil {
        // Handle the error, e.g., return an error response
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to retrieve VMs",
        })
    }

    // Return the list of VMs as JSON
    return c.JSON(http.StatusOK, VMWithUsers)
}

// ExportAllVMsHandler exports the list of all VMs as an Excel file.
func ExportAllVMsHandler(c echo.Context) error {
	// Query the database for all VMs
	vms, err := DB.GetAllVMs()
	if err != nil {
		// Handle the error, e.g., return an error response
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to retrieve VMs",
		})
	}

	// Create a new Excel file
	file := xlsx.NewFile()
	sheet, err := file.AddSheet("VMs")
	if err != nil {
		// Handle the error, e.g., return an error response
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to create Excel sheet",
		})
	}

	// Define the header row
	headers := []string{"UserID", "Username", "ID", "Name", "OSType", "RAM", "CPU", "Status", "IsDeleted"}

	// Create a new row for headers
	headerRow := sheet.AddRow()
	for _, header := range headers {
		cell := headerRow.AddCell()
		cell.Value = header
	}

	// Write VM data to the Excel sheet
	for _, vm := range vms {
		dataRow := sheet.AddRow()
		dataRow.AddCell().SetInt(int(vm.UserID))
		dataRow.AddCell().Value = vm.Username
		dataRow.AddCell().SetInt(int(vm.ID))
		dataRow.AddCell().Value = vm.Name
		dataRow.AddCell().Value = vm.OSType
		dataRow.AddCell().SetInt(vm.RAM)
		dataRow.AddCell().SetInt(vm.CPU)
		dataRow.AddCell().Value = vm.Status
		dataRow.AddCell().SetBool(vm.IsDeleted)
	}

	// Set the response headers for Excel download
	c.Response().Header().Set("Content-Type", "application/vnd.openxmlformats-officedocument.spreadsheetml.sheet")
	c.Response().Header().Set("Content-Disposition", "attachment; filename=vm_export.xlsx")

	// Write the Excel file to the response
	err = file.Write(c.Response().Writer)
	if err != nil {
		// Handle the error, e.g., return an error response
		return c.JSON(http.StatusInternalServerError, map[string]string{
			"error": "Failed to export VMs to Excel",
		})
	}

	return nil
}

func GetProfileDataHandler(c echo.Context) error {
	userID := c.Get("user_id").(uint)
    userData, err := DB.GetUserDataWithVMCounts(userID)

    if err != nil {
        return c.JSON(http.StatusInternalServerError, map[string]string{
            "error": "Failed to retrieve user profile data",
        })
    }
    // Return the list of non-admin users in JSON format
    return c.JSON(http.StatusOK, userData)
}	