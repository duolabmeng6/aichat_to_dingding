package claude

import (
	"aichat/global"
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/etool"
	"testing"
)

func Test_提问(t *testing.T) {
	id := ""
	id, ret := 提问(id, "给我5个随机的数字 1-100")
	println(id, ret)
	id, ret = 提问(id, "谁是最大的数")
	println(id, ret)
	id, ret = 提问(id, "谁是最小的数")

	//println(id, ret)
	//id, ret = 提问(id, "写200个字的作文关于游玩广州北京路")
	//
	//ret := 获取回答内容("1686074473.395889")
	//println(ret)

	//jsondata := `{"ok":true,"channel":"C056D1F9ZJS","ts":"1686073833.557589","message":{"type":"message","text":"<@U056KFS7SSF> abc","user":"U0579ABV1C0","ts":"1686073833.557589","blocks":[{"type":"rich_text","block_id":"zAsh","elements":[{"type":"rich_text_section","elements":[{"type":"user","user_id":"U056KFS7SSF"},{"type":"text","text":" abc"}]}]}],"team":"T056Y8MB8U9"}}`
	//m_ts := etool.Json解析文本(jsondata, "message.thread_ts")
	//if m_ts == "" {
	//	m_ts = etool.Json解析文本(jsondata, "message.ts")
	//}
	//println(m_ts)
}
func Test_获取回答内容(t *testing.T) {
	jsondata := `{"ok":true,"messages":[{"type":"message","text":"<@U056KFS7SSF> \u4f60\u597d","user":"U0579ABV1C0","ts":"1686074473.395889","blocks":[{"type":"rich_text","block_id":"z\/W","elements":[{"type":"rich_text_section","elements":[{"type":"user","user_id":"U056KFS7SSF"},{"type":"text","text":" \u4f60\u597d"}]}]}],"team":"T056Y8MB8U9","thread_ts":"1686074473.395889","reply_count":1,"reply_users_count":1,"latest_reply":"1686074474.855379","reply_users":["U056KFS7SSF"],"is_locked":false,"subscribed":true,"last_read":"1686074473.395889"},{"bot_id":"B056Y8PCTED","type":"message","text":" \u4f60\u597d!\u6211\u662f Claude,\u5f88\u9ad8\u5174\u8ba4\u8bc6\u4f60\u3002","user":"U056KFS7SSF","ts":"1686074474.855379","app_id":"A04KGS7N9A8","blocks":[{"type":"rich_text","block_id":"Hc1","elements":[{"type":"rich_text_section","elements":[{"type":"text","text":"\u4f60\u597d!\u6211\u662f Claude,\u5f88\u9ad8\u5174\u8ba4\u8bc6\u4f60\u3002"}]}]}],"team":"T056Y8MB8U9","bot_profile":{"id":"B056Y8PCTED","deleted":false,"name":"Claude","updated":1684461846,"app_id":"A04KGS7N9A8","icons":{"image_36":"https:\/\/avatars.slack-edge.com\/2023-01-25\/4682316783575_bbab0cdcdb3685eb5c87_36.png","image_48":"https:\/\/avatars.slack-edge.com\/2023-01-25\/4682316783575_bbab0cdcdb3685eb5c87_48.png","image_72":"https:\/\/avatars.slack-edge.com\/2023-01-25\/4682316783575_bbab0cdcdb3685eb5c87_72.png"},"team_id":"T056Y8MB8U9"},"edited":{"user":"B056Y8PCTED","ts":"1686074476.000000"},"thread_ts":"1686074473.395889","parent_user_id":"U0579ABV1C0"}],"has_more":false}`
	数组长度 := etool.Json解析文本(jsondata, "messages.#")
	数字1 := ecore.E到文本(ecore.E到整数(数组长度) - 1)
	数字2 := ecore.E到文本(ecore.E到整数(数组长度) - 2)

	user1 := etool.Json解析文本(jsondata, "messages."+数字1+".user")
	user2 := etool.Json解析文本(jsondata, "messages."+数字2+".user")
	text1 := etool.Json解析文本(jsondata, "messages."+数字1+".text")
	text2 := etool.Json解析文本(jsondata, "messages."+数字2+".text")
	println("数字1", 数字1, "数字2", 数字2)
	println("user1", user1)
	println("user2", user2)
	println("text1", text1)
	println("text2", text2)

}
func Test_连续聊天3(t *testing.T) {
	global.E加载环境变量("../../.env")
	Claude加载配置()
	chatbot := New机器人连续聊天()
	//chatbot.E清空对话()
	//回答 := chatbot.E发送消息("给我5个随机的数字 1-100")
	//fmt.Println(回答)
	//回答 = chatbot.E发送消息("谁是最大的数？")
	//fmt.Println(回答)
	//回答 = chatbot.E发送消息("谁是最小的数？")
	//fmt.Println(回答)
	//for i := 0; i < 10; i++ {
	//	回答 := chatbot.E发送消息(strconv.Itoa(i) + "你好")
	//	fmt.Println(回答)
	//}

	问题 := `
func (c *E机器人连续聊天) Clone() aichat.E机器人连续聊天接口 {
	newChatbot := *c
	return &newChatbot
}
解释一下这段golang代码
`

	回答 := chatbot.E发送消息(问题)
	fmt.Println(回答)

}
