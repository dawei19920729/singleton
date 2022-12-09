package http

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

}
func get() {
	//get
	url := "https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1"

	req, _ := http.NewRequest("GET", url, nil)

	req.Header.Add("Authorization", "xxxx")

	response, _ := http.DefaultClient.Do(req)

	fmt.Println(response)
}
func post() {
	targetUrl := "https://b959e645-00ae-4bc3-8a55-7224d08b1d91.mock.pstmn.io/user/1"

	payload := strings.NewReader("{\"name\":\"张瑀楠\"}")

	req, _ := http.NewRequest("POST", targetUrl, payload)

	req.Header.Add("Content-Type", "application/json")

	response, _ := http.DefaultClient.Do(req)
	fmt.Println(response)
	response, err := http.DefaultClient.Do(req)
	if err != nil {
		// 错误逻辑处理
	}

	defer response.Body.Close()              // 这步是必要的，防止以后的内存泄漏，切记
	fmt.Println(response.StatusCode)         // 获取状态码
	fmt.Println(response.Status)             // 获取状态码对应的文案
	fmt.Println(response.Header)             // 获取响应头
	body, _ := ioutil.ReadAll(response.Body) // 读取响应 body, 返回为 []byte
	fmt.Println(string(body))                // 转成字符串看一下结果

}
