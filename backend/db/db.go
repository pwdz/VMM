package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	"github.com/pwdz/VMM/backend/configs"
	"github.com/pwdz/VMM/backend/models"
_ 	"github.com/jinzhu/gorm/dialects/mysql" // Import MySQL driver
)

type Database struct {
    connection *gorm.DB
}

func NewDatabase(config *configs.DBConfig) (*Database, error) {
    // Update the database connection string to use the MySQL container's IP and port
    dbURI := fmt.Sprintf(
        "%s:%s@tcp(%s:%s)/%s?charset=utf8&parseTime=True&loc=Local",
        config.User, config.Password, config.Host, config.Port, config.DBName,
    )

	fmt.Println(dbURI)

    db, err := gorm.Open("mysql", dbURI)
	// if db != nil{
		// fmt.Println("ADASDASD")
		// (&Database{connection: db}).AutoMigrate()
	// }
    if err != nil {
        return nil, err
    }
    return &Database{connection: db}, nil
}

func (db *Database) Close() {
    db.connection.Close()
}

func (db *Database) AutoMigrate() {
    db.connection.AutoMigrate(&models.User{})
}

func (db *Database) FindUserByUsername(username string) *models.User {
    var user models.User
    if db.connection.Where("username = ?", username).First(&user).RecordNotFound() {
        return nil
    }
    return &user
}

func (db *Database) CreateUser(user *models.User) error {
    return db.connection.Create(user).Error
}
