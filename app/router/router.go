package router

import (
	"blogy/app/controllers"
	"github.com/gin-gonic/gin"
	"github.com/jinzhu/gorm"
)

func Route(r *gin.Engine, db *gorm.DB) {
	home := controllers.Home{}
	home.DB = db
	r.GET("/", home.Index)
}
