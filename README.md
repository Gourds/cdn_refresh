
华为CDN刷新及预热工具

>**注意**
- 输入URL必须带有“http://”或“https://”,多个URL用逗号分隔,单个url的长度限制为4096字符,单次最多输入1000个url。
- 输入URL必须带有“http://”或“https://”,多个URL用逗号分隔,目前不支持对目录的预热,单个url的长度限制为4096字符,单次最多输入1000个url。


使用
- 默认 -t 为Refresh
- 不指定`--Platform`会默认使用华为的配置

```bash
# 华为云
./CDNManager.mac -i xxxx -k xxxx  -r urls.txt  -p xxxx
./CDNManager.mac -i xxxx -k xxxx  -r urls.txt  -p xxxx -t Preheating
./CDNManager.mac --Platform HuaWei -i xxxx -k xxxx  -r urls.txt  -p xxxx -t Preheating

# 白山云
./CDNManager.mac --Platform BaiShan -r urls.txt --accessKeySecret xxxxxxx   -t Preheating

# 阿里云
./CDNManager.mac --Platform AliYun -r urls.txt_aliyun -i xxxxxx -k xxxxx
./CDNManager.mac --Platform AliYun -r urls.txt_aliyun -i xxxxxx -k xxxxx -t Preheating
```

刷新URL的内容文本，如`urls.txt`，每行一个url
```
https://gourds.site/foo1.json
https://gourds.site/bar2.json
```

Mac下编译
```bash
#运行环境是mac
go build -ldflags "-X github.com/gourds/cdn_refresh/version.Version=2.1.0" -o ./CDNManager.mac
#运行环境是linux
CGO_ENABLED=0  GOOS=linux  GOARCH=amd64 go build -ldflags "-X github.com/gourds/cdn_refresh/version.Version=2.1.0" -o ./CDNManager.linux
```

TODO
- [x] 编译指定版本号
- [x] 支持华为云CDN刷新预热
- [x] 支持阿里云CDN刷新预热
- [x] 支持白山云CDN刷新预热