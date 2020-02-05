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
    - /worker
    - /common

 	
