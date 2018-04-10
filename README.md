# cqhttp-go-sdk

本项目为酷 Q 的 CoolQ HTTP API 插件的 Golang SDK，封装了 web server 相关的代码，让使用 Golang 的开发者能方便地开发插件。仅支持插件 3.0.0 或更新版本。

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

...
}
```
所有方法名字与http插件api文档中的一样，格式改为驼峰命名。
参数顺序也基本一样(如果不一样，那就是我写错了)。
所有方法返回值都为一个map，结构与http插件响应的json相同。

所有方法均有并发版本(函数名字前添加Concurrent)，如:ConcurrentSendPrivateMsg
此类方法只返回一个chan，存放着一个map。error则为map["error"]
