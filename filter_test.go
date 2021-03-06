package main

import (
	"reflect"
	"testing"
)

type user struct {
	Name string
	Surname *string
	Phone int
}

type pc struct {
	Name string
	GPU *string
	CPU string
}

var surname1 = "GoodfellowДл"
var surname2 = "Goodfellow"
var user1 = user{
	Name: "ElonИ",
	Surname: &surname1,
	Phone: 987456,
}
var user2 = user{
	Name: "Elon",
	Surname: &surname2,
	Phone: 987456,
}

var GeF1 = "GeForceТитан"
var GeF2 = "GeForce"
var pc1 = pc{
	Name: "MacbookЭйр",
	GPU: &GeF1,
	CPU: "M1",
}
var pc2 = pc{
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
		result := DeleteCyrillic(testCase.myStruct)           // Passing struct
		v := reflect.ValueOf(result).Elem()                   // Got from func
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

var Expect = []User{
	{1, Address{5, "Satbayev"}, 20},
	{1, Address{6, "Al-Farabi"}, 32},
}
func TestUnmarshal(t *testing.T) {
	testTable := []struct {
		myStruct []byte
		expected []User
	}{
		{
			myStruct: []byte(`[
				  {
					"id": 1,
					"address": {
					  "city_id": 5,
					  "street": "Satbayev"
					},
					"Age": 20
				  },
				  {
					"id": 1,
					"address": {
					  "city_id": "6",
					  "street": "Al-Farabi"
					},
					"Age": "32"
				  }
				]`),
			expected: Expect,
		},
	}

	for _, testCase := range testTable {
		result := MyUnmarshal(testCase.myStruct)
		if !reflect.DeepEqual(result, testCase.expected) {
			t.Errorf("Incorrect result, Expect %v, got %v", testCase.expected, result)
		}
	}
}