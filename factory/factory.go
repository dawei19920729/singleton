package main

import "fmt"

type ConfigParse interface {
	Parse(path string)
}
type JsonConfigParse struct {
}

func (*JsonConfigParse) Parse(path string) {
	fmt.Println("JsonConfigParse " + path)
}

type YarmConfigParse struct {
}

func (*YarmConfigParse) Parse(path string) {
	fmt.Println("YarmConfigParse" + path)
}

// 简单工厂
type SimpleFacory struct {
}

func (*SimpleFacory) CreateParse(str string) ConfigParse {
	switch str {
	case "json":
		return &JsonConfigParse{}
	case "yarm":
		return &YarmConfigParse{}
	}
	return nil
}

// 工厂模式
type NormaleFactory interface {
	Create() ConfigParse
}
type JsonNormaleFactory struct {
}

func (*JsonNormaleFactory) Create() ConfigParse {
	return &JsonConfigParse{}
}
func main() {
	factory := &SimpleFacory{}
	jsonParse := factory.CreateParse("json")
	jsonParse.Parse("user/json")
	yarmParse := factory.CreateParse("yarm")
	yarmParse.Parse("user/yarm")

	factory1 := &JsonNormaleFactory{}
	jsonParse1 := factory1.Create()
	jsonParse1.Parse("user/json")

}
