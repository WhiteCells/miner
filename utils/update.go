package utils

import (
	"fmt"
	"reflect"
)

// 使用反射修改对象
func UpdateStructObjFromMap(obj any, updates map[string]any) error {
	// 递归更新字段的函数
	var updateFields func(reflect.Value, map[string]any) error

	updateFields = func(val reflect.Value, updates map[string]any) error {
		typ := val.Type()
		for key, value := range updates {
			for i := 0; i < val.NumField(); i++ {
				tagVal := typ.Field(i).Tag.Get("json")
				if tagVal == key { // 匹配 JSON 标签
					fieldValue := val.Field(i)
					if fieldValue.Kind() == reflect.Struct {
						// 如果字段是结构体，则递归更新
						if subUpdates, ok := value.(map[string]any); ok {
							err := updateFields(fieldValue, subUpdates)
							if err != nil {
								return err
							}
						} else {
							return fmt.Errorf("field %s expects a nested object", key)
						}
					} else {
						// 更新普通字段
						if fieldValue.CanSet() {
							fieldValue.Set(reflect.ValueOf(value))
						} else {
							return fmt.Errorf("field %s cannot be set", key)
						}
					}
					break
				}
			}
		}
		return nil
	}

	return updateFields(reflect.ValueOf(obj).Elem(), updates)
}

// func UpdateStructObjFromStruct(obj any, from any) {

// }
