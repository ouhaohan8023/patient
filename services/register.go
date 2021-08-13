package services

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"github.com/kataras/iris/v12"
	"github.com/ouhaohan8023/patientRegistration/database"
	"github.com/ouhaohan8023/patientRegistration/lib"
	"github.com/ouhaohan8023/patientRegistration/model"
	"github.com/spf13/viper"
	"math/rand"
	"path/filepath"
	"strconv"
	"time"
)

func Submit(ctx iris.Context) {
	r := ctx.Request()
	ip, _ := lib.GetIP(r)

	client := database.ConnectRedis()
	cct := client.Context()
	val, err := client.Get(cct, ip).Result()
	if err != nil {
		fmt.Println(ip + " not exist")
		err = client.Set(cct, ip, 1, 3600*time.Second).Err()
		if err != nil {
			panic(err)
		}
	}
	nums, err := strconv.Atoi(val)

	requestLimitStr := viper.GetString(`requestLimit`)
	requestLimit, err := strconv.Atoi(requestLimitStr)
	if nums >= requestLimit {
		ctx.JSON(lib.CommonResponse(500, "submit too quickly", nil))
	} else {
		add := nums + 1
		err = client.Set(cct, ip, add, 60*time.Second).Err()
		if err != nil {
			panic(err)
		}

		name := ctx.PostValue("name")
		birthday := ctx.PostValue("birthday")
		phone := ctx.PostValue("phone")
		email := ctx.PostValue("email")
		photo := ctx.PostValue("photo")
		appointment := ctx.PostValue("appointment")
		currentTime := time.Now()
		M := model.Users{
			Name:        name,
			Birthday:    birthday,
			Phone:       phone,
			Email:       email,
			Photo:       photo,
			Appointment: appointment,
			CreatedAt:   currentTime,
			UpdatedAt:   currentTime,
		}
		validate := validator.New()
		err := validate.Struct(M)
		if err != nil {
			for _, err := range err.(validator.ValidationErrors) {
				ctx.JSON(lib.CommonResponse(500, "submit error", iris.Map{
					"key": err.Field(),
					"msg": err.Tag(),
				}))
			}
			return
		}
		isUsPhone := lib.ValidateUsPhoneNumber(phone)
		if !isUsPhone {
			ctx.JSON(lib.CommonResponse(500, "not validate phone", nil))
		} else {
			connect := database.ConnectMysql()
			create := connect.Create(&M)
			ctx.JSON(iris.Map{
				"status":  200,
				"message": "submit success",
				"content": create.Value,
			})
		}

	}
}

func Upload(ctx iris.Context) {
	f, fh, err := ctx.FormFile("file")
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("1Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	defer f.Close()

	filename := GetRand(10) + "_" + fh.Filename
	_, err = ctx.SaveFormFile(fh, filepath.Join("./uploads", filename))
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("2Error while uploading: <b>" + err.Error() + "</b>")
		return
	}
	ctx.JSON(lib.CommonResponse(200, "upload success", nil))
}

func GetRand(n int) string {
	rand.Seed(time.Now().UnixNano())
	letterRunes := []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}
