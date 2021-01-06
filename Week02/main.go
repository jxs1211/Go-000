package main

/*
#学号:G20200607011001
#小组: 6组
#作业链接:https://github.com/jxs1211/Go-000/tree/main/Week02
题目:
我们在数据库操作的时候，比如 dao 层中当遇到一个 sql.ErrNoRows 的时候，是否应该 Wrap 这个 error，抛给上层。为什么，应该怎么做请写出代码？
*/

import (
	"Go-000/Week02/code"
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
		log.Fatal("error found: ", err)
		return
	}
	fmt.Println("service ok")
}

/*透传异常*/
func biz() error {
	err := dao()
	if (errors.Is(err, code.NotFound)) {
		return err
	}
	return nil
}

/*捕获底层异常并封装在wrap，向上传递*/
func dao() error {
	sql := "SELECT * FROM user WHERE id = 1"
	return pkg_errors.Wrapf(code.NotFound, fmt.Sprintf("sql: %s err: %v"), sql, errors.New("test"))
}
