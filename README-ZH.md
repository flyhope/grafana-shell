# Grafana-shell

通过Grafana的Alert channel webhook调用指定机器的Shell指令

## 使用方法

1. 下载后使用`go build`编译。
2. 复制`config.simple.xml`至`config.xml`修改相应配置。
3. Grafana Alert channel新增webhook，地址为：`http://{ip or domain}:8850/webhook/`
