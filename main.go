package main

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/ouhaohan8023/patientRegistration/database"
	"github.com/ouhaohan8023/patientRegistration/model"
	"github.com/ouhaohan8023/patientRegistration/services"
	"github.com/spf13/viper"
	"log"
)

var connect *gorm.DB

func init() {
	viper.SetConfigName("app")
	viper.AddConfigPath("./configs")
	viper.AddConfigPath("../configs")
	viper.AddConfigPath(".")
	err := viper.ReadInConfig()
	if err != nil {
		log.Fatal(err)
	}
}

var (
	sigKey = []byte("signature_hmac_secret_shared_key")
)

func main() {
	connect = database.ConnectMysql()
	database.Migrate(connect)

	app := iris.New()

	verifier := jwt.NewVerifier(jwt.HS256, sigKey)
	verifier.WithDefaultBlocklist()
	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(model.Admin)
	})
	protectedAPI := app.Party("/lists")
	protectedAPI.Use(verifyMiddleware)
	protectedAPI.Get("/", services.List)

	app.Post("/submit", services.Submit)
	app.Post("/login", services.Login)

	app.Listen(":8080")
}