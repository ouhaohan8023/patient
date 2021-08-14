package main

import (
	"github.com/iris-contrib/middleware/cors"
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

const maxSize = 8 * iris.MB

func main() {
	connect = database.ConnectMysql()
	database.Migrate(connect)

	app := iris.New()

	// Backed
	verifier := jwt.NewVerifier(jwt.HS256, sigKey)
	verifier.WithDefaultBlocklist()
	verifyMiddleware := verifier.Verify(func() interface{} {
		return new(model.Admin)
	})
	protectedAPI := app.Party("/lists")
	protectedAPI.Use(verifyMiddleware)
	protectedAPI.Post("/{id}", services.Update)

	protectedAPI.Get("/", services.List)

	app.Post("/upload", services.Upload)
	//app.Post("/lists/{id}", services.Update)
	app.Post("/login", services.Login)

	// Front
	app.Post("/submit", services.Submit)

	crs := cors.New(cors.Options{
		AllowedOrigins:   []string{"*"},
		AllowCredentials: true,
	})
	app.UseRouter(crs)

	app.HandleDir("/uploads", iris.Dir("./uploads"))
	app.Listen(":8080", iris.WithPostMaxMemory(maxSize))
}
