package common

import (
	"fmt"
	"reflect"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
)

var db *sqlx.DB

func SqlxInit(yml Yaml) (*sqlx.DB, error) {

	var err error
	if yml.Kind == "mysql" {
		db, err = sqlx.Connect(yml.Kind, yml.Mysql.ConnectionUri)

		// https://github.com/jmoiron/sqlx/issues/300
		db.DB.SetMaxOpenConns(yml.Mysql.SetMaxOpenConns)       // The default is 0 (unlimited)
		db.DB.SetMaxIdleConns(yml.Mysql.SetMaxIdleConns)       // defaultMaxIdleConns = 2
		db.DB.SetConnMaxLifetime(yml.Mysql.SetConnMaxLifetime) // 0, connections are reused forever.

		if err != nil {
			t := reflect.ValueOf(err).Elem()
			typ := t.Type()
			fmt.Println("====================================================")
			fmt.Printf("error type : [%s]\n", typ.Name())
			fmt.Println("====================================================")
		}

	} else if yml.Kind == "mysql" {
		db, err = sqlx.Connect(yml.Kind, yml.Mssql.Server)

		// https://github.com/jmoiron/sqlx/issues/300
		db.DB.SetMaxOpenConns(yml.Mssql.SetMaxOpenConns)       // The default is 0 (unlimited)
		db.DB.SetMaxIdleConns(yml.Mssql.SetMaxIdleConns)       // defaultMaxIdleConns = 2
		db.DB.SetConnMaxLifetime(yml.Mssql.SetConnMaxLifetime) // 0, connections are reused forever.

		if err != nil {
			t := reflect.ValueOf(err).Elem()
			typ := t.Type()
			fmt.Println("====================================================")
			fmt.Printf("error type : [%s]\n", typ.Name())
			fmt.Println("====================================================")
		}
	} else {
		//ToDo : error
	}

	return db, err
}

func SqlxClose() {
	if db != nil {
		db.Close()
	}
}

func SqlxConnect() *sqlx.DB {
	return db
}
