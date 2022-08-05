package main

import (
	"fmt"
	"singleton/single"
)

func main() {
	singleName := single.Getinstance()
	fmt.Println(singleName)
}
