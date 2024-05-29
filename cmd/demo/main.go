package main

import (
	"common_pkg/cmd/demo/config"
	"common_pkg/initialize"
	"fmt"
)

func main() {

	cusType := &config.CustomizeType{}
	initialize.InitialViperWithCustomize(cusType)
	fmt.Printf("%+v\n", cusType)
}
