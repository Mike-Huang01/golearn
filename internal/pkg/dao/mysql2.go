package dao

import (
	"database/sql"
	"errors"
	"fmt"
	"os"
)

var sqlDB *sql.DB

func init() {
	// 密码在环境变量里
	var ps string
	if ps = os.Getenv("password"); ps == "" {
		panic(errors.New("env no password"))
	}

	var err error
	sqlDB, err = sql.Open("mysql", fmt.Sprintf("root:%s@tcp(127.0.0.1:3306)/%s?charset=utf8mb4&parseTime=True&loc=Local", ps, "test"))
	if err != nil {
		panic(err)
	}
}

func GetSqlDB() *sql.DB {
	return sqlDB
}
