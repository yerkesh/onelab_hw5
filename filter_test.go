package main

import "testing"

func TestFilter(t *testing.T)  {
	var surname1 = "Goodfellow"
	var user1 = User{
		Name: "Elon",
		Surname: &surname1,
		Phone: 987456,
	}
	//Arrange
	testTable := []struct{
		user *User
		expected User
	} {
		{
			user: &user1,
			expected:
		},
	}

	//Act


	//Assert
}