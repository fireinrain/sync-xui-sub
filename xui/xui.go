package xui

import (
	"bytes"
	"encoding/json"
	"fmt"
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
	return "session=" + value
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
	req2.Header.Set("origin", GetBaseUrlFromUrl(listUrl))
	req2.Header.Set("referrer", strings.ReplaceAll(listUrl, "/list", ""))
	resp2, err := client.Do(req2)
	if err != nil {
		return nil, err
	}
	defer resp2.Body.Close()
	jsonResp, err2 := io.ReadAll(resp2.Body)
	if err2 != nil {
		panic(err2)
	}
	//json
	var response = NodeListResp{}
	err2 = json.Unmarshal(jsonResp, &response)

	//err = json.NewDecoder(resp2.Body).Decode(&response)
	//if err != nil {
	//	log.Println(err)
	//	return nil, err
	//}
	fmt.Println(string(jsonResp))
	return &response, nil
}

func GetAllServerNodeList(config *config.Settings) []NodeListResp {
	nodeDetails := config.Servers.NodeDetail
	var result []NodeListResp
	for _, detail := range nodeDetails {
		list, err := GetServerNodeList(detail.BaseUrl, detail.Cookie)
		if err == nil {
			result = append(result, *list)
		} else {
			log.Println("获取节点信息失败: ", err)
		}
	}
	return result
}

//func FilterEnabledNodes(protocol string,nodes []NodeListResp) []NodeListResp {
//
//}

// GetBaseUrlFromUrl
//
//	@Description: 获取baseUrl
//	@param url
//	@return string
func GetBaseUrlFromUrl(urlStr string) string {
	urlString := urlStr
	parsedUrl, err := url.Parse(urlString)
	if err != nil {
		fmt.Println("Error parsing URL:", err)
		return ""
	}
	baseUrl := &url.URL{
		Scheme: parsedUrl.Scheme,
		Host:   parsedUrl.Host,
	}
	//fmt.Println("Base URL:", baseUrl.String())
	return baseUrl.String()
}
