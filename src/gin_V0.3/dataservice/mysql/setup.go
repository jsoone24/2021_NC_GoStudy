package mysql

import (
	"fmt"

	"gin_V0.3/domain/model"
	"github.com/jinzhu/gorm"
)

const (
	dbms = "mysql"
	user = "root"
	pass = "password"
	db   = "webtuto"
)

var dbConn *gorm.DB

func Setup() {
	var err error
	conn := fmt.Sprintf("%s:%s@/%s?parseTime=true", user, pass, db)
	dbConn, err = gorm.Open(dbms, conn)
	if err != nil {
		panic(err)
	}

	dbConn.AutoMigrate(
		&model.Article{},
		&model.User{},
	)
	dbConn.Model(&model.Article{}).AddForeignKey("writer_id", "users(id)", "CASCADE", "CASCADE")
}
