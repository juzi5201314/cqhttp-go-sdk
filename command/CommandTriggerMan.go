package command

import "github.com/juzi5201314/cqhttp-go-sdk"

const (
	GROUP   = 0
	PRIVATE  = 1
	DISCUSS = 2
)

type CommandTriggerMan struct {
	user_id    float64
	message_id float64
	origin     int
	origin_id  float64
}

func (ctm *CommandTriggerMan) GetOrigin() int {
	return ctm.origin
}

func (ctm *CommandTriggerMan) GetId() float64 {
	return ctm.user_id
}

func (ctm *CommandTriggerMan) GetOriginId() float64 {
	return ctm.origin_id
}

func (ctm *CommandTriggerMan) GetMessageId() float64 {
	return ctm.message_id
}

func (ctm *CommandTriggerMan) Reply(message string, api cqhttp_go_sdk.API) {
	switch ctm.GetOrigin() {
	case GROUP:
		api.SendGroupMsg(ctm.GetOriginId(), message, false)
		break
	case PRIVATE:
		api.SendPrivateMsg(ctm.GetOriginId(), message, false)
		break
	case DISCUSS:
		api.SendDiscussMsg(ctm.GetOriginId(), message, false)
		break
	}
}
