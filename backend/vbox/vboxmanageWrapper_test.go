package vboxWrapper

import (
	"testing"
)

func TestCreateVM(t *testing.T) {
	// Replace these values with actual test data
	vmName := "TestVM"
	osType := "Linux"
	amountInMB := "2048"
	numCPUs := "2"
	isoPath := "/home/user/Desktop/BachelorProject/ISO/ubuntu-22.04-desktop-amd64.iso"

	err := CreateVM(vmName, osType, amountInMB, numCPUs, isoPath)
	if err != nil {
		t.Errorf("CreateVM failed: %v", err)
	}
}

func TestDeleteVM(t *testing.T) {
	// Replace this value with the actual VM name to delete
	vmName := "TestVM"

	err := DeleteVM(vmName)
	if err != nil {
		t.Errorf("DeleteVM failed: %v", err)
	}
}

func TestCloneVM(t *testing.T) {
	// Replace these values with actual test data
	sourceVMName := "SourceVM"
	newVMName := "CloneVM"

	err := CloneVM(sourceVMName, newVMName)
	if err != nil {
		t.Errorf("CloneVM failed: %v", err)
	}
}

func TestChangeVMSettings(t *testing.T) {
	// Replace these values with actual test data
	vmName := "Initial-Test"
	settingName := "memory"
	settingValue := "1024"

	err := ChangeVMSettings(vmName, settingName, settingValue)
	if err != nil {
		t.Errorf("ChangeVMSettings failed: %v", err)
	}
}

func TestPowerOffVM(t *testing.T) {
	// Replace this value with the actual VM name to power off
	vmName := "TestVM"

	err := PowerOffVM(vmName)
	if err != nil {
		t.Errorf("PowerOffVM failed: %v", err)
	}
}

func TestPowerOnVM(t *testing.T) {
	// Replace this value with the actual VM name to power on
	vmName := "TestVM"

	err := PowerOnVM(vmName)
	if err != nil {
		t.Errorf("PowerOnVM failed: %v", err)
	}
}

func TestGetVMStatus(t *testing.T) {
	// Replace this value with the actual VM name to get its status
	vmName := "TestVM"

	err := GetVMStatus(vmName)
	if err != nil {
		t.Errorf("GetVMStatus failed: %v", err)
	}
}

func TestGetAvailableVMs(t *testing.T) {
	err := GetAvailableVMs()
	if err != nil {
		t.Errorf("GetAvailableVMs failed: %v", err)
	}
}

func TestUploadFileToVM(t *testing.T) {
	// Replace these values with actual test data
	vmName := "TestVM"
	guestFilePath := "/path/on/guest"
	fileContent := []byte("Test file content")

	err := UploadFileToVM(vmName, guestFilePath, fileContent)
	if err != nil {
		t.Errorf("UploadFileToVM failed: %v", err)
	}
}

func TestTransferFileBetweenVMs(t *testing.T) {
	// Replace these values with actual test data
	sourceVMName := "SourceVM"
	sourceFile := "/path/to/source/file"
	destVMName := "DestVM"
	destFile := "/path/to/destination/file"

	err := TransferFileBetweenVMs(sourceVMName, sourceFile, destVMName, destFile)
	if err != nil {
		t.Errorf("TransferFileBetweenVMs failed: %v", err)
	}
}

func TestExecuteCommandOnVM(t *testing.T) {
	// Replace these values with actual test data
	vmName := "TestVM"
	command := "mkdir test"

	err := ExecuteCommandOnVM(vmName, command)
	if err != nil {
		t.Errorf("ExecuteCommandOnVM failed: %v", err)
	}
}
