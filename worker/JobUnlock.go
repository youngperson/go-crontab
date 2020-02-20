package worker

import (
	"fmt"
	"sync"
	"time"
)

type JobUnlock struct {
	lock     *JobLock
	nextTime int64 // 锁在下一次执行前释放
}

type SafeMap struct {
	sync.RWMutex                         // 解决对map并发读写的报错问题
	jobUnlockTable map[string]*JobUnlock // jobname释放锁列表(轮询里面的锁,过期则释放)
}

var (
	G_safemap *SafeMap
)

func (safemap *SafeMap) writeMap(key string, value *JobUnlock) {
	safemap.Lock()
	safemap.jobUnlockTable[key] = value
	safemap.Unlock()
}

func (safemap *SafeMap) deleteMap(key string) {
	safemap.Lock()
	delete(safemap.jobUnlockTable, key)
	safemap.Unlock()
}

func (safemap *SafeMap) readMap(key string) *JobUnlock {
	safemap.RLock()
	value := safemap.jobUnlockTable[key]
	safemap.RUnlock()
	return value
}

// 释放锁协程
func (safemap *SafeMap) jobUnlockLoop() {
	var (
		jobUnlock *JobUnlock
		nextTime  int64
		jobName   string
	)
	ticker := time.NewTicker(1 * time.Millisecond)
	for {
		select {
		case <-ticker.C:
			// 列表很大可能有性能问题,需要改为延时队列的做法、也可以是轮询+Redis的有序集合方式
			for _, jobUnlock = range safemap.jobUnlockTable {
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
	sm := new(SafeMap)
	sm.jobUnlockTable = make(map[string]*JobUnlock)
	G_safemap = sm
	// 启动释放锁协程
	go G_safemap.jobUnlockLoop()
	return
}
