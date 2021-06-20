package main

import (
	"fmt"
	"reflect"
)

type User struct {
	Name string
	Surname *string
	Phone int
}

var Surname2 string = "GooЙыЙd"

func main() {
	user := User{
		Name: "YYYвакV",
		Surname: &Surname2,
		Phone: 98599,
	}
	DeleteCyrillic(&user)
}

// DeleteCyrillic deletes cyrillic alpha from User's field
func DeleteCyrillic(par *User)  {
	v := reflect.ValueOf(par)
	// Getting each type in struct
	for i := 0; i < v.NumField(); i++ {
		typ := v.Field(i).Type().String() //selected type of struct
		val := v.Field(i) //value of struct
		switch typ {
		case "string":
			var str = val.Interface().(string) //changing reflect.type --> string
			str = filter(str)
			fmt.Println("string", str)
		case "*string":
			var str = *val.Interface().(*string) //changing reflect.type --> string
			str = filter(str)
			fmt.Println("*string ", str)
		default:
			fmt.Printf("unknown type %T\n", val)
		}
	}
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




