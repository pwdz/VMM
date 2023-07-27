package vboxWrapper


import (
	"fmt"
	"os/exec"
	"strings"
)

// Define commands as a static map
var commands = map[string]string{
	"createVM": "VBoxManage createvm --name <VM_Name> --ostype <OS_Type> --register",
	"memory":   "VBoxManage modifyvm <VM_Name> --memory <Amount_in_MB>",
	"cpus":     "VBoxManage modifyvm <VM_Name> --cpus <Number_of_CPUs>",
	"hdd": `VBoxManage createhd --filename "<Path_to_VDI>" --size 20480
VBoxManage storagectl <VM_Name> --name "SATA Controller" --add sata --controller IntelAHCI
VBoxManage storageattach <VM_Name> --storagectl "SATA Controller" --port 0 --device 0 --type hdd --medium "<Path_to_VDI>"`,
	"iso":     `VBoxManage storageattach <VM_Name> --storagectl "IDE Controller" --port 0 --device 0 --type dvddrive --medium "<Path_to_ISO>"`,
	"network": "VBoxManage modifyvm <VM_Name> --nic1 nat",
	"cloneVM": "VBoxManage clonevm <Source_VM_Name> --name <New_VM_Name> --register",
	"deleteVM": "VBoxManage unregistervm <VM_Name> --delete",
	"upload":    "VBoxManage guestcontrol <VM_Name> copyto --source <Local_File> --target <Guest_File>",
	"transfer":  "VBoxManage guestcontrol <Source_VM_Name> copyfrom --target <Dest_VM_Name> --source <Source_File> --destination <Dest_File>",
	"execute":   "VBoxManage guestcontrol <VM_Name> run --exe <Path_to_Exe> -- <Arguments>",
	"change":    "VBoxManage modifyvm <VM_Name> --<Setting_Name> <Value>",
	"poweroff":  "VBoxManage controlvm <VM_Name> poweroff",
	"poweron":   "VBoxManage startvm <VM_Name> --type headless",
	"getStatus": "VBoxManage showvminfo <VM_Name> --machinereadable",
	"listVMs":   "VBoxManage list vms",
}

// CreateVM function to create a Virtual Machine
func CreateVM(vmName, osType, amountInMB, numCPUs, vdiPath, isoPath string) error {
	requiredCommands := map[string]string{
		"createVM": commands["createVM"],
		"memory":   commands["memory"],
		"cpus":     commands["cpus"],
		"hdd":      commands["hdd"],
		"iso":      commands["iso"],
	}

	// Replace placeholders with actual values
	for key, cmd := range requiredCommands {
		cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
		cmd = strings.ReplaceAll(cmd, "<OS_Type>", osType)
		cmd = strings.ReplaceAll(cmd, "<Amount_in_MB>", amountInMB)
		cmd = strings.ReplaceAll(cmd, "<Number_of_CPUs>", numCPUs)
		cmd = strings.ReplaceAll(cmd, "<Path_to_VDI>", vdiPath)
		cmd = strings.ReplaceAll(cmd, "<Path_to_ISO>", isoPath)
		requiredCommands[key] = cmd
	}

	// Execute the commands to create the VM
	for _, cmd := range requiredCommands {
		args := strings.Fields(cmd)
		command := exec.Command(args[0], args[1:]...)
		output, err := command.CombinedOutput()
		if err != nil {
			return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
		}
	}

	return nil
}

// DeleteVM function to delete a Virtual Machine
func DeleteVM(vmName string) error {
	// Use the "deleteVM" command to delete the VM
	cmd := commands["deleteVM"]

	// Replace placeholder with the actual VM name
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
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

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

	return nil
}

// PowerOnVM function to power on a Virtual Machine
func PowerOnVM(vmName string) error {
	// Use the "poweron" command to power on the VM
	cmd := commands["poweron"]

	// Replace placeholder with the actual VM name
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
	}

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
func UploadFileToVM(vmName, localFile, guestFile string) error {
	// Use the "upload" command to upload a file to the VM
	cmd := commands["upload"]

	// Replace placeholders with actual values
	cmd = strings.ReplaceAll(cmd, "<VM_Name>", vmName)
	cmd = strings.ReplaceAll(cmd, "<Local_File>", localFile)
	cmd = strings.ReplaceAll(cmd, "<Guest_File>", guestFile)

	args := strings.Fields(cmd)
	command := exec.Command(args[0], args[1:]...)
	output, err := command.CombinedOutput()
	if err != nil {
		return fmt.Errorf("failed to execute command: %s\nOutput: %s", cmd, string(output))
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


// ////////////////////////////////////////
/*
const(
	VBoxCommand = "vboxmanage"
)

func printCommand(cmd *exec.Cmd) {
  log.Printf("==> Executing: %s\n", strings.Join(cmd.Args, " "))
}

func printError(err error) {
  if err != nil {
    os.Stderr.WriteString(fmt.Sprintf("==> Error: %s\n", err.Error()))
  }
}

func printOutput(outs []byte) {
  if len(outs) > 0 {
    log.Printf("==> Output: %s\n", string(outs))
  }
}

func GetStatus(vmName string) (string, error){
	cmd := exec.Command(VBoxCommand, "showvminfo",vmName,"--machinereadable")

	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	if err != nil{
		printError(err)
		return "", err
	}
	regex, _ := regexp.Compile("VMState=\"[a-zA-Z]+\"")
	status := regex.FindString(string(output))
	log.Println(status)
	status = strings.Split(status, "=")[1]
	status = strings.Trim(status, "\"")
	return status, nil
}
func GetVmNames() ([]string, error){
	cmd := exec.Command(VBoxCommand, "list","vms")

	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	
	if err != nil{
		printError(err)
		return nil, err
	}
	regex, _ := regexp.Compile("\"[A-Za-z0-9]+\"")
	vmNames := regex.FindAllString(string(output), -1)

	for index, vmName := range vmNames{
		vmNames[index] = strings.Trim(vmName, "\"")
	}
	return vmNames, nil
}
func PowerOn(vmName string) (string, error){
	status, err := GetStatus(vmName)
	if err != nil{
		return "", err
	}
	
	if status == "poweroff" {
		cmd := exec.Command(VBoxCommand, "startvm",vmName,"--type","headless")

		printCommand(cmd)
		output, err := cmd.CombinedOutput()
		printOutput(output)
		if err != nil{
			printError(err)
			return "", err
		}
		return "Powering on", nil
	}
	return "", fmt.Errorf(vmName + ">> current status: " + status)
}
func PowerOff(vmName string)(string, error){
	status, err := GetStatus(vmName)
	if err != nil{
		return "", err
	}
	
	if status == "running" {
		cmd := exec.Command(VBoxCommand, "controlvm",vmName,"poweroff")

		printCommand(cmd)
		output, err := cmd.CombinedOutput()
		printOutput(output)
		printError(err)
		if err != nil{
			printError(err)
			return "", err
		}
		return "Powering off", nil
	}

	return "", fmt.Errorf(vmName + ">> current status: " + status)
}

func ChangeSetting(vmName string, cpu, ram int)(string, error){
	args := []string{"modifyvm", vmName}

	if cpu > 0{
		args = append(args, "--cpus", strconv.Itoa(cpu))
	}
	if ram > 0{
		args = append(args, "--memory",strconv.Itoa(ram))
	}
	
	cmd := exec.Command(VBoxCommand, args...)	
	printCommand(cmd)

	output, err := cmd.CombinedOutput()
	printOutput(output)
	if err != nil{
		printError(err)
		return "", fmt.Errorf(string(output))
	}

	return "Ok", nil
}
func Clone(vmSrc, vmDst string)(string, error){
	cmd := exec.Command(VBoxCommand, "clonevm",vmSrc,"--name",vmDst, "--register")
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printOutput(output)
	if err != nil{
		printError(err)
		return "", fmt.Errorf(string(output))
	}

	return "Ok", nil
}
func Delete(vmName string)(string, error){
	cmd := exec.Command(VBoxCommand, "unregistervm",vmName,"--delete")
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printOutput(output)
	if err != nil{
		printError(err)
		return "", fmt.Errorf(string(output))
	}

	return "Ok", nil
}
func Execute(vmName, input string)(string, string, error){
	cmd := exec.Command(VBoxCommand, "guestcontrol",vmName,"run","bin/sh","--username","pwdz","--password", "pwdz", "--wait-stdout", "--wait-stderr", "--","-c",input)
	printCommand(cmd)
	output, err := cmd.CombinedOutput()
	printOutput(output)
	if err != nil{
		printError(err)
		return "", "", fmt.Errorf(string(output))
	}

	return "Ok", string(output), nil
}
func Transfer(vmSrc, vmDst, originPath, dstPath string)(string, error){
	 _, err := os.Stat("./temp")
    if err != nil && os.IsNotExist(err){
		os.Mkdir("temp", 666)
	}

	paths := strings.Split(originPath, "/")
	fileName := paths[len(paths) - 1]

	internalPath :=   "./temp/" 
	log.Println(vmSrc, vmDst)

	copyFromCommand := exec.Command(VBoxCommand, "guestcontrol",vmSrc,"copyfrom","--target-directory",internalPath , originPath,"--username","pwdz","--password", "pwdz","--verbose")
	printCommand(copyFromCommand)
	copyFromOutput, err := copyFromCommand.CombinedOutput()
	printOutput(copyFromOutput)
	if err != nil{
		printError(err)
		return "", fmt.Errorf(string(copyFromOutput))
	}


	internalPath += fileName
	log.Println(internalPath)

	copyToCommand := exec.Command(VBoxCommand, "guestcontrol",vmDst,"copyto","--target-directory",dstPath, internalPath,"--username","pwdz","--password", "pwdz", "--verbose")
	printCommand(copyToCommand)
	copyToOutput, err := copyToCommand.CombinedOutput()
	printOutput(copyToOutput)
	if err != nil{
		printError(err)
		return "", fmt.Errorf(string(copyToOutput))
	}
	os.Remove(internalPath)

	return "Ok", nil
}
func Upload(vmDst, dstPath, originPath string)(string, error){
	copyToCommand := exec.Command(VBoxCommand, "guestcontrol",vmDst,"copyto","--target-directory",dstPath, "/home/user/BachelorProject/TestFile.txt","--username","pwdz","--password", "pwdz", "--verbose")
	printCommand(copyToCommand)
	copyToOutput, err := copyToCommand.CombinedOutput()
	printOutput(copyToOutput)
	if err != nil{
		printError(err)
		return "", fmt.Errorf(string(copyToOutput))
	}

	return "Ok", nil
}

*/