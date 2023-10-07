package db

import (
	"fmt"
	"strconv"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql" // Import MySQL driver
	"github.com/pwdz/VMM/code/backend/configs"
	"github.com/pwdz/VMM/code/backend/models"
	"golang.org/x/crypto/bcrypt"
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
	db.connection.AutoMigrate(&models.User{}, &models.VM{}, &models.PriceConfig{})

	// Add a foreign key constraint to link VMs to Users
	db.connection.Model(&models.VM{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
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
	if hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost); err != nil{
		fmt.Println(err,"WRONG")
		return err
	}else{
		user.Password = string(hashedPassword)
	}
	return db.connection.Create(user).Error
}

func (db *Database) GetVMsByUserID(userID uint) ([]models.VM, error) {
    var vms []models.VM
    if err := db.connection.Where("!is_deleted and user_id = ?", userID).Find(&vms).Error; err != nil {
        return nil, err
    }
    return vms, nil
}

func (db *Database) GetVMStatus(vmID uint) (string, error) {
    var vm models.VM
    if err := db.connection.Select("status").Where("id = ?", vmID).First(&vm).Error; err != nil {
        return "", err
    }
    return vm.Status, nil
}

// CreateVM creates a new VM record in the database.
func (db *Database) CreateVM(vm *models.VM) error {
    if err := db.connection.Create(vm).Error; err != nil {
        return err
    }
    return nil
}

func (db *Database) DeleteVM(vmID uint) error {
	fmt.Println(vmID)
	var vm models.VM
	if db.connection.Where("id = ?", vmID).First(&vm).RecordNotFound() {
		return fmt.Errorf("VM not found")
	}

	// Set IsDeleted to true
	vm.IsDeleted = true

	if err := db.connection.Save(&vm).Error; err != nil {
		return err
	}

	return nil
}

func (db *Database) FindVMByName(vmName string) *models.VM {
	var vm models.VM
	if db.connection.Where("name = ?", vmName).First(&vm).RecordNotFound() {
		return nil
	}
	return &vm
}

func (db *Database) UpdateVMSetting(vmID uint, settingName string, settingValue int) error {
	// Update the specified VM setting in the database
	fmt.Println(settingName, settingValue)
	return db.connection.Model(&models.VM{}).Where("id = ?", vmID).Update(settingName, settingValue).Error
}

func (db *Database) FindVMByID(vmID uint) *models.VM {
	var vm models.VM
	if db.connection.Where("id = ?", vmID).First(&vm).RecordNotFound() {
		return nil
	}
	return &vm
}

func (db *Database) UpdateVMStatus(vmID uint, newStatus string) error {
	// Find the VM by its ID
	var vm models.VM
	if db.connection.Where("id = ?", vmID).First(&vm).RecordNotFound() {
		return fmt.Errorf("VM not found")
	}

	// Update the status
	vm.Status = newStatus

	// Save the updated VM back to the database
	if err := db.connection.Save(&vm).Error; err != nil {
		return err
	}

	return nil
}

func (db *Database) GetNonAdminUsersWithVMCounts() ([]models.UserWithVMCounts, error) {
    var usersWithVMCounts []models.UserWithVMCounts

	query := `
	select *
	from mydb.users u
			 left join (select user_id,
							   sum(if(status = 'on', 1, 0))  active_vm_count,
							   sum(if(status = 'off', 1, 0)) inactive_vm_count,
							   sum(if(!is_deleted, cost, 0)) total_cost
						from mydb.vms v
						group by user_id) v on v.user_id = u.id
	where u.role != 'admin'`

	// Execute the custom SQL query
	if err := db.connection.Raw(query).Scan(&usersWithVMCounts).Error; err != nil {
		return nil, err
	}

    return usersWithVMCounts, nil
}

// GetAllVMs retrieves a list of all VMs from the database.
func (db *Database) GetAllVMs() ([]models.VMWithUser, error) {
    var vms []models.VMWithUser
	query := "select v.*, u.username from mydb.vms v inner join mydb.users u on u.id = v.user_id"

    if err := db.connection.Raw(query).Scan(&vms).Error; err != nil {
        return nil, err
    }
    return vms, nil
}

func (db *Database) GetUserDataWithVMCounts(userID uint) (models.UserWithVMCounts, error) {
    var userWithVMCount models.UserWithVMCounts

	query := `
	select *
	from mydb.users u
         left join (select user_id,
                           sum(if(status = 'on', 1, 0)) active_vm_count,
                           sum(if(status = 'off', 1, 0))  inactive_vm_count,
						   sum(if(!is_deleted, cost, 0)) total_cost
                    from mydb.vms v where !is_deleted
                    group by user_id) v on v.user_id = u.id
	where u.id = ` 
	
	// Execute the custom SQL query
	err := db.connection.Raw(query + strconv.Itoa(int(userID))).Scan(&userWithVMCount).Error

    return userWithVMCount, err
}
func (db *Database) GetPriceConfigs() ([]models.PriceConfig, error) {
    var priceConfigs []models.PriceConfig
    if err := db.connection.Find(&priceConfigs).Error; err != nil {
        return nil, err
    }
    return priceConfigs, nil
}
func (db *Database) UpdatePriceConfig(priceConfig models.PriceConfig) error {
    // Update the specified VM setting in the database
    fmt.Println(priceConfig.ID, priceConfig.CostPerUnit, priceConfig.Type)
    
    // Specify the model type and the condition in the Where clause
    return db.connection.Model(&models.PriceConfig{}).Where("id = ?", priceConfig.ID).Update("cost_per_unit", priceConfig.CostPerUnit).Error
}
