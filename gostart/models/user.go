package models

import (
	"github.com/astaxie/beego/orm"
	_ "github.com/go-sql-driver/mysql" // import your used driver
	"fmt"
	"github.com/astaxie/beego/config"
)

// Model Struct
type User struct {
	Id   int
	Name string `orm:"size(100)"`
}

func init() {

	// 读取配置文件
	cfg, err := config.NewConfig("ini", "conf/app.conf")
	if err != nil {
		panic(err)
	}
	dbUser := cfg.String("mysql::user")
	dbPwd := cfg.String("mysql::password")
	dbHost := cfg.String("mysql::host")
	dbPort := cfg.String("mysql::port")
	dbName := cfg.String("mysql::dbname")
	// root:root@tcp(127.0.0.1:3306)/golang?charset=utf8
	dbUrl := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", dbUser, dbPwd, dbHost, dbPort, dbName)

	// set default database
	err = orm.RegisterDataBase("default", "mysql", dbUrl, 30, 30)
	if err != nil {
		panic(err)
	}

	// register model
	orm.RegisterModel(new(User))
	// create table
	orm.RunSyncdb("default", false, true)
}



