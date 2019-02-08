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

package mysql

import "errors"

type valueCount struct {
	CountValue int
}

func SearchTableCount(tableName, countField, wheres string) (int, error) {
	if tableName == "" {
		return 0, errors.New("param[tableName] is empty")
	}
	if countField == "" {
		return 0, errors.New("param[countField] is empty")
	}

	sqlStr := `select count(` + countField + `) as count_value from ` + tableName
	if wheres != "" {
		sqlStr = sqlStr + ` where ` + wheres
	}

	dbr := GetDB()
	dbr = dbr.Raw(sqlStr)

	var out valueCount
	err := dbr.Find(&out).Error
	if err != nil {
		return 0, err
	}

	return out.CountValue, nil
}

type valueSum struct {
	SumValue float64 `json:"sum_value"`
}

func SearchTableSum(tableName, countField, wheres string) (float64, error) {
	if tableName == "" {
		return 0, errors.New("param[tableName] is empty")
	}
	if countField == "" {
		return 0, errors.New("param[countField] is empty")
	}

	sqlStr := `select sum(` + countField + `) as sum_value from ` + tableName
	if wheres != "" {
		sqlStr = sqlStr + ` where ` + wheres
	}

	dbr := GetDB()
	dbr = dbr.Raw(sqlStr)

	var out valueSum
	err := dbr.Find(&out).Error
	if err != nil {
		return 0, err
	}

	return out.SumValue, nil
}
