package models

type VMRequest struct {
	VMID         uint   `json:"vm_id"`
	VMName       string `json:"vm_name"`
	NewVMName    string `json:"new_vm_name"`
	OSType       string `json:"os_type"`
	RamInMB      string `json:"ram_in_mb"`
	NumCPUs      string `json:"num_cpus"`
	VDIPath      string `json:"vdi_path"`
	ISOPath      string `json:"iso_path"`
	FileContent  string `json:"file_content"` // Base64-encoded file content
	GuestFilePath string `json:"guest_file_path"`
}

type VMResponse struct {
    Message string `json:"message"`
    Error   string `json:"error,omitempty"`
    Data    string `json:"data,omitempty"`
}

type Role string

const (
    UserRole  Role = "user"
    AdminRole Role = "admin"
)

type User struct {
    ID       uint   `gorm:"primary_key json:id"`
    Username string `gorm:"unique_index" json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     Role   `gorm:"default:'user'" json:"role" sql:"column:role"`
}

type UserWithVMCounts struct {
    User
    ActiveVMCount   int `json:"active_vm_count"`
    InactiveVMCount int `json:"inactive_vm_count"`
}

type VM struct {
    ID        uint   `gorm:"primary_key"`
    UserID    uint   `json:"user_id"` // User ID to establish the relationship
    Name      string `json:"name"`
    OSType    string `json:"os_type"`
    RAM       int    `json:"ram"`
    CPU       int    `json:"cpu"`
    Status    string `json:"status"` // You can define an enum for possible statuses
    IsDeleted bool   `json:"is_deleted" gorm:"default:false"`
}
type VMWithUser struct{
    VM
    Username string `json:"username"`
}