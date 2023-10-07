package vboxWrapper

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
	"regexp"
	"strings"
	"time"
)

// Define commands as a static map
var commands = map[string]string{
	"createVM": "VBoxManage createvm --name '<VM_Name>' --ostype '<OS_Type>' --register",
	"memory":   "VBoxManage modifyvm '<VM_Name>' --memory <Amount_in_MB>",
	"cpus":     "VBoxManage modifyvm '<VM_Name>' --cpus <Number_of_CPUs>",
	"hdd1":     "VBoxManage createhd --filename '<Path_to_VDI>' --size 20480",
	"hdd2":     "VBoxManage storagectl '<VM_Name>' --name 'SATA Controller' --add sata --controller IntelAHCI",
	"hdd3":     "VBoxManage storageattach '<VM_Name>' --storagectl 'SATA Controller' --port 0 --device 0 --type hdd --medium '<Path_to_VDI>'",
	"iso":      "VBoxManage storageattach '<VM_Name>' --storagectl 'SATA Controller' --port 0 --device 0 --type dvddrive --medium '<Path_to_ISO>'",
	"network":  "VBoxManage modifyvm '<VM_Name>' --nic1 nat",
	"cloneVM":  "VBoxManage clonevm '<Source_VM_Name>' --name '<New_VM_Name>' --register",
	"deleteVM":   "VBoxManage unregistervm '<VM_Name>' --delete",
	"upload":   "VBoxManage guestcontrol '<VM_Name>' copyto --source '<Local_File>' --target '<Guest_File>'",
	"transfer": "VBoxManage guestcontrol '<Source_VM_Name>' copyfrom --target '<Dest_VM_Name>' --source '<Source_File>' --destination '<Dest_File>'",
	"executeLinux":  "VBoxManage guestcontrol '<VM_Name>' run --exe '/bin/bash' --username <username> --password <password> -- -c <command>",
	"executeWindows":  "VBoxManage guestcontrol '<VM_Name>' run --exe <command> --username <username> --password <password>",
	"change":   "VBoxManage modifyvm '<VM_Name>' --<Setting_Name> <Value>",
	"poweroff": "VBoxManage controlvm '<VM_Name>' poweroff",
	"poweron":  "VBoxManage startvm '<VM_Name>' --type headless",
	"getStatus": "VBoxManage showvminfo '<VM_Name>' --machinereadable",
	"listVMs":  "VBoxManage list vms",
}

func executeCommand(cmd string) error {
	command := exec.Command("bash","-c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command '%s': \nOutput: %s", cmd, string(output))
	}
	return nil
}

func CreateVM(vmName, osType, amountInMB, numCPUs string) error {
	sourceVMName := ""
	if osType == "ubuntu"{
		sourceVMName = "Base-Ubuntu" 
	}else if osType == "windows10"{
		sourceVMName = "Base-Windows"
	}

	fmt.Println("------------------------")
	fmt.Println(vmName, osType, sourceVMName, amountInMB, numCPUs)
	if err := CloneVM(sourceVMName, vmName); err == nil{
		fmt.Println("???????????")
		if err = ChangeVMSettings(vmName, "memory", amountInMB); err == nil{
			if err = ChangeVMSettings(vmName, "cpus", numCPUs); err == nil{
				return nil
			}
			return err
		}
		return err
	}else{
		return err
	}
}

// DeleteVM function to delete a Virtual Machine along with its VDI and network settings
func DeleteVM(vmName string) error {
	// Replace placeholder with the actual VM name
	cmd := strings.ReplaceAll(commands["deleteVM"], "<VM_Name>", vmName)
	fmt.Println(cmd)
	command := exec.Command("bash", "-c", cmd)

	// Start the command
	err := command.Start()
	if err != nil {
		return fmt.Errorf("failed to start command: %s", err)
	}

	// Wait for the command to complete
	err = command.Wait()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s", err)
	}

	return nil
}

// CloneVM function to clone a Virtual Machine
func CloneVM(sourceVMName, newVMName string) error {
	// Use the "cloneVM" command to clone the VM
	cmd := commands["cloneVM"]

	// Replace placeholders with actual values
	cmd = strings.ReplaceAll(cmd, "<Source_VM_Name>", sourceVMName)
	cmd = strings.ReplaceAll(cmd, "<New_VM_Name>", newVMName)

	fmt.Println(cmd)
	command := exec.Command("bash", "-c", cmd)

	// Start the command
	err := command.Start()
	if err != nil {
		return fmt.Errorf("failed to start command: %s", err)
	}

	// Wait for the command to complete
	err = command.Wait()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s", err)
	}

	fmt.Println(sourceVMName, newVMName)

	return nil
}


// ChangeVMSettings function to change settings of a Virtual Machine
func ChangeVMSettings(vmName, settingName, settingValue string) error {
	// Use the "change" command to modify VM settings
	cmd := commands["change"]

	// Replace placeholders with actual values
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
	cmd = strings.ReplaceAll(cmd, "<Setting_Name>", settingName)
	cmd = strings.ReplaceAll(cmd, "<Value>", settingValue)

	fmt.Println(cmd)
	command := exec.Command("bash","-c", cmd)
		// Start the command
	err := command.Start()
	if err != nil {
		return fmt.Errorf("failed to start command: %s", err)
	}

	// Wait for the command to complete
	err = command.Wait()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s", err)
	}

	return nil
}

// PowerOffVM function to power off a Virtual Machine
func PowerOffVM(vmName string) error {
	// Use the "poweroff" command to power off the VM
	cmd := commands["poweroff"]

	// Replace placeholder with the actual VM name
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
	command := exec.Command("bash","-c", cmd)
		// Start the command
	err := command.Start()
	if err != nil {
		return fmt.Errorf("failed to start command: %s", err)
	}

	// Wait for the command to complete
	err = command.Wait()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s", err)
	}

	return nil
}

// PowerOnVM function to power on a Virtual Machine
func PowerOnVM(vmName string) error {
	// Use the "poweron" command to power on the VM
	cmd := commands["poweron"]

	// Replace placeholder with the actual VM name
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
	command := exec.Command("bash","-c", cmd)
		// Start the command
	err := command.Start()
	if err != nil {
		return fmt.Errorf("failed to start command: %s", err)
	}

	// Wait for the command to complete
	err = command.Wait()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s", err)
	}

	return nil
}

// GetVMStatus function to get the status of a Virtual Machine
func GetVMStatus(vmName string) (string, error) {
	// Use the "getStatus" command to get VM status
	cmd := commands["getStatus"]

	// Replace placeholder with the actual VM name
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)

	command := exec.Command("bash", "-c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

	fmt.Println("VM Status:\n", string(output))
    pattern := `VMState="([^"]+)"`
    // Compile the regular expression
    regex := regexp.MustCompile(pattern)
    matches := regex.FindAllStringSubmatch(string(output), -1)

    if len(matches) > 0 {
        // Extract the captured text
        capturedText := matches[0][1]
        return capturedText, nil
    } else {
        return "", fmt.Errorf("No match found.")
    }
}

// GetAvailableVMs function to get the list of available Virtual Machines
func GetAvailableVMs() error {
	// Use the "listVMs" command to list available VMs
	cmd := commands["listVMs"]

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

	fmt.Println("Available VMs:\n", string(output))

	return nil
}

// UploadFileToVM function to upload a file to a Virtual Machine
func UploadFileToVM(vmName, guestFilePath string, fileContect []byte) error {
    // Generate a unique filename for the saved file
    uniqueFilename := generateUniqueFilename()

    // Define the local folder where the file will be saved
    localFolder := "./"

    // Combine the local folder and unique filename to create the full local path
    localFilePath := filepath.Join(localFolder, uniqueFilename)

    // Save the received local file content to the local path
    if err := saveFileContent(localFilePath, fileContect); err != nil {
        return err
    }

    // Use the "upload" command to upload the saved file to the VM
    cmd := commands["upload"]

    // Replace placeholders with the actual VM name, local file path, and guest file path
    cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
    cmd = strings.ReplaceAll(cmd, "<Local_File>", localFilePath) // Use the saved local file path
    cmd = strings.ReplaceAll(cmd, "<Guest_File>", guestFilePath)

    args := strings.Fields(cmd)
    command := exec.Command(args[0], args[1:]...)
    output, err := command.CombinedOutput()
    if err != nil {
        return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
    }

    return nil
}

func generateUniqueFilename() string {
    // Generate a unique filename, you can use a timestamp or UUID, for example
    // Here, we use a timestamp as a simple example
    timestamp := time.Now().Unix()
    return fmt.Sprintf("file_%d", timestamp)
}

func saveFileContent(filePath string, fileContent []byte ) error {
    // Write the file content to the specified file path
    err := ioutil.WriteFile(filePath, fileContent, 0644)
    if err != nil {
        return fmt.Errorf("failed to save file content: %v", err)
    }
    return nil
}

// TransferFileBetweenVMs function to transfer a file between two Virtual Machines
func TransferFileBetweenVMs(sourceVMName, sourceFile, destVMName, destFile string) error {
	// Use the "transfer" command to transfer a file between VMs
	cmd := commands["transfer"]

	// Replace placeholders with actual values
	cmd = strings.ReplaceAll(cmd, "<Source_VM_Name>", sourceVMName)
	cmd = strings.ReplaceAll(cmd, "<Source_File>", sourceFile)
	cmd = strings.ReplaceAll(cmd, "<Dest_VM_Name>", destVMName)
	cmd = strings.ReplaceAll(cmd, "<Dest_File>", destFile)

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

	return nil
}

// ExecuteCommandOnVM function to execute a command on a Virtual Machine
func ExecuteCommandOnVM(vmName, userCommand string) (string, error) {
	// Use the "execute" command to execute a command on the VM
	cmd := commands["executeLinux"]

	// Replace placeholders with actual values
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
	cmd = strings.ReplaceAll(cmd, "<command>", userCommand)
	cmd = strings.ReplaceAll(cmd, "<username>", "defaultuser")
	cmd = strings.ReplaceAll(cmd, "<password>", "defaultuser")
	
	command := exec.Command("bash","-c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		return "", fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

	return string(output), nil
}