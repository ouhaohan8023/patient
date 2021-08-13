package database

import (
	"fmt"
	"github.com/jinzhu/gorm"
	"github.com/ouhaohan8023/patientRegistration/model"
	"github.com/spf13/viper"
)

func ConnectMysql() *gorm.DB {
	environment := "dev"
	host := viper.GetString(environment + `.database.host`)
	database := viper.GetString(environment + ".database.db")
	charset := viper.GetString(environment + ".database.charset")
	username := viper.GetString(environment + ".database.username")
	secret := viper.GetString(environment + ".database.password")
	db, err := gorm.Open("mysql", ""+username+":"+secret+"@("+host+")/"+database+"?charset="+charset+"&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("connect database fail")
		fmt.Println(err)
	} else {
		//fmt.Println("connect database success")
	}
	return db
}

func Migrate(db *gorm.DB) {

	db.AutoMigrate(&model.Users{}, &model.Admin{})
}
