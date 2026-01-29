package connection

import (
	"fmt"
	"log"
	"net/url"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"shellrean.id/back-end/internal/config"
)

func GetDatabase(conf config.Database) *gorm.DB {
	tz := url.QueryEscape(conf.Tz)

	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=%s",
		conf.User,
		conf.Pass,
		conf.Host,
		conf.Port,
		conf.Name,
		tz,
	)

	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("failed to connect database: ", err.Error())
	}

	return db
}
