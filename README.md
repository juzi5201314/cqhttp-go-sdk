# cqhttp-go-sdk

本项目为酷 Q 的 CoolQ HTTP API 插件的 Golang SDK，(即将封装 web server 相关的代码)，让使用 Golang 的开发者能方便地开发插件。仅支持插件 3.0.0 或更新版本。

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
更多cq码请参考[https://d.cqp.me/Pro/CQ码]
