package main

import "fmt"

type User struct {
	Name string
	Age  int
}

type Product struct {
	Name  string
	Price int
}

func NewDiscountedProduct(name string, price int, discount int) *Product {

	if discount >= 0 && discount < 101 {
		total := float64(price) * (1.0 - (float64(discount) / 100.0))
		return &Product{Name: name, Price: int(total)}
	} else if discount < 0 {
		return &Product{Name: name, Price: price}
	} else {
		return &Product{Name: name, Price: 0}
	}
}

func NewUser(name string, age int) User {
	return User{Name: name, Age: age}
}

func NewUserPtr(name string, age int) *User {
	return &User{Name: name, Age: age}

}

func main() {
	user := NewUser("Alice", 30)
	fmt.Println(user.Name) // Alice
	fmt.Println(user.Age)  // Alice
	userPtr := NewUserPtr("Bob", 25)
	fmt.Println(userPtr.Age) // => 26
	userPtr.Age++
	fmt.Println(userPtr.Age) // => 26

	p := NewDiscountedProduct("Laptop", 1000, 10)
	fmt.Println(p.Price) // 900
	p1 := NewDiscountedProduct("Laptop", 1000, 0)
	fmt.Println(p1)
}
