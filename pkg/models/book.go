package models

import(
    "github.com/jinzhu/gorm"
	"github.com/luminous44/ReadSphere/pkg/config"
)

var db *gorm.DB

type Book struct{
	gorm.Model

	Name string `gorm:"primaryKey" json:"name"`
	Author string `json:"author"`
	Publication string `json:"publication"`
}

func init(){
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}