package config

import (
	"gopkg.in/yaml.v3"
	"os"
	"strings"
	"sync-xui-sub/version"
)

var GlobalConfig *Settings

const ConfigFileName = "settings.yaml"

type Settings struct {
	AppName string `yaml:"appName"`
	Version string `yaml:"version"`
	Servers struct {
		Nodes      []string `yaml:"nodes"`
		NodeDetail []Node   `yaml:"nodeDetail"`
	} `yaml:"servers"`
	NodeProtocol string `yaml:"nodeProtocol"`
}
type Node struct {
	LoginUrl string `yaml:"loginUrl"`
	BaseUrl  string `yaml:"baseUrl"`
	Username string `yaml:"username"`
	Password string `yaml:"password"`
	Cookie   string `yaml:"cookie"`
}

func init() {
	GlobalConfig = newConfig()
}

func newConfig() *Settings {
	s := &Settings{
		AppName: version.AppName,
		Version: version.Version,
		Servers: struct {
			Nodes      []string `yaml:"nodes"`
			NodeDetail []Node   `yaml:"nodeDetail"`
		}{},
		NodeProtocol: "",
	}
	// pattern is xxxx,xxx,xxx|xxx,xxx,xxx
	nodes := os.Getenv("XUI_NODES")
	if nodes != "" {
		nodesList := strings.Split(nodes, "|")
		var results []Node
		for _, node := range nodesList {
			nodeList := strings.Split(node, ",")
			loginUrl := nodeList[0]
			userName := nodeList[1]
			password := nodeList[2]
			n := Node{
				LoginUrl: loginUrl,
				BaseUrl:  strings.ReplaceAll(loginUrl, "/login", ""),
				Username: userName,
				Password: password,
			}
			results = append(results, n)
		}
		s.Servers.Nodes = nodesList
		s.Servers.NodeDetail = results
	} else {
		//read yaml
		yamlFile, err := os.ReadFile(ConfigFileName)
		if err != nil {
			panic(err)
		}
		// Unmarshal the YAML into a slice of Login structs
		var settings Settings
		err = yaml.Unmarshal(yamlFile, &settings)
		if err != nil {
			panic(err)
		}
		s = &settings
		nodeStr := s.Servers.Nodes
		var results []Node
		for _, node := range nodeStr {
			split := strings.Split(node, ",")
			n := Node{
				LoginUrl: split[0],
				BaseUrl:  strings.ReplaceAll(split[0], "/login", ""),
				Username: split[1],
				Password: split[2],
			}
			results = append(results, n)

		}
		s.Servers.NodeDetail = results
	}

	protocol := os.Getenv("XUI_PROTOCOL")
	if protocol != "" {
		s.NodeProtocol = protocol
	} else {
		s.NodeProtocol = "vmess"
	}
	return s
}
