package gcron

import (
	"fmt"
	"time"
)

//调度器
type Scheduler struct {
	Ticker     *time.Ticker
	Stop       chan int
	JobManager *JobManager
}

//创建一个调度器
func NewScheduler() *Scheduler {
	s := &Scheduler{
		JobManager: NewJobManager(),
	}
	return s
}

//开始调度
func (s *Scheduler) Start() {
	go s.JobManager.StartHandleJob()
	s.Ticker = time.NewTicker(time.Second)
	go func() {
		for {
			select {
			case <-s.Ticker.C:
				fmt.Println(time.Now().Unix())
				//每一分钟触发一次任务调度
				if time.Now().Second() % 10 == 0 {
					s.JobManager.WaitHandle <- time.Now().Unix()
				}
			case <-s.Stop:
				s.Ticker.Stop()
			}
		}
	}()
	<-s.Stop
}
