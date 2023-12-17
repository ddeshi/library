package database

import (
	"fmt"
	"github.com/ddeshi/library/model"
	"github.com/ddeshi/library/pkg/config"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"github.com/sirupsen/logrus"
	"log"
)

var db *gorm.DB

func DBInit() error {
	logrus.Info("DB init")
	url := fmt.Sprintf("%s:%s@tcp(%s)/%s?charset=utf8&parseTime=True&loc=Local",
		config.DBCfg.UserName,
		config.DBCfg.Password,
		config.DBCfg.Host,
		config.DBCfg.Name)

	var err error
	//var db *gorm.DB
	db, err = gorm.Open(config.DBCfg.Type, url)
	if err != nil {
		log.Fatalf("models.Setup err: %v", err)
	}

	//gorm.DefaultTableNameHandler = func(db *gorm.DB, defaultTableName string) string {
	//	return config.DBCfg.TablePrefix + defaultTableName
	//}
	db.SingularTable(true)
	db.DB().SetMaxIdleConns(10)
	db.DB().SetMaxOpenConns(100)

	if !db.HasTable(&model.User{}) {
		reader := model.User{}
		logrus.Info("database has not table %s, try to create it", reader.TableName())
		if err != nil {
			logrus.Info("create new table error: %v", err)
			return fmt.Errorf("create new table error: %v", err)
		}
	}

	return nil
}
