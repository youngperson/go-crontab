package worker

import (
	"os/exec"
	"time"

	"github.com/go-crontab/common"
)

// 任务执行器
type Executor struct {
}

var (
	G_executor *Executor
)

// 执行一个任务
func (executor *Executor) ExecuteJob(info *common.JobExecuteInfo) {
	go func() {
		var (
			cmd       *exec.Cmd
			err       error
			output    []byte
			result    *common.JobExecuteResult
			jobLock   *JobLock
			jobUnlock *JobUnlock
			nextTime  time.Time
		)

		// 任务结果
		result = &common.JobExecuteResult{
			ExecuteInfo: info,
			Output:      make([]byte, 0),
		}

		// 初始化分布式锁
		jobLock = G_jobMgr.CreateJobLock(info.Job.Name)

		// 记录任务开始时间
		result.StartTime = time.Now()

		// 上锁
		// 机器上会部署ntp对时间进行校验统一
		err = jobLock.TryLock()

		if err != nil { // 上锁失败
			result.Err = err
			result.EndTime = time.Now()
			jobLock.Unlock()
		} else {
			// 上锁成功后，重置任务启动时间
			result.StartTime = time.Now()

			// 执行shell命令
			cmd = exec.CommandContext(info.CancelCtx, "/bin/bash", "-c", info.Job.Command)

			// 执行并捕获输出
			output, err = cmd.CombinedOutput()

			// 记录任务结束时间
			result.EndTime = time.Now()
			result.Output = output
			result.Err = err

			// 不马上释放锁，在下一个任务开始之前释放(防止被其它worker执行多次)
			nextTime = G_scheduler.jobPlanTable[info.Job.Name].NextTime
			jobUnlock = &JobUnlock{
				lock:     jobLock,
				nextTime: nextTime.Unix(),
			}
			G_scheduler.jobUnlockTable[info.Job.Name] = jobUnlock
		}
		// 任务执行完成后，把执行的结果返回给Scheduler，Scheduler会从executingTable中删除掉执行记录
		G_scheduler.PushJobResult(result)
	}()
}

//  初始化执行器
func InitExecutor() (err error) {
	G_executor = &Executor{}
	return
}
