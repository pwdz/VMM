package vboxWrapper

import (
	"testing"
	"time"
)
func ensureVMOff(vmName string, t *testing.T) {
    // Check if the VM is running, and if so, turn it off
    status, err := GetVMStatus(vmName)
    if err != nil {
        t.Fatalf("Failed to get VM status: %v", err)
    }

    if status == "running" {
        if err := PowerOffVM(vmName); err != nil {
            t.Fatalf("Failed to power off VM: %v", err)
        }
    }
}

func ensureVMOn(vmName string, t *testing.T) {
    // Check if the VM is not running, and if so, turn it on
    status, err := GetVMStatus(vmName)
    if err != nil {
        t.Fatalf("Failed to get VM status: %v", err)
    }

    if status != "running" {
        if err := PowerOnVM(vmName); err != nil {
            t.Fatalf("Failed to power on VM: %v", err)
        }
    }
}
// Helper function to check if a VM exists and create it if needed
func ensureTestVMExists(t *testing.T) {
    vmName := "TestVM"

    // Check if the VM exists
    vmStatus, err := GetVMStatus(vmName)
    if err != nil || vmStatus == "" {
        // VM doesn't exist, create it
        err := CreateVM(vmName, "ubuntu", "2048", "2")
        if err != nil {
            t.Fatalf("Failed to create VM: %v", err)
        }
    }
}


func TestCreateVM(t *testing.T) {
    // Define test input values
    vmName := "TestVMCreate" + time.DateTime
    osType := "ubuntu"
    amountInMB := "2048"
    numCPUs := "2"

    // Call the CreateVM function
    err := CreateVM(vmName, osType, amountInMB, numCPUs)

    // Check if the function executed without errors
    if err != nil {
        t.Errorf("CreateVM failed with error: %v", err)
    }
}

func TestDeleteVM(t *testing.T) {
    // Ensure that the test VM exists or create it
    ensureTestVMExists(t)
	ensureVMOff("TestVM", t)

    // Define test input value
    vmName := "TestVM"

    // Call the DeleteVM function
    err := DeleteVM(vmName)

    // Check if the function executed without errors
    if err != nil {
        t.Errorf("DeleteVM failed with error: %v", err)
    }
}

func TestChangeVMSettings(t *testing.T) {
    // Ensure that the test VM exists or create it
    ensureTestVMExists(t)
	ensureVMOff("TestVM", t)

    // Define test input values
    vmName := "TestVM"
    settingName := "memory"
    settingValue := "4096"

    // Call the ChangeVMSettings function
    err := ChangeVMSettings(vmName, settingName, settingValue)

    // Check if the function executed without errors
    if err != nil {
        t.Errorf("ChangeVMSettings failed with error: %v", err)
    }
}

func TestPowerOffVM(t *testing.T) {
    // Ensure that the test VM exists or create it
    ensureTestVMExists(t)
	ensureVMOn("TestVM", t)

    // Define test input value
    vmName := "TestVM"

    // Call the PowerOffVM function
    err := PowerOffVM(vmName)

    // Check if the function executed without errors
    if err != nil {
        t.Errorf("PowerOffVM failed with error: %v", err)
    }
}

func TestPowerOnVM(t *testing.T) {
    // Ensure that the test VM exists or create it
    ensureTestVMExists(t)
	ensureVMOff("TestVM", t)

    // Define test input value
    vmName := "TestVM"

    // Call the PowerOnVM function
    err := PowerOnVM(vmName)

    // Check if the function executed without errors
    if err != nil {
        t.Errorf("PowerOnVM failed with error: %v", err)
    }
}

func TestGetVMStatus(t *testing.T) {
    // Ensure that the test VM exists or create it
    ensureTestVMExists(t)

    // Define test input value
    vmName := "TestVM"

    // Call the GetVMStatus function
    status, err := GetVMStatus(vmName)

    // Check if the function executed without errors
    if err != nil {
        t.Errorf("GetVMStatus failed with error: %v", err)
    }

    // Check if the status is not empty
    if status == "" {
        t.Error("GetVMStatus returned an empty status")
    }
}

func TestGetAvailableVMs(t *testing.T) {
    // Call the GetAvailableVMs function
    err := GetAvailableVMs()

    // Check if the function executed without errors
    if err != nil {
        t.Errorf("GetAvailableVMs failed with error: %v", err)
    }
}

func TestUploadFileToVM(t *testing.T) {
    // Ensure that the test VM exists or create it
    ensureTestVMExists(t)
	ensureVMOn("TestVM",t)

    // Define test input values
    vmName := "TestVM"
    guestFilePath := "/home/user/testfile.txt"
    fileContent := []byte("This is a test file content.")

    // Call the UploadFileToVM function
    err := UploadFileToVM(vmName, guestFilePath, fileContent)

    // Check if the function executed without errors
    if err != nil {
        t.Errorf("UploadFileToVM failed with error: %v", err)
    }
}

func TestExecuteCommandOnVM(t *testing.T) {
    // Ensure that the test VM exists or create it
    ensureTestVMExists(t)
	ensureVMOn("TestVM", t)
    // Define test input values
    vmName := "TestVM"
    userCommand := "cd ~/ && ls -l"

    // Call the ExecuteCommandOnVM function
    output, err := ExecuteCommandOnVM(vmName, userCommand)

    // Check if the function executed without errors
    if err != nil {
        t.Errorf("ExecuteCommandOnVM failed with error: %v", err)
    }

    // Check if the output is not empty
    if output == "" {
        t.Error("ExecuteCommandOnVM returned an empty output")
    }
}