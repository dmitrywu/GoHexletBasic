package main

import (
	"fmt"
	"strings"
)

func ModifySpaces(s, mode string) string {
	switch mode {
	case "dash":
		return strings.ReplaceAll(s, " ", "-")
	case "underscore":
		return strings.ReplaceAll(s, " ", "_")
	default:
		return strings.ReplaceAll(s, " ", "*")
	}

}

func main() {
	x := 10

	switch {
	case x == 10:
		fmt.Println("equal 10 case")
		fallthrough
	case x < 20:
		fmt.Println("less than 20 case")
	}
}
