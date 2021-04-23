# 创建项目
```
# 我们将使用 tutorial.kubebuilder.io 域，
# 所以所有的 API 组将是<group>.tutorial.kubebuilder.io.
kubebuilder init --domain tutorial.kubebuilder.io
```
`go mod tidy`
# 创建`API`
```
kubebuilder create api --group batch --version v1 --kind CronJob
```
*在此期间可能会发现缺少依赖，分析报错日志可以发现实际上是执行`make`缺少了依赖*
根据踩坑经历执行以下命令安装依赖：
```
go get github.com/go-logr/logr@v0.1.0 
go get github.com/onsi/ginkgo@v1.11.0
go get github.com/onsi/gomega@v1.8.1
```
# 设计`API`

字段可以使用大多数的基本类型。数字是个例外：出于 `API` 兼容性的目的，我们只允许三种数字类型。对于整数，需要使用 `int32` 和 `int64` 类型；对于小数，使用 `resource.Quantity` 类型。
## `CronJob`
* 组成部分：
1. 一个时间表(`CronJob`中的`cron`)
2. 要运行的`Job`模板(`CronJob`中的`job`)
* 拓展部分：
1. 一个已经启动的 `Job` 的超时时间（如果该 `Job` 执行超时，那么我们会将在下次调度的时候重新执行该 `Job`）。
2. 如果多个 `Job` 同时运行，该怎么办（我们要等待吗？还是停止旧的 `Job` ？
3. 暂停 `CronJob` 运行的方法，以防出现问题。
4. 对旧 `Job` 历史的限制
* 修改`CronJobSpec`
```
#TODO
```
* 添加并发策略及常量定义
```
#TODO
```
* 设计`status`
```
#TODO
```
**更改完此文件之后需要执行`make`命令**，期间还是会遇到依赖问题，这个时候需要手动执行
`go get k8s.io/api/batch/v1beta1@v0.17.2 `
# 实现一个控制器
1.根据名称加载定时任务

2.列出所有有效的 job，更新其状态

3.根据保留的历史版本数清理版本过旧的 job

4.检查当前 CronJob 是否被挂起(如果被挂起，则不执行任何操作)

5.计算 job 下一个定时执行时间

6.如果 job 符合执行时机，没有超出截止时间，且不被并发策略阻塞，执行该 job

7.当任务进入运行状态或到了下一次执行时间， job 重新排队

# 实现默认/验证`webhook`
执行如下命令

`kubebuilder create webhook --group batch --version v1 --kind CronJob --defaulting --programmatic-validation
`
* 实现 `api/v1/cronjob_webhook.go` 中 `Defautl`方法的实现

