package s3_test

import (
	"fmt"
	"os"

	"gitlab.artofthings.org/platform/ground/pkg/aws"
)

func ExampleUploadFile() {
	sess, err := aws.Connect()

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	f, err := os.Open("/Users/hackartist/go/src/gitlab.artofthings.org/platform/ground/etc/ground.yml")

	if err != nil {
		fmt.Println(err)
		os.Exit(0)
	}

	res, err := aws.UploadFile(sess, "ground-platform", "/interior", "testfile2.yml", f)

	if err != nil {
		fmt.Println(res, " : ", err)
	}

	fmt.Println("ok")

	// Output: ok

}
