package models

type UserDB struct{
	ID uint `gorm:"primary_key AUTO_INCREMENT"`
	Nama string `gorm:"column:nama"`
	Email string `gorm:"column:email"`
}

func (db *UserDB) TableName() string{
	return "user"
}