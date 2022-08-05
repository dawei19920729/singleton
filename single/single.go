package single

import (
	"fmt"
	"sync"
)

type Singleton struct {
	Name string `json:name`
}

var singleton *Singleton
var once sync.Once

func Getinstance() *Singleton {
	once.Do(func() {
		singleton = &Singleton{
			Name: "aaa",
		}
	})
	fmt.Println("this is github ")
	return singleton
}
