/*
 *
 *
 *  * Copyright 2019 koalaone@163.com
 *  *
 *  * Licensed under the Apache License, Version 2.0 (the "License");
 *  * you may not use this file except in compliance with the License.
 *  * You may obtain a copy of the License at
 *  *
 *  *       http://www.apache.org/licenses/LICENSE-2.0
 *  *
 *  * Unless required by applicable law or agreed to in writing, software
 *  * distributed under the License is distributed on an "AS IS" BASIS,
 *  * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 *  * See the License for the specific language governing permissions and
 *  * limitations under the License.
 *
 */

package main

import (
	"html/template"
	"strings"
)

func Tags(columnName string) template.HTML {
	return template.HTML("`json:" + `"` + columnName + "\"`")
}

func Unescaped(x string) interface{} {
	return template.HTML(x)
}

func IsUUID(str string) bool {
	return "uuid" == str
}

func FirstCharLower(str string) string {
	if len(str) > 0 {
		return strings.ToLower(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

func FirstCharUpper(str string) string {
	if len(str) > 0 {
		return strings.ToUpper(str[0:1]) + str[1:]
	} else {
		return ""
	}
}

func ExportColumn(columnName string) string {
	columnItems := strings.Split(columnName, "_")
	columnItems[0] = FirstCharUpper(columnItems[0])
	for i := 0; i < len(columnItems); i++ {
		item := strings.Title(columnItems[i])

		if strings.ToUpper(item) == "ID" {
			item = "ID"
		}

		columnItems[i] = item
	}

	return strings.Join(columnItems, "")

}

func TypeConvert(str string) string {

	switch str {
	case "smallint", "tinyint":
		return "int8"

	case "varchar", "text", "longtext", "char":
		return "string"

	case "date":
		return "string"

	case "int":
		return "int"

	case "timestamp", "datetime":
		return "time.Time"

	case "bigint":
		return "int" // "int64"

	case "float", "double":
		return "float64"

	case "decimal":
		return "decimal.Decimal"

	case "uuid":
		return "gocql.UUID"

	default:
		return str
	}
}

func Join(a []string, sep string) string {
	return strings.Join(a, sep)
}

func ColumnAndType(table_schema []TableSchema) string {
	result := make([]string, 0, len(table_schema))
	for _, t := range table_schema {
		result = append(result, t.ColumnName+" "+TypeConvert(t.DataType))
	}
	return strings.Join(result, ",")
}

func ColumnWithPostfix(columns []string, Postfix, sep string) string {
	result := make([]string, 0, len(columns))
	for _, t := range columns {
		result = append(result, t+Postfix)
	}
	return strings.Join(result, sep)
}

func MakeQuestionMarkList(num int) string {
	a := strings.Repeat("?,", num)
	return a[:len(a)-1]
}

func HumpStructName(tableName string) string {
	humps := strings.Split(tableName, "_")
	value := ""
	for _, item := range humps {
		value = value + strings.Title(item)
	}

	if value[len(value)-1:] == "s" {
		value = value[:len(value)-1]
	}

	return value
}
