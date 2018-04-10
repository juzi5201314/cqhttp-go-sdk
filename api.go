package cqhttp_go_sdk

import (
"net/http"
"io/ioutil"
"strings"
"encoding/json"
)

func Api(url string, token string) API {
return API{url, token}
}

type API struct {
url string
token string
}

func (api *API) post(path string, data map[string]interface{}) (map[string]interface{}, error) {
client := &http.Client{} 
jo, _ := json.Marshal(data)
request, err := http.NewRequest("POST", api.url + path, strings.NewReader(string(jo)))
if err != nil {
panic(err)
}
request.Header.Set("Authorization", "Token " + api.token)
request.Header.Set("Content-Type", "application/json")
rb, err := client.Do(request)
defer rb.Body.Close()
if err != nil {
panic(err)
}
buffer, err := ioutil.ReadAll(rb.Body)
rd := map[string]interface{}{}
json.Unmarshal(buffer, &rd)
return rd, err
}

func (api *API) concurrentPost(path string, data map[string]interface{}, c chan<- map[string]interface{}) {
m, err := api.post(path, data)
m["error"] = err
c <- m
}

func (api *API) SendPrivateMsg(user_id int, message string, auto_escape bool) (map[string]interface{}, error) {
return api.post("/send_private_msg", map[string]interface{}{
"user_id": user_id,
"message": message,
"auto_escape": auto_escape,
})
}

func (api *API) ConcurrentSendPrivateMsg(user_id int, message string, auto_escape bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/send_private_msg", map[string]interface{}{
"user_id": user_id,
"message": message,
"auto_escape": auto_escape,
}, c)
return c
}

func (api *API) SendGroupMsg(group_id int, message string, auto_escape bool) (map[string]interface{}, error) {
return api.post("/send_group_msg", map[string]interface{}{
"group_id": group_id,
"message": message,
"auto_escape": auto_escape,
})
}

func (api *API) ConcurrentSendGroupMsg(group_id int, message string, auto_escape bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/send_group_msg", map[string]interface{}{
"group_id": group_id,
"message": message,
"auto_escape": auto_escape,
}, c)
return c
}

func (api *API) SendDiscussMsg(discuss_id int, message string, auto_escape bool) (map[string]interface{}, error) {
return api.post("/send_discuss_msg", map[string]interface{}{
"discuss_id": discuss_id,
"message": message,
"auto_escape": auto_escape,
})
}

func (api *API) ConcurrentSendDiscussMsg(discuss_id int, message string, auto_escape bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/send_discuss_msg", map[string]interface{}{
"discuss_id": discuss_id,
"message": message,
"auto_escape": auto_escape,
}, c)
return c
}

func (api *API) DeleteMsg(message_id int) (map[string]interface{}, error) {
return api.post("/delete_msg", map[string]interface{}{
"message_id": message_id,
})
}

func (api *API) ConcurrentDeleteMsg(message_id int) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/delete_msg", map[string]interface{}{
"message_id": message_id,
}, c)
return c
}

func (api *API) SendLike(user_id int, times int) (map[string]interface{}, error) {
return api.post("/send_like", map[string]interface{}{
"user_id": user_id,
"times": times,
})
}

func (api *API) ConcurrentSendLike(user_id int, times int) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/send_like", map[string]interface{}{
"user_id": user_id,
"times": times,
}, c)
return c
}

func (api *API) SetGroupKick(group_id int,user_id int, reject_add_request bool) (map[string]interface{}, error) {
return api.post("/set_group_kick", map[string]interface{}{
"group_id": group_id,
"user_id": user_id,
"reject_add_request": reject_add_request,
})
}

func (api *API) ConcurrentSetGroupKick(group_id int, user_id int, reject_add_request bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_kick", map[string]interface{}{
"user_id": user_id,
"group_id": group_id,
"reject_add_request": reject_add_request,
}, c)
return c
}

func (api *API) SetGroupBan(group_id int, user_id int, duration int) (map[string]interface{}, error) {
return api.post("/set_group_ban", map[string]interface{}{
"user_id": user_id,
"group_id": group_id,
"duration": duration,
})
}

func (api *API) ConcurrentSetGroupBan(group_id int, user_id int, duration int) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_ban", map[string]interface{}{
"user_id": user_id,
"group_id": group_id,
"duration": duration,
}, c)
return c
}

func (api *API) SetGroupAnonymousBan(group_id int, flag string, duration int) (map[string]interface{}, error) {
return api.post("/set_group_anonymous_ban", map[string]interface{}{
"group_id": group_id,
"flag": flag,
"duration": duration,
})
}

func (api *API) ConcurrentSetGroupAnonymousBan(group_id int, flag string, duration int) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_anonymous_ban", map[string]interface{}{
"group_id": group_id,
"flag": flag,
"duration": duration,
}, c)
return c
}

func (api *API) SetGroupWholeBan(group_id int, enable bool) (map[string]interface{}, error) {
return api.post("/set_group_whole_ban", map[string]interface{}{
"group_id": group_id,
"enable": enable,
})
}

func (api *API) ConcurrentSetGroupWholeBan(group_id int, enable bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_whole_ban", map[string]interface{}{
"group_id": group_id,
"enable": enable,
}, c)
return c
}


func (api *API) SetGroupAdmin(group_id int, user_id int, enable bool) (map[string]interface{}, error) {
return api.post("/set_group_admin", map[string]interface{}{
"user_id": user_id,
"group_id": group_id,
"enable": enable,
})
}

func (api *API) ConcurrentSetGroupAdmin(group_id int, user_id int, enable bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_admin", map[string]interface{}{
"user_id": user_id,
"group_id": group_id,
"enable": enable,
}, c)
return c
}

func (api *API) SetGroupAnonymous(group_id int, enable bool) (map[string]interface{}, error) {
return api.post("/set_group_anonymous", map[string]interface{}{
"group_id": group_id,
"enable": enable,
})
}

func (api *API) ConcurrentSetGroupAnonymous(group_id int, enable bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_anonymous", map[string]interface{}{
"group_id": group_id,
"enable": enable,
}, c)
return c
}

func (api *API) SetGroupCard(group_id int, user_id int, card string) (map[string]interface{}, error) {
return api.post("/set_group_card", map[string]interface{}{
"user_id": user_id,
"group_id": group_id,
"card": card,
})
}

func (api *API) ConcurrentSetGroupCard(group_id int, user_id int, card string) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_card", map[string]interface{}{
"user_id": user_id,
"group_id": group_id,
"card": card,
}, c)
return c
}

func (api *API) SetGroupLeave(group_id int, is_dismiss bool) (map[string]interface{}, error) {
return api.post("/set_group_leave", map[string]interface{}{
"group_id": group_id,
"is_dismiss": is_dismiss,
})
}

func (api *API) ConcurrentSetGroupLeave(group_id int, is_dismiss bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_leave", map[string]interface{}{
"group_id": group_id,
"is_dismiss": is_dismiss,
}, c)
return c
}

func (api *API) SetGroupSpecialTitle(group_id int, user_id int, special_title string) (map[string]interface{}, error) {
return api.post("/set_group_special_title", map[string]interface{}{
"user_id": user_id,
"group_id": group_id,
"special_title": special_title,
})
}

func (api *API) ConcurrentSetGroupSpecialTitle(group_id int, user_id int, special_title string) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_special_title", map[string]interface{}{
"user_id": user_id,
"group_id": group_id,
"special_title": special_title,
}, c)
return c
}

func (api *API) SetDiscussLeave(discuss_id int) (map[string]interface{}, error) {
return api.post("/set_discuss_leave", map[string]interface{}{
"discuss_id": discuss_id,
})
}

func (api *API) ConcurrentSetDiscussLeave(discuss_id int) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_discuss_leave", map[string]interface{}{
"discuss_id": discuss_id,
}, c)
return c
}

func (api *API) SetFriendAddRequest(flag string, approve bool, remark string) (map[string]interface{}, error) {
return api.post("/set_friend_add_request", map[string]interface{}{
"flag": flag,
"approve": approve,
"remark": remark,
})
}

func (api *API) ConcurrentSetFriendAddRequest(flag string, approve bool, remark string) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_friend_add_request", map[string]interface{}{
"flag": flag,
"approve": approve,
"remark": remark,
}, c)
return c
}

func (api *API) SetGroupAddRequest(flag string, _type string, approve bool, reason string) (map[string]interface{}, error) {
return api.post("/set_group_add_request", map[string]interface{}{
"flag": flag,
"type": _type,
"approve": approve,
"reason": reason,
})
}

func (api *API) ConcurrentSetGroupAddRequest(flag string, _type string, approve bool, reason string) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_group_add_request", map[string]interface{}{
"flag": flag,
"type": _type,
"approve": approve,
"reason": reason,
}, c)
return c
}

func (api *API) GetLoginInfo() (map[string]interface{}, error) {
return api.post("/get_login_info", map[string]interface{}{})
}

func (api *API) ConcurrentGetLoginInfo() chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_login_info", map[string]interface{}{}, c)
return c
}

func (api *API) GetStrangerInfo(user_id int, no_cache bool) (map[string]interface{}, error) {
return api.post("/get_stranger_info", map[string]interface{}{
"user_id": user_id,
"no_cache": no_cache,
})
}

func (api *API) ConcurrentGetStrangerInfo(user_id int, no_cache bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_stranger_info", map[string]interface{}{
"user_id": user_id,
"no_cache": no_cache,
}, c)
return c
}

func (api *API) GetGroupList() (map[string]interface{}, error) {
return api.post("/get_group_list", map[string]interface{}{})
}

func (api *API) ConcurrentGetGroupList() chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_group_list", map[string]interface{}{}, c)
return c
}

func (api *API) GetGroupMemberInfo(group_id int, user_id int, no_cache bool) (map[string]interface{}, error) {
return api.post("/get_group_member_info", map[string]interface{}{
"group_id": group_id,
"user_id": user_id,
"no_cache": no_cache,
})
}

func (api *API) ConcurrentGetGroupMemberInfo(group_id int, user_id int, no_cache bool) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_group_member_info", map[string]interface{}{
"group_id": group_id,
"user_id": user_id,
"no_cache": no_cache,
}, c)
return c
}

func (api *API) GetGroupMemberList(group_id int) (map[string]interface{}, error) {
return api.post("/get_group_member_list", map[string]interface{}{
"group_id": group_id,
})
}

func (api *API) ConcurrentGetGroupMemberList(group_id int) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_group_member_list", map[string]interface{}{
"group_id": group_id,
}, c)
return c
}

func (api *API) GetCookies() (map[string]interface{}, error) {
return api.post("/get_cookies", map[string]interface{}{})
}

func (api *API) ConcurrentGetCookies() chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_cookies", map[string]interface{}{}, c)
return c
}

func (api *API) GetCsrfToken() (map[string]interface{}, error) {
return api.post("/get_csrf_token", map[string]interface{}{})
}

func (api *API) ConcurrentGetCsrfToken() chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_csrf_token", map[string]interface{}{}, c)
return c
}

func (api *API) GetRecord(file string, out_format string) (map[string]interface{}, error) {
return api.post("/get_record", map[string]interface{}{
"file": file,
"out_format": out_format,
})
}

func (api *API) ConcurrentGetRecord(file string, out_format string) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_record", map[string]interface{}{
"file": file,
"out_format": out_format,
}, c)
return c
}

func (api *API) GetStatus() (map[string]interface{}, error) {
return api.post("/get_status", map[string]interface{}{})
}

func (api *API) ConcurrentGetStatus() chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_status", map[string]interface{}{}, c)
return c
}

func (api *API) GetVersionInfo() (map[string]interface{}, error) {
return api.post("/get_version_info", map[string]interface{}{})
}

func (api *API) ConcurrentGetVersionInfo() chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/get_version_info", map[string]interface{}{}, c)
return c
}

func (api *API) SetRestartPlugin() (map[string]interface{}, error) {
return api.post("/set_restart_plugin", map[string]interface{}{})
}

func (api *API) ConcurrentSetRestartPlugin() chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/set_restart_plugin", map[string]interface{}{}, c)
return c
}

func (api *API) CleanDataDir(data_dir string) (map[string]interface{}, error) {
return api.post("/clean_data_dir", map[string]interface{}{
"data_dir": data_dir,
})
}

func (api *API) ConcurrentCleanDataDir(data_dir string) chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/clean_data_dir", map[string]interface{}{
"data_dir": data_dir,
}, c)
return c
}

//试验性接口
func (api *API) GetFriendList() (map[string]interface{}, error) {
return api.post("/_get_friend_list", map[string]interface{}{})
}

func (api *API) ConcurrentGetFriendList() chan map[string]interface{} {
c := make(chan map[string]interface{}, 1)
go api.concurrentPost("/_get_friend_list", map[string]interface{}{}, c)
return c
}
