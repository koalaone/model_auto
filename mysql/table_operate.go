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
