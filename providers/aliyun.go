package providers

import (
	"bytes"
	"crypto/hmac"
	"crypto/md5"
	"crypto/rand"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"github.com/gourds/cdn_refresh/config"
	"github.com/wonderivan/logger"
	"hash"
	"io"
	"io/ioutil"
	"net/http"
	"net/url"
	"sort"
	"strings"
	"time"
)

func (c *AliYun)Session() (SessionType, error){
	return SessionType{aliyun: "gourds"}, nil
}

func (c *AliYun)Refresh(client SessionType) (err error){
	logger.Info("Aliyun Refresh Job Beginning...")
	urls := strings.Join(config.UrlsList,"\\n")
	composeUrl(c.AccessKeyID, c.AccessKeySecret, "RefreshObjectCaches", urls)
	return
}



func (c *AliYun)Preheating(client SessionType)(err error){
	logger.Info("Aliyun Refresh Job Preheating...")
	urls := strings.Join(config.UrlsList,"\\n")
	composeUrl(c.AccessKeyID, c.AccessKeySecret, "PushObjectCache", urls)
	return
}


func composeUrl(accessKeyId , accessKeySecret ,Action ,refreshUrl string){
	HTTPMethod := "GET"
	params := map[string]string{}
	demostring := "SignatureVersion=1.0&Format=JSON&Timestamp=2015-08-06T02:19:46Z&AccessKeyId=testid&SignatureMethod=HMAC-SHA1&Version=2014-11-11&Action=DescribeCdnService&SignatureNonce=9b7a44b0-3be1-11e5-8c73-08002700c460"
	for _, kv := range strings.Split(demostring, "&") {
		pairs := strings.Split(kv, "=")
		params[pairs[0]] = pairs[1]
	}
	//Timestamp AccessKeyID SignatureNoce
	params["Timestamp"] = time.Now().UTC().Format(time.RFC3339)
	params["AccessKeyId"] = accessKeyId
	params["SignatureNonce"] = getGuid()

	//Refresh Choice
	params["Action"] = Action
	params["ObjectPath"] = refreshUrl
	params["ObjectType"] = "File"

	//Header
	hs := newHeaderSorter(params)
	hs.Sort()

	// GET the CanonicalizedOSSHeaders
	CanonicalizedQueryString := ""
	for i := range hs.Keys {
		CanonicalizedQueryString += percentEncode(hs.Keys[i]) + "=" +
			percentEncode(hs.Values[i]) + "&"
	}
	CanonicalizedQueryString = strings.TrimSuffix(CanonicalizedQueryString, "&")
	logger.Info("CanonicalizedQueryString",CanonicalizedQueryString)
	stringToSign := HTTPMethod + "&" + percentEncode("/") + "&" + percentEncode(CanonicalizedQueryString)
	logger.Debug("stringToSign", stringToSign)

	//computer signature
	h := hmac.New(func() hash.Hash {
		return sha1.New()
	}, []byte(accessKeySecret + "&"))
	io.WriteString(h, stringToSign)
	signedStr := base64.StdEncoding.EncodeToString(h.Sum(nil))
	params["Signature"] = percentEncode(signedStr)

	//make request
	queryStr := ""
	for k, v := range params {
		queryStr += k + "=" + v + "&"
	}
	queryStr = strings.TrimSuffix(queryStr, "&")
	reqUrl := "http://cdn.aliyuncs.com?" + queryStr
	logger.Info(reqUrl)
	res, err := http.DefaultClient.Get(reqUrl)
	if err != nil {
		logger.Error(err)
	}
	bs,_ := ioutil.ReadAll(res.Body)
	defer res.Body.Close()
	logger.Info(string(bs))
	if res.StatusCode > 200 {
		logger.Info(res.Status)
	}
}

type headerSorter struct {
	Keys []string
	Values []string
}

func newHeaderSorter(m map[string]string) *headerSorter {
	hs := &headerSorter{
		Keys: make([]string, 0, len(m)),
		Values: make([]string,0,len(m)),
	}
	for k,v := range m {
		hs.Keys = append(hs.Keys, k)
		hs.Values = append(hs.Values, v)
	}
	return hs
}

// Additional function for function SignHeader.

func (hs *headerSorter) Sort() {
	sort.Sort(hs)
}

func (hs *headerSorter) Len() int {
	return len(hs.Values)
}

func (hs *headerSorter) Less(i, j int) bool{
	return bytes.Compare([]byte(hs.Keys[i]), []byte(hs.Keys[j])) < 0
}

func (hs *headerSorter) Swap(i,j int){
	hs.Values[i], hs.Values[j] = hs.Values[j], hs.Values[i]
	hs.Keys[i], hs.Keys[j] = hs.Keys[j], hs.Keys[i]
}

func percentEncode(s string) string{
	s = url.QueryEscape(s) //测试发现 go QueryEscape 够用了. 下面的转换Replace可选
	return s
	//return strings.Replace(
	//	strings.Replace(
	//		strings.Replace(s, "+", "%20", -1),
	//		"*", "%2A", -1),
	//	"%7E", "~", -1);
}

func getGuid() string{
	b := make([]byte, 48)
	if _, err := io.ReadFull(rand.Reader, b); err != nil{
		return ""
	}
	return getMd5String(base64.URLEncoding.EncodeToString(b))
}

func getMd5String(s string) string{
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}