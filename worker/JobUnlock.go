package worker

import (
	"fmt"
	"time"

	cmap "github.com/orcaman/concurrent-map"
)

type JobUnlock struct {
	lock     *JobLock
	nextTime int64 // 锁在下一次执行前释放
}

type SafeMap struct {
	Map cmap.ConcurrentMap
}

var (
	G_safemap *SafeMap
)

func (safemap *SafeMap) writeMap(key string, value *JobUnlock) {
	safemap.Map.Set(key, value)
}

func (safemap *SafeMap) deleteMap(key string) {
	safemap.Map.Remove(key)
}

func (safemap *SafeMap) readMap(key string) *JobUnlock {
	if tmp, ok := safemap.Map.Get(key); ok {
		return tmp.(*JobUnlock)
	}
	return &JobUnlock{}
}

// 释放锁协程
func (safemap *SafeMap) jobUnlockLoop() {
	var (
		item      interface{}
		jobUnlock *JobUnlock
		nextTime  int64
		jobName   string
	)
	ticker := time.NewTicker(1 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			// 列表很大可能有性能问题,需要改为延时队列的做法、也可以是轮询+Redis的有序集合方式
			for _, item = range safemap.Map.Items() {
				jobUnlock = item.(*JobUnlock)
				nextTime = jobUnlock.nextTime
				if nextTime <= time.Now().Unix() {
					jobName = jobUnlock.lock.jobName
					fmt.Println("释放锁完成：", jobName)
					jobUnlock.lock.Unlock()
					safemap.deleteMap(jobName)
				}
			}
		}
	}
}

// 初始化对map的读写
func InitSafeMap() (err error) {
	G_safemap = &SafeMap{
		Map: cmap.New(),
	}
	// 启动释放锁协程
	go G_safemap.jobUnlockLoop()
	return
}
