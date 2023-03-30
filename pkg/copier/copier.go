package copier

import (
	"fmt"
	"reflect"
)

type srcStruct struct {
	fieldName  string
	fieldType  reflect.Type
	fieldValue reflect.Value
}

// Copier copy source struct to destination struct by tag name
func Copier(src any, dst any, tagString string) error {
	srcObjValue := reflect.ValueOf(src)
	srcObjType := reflect.TypeOf(src)

	dstObjValue := reflect.ValueOf(dst)
	dstObjType := reflect.TypeOf(dst)
	_ = dstObjType
	if srcObjValue.Kind() != reflect.Struct {
		return fmt.Errorf("type of src must be struct but got %s", srcObjType.Kind())
	}

	if dstObjValue.Kind() != reflect.Pointer {
		return fmt.Errorf("type of dst must be struct but got %s", srcObjType.Kind())
	}

	srcMap := map[string]srcStruct{}

	for i := 0; i < srcObjValue.NumField(); i++ {
		element := srcObjType.Field(i)
		srcObjValue.Field(i).Interface()
		srcMap[element.Tag.Get(tagString)] = srcStruct{
			fieldName:  element.Name,
			fieldType:  srcObjValue.Field(i).Type(),
			fieldValue: srcObjValue.Field(i),
		}
	}

	dstIndVal := reflect.Indirect(dstObjValue)

	for i := 0; i < dstIndVal.NumField(); i++ {
		elem := dstIndVal.Field(i)
		if elem.CanSet() {
			tName := dstIndVal.Type().Field(i).Tag.Get(tagString)
			val := srcMap[tName]
			if elem.Type() == val.fieldType {
				elem.Set(val.fieldValue)
			}
		}

	}

	return nil
}
