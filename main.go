package main

import (
	"fmt"
	"sync-xui-sub/config"
	"sync-xui-sub/xui"
)

func main() {
	// Define the form data
	globalConfig := config.GlobalConfig

	config := xui.LoginAllNodeCookies(globalConfig)
	globalConfig = config

	list := xui.GetAllServerNodeList(globalConfig)
	//fmt.Println(list)
	nodes := xui.FilterEnabledNodes(globalConfig.NodeProtocol, globalConfig.Servers.IgnoreNodeFlag, list)
	//fmt.Println(nodes)
	linksStr := xui.GenVmessLinkFromObjs(nodes)

	fmt.Println(linksStr)

}
