package main

import (
	"fmt"
	"github.com/gorhill/cronexpr"
	"time"
)

type CronJob struct {
	expr     *cronexpr.Expression
	nextTime time.Time
}

func main() {
	var (
		cronjob       *CronJob
		expr          *cronexpr.Expression
		now           time.Time
		scheduleTable map[string]*CronJob
	)

	scheduleTable = make(map[string]*CronJob)

	expr = cronexpr.MustParse("*/5 * * * * * *")

	now = time.Now()
	cronjob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}

	scheduleTable["job1"] = cronjob

	expr = cronexpr.MustParse("*/5 * * * * * *")

	now = time.Now()
	cronjob = &CronJob{
		expr:     expr,
		nextTime: expr.Next(now),
	}

	scheduleTable["job2"] = cronjob

	go func() {
		var (
			jobName string
			cronJob *CronJob
			now     time.Time
		)
		for {

			now = time.Now()
			for jobName, cronJob = range scheduleTable {
				if cronJob.nextTime.Before(now) || cronJob.nextTime.Equal(now) {
					//启动携程执行任务
					go func(jobName string) {
						fmt.Println("执行", jobName)
					}(jobName)

					// 计算下次时间
					cronjob.nextTime = cronjob.expr.Next(now)
					fmt.Println("下次执行时间", cronjob.nextTime)

				}
			}

			// 睡眠100毫秒
			select {
			case <-time.NewTimer(100 * time.Millisecond).C:

			}

		}
	}()

	time.Sleep(100 * time.Second)
}
