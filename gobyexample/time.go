package main

import (
	"fmt"
	"time"
)

func main() {
	prt := fmt.Println

	// 获取现在的时间（tz=local）
	now := time.Now()
	prt("now:", now)
	prt("TimeZone:", now.Location())
	prt("weekday:", now.Weekday())

	// 构建一个新的时间
	then := time.Date(2012, 12, 11, 20, 34, 58, 3423, time.UTC)
	prt("then:", then)

	// 比较两个时间的关系，是否是之前，之后或者是同一时刻，精确到秒
	prt("Before:", now.Before(then))
	prt("After:", now.After(then))
	prt("Equal:", now.Equal(then))

	// Sub 返回一个 Duration 来表示两个时间点的间隔时间
	diff := now.Sub(then)
	prt("Sub:", diff)

	// time.Add(Duration) 将时间后移一个时间间隔
	prt("Add:", then.Add(-diff))
}
