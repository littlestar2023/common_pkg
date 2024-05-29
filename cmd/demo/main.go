package main

import (
	"fmt"
	"github.com/littlestar2023/common_pkg/cmd/demo/config"
	"github.com/littlestar2023/common_pkg/initialize"
)

func main() {

	cusType := &config.CustomizeType{}
	initialize.InitialViperWithCustomize(cusType)
	fmt.Printf("%+v\n", cusType)
}
