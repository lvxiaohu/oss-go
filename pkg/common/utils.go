package common

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/spf13/viper"
)

type credentialsConfig struct {
	endpoint string
	accessKeyID string
	accessKeySecret string
}

func GetConfig() credentialsConfig {
	ossConf := credentialsConfig{
		endpoint: viper.GetString("credentialsConfig.endpoint"),
		accessKeyID: viper.GetString("credentialsConfig.accessKeyID"),
		accessKeySecret: viper.GetString("credentialsConfig.accessKeySecret"),
	}
	return ossConf
}

func OssClient() *oss.Client {
	ossClient, err := oss.New(GetConfig().endpoint,GetConfig().accessKeyID,GetConfig().accessKeySecret)
	if err != nil {
		fmt.Println(err)
		return nil
	}
	return ossClient
}
