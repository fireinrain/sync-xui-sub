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

	fmt.Println(list)
	nodes := xui.FilterEnabledNodes(globalConfig.NodeProtocol, globalConfig.Servers.IgnoreNodeFlag, list)
	fmt.Println(nodes)
	linksStr := ""
	switch globalConfig.NodeProtocol {
	case "vmess":
		objs := xui.GenVmessLinkFromObjs(nodes)
		linksStr = objs
	case "vless":
		fmt.Println("processing vless nodes")
	case "trojan":
		fmt.Println("processing trojan nodes")
	default:
		fmt.Println("processing default")
	}
	fmt.Println(linksStr)

	//let obj = {
	//            v: '2',
	//            ps: remark,
	//            add: address,
	//            port: this.port,
	//            id: this.settings.vmesses[0].id,
	//            aid: this.settings.vmesses[0].alterId,
	//            net: network,
	//            type: type,
	//            host: host,
	//            path: path,
	//            tls: this.stream.security,
	//        };
	//        return 'vmess://' + base64(JSON.stringify(obj, null, 2));

}
