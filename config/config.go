package config

import (
	"fmt"
	"io/ioutil"
	"strings"
)

var (
	CommonCfg   Config
	InputConfig Config
	UrlsList []string
)

type Config struct {
	Type            string `toml:"type"`
	ReadFile        string `toml:"readFile"`
	AccessKeyID     string `toml:"accessKeyID"`
	AccessKeySecret string `toml:"accessKeySecret"`
	DomainId        string `toml:"domainId"`
	ProjectId       string `toml:"ProjectId"`
	EndPoint  		string `toml:"endPoint"`
	Region          string `toml:"Region"`
	Platform		string `toml:"Platform"`
}

func getUrlsFromFile(){
	content, err := ioutil.ReadFile(InputConfig.ReadFile)
	if err != nil{
		fmt.Println(err)
	}
	UrlsList = strings.FieldsFunc(string(content), Split)

	for i := range UrlsList {
		UrlsList[i] = strings.TrimSpace(UrlsList[i])
	}
}

func Split(r rune) bool {
	return r == ',' || r == '\n'
}


func init() {
	initConfig()
	getUrlsFromFile()
}