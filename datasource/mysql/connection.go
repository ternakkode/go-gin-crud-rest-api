package mysql

import (
	"database/sql"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/ternakkode/go-gin-crud-rest-api/config"
)

var Mysql *sql.DB

func init() {
	var err error

	Mysql, err = sql.Open(config.Conf.DBDriver, config.Conf.DBSource)
	if err != nil {
		panic(err)
	}

	if err := Mysql.Ping(); err != nil {
		panic(err)
	}

	log.Println("database successfully configured")
}
