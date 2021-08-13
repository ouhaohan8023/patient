package services

import (
	"github.com/jinzhu/gorm"
	"github.com/kataras/iris/v12"
	"github.com/ouhaohan8023/patientRegistration/database"
	"github.com/ouhaohan8023/patientRegistration/lib"
	"github.com/ouhaohan8023/patientRegistration/model"
	"strconv"
)

func List(ctx iris.Context) {
	page := ctx.URLParamDefault("page", "1")
	pageSize := ctx.URLParamDefault("page-size", "1")
	name := ctx.URLParamDefault("name", "")
	phone := ctx.URLParamDefault("phone", "")
	connect := database.ConnectMysql()
	var users []model.Users
	var totalUsers []model.Users
	total := connect
	if name != "" {
		total = total.Where("name LIKE ?", "%" +name+ "%")
	}
	if phone != "" {
		total = total.Where("phone LIKE ?", "%" +phone+ "%")
	}
	result := total.Scopes(Paginate(page, pageSize)).Find(&users)
	total = total.Find(&totalUsers)
	c := iris.Map{
		"lists": result.Value,
		"total": total.RowsAffected,
	}
	ctx.JSON(lib.CommonResponse(200, "success", c))
}

func Paginate(page string, limit string) func(db *gorm.DB) *gorm.DB {
	return func (db *gorm.DB) *gorm.DB {
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
