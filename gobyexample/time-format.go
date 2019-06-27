package main

import (
	"fmt"
	"time"
)

const (
	timeFmt = "2006-01-02 15:04:05"
	tzFmt   = "2006-01-02 15:04:05Z07:00  MST"
)

func main() {
	prt := fmt.Println
	now := time.Now()
	// 时间格式化为字符串
	prt(now.Format(time.RFC3339))
	// 自定义格式化内容
	prt(now.Format(timeFmt))
	prt(now.Format(tzFmt))
	// 纯数字表示的时间，你也可以使用标准的格式化字符
	fmt.Printf("%d-%02d-%02d %02d:%02d:%02d\n",
		now.Year(), now.Month(), now.Day(),
		now.Hour(), now.Minute(), now.Second())

	// 字符串解析为时间类型(tz=UTC)
	t, _ := time.Parse(timeFmt, now.Format(timeFmt))
	prt("parse:", t.Local())

	_, err := time.Parse(timeFmt, now.Format(tzFmt))
	prt("parse error:", err)
}
