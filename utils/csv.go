package main

import (
	"fmt"
	"github.com/gocarina/gocsv"
	"os"
)

type CsvModel struct {
	Index int `form:"index" json:"index" query:"index"`
}
type CsvModelStressUser struct {
	UserId     int64  `csv:"user_id"`
	SessionKey string `csv:"session_key"`
}
type CsvModelStressUserIM struct {
	UserId     int64  `csv:"user_id"`
	SessionKey string `csv:"session_key"`
}
type PipeUserModel struct {
	UserId       int64  `csv:"user_id"`
	SessionKey   string `csv:"session_key"`
	TenantID     int64  `csv:"tenant_id"`
	Mobile       string `csv:"mobile"`
	Did          int64  `csv:"did"`
	CredentialID int64  `csv:"credentialId"`
	Contact      string `csv:"contact"`
}

func main() {
	//getIndex()
	DealWithPipeData()
	//mergeFiles2Csv()
}
func getIndex() {
	index := 10110
	number := 15100
	fileName := fmt.Sprintf("/Users/bytedance/go/src/singleton/index.csv")
	userFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer userFile.Close()
	if err != nil {
		panic(err)
	}
	array := make([]*CsvModel, 0)
	for i := 0; i < number; i++ {
		csv := &CsvModel{
			Index: index + i,
		}
		array = append(array, csv)
	}
	ee := gocsv.MarshalFile(&array, userFile)
	if ee != nil {
		panic(ee)
	}
}

func dealWithCsv() {
	// fileName := fmt.Sprintf("/Users/bytedance/go/src/code/butterfly/sg_user_1000.csv")
	// fileName := fmt.Sprintf("/Users/bytedance/go/src/code/butterfly/va_user_5000.csv")
	fileName := fmt.Sprintf("/Users/bytedance/go/src/singleton/stressUser.csv")
	userFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	array := make([]*CsvModelStressUser, 0)
	ee := gocsv.Unmarshal(userFile, &array)
	if ee != nil {
		panic(ee)
	}
	// writeCsvList := make([]*CsvModel, 0)
	// index := 0
	//
	//	for _, v := range array {
	//		csv := &CsvModel{
	//			Index: 1,
	//		}
	//		writeCsvList = append(writeCsvList, csv)
	//		index++
	//		//time.Sleep(500 * time.Millisecond)
	//		if index%100 == 0 {
	//			fmt.Printf("当前index = %d\n", index)
	//		}
	//	}
	writeCSV(array)
}
func writeCSV(csvModel []*CsvModelStressUser) {
	fileName := fmt.Sprintf("/Users/bytedance/go/src/singleton/psmQps.csv")
	userFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer userFile.Close()
	ee := gocsv.MarshalFile(&csvModel, userFile)
	if ee != nil {
		panic(ee)
	}
}
func DealWithPipeData() {
	fileName := fmt.Sprintf("/Users/bytedance/Downloads/wdxiazai/va_user_5000.csv")
	limit := 10000

	userFile, err := os.OpenFile(fileName, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	array := make([]*CsvModelStressUser, 0)
	arrayIM := make([]*CsvModelStressUserIM, 0)
	ee := gocsv.Unmarshal(userFile, &array)
	if ee != nil {
		panic(ee)
	}
	userRefreshFailList := make([]int64, 0)
	index := 0
	for _, v := range array {
		userIM := &CsvModelStressUserIM{
			UserId:     v.UserId,
			SessionKey: v.SessionKey,
		}
		arrayIM = append(arrayIM, userIM)
		index++
		// time.Sleep(500 * time.Millisecond)
		if index%100 == 0 {
			fmt.Printf("当前index = %d\n", userRefreshFailList, index)
		}
		if limit != 0 && index == limit {
			break
		}
	}

	fileNameNew := fmt.Sprintf("/Users/bytedance/go/src/singleton/user_new.csv")
	userFileNew, err := os.OpenFile(fileNameNew, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		panic(err)
	}
	defer userFile.Close()
	eeNew := gocsv.MarshalFile(&arrayIM, userFileNew)
	if eeNew != nil {
		panic(eeNew)
	}
}
func mergeFiles2Csv() error {
	files := []string{
		"/Users/bytedance/Downloads/wdxiazai/scene_loop_28198.csv",
		"/Users/bytedance/Downloads/wdxiazai/scene_loop_28205.csv",
	}
	res := make([]*PipeUserModel, 0)
	for _, file := range files {
		userFile, err := os.OpenFile(file, os.O_RDWR|os.O_CREATE, os.ModePerm)
		defer userFile.Close()
		if err != nil {
			return err
		}
		array := make([]*PipeUserModel, 0)
		ee := gocsv.Unmarshal(userFile, &array)
		if ee != nil {
			panic(ee)
		}
		res = mergeArray(res, array)
	}

	outFile := fmt.Sprintf("/Users/bytedance/go/src/singleton/user_merge.csv")
	userOutFile, err := os.OpenFile(outFile, os.O_RDWR|os.O_CREATE, os.ModePerm)
	defer userOutFile.Close()
	if err != nil {
		return err
	}
	err = gocsv.MarshalFile(&res, userOutFile)
	if err != nil {
		panic(err)
	}
	return nil
}
func mergeArray(dest, src []*PipeUserModel) []*PipeUserModel {
	if len(dest) == 0 {
		return src
	}
	if len(src) == 0 {
		return dest
	}
	result := make([]*PipeUserModel, len(dest)+len(src))
	copy(result, dest)
	copy(result[len(dest):], src)
	return result
}
