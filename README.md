# cqhttp-go-sdk

本项目为酷 Q 的 CoolQ HTTP API 插件的 Golang SDK，封装了 web server 相关的代码，让 Gopher 能方便地开发插件。仅支持插件 3.0.0 或更新版本。

关于 CoolQ HTTP API 插件，见 [richardchien/coolq-http-api](https://github.com/richardchien/coolq-http-api)。

## 获取
```bash
go get -u github.com/juzi5201314/cqhttp-go-sdk
```
## 使用
```go
import "github.com/juzi5201314/cqhttp_go_sdk"

...

api := cqhttp_go_sdk.Api(url, token)
```
例如给 123456789 发送消息 233
```go
package main

import "github.com/juzi5201314/cqhttp_go_sdk"

func main(){
api := cqhttp_go_sdk.Api("http://localhost:5700", "xxxx")
m, err :=api.SendPrivateMsg(123456789, "233", false)
if err != nil {
panic(err)
}
...
}
```
所有方法名字与http插件api文档中的一样，格式改为驼峰命名。
参数顺序也基本一样(如果不一样，那就是我写错了)。
所有方法返回值都为一个map，结构与http插件响应的json相同。

所有方法均有并发版本(函数名字前添加Concurrent)，如:ConcurrentSendPrivateMsg
此类方法只返回一个chan，存放着一个map。error则为map["error"]
例如:
```go
api := cqhttp_go_sdk.Api("http://localhost:5700", "xxxx")
c := api.ConcurrentSendPrivateMsg(23456789, "msg", false)
m := <-c
err := m["error"]
if err != nil {
panic(err)
}
```
这个有什么用?
假设你代码跟qq机器人不是同一个服务器而延迟200ms
意味着加上处理消息的时间，发送到收到响应信息需要500ms
那么你同时发送5条信息，最后一条甚至要2秒之后才会开始发送
对于一些业务，响应时间长会造成用户体验极差
而得益于go强大的并发能力，如果是并发版本，5条信息几乎是同时发送的。

## CQ码
```go
import "github.com/juzi5201314/cqhttp_go_sdk/cq"
...
cq.At(123456789)//返回一个字符串[CQ:at,qq=123456789]
```
更多cq码请参考[酷q官方CQ说明](https://d.cqp.me/Pro/CQ码)

## webserver
```go
import "github.com/juzi5201314/cqhttp_go_sdk/server"

...
s := server.StartListenServer(5700, "/")
//启动一个http server在5700端口

s.ListenPrivateMessage(server.PrivateMessageListener(pm))
//添加一个函数来监听私聊消息

s.Listen()
//开始监听，此函数使用之后将会阻塞
...

//函数参数与cqhttp插件文档的顺序一样，如下。
//返回的map为响应数据，如返回nil或者空map表示不响应
func pm(sub_type string, message_id float64, user_id float64, message string, font float64) map[string]interface{} {
println("收到消息"+message+"，类型为"+sub_type)
return map[string]interface{}{
"reply": "这是一条自动回复消息",
}
}
```
更多事件请看[cqhttp插件官方文档](https://cqhttp.cc/docs/3.4/#/Post)
