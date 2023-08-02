package main

import (
	"flag"
	"fmt"
	"path/filepath"
)

func main() {
	//map1 := make(map[string]map[int64]int64)
	//mapp := map[int64]int64{1: 1}
	//map1["2"] = mapp
	//fmt.Printf("22")

	//tags := make([]string, 0)
	//fmt.Println(tags)
	//a, _ := json.Marshal(tags)
	//fmt.Println(string(a))
	//fmt.Println(string(a) == "[]")
	//
	//tags = append(tags, "aaa")
	//tags = append(tags, "bbb")
	//tags = append(tags, "ccc")
	//
	//fmt.Println(tags)
	//fmt.Println(string(tags) == "[]")
	//
	//s, _ := json.Marshal(tags)
	//fmt.Println(string(s))
	//
	//var tagsNew []string
	//err := json.Unmarshal(s, &tagsNew)
	//if err != nil {
	//	fmt.Println(err.Error())
	//}
	//fmt.Println(tagsNew)
	//reportTime := time.Now()
	//fmt.Println(reportTime)
	//
	//reportParseTime, e := time.Parse("2006-01-02T15:04:05+08:00", "2020-06-11 14:14:54")
	//if e != nil {
	//	return
	//}
	//fmt.Println(reportParseTime)

	// Declaring layout constant
	//const layout = "2023-03-08T20:24:59+08:00"
	//reportParseTime, _ := time.Parse("2006-01-02T15:04:05+08:00", layout)
	//
	//// Returns output
	//fmt.Println(reportParseTime)

	//打印数组
	//userRefreshFailList := []string{}
	//userRefreshFailList = append(userRefreshFailList, "aaaaa")
	//userRefreshFailList = append(userRefreshFailList, "bbbb")
	//fmt.Print(fmt.Sprintf("压测账号有效性巡检测试: %s", userRefreshFailList))
	//test3()

	//string.contains

	//数组打印
	//mapp := make(map[string]string)
	//mapp["test"] = "test"
	//mapp["test1"] = "test1"
	//mapp["test2"] = "test2"
	//fmt.Printf("map data is : %s", mapp)
	boolVal := flag.Bool("testBool", false, "testBool is bool type.")
	flag.Parse()

	// 如果使用 -testBool作为参数，控制台将会打印 true, 否则打印 false
	fmt.Println(boolVal)
	path := filepath.Join("aa", "bb", "vv")
	println(path)
}
func test1() error {
	fmt.Println(fmt.Sprintf("httpcode=%d, err=%s", 200, "test"))
	return fmt.Errorf(fmt.Sprintf("httpcode=%d, err=%s", 200, "test"))
}
func test2() error {
	err := test1()
	fmt.Println(err)
	fmt.Println(err.Error())
	return fmt.Errorf(fmt.Sprintf("msg = %s ", err.Error()))
}
func test3() error {
	err := test2()
	fmt.Println(err)
	fmt.Println(err.Error())
	return fmt.Errorf(fmt.Sprintf("msg = %s ", err.Error()))

}
