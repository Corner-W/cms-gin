package collection

import (
	"errors"
	"fmt"
	"reflect"

	"golang.org/x/exp/constraints"
)

func GetStructSliceFields[T constraints.Ordered, STRUCT any](input []STRUCT, key string) (values []T, err error) {
	for _, item := range input {
		field, err2 := GetStructField[T, STRUCT](item, key)
		if err2 != nil {
			return values, err2
		}
		values = append(values, field)
	}
	return
}

func GetStructField[T constraints.Ordered, STRUCT any](input STRUCT, key string) (value T, err error) {
	rv := reflect.ValueOf(input)
	switch rv.Kind() {
	case reflect.Ptr:
		rv = rv.Elem()
	case reflect.Struct:
	default:
		return value, errors.New("input must be struct or ptr")
	}

	keyExist := false
	for i := 0; i < rv.NumField(); i++ {
		fieldValue := rv.Field(i)
		fieldType := rv.Type().Field(i)

		if fieldType.Name == key {
			switch fieldValue.Kind() {
			case reflect.String, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Float64, reflect.Float32:
				keyExist = true
				if k, ok := fieldValue.Interface().(T); ok {
					value = k
				} else {
					return value, errors.New("can't convert slice key")
				}
			default:
				return value, errors.New("key must be int float or string")
			}
		}
	}
	if !keyExist {
		return value, fmt.Errorf("key %s not found in %s's field", key, rv)
	}
	return
}

func StructSliceToMap[T constraints.Ordered, STRUCT any](input []STRUCT, key string) (values map[T]STRUCT, err error) {
	values = make(map[T]STRUCT)
	for _, item := range input {
		fieldMap, err2 := StructToMap[T, STRUCT](item, key)
		if err2 != nil {
			return values, err2
		}

		for i, v := range fieldMap {
			values[i] = v
		}
	}
	return
}

func StructToMap[T constraints.Ordered, STRUCT any](input STRUCT, key string) (value map[T]STRUCT, err error) {
	rv := reflect.ValueOf(input)
	switch rv.Kind() {
	case reflect.Ptr:
		rv = rv.Elem()
	case reflect.Struct:
	default:
		return value, errors.New("input must be struct or ptr")
	}

	value = make(map[T]STRUCT)
	keyExist := false
	for i := 0; i < rv.NumField(); i++ {
		fieldValue := rv.Field(i)
		fieldType := rv.Type().Field(i)

		if fieldType.Name == key {
			switch fieldValue.Kind() {
			case reflect.String, reflect.Int64, reflect.Int32, reflect.Int16, reflect.Int8, reflect.Int, reflect.Float64, reflect.Float32:
				keyExist = true
				if k, ok := fieldValue.Interface().(T); ok {
					value[k] = input
				} else {
					return value, errors.New("can't convert slice key")
				}
			default:
				return value, errors.New("key must be int float or string")
			}
		}
	}

	if !keyExist {
		return value, fmt.Errorf("key %s not found in %s's field", key, rv)
	}
	return
}
