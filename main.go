package main

import (
	"fmt"
	"log"

	"github.com/TokiLoshi/gator/internal/config"
)


func main() {
	cfg, err := config.Read()
	if err != nil {
		log.Fatalf("error reading config")
	}
	
	fmt.Printf("Content :%v\n", cfg)

	err = cfg.SetUser("bianca")
	if err != nil {
		log.Fatal("error setting user")
	}

	cfg, err = config.Read()
	if err != nil {
		log.Fatal("error reading config")
	}
	fmt.Printf("More content: %v\n", cfg)
}