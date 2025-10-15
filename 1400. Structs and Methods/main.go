package main

import "fmt"

type Package struct {
	ID        string
	Delivered bool
}

func (p Package) Status() string {
	if p.Delivered == false {
		return "Package " + p.ID + " is still in transit"
	}
	return "Package " + p.ID + " has been delivered"
}

func main() {
	p1 := Package{ID: "PKG-1042", Delivered: true}
	p1.Status() // "Package PKG-1042 has been delivered"
	fmt.Println(p1.Status())

	p2 := Package{ID: "PKG-2048", Delivered: false}
	p2.Status() // "Package PKG-2048 is still in transit"
	fmt.Println(p2.Status())

}
