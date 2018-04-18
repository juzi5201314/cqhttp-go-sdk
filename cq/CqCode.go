package cq

import "strconv"

func At(qq string) string {
	return "[CQ:at,qq=" + qq + "]"
}

func Face(faceid int) string {
	return "[CQ:face,id=" + string(faceid) + "]"
}

func Emoji(emojiid int) string {
	return "[CQ:emoji,id=" + string(emojiid) + "]"
}

func Bface(bfaceid int) string {
	return "[CQ:bface,id=" + string(bfaceid) + "]"
}

func Sface(sfaceid int) string {
	return "[CQ:sface,id=" + string(sfaceid) + "]"
}

func Image(file string) string {
	return "[CQ:image,id=" + file + "]"
}

func Record(file string, magic bool) string {
	return "[CQ:record,file=" + file + ",magic=" + strconv.FormatBool(magic) + "]"
}

func Rps() string {
	return "[CQ:rps]"
}

func Dice() string {
	return "[CQ:dice]"
}

func Shake() string {
	return "[CQ:Shake]"
}

func Anonymous(ignore bool) string {
	return "[CQ:anonymous,ignore=" + strconv.FormatBool(ignore) + "]"
}

func Music(_type string, id int) string {
	return "[CQ:music,type=" + _type + "，id=" + string(id) + "]"
}

func CustomMusic(url string, audio string, title string, content string, image string) string {
	return "[CQ:music,type=custom,url=" + url + ",audio=" + audio + ",title=" + title + ",content=" + content + ",image=" + image + "]"
}

func Share(url string, title string, content string, image string) string {
	return "[CQ:share,url=" + url + ",title=" + title + ",content=" + content + ",image=" + image + "]"
}
