# 打包流程
根据系统环境参数不同
env GOOS=linux GOARCH=amd64 GOARM=7 go build
# 服务器运行流程
chmod 777 IMDemo
nohup ./IMDemo &
