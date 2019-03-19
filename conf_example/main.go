package main

import (
	"bytes"
	"fmt"
	"io/ioutil"

	"github.com/spf13/viper"
)

func main() {

	viper.SetConfigType("yaml") // or viper.SetConfigType("YAML")
	yamlExample, _ := ioutil.ReadFile("./test.yaml")

	// any approach to require this configuration into your program.

	err := viper.ReadConfig(bytes.NewBuffer(yamlExample))
	fmt.Println(err)

	// // this would be "steve"
	fmt.Println(viper.Get("eyes"))
}
