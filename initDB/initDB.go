package initDB

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"log"
)

var Db *sql.DB

func init()  {
	var err error
	log.Print("数据库初始化")
	Db, err = sql.Open("mysql", "root:tan1992527@tcp(127.0.0.1:3306)/ginhello")
	if err != nil {
		log.Panicln("err:",err.Error())
	}
	Db.SetMaxOpenConns(10)
	Db.SetMaxIdleConns(10)
}



