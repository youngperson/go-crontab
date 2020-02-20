# TODO
- 已解决：1分钟会调度60次，多个worker但是只能执行1次。
    - 利用锁不马上释放，放到channel中的处理(改为放到map中去 + 轮询)

# go-crontab
## 传统方法-crontab
- 配置任务时，需要ssh登录脚本服务器进行操作
- 服务器宕机，任务将终止调度，需要人工迁移
- 排查问题低效，无法方便的查看任务状态与错误输出

## 分布式任务调度（自研）
- 可视化web后台，方便进行任务管理
- 分布式架构、集群化调度，不存在单点故障
- 追踪任务执行状态，采集任务输出，可视化log查看

### 国内部分开源项目（都是Java的）
- ELastic Job
- XXL-JOB
- light-task-scheduler(LTS)

## 环境要求
- 依赖存储：MongoDB 3.0以上，etcd 3.0以上
- 生产部署环境：centos7.1

## 项目结构
- crontab
    - /master
        - 搭建go项目框架，配置文件，命令行参数，线程配置
        - 给web后台提供http API，用于管理job
        - 写一个web后台的前端页面，bootstrap+jquery，前后端分离开发
    - /worker
        - 从etcd中把job同步到内存中
        - 实现调度模块，基于cron表达式调用N个job
        - 实现执行模块，并发的执行多个job
        - 对job的分布式锁，防止集群并发
        - 把执行日志保存到MongoDB
    - /common

## 参考文章
- https://blog.csdn.net/CSDN_FlyYoung/article/details/99072350
- https://blog.csdn.net/oXiaoBuDing/article/details/89085351
 	
