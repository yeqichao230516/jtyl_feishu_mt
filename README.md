# jtyl_feishu_mt

### 编译为 Ubuntu x86_64（64位 Intel/AMD）Linux二进制可执行文件
set GOOS=linux
set GOARCH=amd64
go build -o jtyl_feishu_mt main.go

### 服务部署位置
/usr/local/bin/jtyl_feishu_mt

### 后台运行程序
`nohup ./jtyl_feishu_mt`

### 查看所包含程序名的进程PID
`pgrep -f jtyl_feishu_mt`

### 停止进程
`kill <PID>`

### 检查进程是否存在
`ps -p <PID>`

# 查看服务是否运行
ps aux | grep jtyl_feishu_mt