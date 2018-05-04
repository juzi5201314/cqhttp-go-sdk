package server

import (
	"encoding/json"
	"github.com/juzi5201314/cqhttp-go-sdk/command"
	"io/ioutil"
	"net/http"
	"strconv"
)

func StartListenServer(port int, path string) ListenServer {
	return ListenServer{port, path, EventListener{}}
}

type ListenServer struct {
	port     int
	path     string
	listener EventListener
}

func (s *ListenServer) Listen() {
	http.HandleFunc(s.path, func(writer http.ResponseWriter, request *http.Request) {
		defer request.Body.Close()
		buffer, _ := ioutil.ReadAll(request.Body)
		m := map[string]interface{}{}
		json.Unmarshal(buffer, &m)
		rm := map[string]interface{}{}
		switch m["post_type"] {
		case "message":
			rm = s.listener.onMessage(m)
			if b, ok := rm["stop"]; ok && b.(bool) {
				break
			}
			cs := command.Excision(m["message"].(string))
			go command.Exec(cs[0], cs, m)
		case "event":
			rm = s.listener.onEvent(m)
			break
		case "request":
			rm = s.listener.onRequest(m)
			break
		}

		if rm == nil || len(rm) < 1 {
			writer.WriteHeader(204)
		} else {
			jo, _ := json.Marshal(rm)
			writer.Write(jo)
		}
	})
	err := http.ListenAndServe(":"+strconv.Itoa(s.port), nil)
	if err != nil {
		panic(err)
	}
}

func (s *ListenServer) ListenPrivateMessage(listener PrivateMessageListener) {
	s.listener.private_message = listener
}

func (s *ListenServer) ListenGroupMessage(listener GroupMessageListener) {
	s.listener.group_message = listener
}

func (s *ListenServer) ListenDiscussMessage(listener DiscussMessageListener) {
	s.listener.discuss_message = listener
}

func (s *ListenServer) ListenGroupUpload(listener GroupUploadListener) {
	s.listener.group_upload = listener
}

func (s *ListenServer) ListenGroupAdmin(listener GroupAdminListener) {
	s.listener.group_admin = listener
}

func (s *ListenServer) ListenGroupDecrease(listener GroupDecreaseListener) {
	s.listener.group_decrease = listener
}

func (s *ListenServer) ListenGroupIncrease(listener GroupIncreaseListener) {
	s.listener.group_increase = listener
}

func (s *ListenServer) ListenFriendRequest(listener FriendRequestListener) {
	s.listener.friend_request = listener
}

func (s *ListenServer) ListenGroupRequest(listener GroupRequestListener) {
	s.listener.group_request = listener
}
