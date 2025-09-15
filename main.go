// go
package main

import (
	"fmt"
	"log"

	"github.com/Predator792002/WebGator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg)

	if err := cfg.SetUser("Ayush"); err != nil {
		log.Fatal(err)
	}

	cfg2, err := config.Read()
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(cfg2)
}
