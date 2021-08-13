package services

import (
	"github.com/kataras/iris/v12"
	"github.com/kataras/iris/v12/middleware/jwt"
	"github.com/ouhaohan8023/patientRegistration/lib"
	"github.com/ouhaohan8023/patientRegistration/model"
	"time"
)

func init() {

}
var (
	sigKey = []byte("signature_hmac_secret_shared_key")
)

func Login(ctx iris.Context) {
	username := ctx.PostValue("username")
	password := ctx.PostValue("password")
	// todo verify by db
	if username == "admin" && password == "admin" {
		signer := jwt.NewSigner(jwt.HS256, sigKey, 60*time.Minute)
		token := generateToken(signer, ctx)
		context := iris.Map{
			"token": token,
		}
		ctx.JSON(lib.CommonResponse(200, "Login Success", context))
	} else {
		ctx.JSON(lib.CommonResponse(500, "Username / Password ERROR", nil))
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
