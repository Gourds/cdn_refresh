
华为CDN刷新及预热工具

使用
>- 默认 -t 为Refresh

```bash
./hwCDNManager -i xxxx -k xxxx  -r urls.txt  -p xxxx
./hwCDNManager -i xxxx -k xxxx  -r urls.txt  -p xxxx -t Preheating
```

刷新URL的内容文本，如`urls.txt`，每行一个url
```
https://gourds.site/foo1.json
https://gourds.site/bar2.json
```

Mac下编译
```bash
#运行环境是mac
go build -ldflags "-X version.Version=1.0.0" -o ./hwCDNManager.mac
#运行环境是linux
CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -ldflags "-X version.Version=1.0.0" -o ./hwCDNManager.linux
```

TODO
- [ ] 编译指定版本