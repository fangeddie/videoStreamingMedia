package dbops

import (
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
)

var (
	dbConn *sql.DB
	err    error
)

func init() {
	dbConn, err = sql.Open("mysql", "root:123456@tcp(localhost:3306)/vid_stre_med?charset=utf8")
	if err != nil {
		panic(err.Error())
	}

	err = dbConn.Ping()
	if err != nil {
		panic(err.Error())
	}
}

//func openConn() *sql.DB {
//	dbConn, err := sql.Open("mysql", "root:123456@#@tcp(localhost:3306)/vid_stre_med?charset=utf")
//	if err != nil {
//		panic(err.Error())
//	}
//	defer dbConn.Close()
//
//	return dbConn
//}
