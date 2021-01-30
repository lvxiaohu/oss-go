package project

import (
	"fmt"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/oss-go/pkg/common"
)

func ListBucket() {
	bucket, err := common.OssClient().ListBuckets()
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println("[List all bucket]")
	for _,b := range bucket.Buckets {
		fmt.Println(b.Name)
	}
}

func ListObject(bucketName string,project string)  {
	bucket, err := common.OssClient().Bucket(bucketName)
	if err != nil {
		fmt.Println(err)
		return
	}
	lor, err := bucket.ListObjects(oss.Prefix(project))
	if err != nil {
		fmt.Println(err)
		return
	}
	var output string
	for _, object := range lor.Objects {
		output += object.Key + "  "
	}
	fmt.Println(output)
}
