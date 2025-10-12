package main

import "fmt"

func DomainForLocale(domain, locale string) string {
	if locale == "" {
		return fmt.Sprintf("en.%s", domain)
	} else {
		return fmt.Sprintf("%s.%s", locale, domain)
	}
}

// func DomainForLocale(domain, locale string) string {
// 	if locale == "" {
// 		return "en." + domain
// 	} else {
// 		return locale + "." + domain
// 	}
// }

func main() {

}
