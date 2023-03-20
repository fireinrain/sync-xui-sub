package xui

import (
	"bytes"
	"encoding/json"
	"io"
	"log"
	"net/http"
	"net/http/cookiejar"
	"net/url"
	"strings"
	"sync-xui-sub/config"
	"time"
)

type NodeListResp struct {
	Success bool   `json:"success"`
	Msg     string `json:"msg"`
	Obj     []struct {
		ID             int    `json:"id"`
		Up             int64  `json:"up"`
		Down           int64  `json:"down"`
		Total          int    `json:"total"`
		Remark         string `json:"remark"`
		Enable         bool   `json:"enable"`
		ExpiryTime     int    `json:"expiryTime"`
		Autoreset      bool   `json:"autoreset"`
		Ipalert        bool   `json:"ipalert"`
		Iplimit        int    `json:"iplimit"`
		Listen         string `json:"listen"`
		Port           int    `json:"port"`
		Protocol       string `json:"protocol"`
		Settings       string `json:"settings"`
		StreamSettings string `json:"streamSettings"`
		Tag            string `json:"tag"`
		Sniffing       string `json:"sniffing"`
	} `json:"obj"`
}

// LoginForCookie
//
//	@Description: 登录xui获取cookie
//	@param loginUrl
//	@param username
//	@param password
//	@return string
func LoginForCookie(loginUrl string, username, password string) string {
	jar, err := cookiejar.New(nil)
	if err != nil {
		panic(err)
	}
	formData := url.Values{}
	formData.Set("username", username)
	formData.Set("password", password)

	// Create a new request with a POST method
	req, err := http.NewRequest("POST", loginUrl, bytes.NewBufferString(formData.Encode()))
	if err != nil {
		panic(err)
	}

	// Set the content type for the form data
	req.Header.Set("Content-Type", "application/x-www-form-urlencoded")

	client := &http.Client{
		CheckRedirect: nil,
		Jar:           jar,
	}
	// Make the request
	resp, err := client.Do(req)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()

	// Read and print the response
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		log.Println(err)
	}
	if err != nil {
		log.Println("parse string to url failed: ", err.Error())
	}
	cookies := resp.Cookies()
	cookie := cookies[0]
	value := cookie.Value
	log.Println(string(body))
	return value
}

// LoginAllNodeCookies
//
//	@Description: 登入所有节点并获取cookie
//	@param config
//	@return *config.Settings
func LoginAllNodeCookies(config *config.Settings) *config.Settings {
	nodes := config.Servers.Nodes
	for index, node := range nodes {
		split := strings.Split(node, ",")
		loginUrl := strings.TrimSpace(split[0])
		userName := strings.TrimSpace(split[1])
		password := strings.TrimSpace(split[2])
		cookie := LoginForCookie(loginUrl, userName, password)
		config.Servers.NodeDetail[index].Cookie = cookie
	}
	return config
}

// GetServerNodeList
//
//	@Description: 获取节点列表
//	@param baseUrl
//	@param cookie
//	@return *NodeListResp
//	@return error
func GetServerNodeList(baseUrl string, cookie string) (*NodeListResp, error) {
	client := http.Client{
		Timeout: time.Second * 10,
	}
	listUrl := baseUrl + "/xui/inbound/list"
	//使用cookie值 请求接口
	req2, err2 := http.NewRequest(http.MethodPost, listUrl, nil)
	if err2 != nil {
		return nil, err2
	}
	req2.Header.Set("Cookie", cookie)
	resp2, err := client.Do(req2)
	if err != nil {
		return nil, err
	}
	defer resp2.Body.Close()
	//json
	var response NodeListResp
	err = json.NewDecoder(resp2.Body).Decode(&response)
	if err != nil {
		log.Println(err)
		return nil, err
	}
	return &response, nil
}

//func FilterEnabledNodes(protocol string,nodes []NodeListResp) []NodeListResp {
//
//}
