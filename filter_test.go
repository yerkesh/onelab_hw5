package main

import (
	"reflect"
	"testing"
)

var surname1 = "GoodfellowДл"
var surname2 = "Goodfellow"
var user1 = User{
	Name: "ElonИ",
	Surname: &surname1,
	Phone: 987456,
}
var user2 = User{
	Name: "Elon",
	Surname: &surname2,
	Phone: 987456,
}

var GeF1 = "GeForceТитан"
var GeF2 = "GeForce"
var pc1 = PC{
	Name: "MacbookЭйр",
	GPU: &GeF1,
	CPU: "M1",
}
var pc2 = PC{
	Name: "Macbook",
	GPU: &GeF2,
	CPU: "M1",
}

func TestFilter(t *testing.T)  {
	//Arrange
	testTable := []struct{
		myStruct interface{}
		expected interface{}
	} {
		{
			myStruct: &user1,
			expected: &user2,
		},
		{
			myStruct: &pc1,
			expected: &pc2,
		},
	}

	//Act
	for _, testCase := range testTable {
		result := DeleteCyrillic(testCase.myStruct) // Passing struct
		v := reflect.ValueOf(result).Elem() // Got from func
		expected := reflect.ValueOf(testCase.expected).Elem() // expected struct
		for i := 0; i < v.NumField(); i++ {
			typ := v.Field(i).Type().String() //selected type of struct
			val := v.Field(i)                 //value of struct
			expectedVal := expected.Field(i)
			switch typ {
			case "string":
				//Assert
				if p, ex := val.Interface().(string), expectedVal.Interface().(string); ex != p{
					t.Errorf("Incorrect result, Expect %s, got %s", ex, p)
				}

			case "*string":
				//Assert
				if p, ex:= *val.Interface().(*string), *expectedVal.Interface().(*string); ex != p{
					t.Errorf("Incorrect result, Expect %s, got %s", ex, p)
				}
			}
		}
	}


}