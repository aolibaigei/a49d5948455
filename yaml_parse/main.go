package main

import (
	"fmt"
	"log"

	"github.com/kylelemons/go-gypsy/yaml"
)

func main() {

	// config, err := yaml.ReadFile("test.yaml")
	config, err := yaml.ReadFile("test.yaml")
	if err != nil {
		log.Fatalf("readfile(%q): %s", "test.yaml", err)
	}

	val, err := config.Get("eyes")
	if err != nil {
		fmt.Printf("%s\n", err)
	}
	fmt.Println(len(val))
	// for e := range val {
	// 	fmt.Println(e)
	// }

}
