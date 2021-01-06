package main

import (
	"errors"
	"fmt"
	pkg_errors "github.com/pkg/errors"
	"log"
)

func main() {
	service()
}

/*如果有异常，在此统一处理*/
func service() {
	if err := biz(); err != nil {
		//fmt.Println("err found: %v", err)
		log.Fatal("err found: ", err)
		return
	}
	fmt.Println("service ok")
}

/*透传异常*/
func biz() error {
	return dao()
}

/*捕获底层异常并封装在wrap，向上传递*/
func dao() error {
	return pkg_errors.Wrapf(errors.New("sqlNoRows"), "test")
}
