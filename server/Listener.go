package server

type PrivateMessageListener func(string, float64, float64, string, float64) map[string]interface{}
type GroupMessageListener func(string, float64, float64, float64, string, string, string, float64) map[string]interface{}
type DiscussMessageListener func(float64, float64, float64, string, float64) map[string]interface{}

type GroupUploadListener func(float64, float64, map[string]interface{}) map[string]interface{}
type GroupAdminListener func(string, float64, float64) map[string]interface{}
type GroupDecreaseListener func(string, float64, float64, float64) map[string]interface{}
type GroupIncreaseListener func(string, float64, float64, float64) map[string]interface{}

type FriendRequestListener func(float64, string, string) map[string]interface{}
type GroupRequestListener func(string, float64, float64, string, string) map[string]interface{}

type EventListener struct {
	private_message PrivateMessageListener
	group_message   GroupMessageListener
	discuss_message DiscussMessageListener
	group_upload    GroupUploadListener
	group_admin     GroupAdminListener
	group_decrease  GroupDecreaseListener
	group_increase  GroupIncreaseListener
	friend_request  FriendRequestListener
	group_request   GroupRequestListener
}

func (el *EventListener) onMessage(m map[string]interface{}) map[string]interface{} {
	switch m["message_type"] {
	case "private":
		if el.private_message != nil {
			return el.private_message(m["sub_type"].(string), m["message_id"].(float64), m["user_id"].(float64), m["message"].(string), m["font"].(float64))
		}
		break
	case "group":
		if el.group_message != nil {
			return el.group_message(m["sub_type"].(string), m["message_id"].(float64), m["group_id"].(float64), m["user_id"].(float64), m["anonymous"].(string), m["anonymous_flag"].(string), m["message"].(string), m["font"].(float64))
		}
		break
	case "discuss":
		if el.discuss_message != nil {
			return el.discuss_message(m["message_id"].(float64), m["discuss_id"].(float64), m["user_id"].(float64), m["message"].(string), m["font"].(float64))
		}
		break
	}
	return nil
}

func (el *EventListener) onEvent(m map[string]interface{}) map[string]interface{} {
	switch m["event"] {
	case "group_upload":
		if el.group_upload != nil {
			return el.group_upload(m["group_id"].(float64), m["user_id"].(float64), m["file"].(map[string]interface{}))
		}
		break
	case "group_admin":
		if el.group_admin != nil {
			return el.group_admin(m["sub_type"].(string), m["group_id"].(float64), m["user_id"].(float64))
		}
		break
	case "group_decrease":
		if el.group_decrease != nil {
			return el.group_decrease(m["sub_type"].(string), m["group_id"].(float64), m["user_id"].(float64), m["operator_id"].(float64))
		}
		break
	case "group_increase":
		if el.group_increase != nil {
			return el.group_increase(m["sub_type"].(string), m["group_id"].(float64), m["user_id"].(float64), m["operator_id"].(float64))
		}
		break
	}
	return nil
}

func (el *EventListener) onRequest(m map[string]interface{}) map[string]interface{} {
	switch m["request_type"] {
	case "friend":
		if el.friend_request != nil {
			return el.friend_request(m["user_id"].(float64), m["message"].(string), m["flag"].(string))
		}
		break
	case "group":
		if el.group_request != nil {
			return el.group_request(m["sub_type"].(string), m["group_id"].(float64), m["user_id"].(float64), m["message"].(string), m["flag"].(string))
		}
		break
	}
	return nil
}
