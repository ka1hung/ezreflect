package main

import (
	"fmt"
	"reflect"
	"strconv"
)

func wrapTypeOf(data interface{}) reflect.Type {
	var t reflect.Type
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		t = reflect.TypeOf(data)
	} else {
		t = reflect.TypeOf(data).Elem()
	}
	return t
}
func wrapValueOf(data interface{}) reflect.Value {
	var v reflect.Value
	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		v = reflect.ValueOf(data)
	} else {
		v = reflect.ValueOf(data).Elem()
	}
	return v
}

//GetFieldNames input a struct data then return the struct field names
func GetFieldNames(data interface{}) []string {
	t := wrapTypeOf(data)

	fs := []string{}
	num := t.NumField()

	for i := 0; i < num; i++ {
		fs = append(fs, t.Field(i).Name)
	}
	return fs
}

//GetFieldNames input a struct data then return the struct field types
func GetFieldTypes(data interface{}) []string {
	t := wrapTypeOf(data)
	fs := []string{}
	num := t.NumField()

	for i := 0; i < num; i++ {
		fs = append(fs, t.Field(i).Type.Kind().String())
	}
	return fs
}

//GetFieldTypesByMap input a struct data then return the map(key:FieldName val:FieldType) of struct field types.
func GetFieldTypesByMap(data interface{}) map[string]string {
	t := wrapTypeOf(data)
	m := map[string]string{}
	num := t.NumField()

	for i := 0; i < num; i++ {
		m[t.Field(i).Name] = t.Field(i).Type.Kind().String()
	}
	return m
}

//GetFieldData input a struct data then return the map(key:FieldName val:FieldValue) of data.
func GetFieldData(data interface{}) map[string]interface{} {
	t := wrapTypeOf(data)
	v := wrapValueOf(data)

	m := map[string]interface{}{}
	num := reflect.ValueOf(data).NumField()
	for i := 0; i < num; i++ {

		m[t.Field(i).Name] = v.Field(i).Interface()

	}
	return m
}

//GetFieldDataString input a struct data then return the map(key:FieldName val:FieldValue In String) of data.
func GetFieldDataString(data interface{}) map[string]string {
	t := wrapTypeOf(data)
	v := wrapValueOf(data)

	m := map[string]string{}
	num := reflect.ValueOf(data).NumField()
	for i := 0; i < num; i++ {
		s := fmt.Sprintf("%v", v.Field(i).Interface())
		m[t.Field(i).Name] = s

	}
	return m
}

//GetFieldTag input a struct data then return the map(key:FieldName val:FieldValue In String) of data.
func GetFieldTag(data interface{}) map[string]string {
	t := wrapTypeOf(data)

	m := map[string]string{}
	num := reflect.ValueOf(data).NumField()
	for i := 0; i < num; i++ {
		m[t.Field(i).Name] = fmt.Sprintf("%v", t.Field(i).Tag)
	}
	return m
}

//FieldCopyByNames input f=from and t=target then copy the fields by names(fs)
func FieldCopyByNames(f, t interface{}, fs []string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if reflect.TypeOf(f).Kind() != reflect.Ptr || reflect.TypeOf(t).Kind() != reflect.Ptr {
		return fmt.Errorf("input data should be a point")
	}

	from := reflect.ValueOf(f).Elem()
	to := reflect.ValueOf(t).Elem()
	for _, f := range fs {
		// fieldName := f
		t, ok := reflect.TypeOf(t).Elem().FieldByName(f)
		if !ok {
			// return fmt.Errorf("field not exist")
			continue
		}
		fieldType := t.Type.Kind()
		switch fieldType {
		case reflect.String:
			to.FieldByName(f).SetString(from.FieldByName(f).String())
		case reflect.Float64:
			to.FieldByName(f).SetFloat(from.FieldByName(f).Float())
		case reflect.Float32:
			to.FieldByName(f).SetFloat(from.FieldByName(f).Float())
		case reflect.Uint:
			to.FieldByName(f).SetUint(from.FieldByName(f).Uint())
		case reflect.Uint8:
			to.FieldByName(f).SetUint(from.FieldByName(f).Uint())
		case reflect.Uint16:
			to.FieldByName(f).SetUint(from.FieldByName(f).Uint())
		case reflect.Uint32:
			to.FieldByName(f).SetUint(from.FieldByName(f).Uint())
		case reflect.Uint64:
			to.FieldByName(f).SetUint(from.FieldByName(f).Uint())
		case reflect.Int:
			to.FieldByName(f).SetInt(from.FieldByName(f).Int())
		case reflect.Int8:
			to.FieldByName(f).SetInt(from.FieldByName(f).Int())
		case reflect.Int16:
			to.FieldByName(f).SetInt(from.FieldByName(f).Int())
		case reflect.Int32:
			to.FieldByName(f).SetInt(from.FieldByName(f).Int())
		case reflect.Int64:
			to.FieldByName(f).SetInt(from.FieldByName(f).Int())
		case reflect.Bool:
			to.FieldByName(f).SetBool(from.FieldByName(f).Bool())
		default:
			return fmt.Errorf(fieldType.String() + " type not support")
		}

	}
	return nil
}

//FieldParseString parse struct data from  map k=FieldName v=FieldValueString
func FieldParseFromString(data interface{}, fs map[string]string) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return fmt.Errorf("input data should be a point")
	}

	t := wrapTypeOf(data)
	v := wrapValueOf(data)
	for k, s := range fs {

		sf, ok := t.FieldByName(k)
		if !ok {
			continue
		}
		fieldType := sf.Type.Kind()
		switch fieldType {
		case reflect.String:
			v.FieldByName(k).SetString(s)
		case reflect.Float64:
			o, _ := strconv.ParseFloat(s, 64)
			v.FieldByName(k).SetFloat(o)
		case reflect.Float32:
			o, _ := strconv.ParseFloat(s, 32)
			v.FieldByName(k).SetFloat(o)
		case reflect.Uint:
			o, _ := strconv.ParseUint(s, 10, 32)
			v.FieldByName(k).SetUint(o)
		case reflect.Uint8:
			o, _ := strconv.ParseUint(s, 10, 8)
			v.FieldByName(k).SetUint(o)
		case reflect.Uint16:
			o, _ := strconv.ParseUint(s, 10, 16)
			v.FieldByName(k).SetUint(o)
		case reflect.Uint32:
			o, _ := strconv.ParseUint(s, 10, 32)
			v.FieldByName(k).SetUint(o)
		case reflect.Uint64:
			o, _ := strconv.ParseUint(s, 10, 64)
			v.FieldByName(k).SetUint(o)
		case reflect.Int:
			o, _ := strconv.ParseInt(s, 10, 32)
			v.FieldByName(k).SetInt(o)
		case reflect.Int8:
			o, _ := strconv.ParseInt(s, 10, 8)
			v.FieldByName(k).SetInt(o)
		case reflect.Int16:
			o, _ := strconv.ParseInt(s, 10, 16)
			v.FieldByName(k).SetInt(o)
		case reflect.Int32:
			o, _ := strconv.ParseInt(s, 10, 32)
			v.FieldByName(k).SetInt(o)
		case reflect.Int64:
			o, _ := strconv.ParseInt(s, 10, 64)
			v.FieldByName(k).SetInt(o)
		case reflect.Bool:
			o, _ := strconv.ParseBool(s)
			v.FieldByName(k).SetBool(o)
		default:
			return fmt.Errorf(fieldType.String() + " type not support")
		}

	}
	return nil
}

//FieldParse parse struct data from  map k=FieldName v=FieldValue
func FieldParse(data interface{}, fs map[string]interface{}) error {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()

	if reflect.TypeOf(data).Kind() != reflect.Ptr {
		return fmt.Errorf("input data should be a point")
	}

	t := wrapTypeOf(data)
	v := wrapValueOf(data)
	for k, s := range fs {

		sf, ok := t.FieldByName(k)
		if !ok {
			continue
		}
		fieldType := sf.Type.Kind()
		switch fieldType {
		case reflect.String:
			v.FieldByName(k).SetString(s.(string))

		case reflect.Float64:
			v.FieldByName(k).SetFloat(s.(float64))

		case reflect.Float32:
			o, ok := s.(float32)
			if !ok {
				return fmt.Errorf(k + " float32 convert failed")
			}
			v.FieldByName(k).SetFloat(float64(o))

		case reflect.Uint:
			o, ok := s.(uint)
			if !ok {
				return fmt.Errorf(k + " uint convert failed")
			}
			v.FieldByName(k).SetUint(uint64(o))

		case reflect.Uint8:
			o, ok := s.(uint8)
			if !ok {
				return fmt.Errorf(k + " uint8 convert failed")
			}
			v.FieldByName(k).SetUint(uint64(o))

		case reflect.Uint16:
			o, ok := s.(uint16)
			if !ok {
				return fmt.Errorf(k + " uint16 convert failed")
			}
			v.FieldByName(k).SetUint(uint64(o))

		case reflect.Uint32:
			o, ok := s.(uint32)
			if !ok {
				return fmt.Errorf(k + " uint32 convert failed")
			}
			v.FieldByName(k).SetUint(uint64(o))

		case reflect.Uint64:
			v.FieldByName(k).SetUint(s.(uint64))

		case reflect.Int:
			o, ok := s.(int)
			if !ok {
				return fmt.Errorf(k + " int convert failed")
			}
			v.FieldByName(k).SetInt(int64(o))

		case reflect.Int8:
			o, ok := s.(int8)
			if !ok {
				return fmt.Errorf(k + " int8 convert failed")
			}
			v.FieldByName(k).SetInt(int64(o))

		case reflect.Int16:
			o, ok := s.(int16)
			if !ok {
				return fmt.Errorf(k + " int16 convert failed")
			}
			v.FieldByName(k).SetInt(int64(o))

		case reflect.Int32:
			o, ok := s.(int32)
			if !ok {
				return fmt.Errorf(k + " int32 convert failed")
			}
			v.FieldByName(k).SetInt(int64(o))

		case reflect.Int64:
			v.FieldByName(k).SetInt(s.(int64))

		case reflect.Bool:
			v.FieldByName(k).SetBool(s.(bool))

		default:
			return fmt.Errorf(fieldType.String() + " type not support")
		}

	}
	return nil
}
