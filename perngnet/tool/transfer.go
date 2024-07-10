package tool

import (
	"bytes"
	"encoding/binary"
	"fmt"
	"reflect"
)

// BytesToStruct changes a byte array to a network packet struct
func BytesToStruct(data []byte, out any) error {
	buf := bytes.NewReader(data)
	v := reflect.ValueOf(out).Elem()
	if v.Kind() != reflect.Struct {
		return fmt.Errorf("output data is not a struct pointer")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Slice:
			if field.Type().Elem().Kind() == reflect.Uint8 {
				// 假设切片是固定长度的
				length := field.Len()
				if length == 0 {
					length = buf.Len() // 如果长度为0，则读取剩余的所有数据
				}
				slice := make([]byte, length)
				if err := binary.Read(buf, binary.BigEndian, &slice); err != nil {
					return err
				}
				field.SetBytes(slice)
			} else {
				return fmt.Errorf("unsupported slice type")
			}
		case reflect.String:
			// 假设字符串是固定长度的（需要根据实际情况调整）
			length := buf.Len() // 读取剩余的所有数据作为字符串
			str := make([]byte, length)
			if err := binary.Read(buf, binary.BigEndian, &str); err != nil {
				return err
			}
			field.SetString(string(str))
		default:
			if err := binary.Read(buf, binary.BigEndian, field.Addr().Interface()); err != nil {
				return err
			}
		}
	}
	return nil
}

// StructToByte changes a network packet struct to byte array
func StructToByte(data any) ([]byte, error) {
	var buf bytes.Buffer
	v := reflect.ValueOf(data)
	if v.Kind() != reflect.Struct {
		return nil, fmt.Errorf("input data is not a struct")
	}

	for i := 0; i < v.NumField(); i++ {
		field := v.Field(i)
		switch field.Kind() {
		case reflect.Slice:
			if field.Type().Elem().Kind() == reflect.Uint8 {
				buf.Write(field.Bytes())
			} else {
				return nil, fmt.Errorf("unsupported slice type")
			}
		case reflect.String:
			buf.WriteString(field.String())
		default:
			if err := binary.Write(&buf, binary.BigEndian, field.Interface()); err != nil {
				return nil, err
			}
		}
	}
	return buf.Bytes(), nil
}
