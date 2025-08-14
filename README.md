# jtyl_feishu_mt

### 编译为 Ubuntu x86_64（64位 Intel/AMD）Linux二进制可执行文件
go env
set GOOS=linux
set GOARCH=amd64
go build -o jtyl_feishu_mt main.go

### 服务部署位置
cd /usr/local/bin/jtyl_feishu_mt

### 后台运行程序
nohup ./jtyl_feishu_mt

# 发送强制终止信号（SIGKILL）
pkill -9 -f jtyl_feishu_mt

# 再次验证是否全部终止
pgrep -f jtyl_feishu_mt