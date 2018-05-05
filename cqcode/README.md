## CQ码

有关 CQ 码请参考 [酷Q官方CQ码说明](https://d.cqp.me/Pro/CQ码) 以及 [cq-http插件说明](https://cqhttp.cc/docs/3.4/#/CQCode)

解码消息

```go
func pm(sub_type string, message_id float64, user_id float64, message string, font float64) map[string]interface{} {

	m, err := cqcode.ParseMessage(message)
	if err != nil {
		return map[string]interface{}{}
	}

	for _, seg := range m {

		switch seg.Type {
		case "image":
			var image cqcode.Image
			seg.ParseMedia(&image)

			fmt.Print(image.FileID)
		}

	}
	...
}
```

编码消息

```go
...
	m := cqcode.NewMessage()

	face := cqcode.Face{
		FaceID: 170,
	}
	m.Append(&face)

	// 如果消息上报格式为 string 则转换为 string
	// 如果为 array 则直接使用 m 即可
	messageStr := m.CQString()
...
```

命令解析

```go
func pm(sub_type string, message_id float64, user_id float64, message string, font float64) map[string]interface{} {

	// 如果上报格式为 string 可以使用静态方法
	if !cqcode.IsCommand(m.(string)) {
		return map[string]interface{}{}
	}
	cmd, args := cqcode.Command(m.(string))

	// 或者先解码为 Message
	m, err := cqcode.ParseMessage(message)
	if err != nil {
		return map[string]interface{}{}
	}
	if !m.IsCommand() {
		return map[string]interface{}{}
	}
	cmd, args := m.Command()

	// cmd string, args []string
	// 注意：cmd 和 args 仍然为富媒体，可以使用 ParseMessage 解析
	...
}
```
