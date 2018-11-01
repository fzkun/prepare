package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

func main() {
	var (
		expr *cronexpr.Expression
		err error
		now time.Time
		nextTime time.Time
	)

	//秒粒度，年配置(2018-2099)
	//哪一分钟(0-59) 哪一小时(0-23) 哪一天(1-31) 哪月(1-12) 星期几(0-6)

	//每分钟执行一次
	//if expr, err = cronexpr.Parse("* * * * *"); err != nil {
	//	fmt.Println(err)
	//	return
	//}

	//每隔5分钟执行一次
	if expr, err = cronexpr.Parse("*/5 * * * * * *"); err != nil {
		fmt.Println(err)
		return
	}

	now = time.Now()
	nextTime = expr.Next(now)

	fmt.Println(now, nextTime)

	time.AfterFunc(nextTime.Sub(now), func() {
		fmt.Println("Time out", nextTime)
	})

	time.Sleep(6 * time.Second)

}
