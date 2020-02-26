# 嵌入代码
- 在入口文件中加入如下代码
```
import (
	"net/http"
	_ "net/http/pprof"
)

// pprof
go func() {
    log.Println(http.ListenAndServe("0.0.0.0:6060", nil))
}()
```

# 查看性能数据
- http://127.0.0.1:6060/debug/pprof/
```
Types of profiles available:
Count	Profile
150	allocs      # 过去所有内存分配的采样
0	block       # 导致同步原语阻塞的堆栈跟踪
0	cmdline     # 当前程序的命令行调用
41	goroutine   # 当前所有goroutine的堆栈跟踪
150	heap        # 活动对象的内存分配的采样
0	mutex       # 互斥锁跟踪
0	profile     # 生成cpuprofile文件 生成文件使用go tool pprof工具分析
21	threadcreate    # 创建系统线程的堆栈跟踪
0	trace       # 对当前程序执行的跟踪 生成文件使用go tool trace工具分析    
full goroutine stack dump   # 显示所有goroutine的堆栈
```

# 安装Graphviz 
- 用于go tool工具生成的数据中进行在浏览器中画图
- Mac下使用brew install graphviz

## go tool trace trace文件
- 通过上面的性能地址下载的trace文件
- 会通过网页的方式打开如下视图
```
View trace                        # 视图跟踪
Goroutine analysis
Network blocking profile
Synchronization blocking profile  # 同步阻塞
Syscall blocking profile
Scheduler latency profile         # 调度延迟
User-defined tasks
User-defined regions
```

## go tool pprof 
- go tool pprof -http=:8000 http://127.0.0.1:6060/debug/pprof/heap  查看内存使用
- go tool pprof -http=:8000 http://127.0.0.1:6060/debug/pprof/profile 查看CPU使用
- go tool pprof http://127.0.0.1:6060/debug/pprof/block
- go tool pprof http://127.0.0.1:6060/debug/pprof/mutex


