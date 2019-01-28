package mysql

import (
	"errors"
	"strings"
)

func SearchObject(objectPtr interface{}, wheres map[string]interface{}, out interface{}) error {
	if objectPtr == nil {
		return errors.New("param[objectPtr] is empty")
	}
	if out == nil {
		return errors.New("param[outArr] is empty")
	}

	dbr := GetDB().Model(objectPtr)
	if len(wheres) > 0 {
		dbr = dbr.Where(wheres)
	}

	err := dbr.Find(out).Error
	if err != nil {
		return err
	}

	return nil
}

func SearchObjectByIn(objectPtr interface{}, wheres map[string]interface{},
	ins map[string]interface{}, out interface{}) error {

	if objectPtr == nil {
		return errors.New("param[objectPtr] is empty")
	}
	if out == nil {
		return errors.New("param[outArr] is empty")
	}

	dbr := GetDB().Model(objectPtr)
	if len(wheres) > 0 {
		dbr = dbr.Where(wheres)
	}

	if len(ins) > 0 {
		for key, value := range ins {
			dbr = dbr.Where(key, value)
		}
	}

	err := dbr.Find(out).Error
	if err != nil {
		return err
	}

	return nil
}

func SearchObjectByOrder(objectPtr interface{}, wheres map[string]interface{}, ins map[string]interface{},
	orders string, limit, offset int, out interface{}) error {

	if objectPtr == nil {
		return errors.New("param[objectPtr] is empty")
	}
	if out == nil {
		return errors.New("param[outArr] is empty")
	}

	dbr := GetDB().Model(objectPtr)
	if len(wheres) > 0 {
		dbr = dbr.Where(wheres)
	}

	if len(ins) > 0 {
		for key, value := range ins {
			dbr = dbr.Where(key, value)
		}
	}

	if orders != "" {
		orderList := strings.Split(orders, ",")
		if len(orderList) > 1 {
			for _, itemOrder := range orderList {
				dbr = dbr.Order(itemOrder)
			}
		} else {
			dbr = dbr.Order(orders)
		}
	}

	if limit > 0 {
		dbr = dbr.Limit(limit)
	}

	if offset > 0 {
		dbr = dbr.Offset(offset)
	}

	err := dbr.Find(out).Error
	if err != nil {
		return err
	}

	if err != nil {
		return err
	}

	return nil
}

func CreateObject(value interface{}) error {
	if value == nil {
		return errors.New("param[value] is nil")
	}

	err := GetDB().Save(value).Error
	if err != nil {
		return err
	}

	return nil
}

// updates maybe use gorm.Expr()
func UpdateObject(objectPtr interface{}, wheres, updates map[string]interface{}) error {
	if objectPtr == nil {
		return errors.New("param[objectPtr] is empty")
	}
	if len(wheres) == 0 {
		return errors.New("param[wheres] length is zero")
	}
	if len(updates) == 0 {
		return errors.New("param[updates] length is zero")
	}

	err := GetDB().Model(objectPtr).Where(wheres).UpdateColumns(updates).Error
	if err != nil {
		return err
	}

	return nil
}

// note: return check gorm.ErrRecordNotFound
func SearchObjectPreload(objectPtr interface{}, tableName string, wheres map[string]interface{},
	preloads []string, joins []string) error {

	if objectPtr == nil {
		return errors.New("param[objectPtr] is empty")
	}
	if tableName == "" {
		return errors.New("param[tableName] is empty")
	}

	dbr := GetDB().Table(tableName)
	if len(wheres) > 0 {
		dbr = dbr.Where(wheres)
	}

	if len(preloads) > 0 {
		for _, preload := range preloads {
			if preload == "" {
				continue
			}

			dbr = dbr.Preload(preload)
		}
	}

	if len(joins) > 0 {
		for _, join := range joins {
			if join == "" {
				continue
			}

			dbr = dbr.Joins(join)
		}
	}

	err := dbr.Find(objectPtr).Error
	if err != nil {
		return err
	}

	return nil
}

// note: return check gorm.ErrRecordNotFound
func SearchSelectPreload(objectPtr interface{}, tableName, selectValue string,
	wheres map[string]interface{}, preloads []string, joins []string) error {

	if objectPtr == nil {
		return errors.New("param[objectPtr] is empty")
	}
	if tableName == "" {
		return errors.New("param[tableName] is empty")
	}
	if selectValue == "" {
		return errors.New("param[selectValue] is empty")
	}

	dbr := GetDB().Table(tableName).Select(selectValue)
	if len(wheres) > 0 {
		dbr = dbr.Where(wheres)
	}

	if len(preloads) > 0 {
		for _, preload := range preloads {
			if preload == "" {
				continue
			}

			dbr = dbr.Preload(preload)
		}
	}

	err := dbr.Find(objectPtr).Error
	if err != nil {
		return err
	}

	return nil

}
