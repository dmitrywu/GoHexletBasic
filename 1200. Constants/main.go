package main

import "fmt"

const defaultStatus = 200 // package area

func ErrorMessageToCode(msg string) int {
	const (
		OK = iota
		CANCELLED
		UNKNOWN
	)

	switch {
	case msg == "OK":
		return OK
	case msg == "CANCELLED":
		return CANCELLED
	default:
		return UNKNOWN
	}

}

func main() {
	const greeting = "Hello, Hexlet"
	const age = 10

	const (
		StatusOK            = 200
		StatusNotFound      = 404
		statusInternalError = 500
	)

	const (
		RoleGuest      = iota // 0
		RoleUser              // 1
		RoleAdmin             // 2
		RoleSuperAdmin        // 3
	)

	fmt.Println(ErrorMessageToCode("OK"))
	fmt.Println(ErrorMessageToCode("CANCELLED"))
	fmt.Println(ErrorMessageToCode("UNKNOWN"))
	fmt.Println(ErrorMessageToCode("err"))
}
