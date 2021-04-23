**本地环境为Mac**

通过配置 `kubectl` 的 `config` 配置文件访问远端集群，配置文件为`config/config`。

由于每次执行`kubectl`命令都需要带上`--kubeconfig ./config/config`参数，使用`alias`来定义别名
`alias ktf='kubectl --kubeconfig ./config/config'`

