package models

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

type User struct {
    ID       uint   `gorm:"primary_key"`
    Username string `gorm:"unique_index" json:"username"`
    Email    string `json:"email"`
    Password string `json:"password"`
    Role     string `gorm:"default:normal"` 
}