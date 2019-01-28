package mysql

import (
	"database/sql"
	"errors"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	sqlLogger = log.New(os.Stderr, "\r\n[sql]", 0)

	instance *gorm.DB

	sqlMode bool
)

func InitConnDB(dbConn string, dbType string) error {
	if dbConn == "" {
		return errors.New("param dbConn empty")
	}

	if dbType == "" {
		return errors.New("param dbType empty")
	}

	var err error
	instance, err = gorm.Open(dbConn, dbType)
	if err != nil {
		return err
	}

	return nil
}

//SetLogger set log writer
func SetLogger(logger *log.Logger) {
	sqlLogger = logger
}

//SetSQLMode indicate whether log sql or not
func SetSQLMode(enable bool) {
	sqlMode = enable
}

func GetDB() *gorm.DB {
	return instance
}

func Ping() error {
	err := instance.DB().Ping()
	if err != nil {
		sqlLogger.Fatalf("Error connecting to the database: %s\n", err)
		return err
	}

	return nil
}

func QueryRows(sqlStr string) (*sql.Rows, error) {
	go writeLog(sqlStr)

	return instance.DB().Query(sqlStr)
}

//QueryRow excute sql and return row
func QueryRow(sqlStr string) (*sql.Row, error) {
	go writeLog(sqlStr)

	return instance.DB().QueryRow(sqlStr), nil
}

func WriteError(sqlStr string, err error) {
	go writeLog(sqlStr+", error:%v", err.Error())
}

func WriteLog(sqlStr string, args ...interface{}) {
	go writeLog(sqlStr, args)
}

func writeLog(query string, args ...interface{}) {
	if !sqlMode {
		return
	}
	if len(args) == 0 {
		sqlLogger.Println(query)
	} else {
		sqlLogger.Printf(query, args...)
	}
}
