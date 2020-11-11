package main

import (
	"fmt"

	"github.com/spf13/pflag"
)

var conf = pflag.StringP("config", "c", "", "set the config file path~~")

func main() {
	pflag.Parse()
	fmt.Println("config=", *conf)
}
