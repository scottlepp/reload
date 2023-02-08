package main

import (
	"log"
	"os"
	"plugin"
)

func main() {
	path, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}
	pluginPath := path + "/plugin.so"
	p, err := plugin.Open(pluginPath)
	if err != nil {
		panic(err)
	}
	v, err := p.Lookup("V")
	if err != nil {
		panic(err)
	}
	f, err := p.Lookup("F")
	if err != nil {
		panic(err)
	}
	*v.(*int) = 7
	f.(func())() // prints "Hello, number 7"
}
