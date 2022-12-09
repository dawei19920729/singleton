package main

import (
	"fmt"
	"singleton/single"
)

func main() {
	singleName := single.Getinstance()
	fmt.Println(singleName)
	//fmt.Sprintf("BindRequest error, err: %v", singleName)
	//fmt.Sprintf("BindRequest error, err: %v", singleName)

}
