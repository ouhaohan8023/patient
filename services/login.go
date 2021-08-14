package services

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/ouhaohan8023/patientRegistration/lib"
	"github.com/ouhaohan8023/patientRegistration/model"
	"github.com/spf13/viper"
	"time"
)

func init() {

}

var (
	sigKey = []byte("signature_hmac_secret_shared_key")
)

type LoginInput struct {
	Username   string `json:"username"`
	Password   string `json:"password"`
	IsRemember string `json:"is_remember"`
}

func Login(ctx iris.Context) {
	c := &LoginInput{}

	if err := ctx.ReadJSON(c); err != nil {
		panic(err.Error())
	} else {
		username := c.Username
		password := c.Password
		// todo verify by db
		if username == "admin" && password == "admin" {
			signer := jwt.NewSigner(jwt.HS256, sigKey, 14400*time.Minute)
			token := generateToken(signer, ctx)
			environment := viper.GetString(`environment`)

			context := iris.Map{
				"token": token,
				"uuid":  1,
				"info": iris.Map{
					"name":   "Admin",
					"avatar": viper.GetString(environment+`.baseUrl`) + "avatar.jpeg",
				},
			}
			ctx.JSON(lib.CommonResponse(200, "Login Success", context))
		} else {
			ctx.JSON(lib.CommonResponse(500, "Username / Password ERROR", nil))
		}
	}
}

func generateToken(signer *jwt.Signer, ctx iris.Context) string {
	claims := model.Admin{Username: "bar"}

	token, err := signer.Sign(claims)
	if err != nil {
		ctx.StopWithStatus(iris.StatusInternalServerError)
		return ""
	}
	return string(token)
}
