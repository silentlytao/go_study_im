## golang语言打造高并发web即时聊天(IM)应用-支持10万人同时在线 学习源码
这是学习B站上的golang IM聊天项目的学习代码，并非视频中的源码。
### bilibili链接 https://www.bilibili.com/video/BV1CX4y1G7cN

###运行注意事项
#### 1 请将vendor里面的代码 copy到你的 GOPATH 或者 GOROOT src目录下
#### 2 项目启动前 设置 set GO111MODULE =off
#### 3 关于udp发送协程的ip问题 请先查看本机的ip网段，替换ctrl->chat.go 128行的代码 
视频中是用192.168.0.255 但是你要根据你自己的网段进行配置，比如我的网段就是 192.168.10.255
否则你会接收不到udp数据的哦
####
