package project

import (
	"fmt"
	"github.com/oss-go/pkg/common"
)

//type oss struct {
//	Bucket string
//	LocalFilePath string
//	RemoteFilePath string
//}


func PutObject(bucketName string,remoteFilePath string,localFilePath string)  {

	bucket, err := common.OssClient().Bucket(bucketName)
	if err != nil {
		fmt.Println(err)
		return
	}

	err = bucket.PutObjectFromFile(localFilePath, remoteFilePath)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf(" %s upload successfully",localFilePath)
}