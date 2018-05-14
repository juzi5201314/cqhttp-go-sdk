package cqcode

import (
	"testing"
	"encoding/json"
)

func TestMessageSegment_ParseMedia(t *testing.T) {

	seg := MessageSegment{
		Type: "text",
		Data: map[string]interface{}{
			"text": "test text message",
		},
	}

	var text Text
	seg.ParseMedia(&text)

	if text.Text == "test text message" {
		t.Log("Decode text passed")
	} else {
		t.Errorf("Decode text failed: %v", text.Text)
	}

}

func TestParseMessageFromString(t *testing.T) {

	str := "&#91;he&#44;ym[CQ:at,qq=123&#44;456][CQ:face,id=14] \nSee this awesome image, [CQ:image,file=1.jpg] Isn't it cool? [CQ:shake]\n"

	mes, err := ParseMessageFromString(str)

	if err != nil {
		t.Fatalf("Decode text failed: %v", err)
	}

	res, _ := json.Marshal(mes)

	jsonstr := string(res)

	if string(res) == `[{"type":"text","data":{"text":"[he,ym"}},{"type":"at","data":{"qq":"123,456"}},{"type":"face","data":{"id":"14"}},{"type":"text","data":{"text":" \nSee this awesome image, "}},{"type":"image","data":{"file":"1.jpg"}},{"type":"text","data":{"text":" Isn't it cool? "}},{"type":"shake","data":{}},{"type":"text","data":{"text":"\n"}}]` {
		t.Log("Decode text passed")
	} else {
		t.Errorf("Decode text failed: %v", jsonstr)
	}

}

func TestMessage_Append(t *testing.T) {

	music := Music{
		Type:     "custom",
		ShareURL: "http://localhost:8080",
	}

	m := NewMessage()

	err := m.Append(&music)

	if err != nil {
		t.Fatalf("Decode text failed: %v", err)
	}

	res, _ := json.Marshal(m)

	jsonstr := string(res)

	if string(res) == `[{"type":"music","data":{"audio":"","content":"","id":"","image":"","title":"","type":"custom","url":"http://localhost:8080"}}]` {
		t.Log("Append music passed")
	} else {
		t.Errorf("Append music failed: %v", jsonstr)
	}

}

func TestMessageSegment_CQString(t *testing.T) {

	rec := Record{
		FileID: "/data/audio/[,]&",
		Magic:  false,
	}

	seg, _ := NewMessageSegment(&rec)

	f := seg.CQString()

	shake := Shake{}

	seg, _ = NewMessageSegment(&shake)

	s := seg.CQString()

	text := Text{
		Text: "[,]&",
	}

	seg, _ = NewMessageSegment(&text)

	ts := seg.CQString()

	if f == "[CQ:record,file=/data/audio/&#91;&#44;&#93;&amp;,magic=false,url=]" && s == "[CQ:shake]" && ts == "&#91;,&#93;&amp;" {
		t.Log("Format CQString passed")
	} else {
		t.Errorf("Format CQString failed: %v %v %v", f, s, ts)
	}

}

func TestCommand(t *testing.T) {

	m := NewMessage()

	text := Text{
		Text: "/",
	}

	m.Append(&text)

	face := Face{
		FaceID: 170,
	}

	m.Append(&face)

	text = Text{
		Text: ` arg1 'a \'r 
g 2' "a \"r \\\"g 3\\" arg4
argemoji`,
	}

	m.Append(&text)

	emoji := Emoji{
		EmojiID: 10086,
	}

	m.Append(&emoji)

	text = Text{
		Text: ` arg5`,
	}

	m.Append(&text)

	music := Music{
		Content: "Alice\nLove\nBob",
	}

	m.Append(&music)

	StrictCommand = true

	if !m.IsCommand() {
		t.Error("Should be command")
	}

	cmd, args := m.Command()

	res, _ := json.Marshal(args)

	jsonstr := string(res)

	if cmd == "[CQ:face,id=170]" && jsonstr == `["arg1","a 'r \ng 2","a \"r \\\"g 3\\","arg4","argemoji[CQ:emoji,id=10086]","arg5[CQ:music,type=,id=,url=,audio=,title=,content=Alice\nLove\nBob,image=]"]` {
		t.Log("Good command")
	} else {
		t.Errorf("Parse command failed: %v", jsonstr)
	}

}
