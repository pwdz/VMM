package vboxWrapper

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"path/filepath"
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
	"deleteVDI":  "VBoxManage closemedium disk '<VDI_UUID>' --delete",
	"deleteNet":  "VBoxManage modifyvm '<VM_Name>' --nic<Adapter_Number> none",
	"upload":   "VBoxManage guestcontrol '<VM_Name>' copyto --source '<Local_File>' --target '<Guest_File>'",
	"transfer": "VBoxManage guestcontrol '<Source_VM_Name>' copyfrom --target '<Dest_VM_Name>' --source '<Source_File>' --destination '<Dest_File>'",
	"execute":  "VBoxManage guestcontrol '<VM_Name>' run --exe '<Path_to_Exe>' -- <Arguments>",
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

func CreateVM(vmName, osType, amountInMB, numCPUs, isoPath string) error {
	requiredCommands := []string{
		"createVM",
		"memory",
		"cpus",
		"hdd1",
		"hdd2",
		"hdd3",
		"iso",
		"network",
	}

	// Replace placeholders with actual values
	for _, cmdKey := range requiredCommands {
		cmd := commands[cmdKey]
		cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
		cmd = strings.ReplaceAll(cmd, "<OS_Type>", osType)
		cmd = strings.ReplaceAll(cmd, "<Amount_in_MB>", amountInMB)
		cmd = strings.ReplaceAll(cmd, "<Number_of_CPUs>", numCPUs)
		cmd = strings.ReplaceAll(cmd, "<Path_to_VDI>", strings.Join([]string{"/home/user/VirtualBox VMs/",vmName,".vdi"}, ""))
		cmd = strings.ReplaceAll(cmd, "<Path_to_ISO>", isoPath)

		println(cmd)
		if err := executeCommand(cmd); err != nil {
			return err
		}
	}

	return nil
}

// DeleteVM function to delete a Virtual Machine along with its VDI and network settings
func DeleteVM(vmName string) error {
	// Get the UUID of the VM's VDI first
	cmd := "VBoxManage showvminfo " + vmName + " --machinereadable | grep vdi"
	command := exec.Command("bash", "-c", cmd)
	output, err := command.CombinedOutput()
	if err != nil {
		fmt.Printf("failed to get VDI information for VM: %s", vmName)
	}

	// Extract the VDI UUID from the output
	vdiUUID := ""
	lines := strings.Split(string(output), "\n")
	for _, line := range lines {
		if strings.HasPrefix(line, "vdi=") {
			vdiUUID = strings.TrimPrefix(line, "vdi=")
			break
		}
	}

	// Use the "deleteVDI" command to delete the VM's VDI using its UUID
	if vdiUUID != "" {
		cmd = commands["deleteVDI"]
		cmd = strings.ReplaceAll(cmd, "<VDI_UUID>", vdiUUID)
		args := strings.Fields(cmd)
		command = exec.Command(args[0], args[1:]...)
		output, err = command.CombinedOutput()
		if err != nil {
			fmt.Printf("failed to execute command: %s\nOutput: %s", cmd, string(output))
		}
	}

	// Use the "deleteNet" command to remove all the VM's network settings
	cmd = commands["deleteNet"]
	for adapterNumber := 1; adapterNumber <= 8; adapterNumber++ {
		cmdWithAdapter := strings.ReplaceAll(cmd, "<VM_Name>", vmName)
		cmdWithAdapter = strings.ReplaceAll(cmdWithAdapter, "<Adapter_Number>", fmt.Sprintf("%d", adapterNumber))
		args := strings.Fields(cmdWithAdapter)
		command = exec.Command(args[0], args[1:]...)
		output, err = command.CombinedOutput()
		if err != nil {
			// The network adapter might not exist, so ignore the error
			continue
		}
	}

	// Use the "deleteVM" command to delete the VM
	cmd = commands["deleteVM"]

	// Replace placeholder with the actual VM name
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)

	args := strings.Fields(cmd)
	command = exec.Command(args[0], args[1:]...)
	output, err = command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
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

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

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

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

	return nil
}

// PowerOffVM function to power off a Virtual Machine
func PowerOffVM(vmName string) error {
	// Use the "poweroff" command to power off the VM
	cmd := commands["poweroff"]

	// Replace placeholder with the actual VM name
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
	executeCommand(cmd)
	return nil
}

// PowerOnVM function to power on a Virtual Machine
func PowerOnVM(vmName string) error {
	// Use the "poweron" command to power on the VM
	cmd := commands["poweron"]

	// Replace placeholder with the actual VM name
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)

	executeCommand(cmd)
	return nil
}

// GetVMStatus function to get the status of a Virtual Machine
func GetVMStatus(vmName string) error {
	// Use the "getStatus" command to get VM status
	cmd := commands["getStatus"]

	// Replace placeholder with the actual VM name
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

	fmt.Println("VM Status:\n", string(output))

	return nil
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
func ExecuteCommandOnVM(vmName, pathToExe, arguments string) error {
	// Use the "execute" command to execute a command on the VM
	cmd := commands["execute"]

	// Replace placeholders with actual values
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
	cmd = strings.ReplaceAll(cmd, "<Path_to_Exe>", pathToExe)
	cmd = strings.ReplaceAll(cmd, "<Arguments>", arguments)

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

	return nil
}