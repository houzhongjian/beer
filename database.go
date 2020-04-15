package beer

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"log"
	"strings"
)

type dbConfig struct {
	Host     string
	User     string
	Name     string
	Port     int
	Password string
}

type databaseManager struct {
	engine string
	gorm   *gorm.DB
}

var db *databaseManager

func init() {
	db = new(databaseManager)
	db.engine = getDbEngine()
}

func (dbSrv *databaseManager) connDb() {
	if db.engine == "gorm" {
		if err := db.gormConnDb(); err != nil {
			log.Printf("err123:%+v\n", err)
			return
		}
	}
}

//Gorm .
func Gorm() *gorm.DB {
	if db.gorm == nil {
		db.connDb()
	}
	return db.gorm
}

func getDbEngine() string {
	dbEng := strings.ToLower(Config().GetString("database_engine"))
	if len(dbEng) < 1 {
		return "gorm"
	}
	if dbEng != "gorm" {
		msg := "暂不支持 " + dbEng + "这个数据库引擎"
		panic(msg)
	}
	return dbEng
}

func (dbSrv *databaseManager) gormConnDb() error {
	cf := dbConfig{
		Host:     conf.GetString("db_host"),
		User:     conf.GetString("db_user"),
		Name:     conf.GetString("db_name"),
		Port:     conf.GetInt("db_port"),
		Password: conf.GetString("db_password"),
	}

	args := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8&parseTime=True&loc=Local",
		cf.User,
		cf.Password,
		cf.Host,
		cf.Port,
		cf.Name,
	)
	dbConn, err := gorm.Open("mysql", args)
	if err != nil {
		return err
	}

	dbSrv.gorm = dbConn
	dbSrv.SetMaxOpenConns(10000)

	return nil
}

func (dbSrv *databaseManager) SetMaxOpenConns(n int) {
	dbSrv.gorm.DB().SetMaxOpenConns(n)
}
