package authToken

import (
	"crypto/md5"
	"encoding/base64"
	"encoding/json"
	"fmt"
	"strings"
)

type Header struct {
	AppId string `json:Apppid`
	Sign  string `json:Sign`
}
type Data struct {
	Header Header
	Body   string `json:"body"`
}
type AuthToken struct {
	sign string
}

func CreateAuthToken() *AuthToken {
	return &AuthToken{}
}

func (authToken *AuthToken) Create(appId string, appKey string, body string) (sign string) {
	h := md5.New()
	h.Write([]byte(appId + body + appKey))
	authToken.sign = strings.ToUpper(fmt.Sprintf("%x", h.Sum(nil)))
	return authToken.sign
}

func (authToken AuthToken) check(token *AuthToken) bool {
	if authToken.sign == token.sign {
		return true
	}
	return false
}

type ApiRequest struct {
	appId string
	sign  string
	data  *Data
}

func CreateApiRequest() *ApiRequest {
	return &ApiRequest{}
}
func (apiRequest *ApiRequest) Eecode(data []byte) string {
	return base64.StdEncoding.EncodeToString(data)
}

func (apiRequest *ApiRequest) Decode(data string) (appId string, sign string, err error) {
	bytes, err := base64.StdEncoding.DecodeString(data)
	if err != nil {
		return "", "", err
	}
	apiRequest.data = &Data{}
	if err := json.Unmarshal(bytes, apiRequest.data); err != nil {
		return "", "", nil
	}
	apiRequest.appId = apiRequest.data.Header.AppId
	apiRequest.sign = apiRequest.data.Header.Sign
	return apiRequest.appId, apiRequest.sign, nil
}
func (apiRequest *ApiRequest) GetAppid() string {
	return apiRequest.appId
}

func (apiRequest *ApiRequest) GetSign() string {
	return apiRequest.sign
}

type CredentialStorage interface {
	GetAppKeyByAppId(appId string) string
}
type CredentialStorageConfig struct {
}

func (credentialStorageConfig *CredentialStorageConfig) GetAppKeyByAppId(appId string) (appKey string) {
	if appId == "test" {
		return "test"
	} else {
		return "test"
	}
}

type ApiAuthCator struct {
	credentialStorage CredentialStorage
}

func CreateApiAuthCator(cs CredentialStorage) *ApiAuthCator {
	return &ApiAuthCator{}
}
func (apiAuthCator *ApiAuthCator) Auth(data string) (bool, error) {
	apiRequest := &ApiRequest{}
	appId, sign, err := apiRequest.Decode(data)
	if err != nil {
		return false, fmt.Errorf("Decode failed")
	}
	appKey := apiAuthCator.credentialStorage.GetAppKeyByAppId(appId)
	authToken := CreateAuthToken()
	newSign := authToken.Create(appId, appKey, apiRequest.data.Body)
	if sign == newSign {
		return true, nil
	}
	return false, nil
}

func main() {
	appId := "test"
	appKey := "test"
	sendData := &Data{
		Header: Header{
			AppId: appId,
		},
		Body: "this is test",
	}
	authToken := &AuthToken{}
	authToken.Create(appId, appKey, sendData.Body)
	sendData.Header.Sign = authToken.sign
	sendDataMarshal, _ := json.Marshal(sendData)
	sendDataString := CreateApiRequest().Eecode(sendDataMarshal)

	//服务端
	apiAuthenCator := CreateApiAuthCator(new(CredentialStorageConfig))
	auth, err := apiAuthenCator.Auth(sendDataString)
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	if auth == false {
		fmt.Println("auth failed")
		return
	}
	fmt.Println("auth success")
	return
}
