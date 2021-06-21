package main

import (
	"fmt"
	"reflect"
)

// DeleteCyrillic deletes cyrillic alpha from User's field
func DeleteCyrillic(pointer interface{}) interface{} {
	v := reflect.ValueOf(pointer).Elem()
	for i := 0; i < v.NumField(); i++ {
			typ := v.Field(i).Type().String() //selected type of struct
			val := v.Field(i)                 //value of struct
			switch typ {
			case "string":
				var str = val.Interface().(string) //changing reflect.type --> string
				str = filter(str)
				val.SetString(str)
			case "*string":
				var str = *val.Interface().(*string) //changing reflect.type --> string
				str = filter(str)
				val.Set(reflect.ValueOf(&str))
			default:
				fmt.Printf("unknown type %v\n", val.Kind())
			}
	}
	return pointer
}
//filter is used to remove cyrillic alpha from given a string
func filter(str string) string{
	var res []rune
	for _, val := range str {
		if val < 1024 || val > 1279 { //Range of cyrillic alpha
			res = append(res, val)
		}
	}
	return string(res)
}




