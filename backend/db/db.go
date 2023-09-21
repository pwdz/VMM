package db

import (
	"fmt"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL driver
	"github.com/pwdz/VMM/code/backend/configs"
	"github.com/pwdz/VMM/code/backend/models"
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
	if err != nil {
		return nil, err
	}

	DB := &Database{connection: db}
	DB.AutoMigrate()

	return DB, nil
}

func (db *Database) Close() {
	db.connection.Close()
}

func (db *Database) AutoMigrate() {
	db.connection.AutoMigrate(&models.User{})
}

func (db *Database) FindUserByUsername(username string) *models.User {
	var user models.User
	fmt.Println(username)
	if db.connection.Where("username = ?", username).First(&user).RecordNotFound() {
		return nil
	}
	return &user
}

func (db *Database) CreateUser(user *models.User) error {
	return db.connection.Create(user).Error
}
