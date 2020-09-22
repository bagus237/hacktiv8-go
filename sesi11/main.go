package main

import (
	"fmt"

	"sesi11/models"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	username := "root"
	password := ""
	dbname := "golang"
	host := "localhost:3306"

	db, err := gorm.Open("mysql", fmt.Sprintf("%s:%s@(%s)/%s?charset=utf8&parseTime=True&loc=Local", username, password, host, dbname))
	defer db.Close()
	if err != nil {
		fmt.Println("Failed to connect DB")
		return
	}
	db.LogMode(true)
	db.AutoMigrate(&models.UserDB{})

	dataInsert := models.UserDB{
		Nama:  "Pratama",
		Email: "p@gmail.com",
	}
	InsertNewUser(db, dataInsert)

	modelUser1 := models.UserDB{ID: 1}
	UpdateUser1(db, modelUser1, "nama", "dika")

	UpdateNamaUser(db, 2, "dicky")

	user, _ := GetUser(db, "Dicky")
	fmt.Println(user)

	userList, _ := GetAllUser(db, "Dicky Pratama", "dicky@gmail.com")
	fmt.Println(userList)

	users, _ := RawGetData(db)
	fmt.Println(users)

	DeleteUser(db, models.UserDB{ID: 3})
	DeleteUserByID(db, 2)

	fmt.Println("Connected to mysql")
}

func InsertNewUser(db *gorm.DB, req models.UserDB) error {
	err := db.Create(&req).Error
	if err != nil {
		fmt.Println("Failed to insert new user: ", err)
		return err
	}
	return nil
}

func GetUser(db *gorm.DB, nama string) (models.UserDB, error) {
	res := models.UserDB{}
	nama = "%" + nama + "%"
	err := db.Where(`nama like ?`, nama).First(&res).Error
	if err != nil {
		fmt.Println("Failed to get user: ", err)
		return res, err
	}
	return res, nil
}

func GetAllUser(db *gorm.DB, nama string, email string) ([]models.UserDB, error) {
	res := []models.UserDB{}
	err := db.Where("nama = ? and email = ?", nama, email).Limit(1).Order("id DESC").Find(&res).Error
	if err != nil {
		fmt.Println("Failed to get user: ", err)
		return res, err
	}
	return res, nil
}

func UpdateUser1(db *gorm.DB, model models.UserDB, key string, value string) error {
	err := db.Model(&model).Update(key, value).Error
	if err != nil {
		fmt.Println("Failed to update user: ", err)
		return err
	}
	return nil
}

func UpdateNamaUser(db *gorm.DB, id int, nama string) error {
	err := db.Exec(`UPDATE user SET nama = ? WHERE id = ?`, nama, id).Error
	if err != nil {
		fmt.Println("Failed to update user: ", err)
		return err
	}
	return nil
}

func RawGetData(db *gorm.DB) ([]models.UserDB, error) {
	res := []models.UserDB{}
	err := db.Raw(`SELECT * FROM user ORDER BY id DESC LIMIT 2`).Scan(&res).Error
	if err != nil {
		fmt.Println("Failed to get user: ", err)
		return res, err
	}
	return res, nil
}

func DeleteUser(db *gorm.DB, data models.UserDB) error {
	err := db.Delete(&data).Error
	if err != nil {
		fmt.Println("Failed to delete user: ", err)
		return err
	}
	return nil
}

func DeleteUserByID(db *gorm.DB, id int) error {
	err := db.Exec(`DELETE FROM user WHERE id = ?`, id).Error
	if err != nil {
		fmt.Println("Failed to delete user: ", err)
		return err
	}
	return nil
}
