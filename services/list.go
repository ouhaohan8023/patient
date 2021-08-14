package services

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/ouhaohan8023/patientRegistration/database"
	"github.com/ouhaohan8023/patientRegistration/lib"
	"github.com/ouhaohan8023/patientRegistration/model"
	"strconv"
	"time"
)

func List(ctx iris.Context) {
	page := ctx.URLParamDefault("page", "1")
	pageSize := ctx.URLParamDefault("size", "1")
	name := ctx.URLParamDefault("name", "")
	phone := ctx.URLParamDefault("phone", "")
	email := ctx.URLParamDefault("email", "")
	status := ctx.URLParamDefault("status", "")
	orderKey := ctx.URLParamDefault("sortKey", "id")
	orderType := ctx.URLParamDefault("sortType", "desc")
	connect := database.ConnectMysql()
	var users []model.Users
	var totalUsers []model.Users
	var unTouchedUsers []model.Users
	total := connect
	if name != "" {
		total = total.Where("name LIKE ?", "%"+name+"%")
	}
	if phone != "" {
		total = total.Where("phone LIKE ?", "%"+phone+"%")
	}
	if email != "" {
		total = total.Where("email LIKE ?", "%"+email+"%")
	}
	if status != "" {
		total = total.Where("status = ?", status)
	}
	result := total.Scopes(Paginate(page, pageSize)).Order(orderKey + " " + orderType).Find(&users)
	unTouched := total.Where("status = 0").Find(&unTouchedUsers)
	total = total.Find(&totalUsers)
	c := iris.Map{
		"lists":     result.Value,
		"total":     total.RowsAffected,
		"untouched": unTouched.RowsAffected,
	}
	ctx.JSON(lib.CommonResponse(200, "success", c))
}

func Update(ctx iris.Context) {
	id := ctx.Params().Get("id")

	connect := database.ConnectMysql()
	connect.Model(&model.Users{}).Where("id = ?", id).Update(model.Users{UpdatedAt: time.Time{}, Status: 1})
	ctx.JSON(lib.CommonResponse(200, "success", nil))
}

func Paginate(page string, limit string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		page, _ := strconv.Atoi(page)
		if page == 0 {
			page = 1
		}

		pageSize, _ := strconv.Atoi(limit)
		switch {
		case pageSize > 100:
			pageSize = 100
		case pageSize <= 0:
			pageSize = 10
		}

		offset := (page - 1) * pageSize
		return db.Offset(offset).Limit(pageSize)
	}
}
