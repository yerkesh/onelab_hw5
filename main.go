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
	fmt.Println(user.Name)
}

// DeleteCyrillic deletes cyrillic alpha from User's field
func DeleteCyrillic(par interface{}) interface{} {
	dgf := reflect.ValueOf(par).Elem()
	for i := 0; i < dgf.NumField(); i++ { ///тут кирилица не обновляется, думаю из-за метода Elem(). Хотелось бы юзать пойнтер как "*dgf.NumField()"
			typ := dgf.Field(i).Type().String() //selected type of struct
			val := dgf.Field(i) //value of struct
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
	return dgf
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




