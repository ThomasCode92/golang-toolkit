package main

import (
	"fmt"

	"github.com/ThomasCode92/toolkit"
)

func main() {
	var tools toolkit.Tools

	s := tools.RandomString(10)
	fmt.Println("Random string:", s)

	tools.CreateDirIfNotExist("./testdir")
	tools.CreateDirIfNotExist("./tmp/subdir")
}
