package command

import (
	"regexp"
)

type Executant func(string, []string, CommandTriggerMan)

var commands map[string]Executant = map[string]Executant{}

func Excision(cmd string) []string {
	return regexp.MustCompile(`[\S]+`).FindAllString(cmd, -1)
}

func Register(cmd string, ef Executant) {
	commands[cmd] = ef
}

func Exec(cmd string, args []string, info map[string]interface{}) {
	var (
		origin    int
		origin_id float64
	)
	switch info["message_type"] {
	case "group":
		origin = GROUP
		origin_id = info["group_id"].(float64)
		break
	case "private":
		origin = PRIVATE
		origin_id = info["user_id"].(float64)
		break
	case "discuss":
		origin = DISCUSS
		origin_id = info["discuss_id"].(float64)
		break
	}
	if c, ok := commands[cmd]; ok {
		c(cmd, args, CommandTriggerMan{info["user_id"].(float64), info["message_id"].(float64), origin, origin_id})
	}
}
