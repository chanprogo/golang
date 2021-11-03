package model

import (
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/go-xorm/xorm"
)

func NewMysqlOrm(user, password, host, port, database string,
	maxOpen, maxIdle int) {

	var err error
	conn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8", user, password, host, port, database)
	Orm, err = xorm.NewEngine("mysql", conn)
	if err != nil {
		log.Fatal(err.Error())
	}

	if err := Orm.Ping(); err != nil {
		panic("数据库连接失败，err：" + err.Error())
	}

	Orm.SetMaxOpenConns(maxOpen)
	Orm.SetMaxIdleConns(maxIdle)

	// Orm.ShowSQL(true)

}
