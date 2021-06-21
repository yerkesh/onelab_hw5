package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"strconv"
)
type StringInt int

// UnmarshalJSON create a custom unmarshal for the StringInt
/// this helps us check the type of our value before unmarshalling it

func (st *StringInt) UnmarshalJSON(b []byte) error {
	//convert the bytes into an interface
	//this will help us check the type of our value
	//if it is a string that can be converted into an int we convert it
	///otherwise we return an error
	var item interface{}
	if err := json.Unmarshal(b, &item); err != nil {
		return err
	}
	switch v := item.(type) {
	case int:
		*st = StringInt(v)
	case float64:
		*st = StringInt(int(v))
	case string:
		///here convert the string into
		///an integer
		i, err := strconv.Atoi(v)
		if err != nil {
			///the string might not be of integer type
			///so return an error
			return err

		}
		*st = StringInt(i)

	}
	return nil
}

var rawJson = []byte(`[
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
]`)

var rawXML = []byte(`
	<users>
		<user>
			<id>1</id>
			<address>
				<city_id>1</city_id>
				<street>Satbayev</street>
			</address>
			<age>20</age>
		</user>
		<user>
			<id>1</id>
			<address>
				<city_id>6</city_id>
				<street>Al-Farabi</street>
			</address>
			<age>32</age>
		</user>
	</users>
	`)

func main() {
	var users []User
	if err := json.Unmarshal(rawJson, &users); err != nil {
		panic(err)
	}

	for _, user := range users {
		fmt.Printf("%#v\n", user)
	}
	//XML
	var userss Users
	if err := xml.Unmarshal(rawXML, &userss); err != nil {
		panic(err)
	}

	for _, user := range userss.Users {
		fmt.Printf("%#v\n", user)
	}

}

type User struct {
	ID      StringInt   `xml:"id" json:"id"`
	Address Address `xml:"address" json:"address"`
	Age     StringInt     `xml:"age" json:"age"`
}

type Address struct {
	CityID StringInt  `xml:"city_id" json:"city_id"`
	Street string `xml:"street" json:"street"`
}

type Users struct {
	XMLName xml.Name `xml:"users"`
	Users   []User   `xml:"user"`
}
