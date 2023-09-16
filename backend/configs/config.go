package configs 

// DBConfig represents the database configuration.
type DBConfig struct {
    Host     string
    Port     string
    User     string
    Password string
    DBName   string
}

func GetDBConfig() *DBConfig {
    return &DBConfig{
        Host:     "127.0.0.1",
        Port:     "3306",
        User:     "pwdz",
        Password: "1234567890",
        DBName:   "mydb",
    }
}

type ServerConfig struct{
	Port	string	`env:"Server_Port" env-default:"8000"`
	Host	string	`env:"Server_Host" env-default:"localhost"`	
}
