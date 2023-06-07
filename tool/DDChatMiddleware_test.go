package tool

import (
	"aichat/aichat"
	"aichat/aichat/claude"
	"fmt"
	"testing"
)

func testxx(chatbot *aichat.E机器人连续聊天, k string) {
	var chatbotNew *aichat.E机器人连续聊天
	v, _ := Global_bot.Load(chatbot.GetName() + k)
	if v == nil {
		chatbotNew = chatbot.Clone()
		Global_bot.Store(chatbot.GetName()+k, chatbotNew)
	} else {
		chatbotNew = v.(*aichat.E机器人连续聊天)
	}

	机器人回答 := chatbotNew.E发送消息("收到消息")

	fmt.Println(k, "======", v, &chatbot.E机器人, "机器人回答", 机器人回答)
}
func TestDDChatMiddleware(t *testing.T) {
	//bot1 := aichat.New机器人连续聊天(chatgpt.New机器人连续聊天(global.G.Secret_id, global.G.Secret_key))
	bot2 := aichat.New机器人连续聊天(claude.New机器人连续聊天())

	//fmt.Println("bot1", bot1.GetName())
	fmt.Println("bot2", bot2.GetName())
	testxx(bot2, "k1")
	testxx(bot2, "k2")
	testxx(bot2, "k1")
	testxx(bot2, "k2")
	testxx(bot2, "k1")
	testxx(bot2, "k2")
	testxx(bot2, "k1")
	testxx(bot2, "k2")
}
