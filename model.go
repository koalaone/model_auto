package main

import (
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"path"
	"path/filepath"
	"strings"
	"text/template"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var instance *gorm.DB
var instanceMysql *gorm.DB

var DbName string

func DBInit(dbName, user, password, host, port string) (err error) {
	conn_str := user + ":" + password + "@tcp(" + host + ":" + port + ")/" + dbName + "?charset=utf8mb4&parseTime=True"
	instance, err = gorm.Open("mysql", conn_str)
	if err != nil {
		return err
	}

	mysql_conn_str := user + ":" + password + "@tcp(" + host + ":" + port + ")/information_schema?charset=utf8mb4&parseTime=True"
	instanceMysql, err = gorm.Open("mysql", mysql_conn_str)
	if err != nil {
		return err
	}

	DbName = dbName

	return nil
}

func ModelGenerate(importName, tableName string) error {
	type TableInfo struct {
		TableName string `json:"table_name"`
	}
	var tablaNames []TableInfo
	err := instanceMysql.Raw(`SELECT table_name from tables where table_schema = ? `, DbName).Find(&tablaNames).Error
	if err != nil {
		return err
	}

	modelData, err := ioutil.ReadFile("./model.tmpl")
	if err != nil {
		fmt.Println("read tplFile err:", err)
		return err
	}

	render := template.Must(template.New("model").
		Funcs(template.FuncMap{
			"FirstCharUpper":       FirstCharUpper,
			"TypeConvert":          TypeConvert,
			"Tags":                 Tags,
			"ExportColumn":         ExportColumn,
			"Join":                 Join,
			"MakeQuestionMarkList": MakeQuestionMarkList,
			"ColumnAndType":        ColumnAndType,
			"ColumnWithPostfix":    ColumnWithPostfix,
		}).Parse(string(modelData)))

	for _, table := range tablaNames {
		if (tableName == "") || (tableName != "" && tableName == table.TableName) {
			err := genModelFile(render, importName, table.TableName)
			if err != nil {
				fmt.Println("genModelFile err:", err)
				return err
			}
		}
	}

	return nil
}

type ModelInfo struct {
	BDName          string
	DBConnection    string
	TableName       string
	ExportModelName string
	ImportName      string
	PackageName     string
	ModelName       string
	TableSchema     *[]TableSchema
}

type TableSchema struct {
	ColumnName    string `db:"column_name" json:"column_name"`
	DataType      string `db:"data_type" json:"data_type"`
	ColumnKey     string `db:"column_key" json:"column_key"`
	ColumnComment string `db:"column_comment", json:"column_comment"`
}

func genModelFile(render *template.Template, importName, tableName string) error {
	if tableName == "" {
		return nil
	}

	var tableSchema []TableSchema
	err := instanceMysql.Raw(`SELECT column_name, data_type,column_key,column_comment from COLUMNS `+
		` where TABLE_NAME= ? and table_schema = ? `, tableName, DbName).Find(&tableSchema).Error
	if err != nil {
		fmt.Println(err)
		return err
	}

	if len(tableSchema) <= 0 {
		fmt.Println(tableName, "tableSchema ["+tableName+"] is null")
		return errors.New("tableSchema [" + tableName + "] is null")
	}

	packageName := "local"
	if importName != "" {
		packageName = path.Base(importName)
	}

	dirPath := path.Base("") + string(filepath.Separator) + packageName + string(filepath.Separator)
	if !IsExist(dirPath) {
		err = os.Mkdir(dirPath, os.ModePerm)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	fileName := path.Base("") + string(filepath.Separator) + packageName + string(filepath.Separator) +
		strings.ToLower(tableName) + "_auto.go"
	if IsExist(fileName) {
		err = os.Remove(fileName)
		if err != nil {
			fmt.Println(err)
			return err
		}
	}

	f, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		fmt.Println("os.Create error:", err.Error())
		return err
	}
	defer f.Close()

	model := &ModelInfo{
		ImportName:      importName,
		PackageName:     packageName,
		BDName:          DbName,
		ExportModelName: HumpStructName(tableName),
		TableName:       tableName,
		ModelName:       tableName,
		TableSchema:     &tableSchema,
	}

	if err := render.Execute(f, model); err != nil {
		log.Fatal(err)
	}
	cmd := exec.Command("goimports", "-w", fileName)
	err = cmd.Run()
	if err != nil {
		fmt.Println("format go code error:", err.Error())
		return err
	}

	return nil
}

//判断是否存在文件或者目录
func IsExist(path string) bool {
	_, err := os.Stat(path)
	return err == nil || os.IsExist(err)
}
