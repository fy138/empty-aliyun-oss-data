package main

import (
	"fmt"
	"log"
	"os"

	"github.com/aliyun/aliyun-oss-go-sdk/oss"
)

func HandleError(err error) {
	log.Println(err)
}

func main() {
	ak := os.Getenv("ALICLOUD_ACCESS_KEY")
	as := os.Getenv("ALICLOUD_SECRET_KEY")
	//log.Println(ak, " <-> ", as)
	client, err := oss.New("oss-cn-shenzhen.aliyuncs.com", ak, as)
	if err != nil {
		HandleError(err)
	}

	bucket, err := client.Bucket("mailstoresz")
	if err != nil {
		HandleError(err)
	}

	for {
		lsRes, err := bucket.ListObjects()
		if err != nil {
			HandleError(err)
			break
		}
		if len(lsRes.Objects) == 0 {
			break
		}
		for _, object := range lsRes.Objects {
			fmt.Println("Objects:", object.Key)
			err = bucket.DeleteObject(object.Key)
			if err != nil {
				HandleError(err)
			}
		}
	}
}
