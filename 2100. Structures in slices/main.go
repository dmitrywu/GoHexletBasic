package main

import "fmt"

type User struct {
	Name  string
	Email string
}

func main() {
	users := []User{
		{Name: "Семен", Email: "semen@example.com"},
		{Name: "Юля", Email: "yulia@example.com"},
	}

	users = append(users, User{Name: "Алексей", Email: "alex@example.com"})
	users[0].Name = "Семен Иванов"

	for _, u := range users {
		fmt.Println(u.Name, u.Email)
	}
}
