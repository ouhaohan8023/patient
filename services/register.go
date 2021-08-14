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

type Input struct {
	Name        string `json:"username"`
	Address     string `json:"address"`
	Appointment string `json:"appointment"`
	Birthday    string `json:"birthday"`
	Email       string `json:"email"`
	Img         string `json:"img"`
	Phone       string `json:"mobile"`
	Photo       string `json:"photo"`
}

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
		ctx.JSON(lib.CommonResponse(500, "Submit Too Quickly", nil))
	} else {
		add := nums + 1
		err = client.Set(cct, ip, add, 60*time.Second).Err()
		if err != nil {
			panic(err)
		}

		c := &Input{}

		if err := ctx.ReadJSON(c); err != nil {
			panic(err.Error())
		} else {
			now := time.Now()
			M := model.Users{
				Name:        c.Name,
				Birthday:    c.Birthday,
				Phone:       c.Phone,
				Email:       c.Email,
				Photo:       c.Photo,
				Appointment: c.Appointment,
				Address:     c.Address,
				CreatedAt:   now,
				UpdatedAt:   now,
			}
			validate := validator.New()
			err := validate.Struct(M)
			if err != nil {
				for _, err := range err.(validator.ValidationErrors) {
					ctx.JSON(lib.CommonResponse(500, "submit error", iris.Map{
						"msg": err.Field() + " is " + err.Tag(),
					}))
					break
				}
				return
			}
			isUsPhone := lib.ValidateUsPhoneNumber(c.Phone)
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
	environment := viper.GetString(`environment`)

	uploadPath := viper.GetString(environment + `.uploadPath`)
	_, err = ctx.SaveFormFile(fh, filepath.Join(uploadPath, filename))
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.HTML("2Error while uploading: <b>" + err.Error() + "</b>")
		return
	}

	rtx := iris.Map{
		"url": viper.GetString(environment+`.baseUrl`) + filename,
	}
	ctx.JSON(lib.CommonResponse(200, "upload success", rtx))
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
