package models

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

// User represents a simple user model
type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}


type ConfigSet struct{
	Port	string	`env:"Cloud_Port" env-default:"8000"`
	Host	string	`env:"Cloud_Host" env-default:"localhost"`
	
}
