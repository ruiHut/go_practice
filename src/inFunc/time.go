package main

import (
	"fmt"
	"time"
)

// go time  时间原点	2006-01-02 15:04:05
func main() {
	// toTime 格式化  方法一 使用毫秒数
	t := time.Unix(1362984425, 0) // 默认为美国时间格式
	fmt.Println(t)                // 2013-03-11 14:47:05 +0800 CST
	// 格式化
	tf := t.Format("2006-01-02 15:04:05") // 需要对应时间原点格式化
	fmt.Println(tf)
	// toTime 格式化    用法二： 使用固定的字符串进行格式化获得time 对象
	TestXYZ()
	addDateTest()
}

const TimeFormat = "2006-01-02 15:04:05"

func TestXYZ() {

	t, _ := time.Parse(TimeFormat, "2013-08-11 11:18:46")
	fmt.Println(t) // 2013-08-11 11:18:46 +0000 UTC
}

// addDate 返回与将给定的年数，月数和天数添加到t相对应的时间
// 例如，应用于2011年1月1日的AddDate（-1,2,3）将于2010年3月4日返回。

func addDateTest() {
	start := time.Date(2009, 1, 1, 0, 0, 0, 0, time.UTC)
	fmt.Println("before is ", start)
	after := start.AddDate(0,-1,0)
	fmt.Println("after is ",after)
}
