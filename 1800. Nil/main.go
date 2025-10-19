package main

import "fmt"

type Config struct {
	Host *string
	Port int
}

func PrintConfig(cfg *Config) {
	if cfg.Host == nil {
		fmt.Println("Host is not set")
		fmt.Println("Port:", cfg.Port)
	} else {
		fmt.Println("Host:", *cfg.Host)
		fmt.Println("Port:", cfg.Port)
	}

}

func main() {
	host := "localhost"
	cfg := &Config{Host: &host, Port: 8080}
	PrintConfig(cfg)
	// Вывод:
	// Host: localhost
	// Port: 8080

	cfg2 := &Config{Port: 3000}
	PrintConfig(cfg2)
	// Вывод:
	// Host is not set
	// Port: 3000
}
