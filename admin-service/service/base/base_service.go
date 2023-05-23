package base

import (
	"admin-service/entity"
	"admin-service/global"
	"encoding/json"
	"gorm.io/gorm"
	"reflect"
	"strings"
)

type BaseService[T any] struct{}

func (m *BaseService[T]) Insert(model *T) (err error) {
	result := global.GCS_DB.Create(&model)
	if result.Error != nil {
		err = result.Error
	}
	return err
}

func (m *BaseService[T]) Delete(model *T) (err error) {
	result := global.GCS_DB.Delete(model)
	if result.Error != nil {
		err = result.Error
	}
	return err
}

func (m *BaseService[T]) Update(model *T) (err error) {
	result := global.GCS_DB.Save(model)
	if result.Error != nil {
		err = result.Error
	}
	return err
}

func (m *BaseService[T]) FindList(model *T) (data *[]T, totalCount int64, err error) {
	data = new([]T)
	props := reflect.ValueOf(*model).FieldByName("ExpandProps").Interface().(map[string]interface{})
	pageValue, p := props["page"].(float64)
	pageSizeValue, ps := props["pageSize"].(float64)
	searchConditions := props["searchCondition"]
	preload, pl := props["preload"].(string)
	db := m.SetSearchCondition(global.GCS_DB, searchConditions)
	if pl {
		db = db.Preload(preload)
	}
	var result *gorm.DB
	if p && ps {
		page := int(pageValue)
		pageSize := int(pageSizeValue)
		result = db.Limit(pageSize).Offset((page-1)*pageSize).Find(data, model).Offset(-1).Limit(-1).Count(&totalCount)
		//result = global.GCS_DB.Limit(pageSize).Offset((page-1)*pageSize).Find(data, model).Offset(-1).Limit(-1).Count(&totalCount)
	} else {
		result = db.Find(data, model).Count(&totalCount)
		//result = global.GCS_DB.Find(data, model).Count(&totalCount)
	}
	if result.Error != nil {
		return nil, 0, result.Error
	}
	return data, totalCount, nil
}

func (m *BaseService[T]) FindOne(model *T) (res *T, err error) {
	result := global.GCS_DB.First(&res, model)
	if result.Error != nil {
		return nil, result.Error
	}
	return res, nil
}

func (m *BaseService[T]) SetSearchCondition(db *gorm.DB, searchConditions interface{}) (tx *gorm.DB) {
	var conditions []entity.SearchCondition
	if jsonData, err := json.Marshal(searchConditions); err == nil {
		if err := json.Unmarshal([]byte(jsonData), &conditions); err == nil {
			for _, c := range conditions {
				switch c.Condition {
				case entity.Equal:
					if strings.TrimSpace(c.Value) != "" {
						db = db.Where(c.Field+" = ?", c.Value)
					}
				case entity.NotEqual:
					if strings.TrimSpace(c.Value) != "" {
						db = db.Where(c.Field+" != ?", c.Value)
					}
				case entity.GreaterThan:
					if strings.TrimSpace(c.Value) != "" {
						db = db.Where(c.Field+" < ?", c.Value)

					}
				case entity.GreaterThanOrEqual:
					if strings.TrimSpace(c.Value) != "" {
						db = db.Where(c.Field+" <= ?", c.Value)
					}
				case entity.LessThan:
					if strings.TrimSpace(c.Value) != "" {
						db = db.Where(c.Field+" > ?", c.Value)
					}
				case entity.LessThanOrEqual:
					if strings.TrimSpace(c.Value) != "" {
						db = db.Where(c.Field+" >= ?", c.Value)
					}
				case entity.Like:
					if strings.TrimSpace(c.Value) != "" {
						db = db.Where(c.Field+" like ?", "%"+c.Value+"%")
					}
				}
			}
		}
	}
	return db
}
