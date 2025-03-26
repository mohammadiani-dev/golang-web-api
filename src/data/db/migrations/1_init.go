package migrations

import (
	"golang-web-api/config"
	"golang-web-api/data/db"
	"golang-web-api/data/models"
	"golang-web-api/pkg/logging"

	"gorm.io/gorm"
	"golang.org/x/crypto/bcrypt"
	// "golang-web-api/logging"
)

var logger = logging.NewLogger(&config.GetConfig().Logger)

func Up_1() {
	db := db.GetDbClient()

	createTables(db)
	createDefaultData(db)
}

func createTables(db *gorm.DB) {
	tables := []interface{}{}

	country := models.Country{}
	city := models.City{}
	user := models.User{}
	role := models.Role{}
	userRole := models.UserRole{}

	tables = addNewTables(db, country, tables)
	tables = addNewTables(db, city, tables)
	tables = addNewTables(db, user, tables)
	tables = addNewTables(db, role, tables)
	tables = addNewTables(db, userRole, tables)

	db.Migrator().CreateTable(tables...)
	logger.Info(logging.Postgres, logging.Migration, "Migrating tables", nil)

}

func addNewTables(db *gorm.DB, model interface{}, tables []interface{}) []interface{} {
	if !db.Migrator().HasTable(&model) {
		tables = append(tables, &model)
	}
	return tables
}


func createDefaultData(db *gorm.DB) {
	role := models.Role{Name: "admin"}
	createRoleIfNotExists(db, &role)

	user := models.User{Username: "admin" , FirstName: "admin" , LastName: "admin" , Email: "admin@admin.com" , MobileNumber: "09123456789"}
	password := "admin"
	hashedPassword, _ := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	user.Password = string(hashedPassword)
	createUserIfNotExists(db, &user , role.ID)
}

func createRoleIfNotExists(db *gorm.DB , role *models.Role) {
	var exists int64
	db.Model(&models.Role{}).Where("name = ?", role.Name).Count(&exists)
	if exists == 0 {
		db.Create(role)
	}
}

func createUserIfNotExists(db *gorm.DB , user *models.User , roleId uint) {
	var exists int64
	db.Model(&models.User{}).Where("username = ?", user.Username).Count(&exists)
	if exists == 0 {
		db.Create(user)
		userRole := models.UserRole{UserId: user.ID, RoleId: roleId}
		createUserRoleIfNotExists(db, &userRole)
	}
}

func createUserRoleIfNotExists(db *gorm.DB , userRole *models.UserRole) {
	var exists int64
	db.Model(&models.UserRole{}).Where("user_id = ?", userRole.UserId).Where("role_id = ?", userRole.RoleId).Count(&exists)
	if exists == 0 {
		db.Create(userRole)
	}
}

func Down_1() {
	// db := db.GetDbClient()
	// db.Migrator().DropTable(&models.Country{}, &models.City{})
}
