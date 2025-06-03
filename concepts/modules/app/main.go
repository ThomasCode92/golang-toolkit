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

	toSlug := "now is the time 123"
	slugified, err := tools.Slugify(toSlug)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println("Slugified string:", slugified)
}
