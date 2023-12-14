package main

import (
	"fmt"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type UserTest struct {
	ID    int    `gorm:"primarykey"`
	Name  string `gorm:"<:create"`
	Email string `gorm:"<:create"`
}

var db *gorm.DB
var err error

func init() {

	dsn := "host=localhost user=postgres password=iamsangeeth dbname=test port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{Logger: logger.Default.LogMode(logger.Info)})
	fmt.Println(db, err)
	if err != nil {
		return
	}
	db.AutoMigrate(&UserTest{})

}

func main() {
	// createDbRecord(db)
	// fetchwithConditionRecord(db)
	fetchAllRecords(db)
	// updateRecord(db)
	// deleteRecord(db)

}

func createDbRecord(db *gorm.DB) {
	u := UserTest{Name: "Sarang", Email: "sangeethsajeev518@gmail.com"}
	db.Create(&u)

}

func fetchwithConditionRecord(db *gorm.DB) {

	var u UserTest
	tc := db.Where("id=?", 2).Find(&u)
	fmt.Println(tc.RowsAffected)
	fmt.Println(u)
}

func fetchAllRecords(db *gorm.DB) {
	var u []UserTest
	db.Find(&u)
	fmt.Println(u)

}

func updateRecord(db *gorm.DB) {
	db.Table("user_tests").Where("id=?", 2).Updates(map[string]interface{}{"email": "sarangsajeev@gmail.com"})
}

func deleteRecord(db *gorm.DB) {
	db.Where("id=?", 2).Delete(&UserTest{})
}
