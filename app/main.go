package main

import (
	"fmt"
	toolkit "toolkit_module"
)

func main() {
	var tools toolkit.Tools 

	s := tools.RandomString(10)
	fmt.Println("Generating string:", s )
}
