# xui-sub-sync
xui-sub-sync is a tool for sync node setting to subconverter for clash or v2rayN subscribe

# 需求
我平时有好几台vps安装了x-ui(xray内核),我使用subconverter来配置这些节点的订阅，但是随着x-ui运行时间的加长
很难保证有一些vmess/trojan/vless端口被ban，出现tcp阻断，这个时候，我需要咋办呢。
我需要知道那些端口不能使用，然后在各台服务器的管理后台修改x-ui面板的节点端口，修改后我需要把新的节点配置字符串
复制出来，然后在subconverter中修改其中的profile 然后将新的节点配置替换原来的，最后还要重启subconverter，然后
在clash for windows中更新配置

# 方案
配置各个面板的登录地址账号密码，然后获取节点，解析出节点信息，最后替换subconverter的profile 并重启subconverter

# 使用
```bash
准备条件: 安装x-ui的vps 若干, 部署的机器上安装了subconverter

```

# TODO
- [x] 完成基本功能
- [ ] 自动发现节点端口被gfw tcp阻断
- [ ] 邮件通知 



# 最后
如果本项目对你有帮助，动动小手star是对我最大的鼓励
