package cqcode

import (
	"strings"
	"fmt"
	"github.com/pkg/errors"
	"github.com/mitchellh/mapstructure"
	"regexp"
)

var StrictCommand = false

type Message []MessageSegment

type MessageSegment struct {
	Type string        `json:"type" cq:"type"`
	Data CQKeyValueMap `json:"data" cq:"data"`
}

type CQKeyValueMap map[string]interface{}

func Decode(input, output interface{}) error {
	config := &mapstructure.DecoderConfig{
		Metadata:         nil,
		Result:           output,
		WeaklyTypedInput: true,
		TagName:          "cq",
	}

	decoder, err := mapstructure.NewDecoder(config)
	if err != nil {
		return err
	}

	return decoder.Decode(input)
}

func NewMessage() (Message) {
	return make(Message, 0)
}

func ParseMessage(msg interface{}) (Message, error) {
	switch x := msg.(type) {
	case string:
		return ParseMessageFromString(x)
	default:
		return ParseMessageFromArray(x)
	}
}

func ParseMessageFromArray(msg interface{}) (Message, error) {
	message := Message{}
	err := Decode(msg, message)
	if err != nil {
		return message, err
	}
	return message, nil
}

func ParseMessageFromString(str string) (Message, error) {
	message := make(Message, 0)
	res := regexp.MustCompile(`\[CQ:.*?\]`).FindAllStringSubmatchIndex(str, -1)
	i := 0
	for _, cqc := range res {
		if cqc[0] > i {
			// There is a text message before this cqc
			seg := MessageSegment{
				Type: "text",
				Data: map[string]interface{}{
					"text": str[i:cqc[0]],
				},
			}
			message = append(message, seg)
		}
		i = cqc[1]
		seg, err := NewMessageSegmentFromCQCode(str[cqc[0]:cqc[1]])
		if err != nil {
			continue
		}
		message = append(message, seg)
	}
	if len(str) > i {
		// There is a text message after all cqc
		seg := MessageSegment{
			Type: "text",
			Data: map[string]interface{}{
				"text": str[i:],
			},
		}
		message = append(message, seg)
	}
	return message, nil
}

func (m *Message) IsCommand() bool {
	str := m.CQString()
	return IsCommand(str)
}

func (m *Message) Command() (cmd string, args []string) {
	str := m.CQString()
	return Command(str)
}

func IsCommand(str string) bool {
	if len(str) == 0 {
		return false
	}
	if StrictCommand && str[:1] != "/" {
		return false
	}
	return true
}

func Command(str string) (cmd string, args []string) {
	strs := regexp.MustCompile(`'.*?'|".*?"|\S*\[CQ:.*?\]\S*|\S+`).FindAllString(str, -1)
	if len(strs) == 0 || len(strs[0]) == 0 {
		return
	}
	if StrictCommand {
		if strs[0][:1] != "/" {
			return
		}
		cmd = strs[0][1:]
	} else {
		cmd = strs[0]
	}
	args = strs[1:]
	return
}

func (m *Message) CQString() string {
	var str string
	for _, seg := range *m {
		str += seg.CQString()
	}
	return str
}

func (m *Message) Append(media Media) error {
	seg, err := NewMessageSegment(media)
	if err != nil {
		return err
	}
	*m = append(*m, seg)
	return nil
}

func NewMessageSegment(media Media) (MessageSegment, error) {
	seg := MessageSegment{}
	seg.Type = media.FunctionName()
	seg.Data = make(CQKeyValueMap)
	err := Decode(media, &seg.Data)
	return seg, err
}

func NewMessageSegmentFromCQCode(str string) (MessageSegment, error) {
	seg := MessageSegment{}
	seg.Data = make(CQKeyValueMap)
	l := len(str)
	if l <= 5 || str[:4] != "[CQ:" || str[len(str)-1:] != "]" {
		err := errors.New("invalid")
		return seg, err
	}
	str = str[4 : len(str)-1]
	strs := strings.Split(str, ",")
	for i, v := range strs {
		if i == 0 {
			seg.Type = strs[0]
		} else {
			kvstrs := strings.Split(v, "=")
			if len(kvstrs) == 0 {
				continue
			}
			seg.Data[kvstrs[0]] = strings.Join(kvstrs[1:], "=")
		}
	}
	return seg, nil
}

func (seg *MessageSegment) IsMedia(mediaType string) bool {
	return seg.Type == mediaType
}

func (seg *MessageSegment) ParseMedia(media Media) error {
	err := Decode(seg.Data, media)
	return err
}

func (seg *MessageSegment) CQString() string {
	if seg.Type == "text" {
		t, ok := seg.Data["text"]
		if !ok {
			return ""
		}
		text := fmt.Sprint(t)
		text = EncodeCQText(text)
		return text
	}
	strs := make([]string, 0)
	strs = append(strs, seg.Type)
	for k, v := range seg.Data {
		text := fmt.Sprint(v)
		text = EncodeCQCodeText(text)
		kvs := fmt.Sprintf("%s=%s", k, text)
		strs = append(strs, kvs)
	}
	str := strings.Join(strs, ",")
	str = fmt.Sprintf("[CQ:%s]", str)
	return str
}

type Media interface {
	FunctionName() string // 功能名
}

type Text struct {
	Text string `cq:"text"`
}

func (t *Text) FunctionName() string {
	return "text"
}

// Mention @
type At struct {
	QQ string `cq:"qq"` // Someone's QQ号, could be "all"
}

func (a *At) FunctionName() string {
	return "at"
}

// QQ表情
type Face struct {
	FaceID int `cq:"id"` // 1-170 (旧版), >170 (新表情)
}

func (f *Face) FunctionName() string {
	return "face"
}

// Emoji
type Emoji struct {
	EmojiID int `cq:"id"` // Unicode Dec
}

func (e *Emoji) FunctionName() string {
	return "emoji"
}

// 原创表情
type Bface struct {
	BfaceID int `cq:"id"`
}

func (b *Bface) FunctionName() string {
	return "bface"
}

// 小表情
type Sface struct {
	SfaceID int `cq:"id"`
}

func (s *Sface) FunctionName() string {
	return "sface"
}

// Image
type Image struct {
	FileID string `cq:"file"`
}

func (i *Image) FunctionName() string {
	return "image"
}

type Record struct {
	FileID string `cq:"file"`
	Magic  bool   `cq:"magic"`
}

func (r *Record) FunctionName() string {
	return "record"
}

const (
	Rock     = 1
	Paper    = 2
	Scissors = 3
)

// 猜拳魔法表情
type Rps struct {
	Type int `cq:"type"`
}

func (rps *Rps) FunctionName() string {
	return "rps"
}

// 掷骰子魔法表情
type Dice struct {
	Type int `cq:"type"` // 1-6
}

func (d *Dice) FunctionName() string {
	return "dice"
}

// 戳一戳
type Shake struct {
}

func (s *Shake) FunctionName() string {
	return "shake"
}

// 音乐
type Music struct {
	Type string `cq:"type"` // qq, 163, xiami
	// non-custom music
	MusicID string `cq:"id"` // id
	// custom music
	ShareURL string `cq:"url"`     // Link open on click
	AudioURL string `cq:"audio"`   // Link of audio
	Title    string `cq:"title"`   // Title
	Content  string `cq:"content"` // Description
	Image    string `cq:"image"`   // Link of cover image
}

func (m *Music) FunctionName() string {
	return "music"
}

func (m *Music) IsCustomMusic() bool {
	return m.Type == "custom"
}

// 分享链接
type Share struct {
	URL     string `cq:"url"`
	Title   string `cq:"title"`   // In 12 words
	Content string `cq:"content"` // In 30 words
	Image   string `cq:"image"`   // Link of cover image
}

func (s *Share) FunctionName() string {
	return "share"
}

// 位置
type Location struct {
}

func (l *Location) FunctionName() string {
	return "location"
}

// 厘米秀
type Show struct {
}

func (s *Show) FunctionName() string {
	return "show"
}

// 签到
type Sign struct {
}

func (s *Sign) FunctionName() string {
	return "sign"
}

// 其他富媒体
type Rich struct {
}

func (r *Rich) FunctionName() string {
	return "rich"
}

func EncodeCQText(str string) string {
	str = strings.Replace(str, "&", "&amp;", -1)
	str = strings.Replace(str, "[", "&#91;", -1)
	str = strings.Replace(str, "]", "&#93;", -1)
	return str
}

func DecodeCQText(str string) string {
	str = strings.Replace(str, "&#93;", "]", -1)
	str = strings.Replace(str, "&#91;", "[", -1)
	str = strings.Replace(str, "&amp;", "&", -1)
	return str
}

func EncodeCQCodeText(str string) string {
	str = strings.Replace(str, "&", "&amp;", -1)
	str = strings.Replace(str, "[", "&#91;", -1)
	str = strings.Replace(str, "]", "&#93;", -1)
	str = strings.Replace(str, ",", "&#44;", -1)
	return str
}

func DecodeCQCodeText(str string) string {
	str = strings.Replace(str, "&#44;", ",", -1)
	str = strings.Replace(str, "&#93;", "]", -1)
	str = strings.Replace(str, "&#91;", "[", -1)
	str = strings.Replace(str, "&amp;", "&", -1)
	return str
}
