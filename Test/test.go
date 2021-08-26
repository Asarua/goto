package test

import (
	"fmt"
)

type log interface {
	success()
}

type Log struct {
	msg string
}

func (logger Log) success() {
	fmt.Printf("要输出的文字是 %s", logger.msg)
}
