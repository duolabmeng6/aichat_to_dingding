package chatgpt

import (
	"aichat/global"
	"fmt"
	"github.com/duolabmeng6/goefun/etool"
	"testing"
)

func Test_连续聊天2(t *testing.T) {

	ret := `{"api_call_status":true,"choices":[{"finish_reason":"stop","index":0,"message":{"content":"\u4f60\u597d\uff01","role":"assistant"}}],"created":1686055304,"id":"chatcmpl-7OQIKNfLuTzmGjGiLaV9sbLfvA2yy","model":"gpt-3.5-turbo-0301","object":"chat.completion","usage":{"completion_tokens":3,"prompt_tokens":10,"total_tokens":13}}`
	回答 := etool.Json解析文本(ret, "choices.0.message.content")
	fmt.Println("回答??", 回答)

}

func Test_连续聊天3(t *testing.T) {
	global.E加载环境变量("../../.env")
	chatbot := New机器人连续聊天(global.G.Secret_id, global.G.Secret_key)
	chatbot.E清空对话()
	回答 := chatbot.E发送消息("给我5个随机的数字 1-100")
	fmt.Println(回答)
	回答 = chatbot.E发送消息("谁是最大的数？")
	fmt.Println(回答)
	回答 = chatbot.E发送消息("谁是最小的数？")
	fmt.Println(回答)
	//for i := 0; i < 10; i++ {
	//	回答 := chatbot.E发送消息(strconv.Itoa(i) + "你好")
	//	fmt.Println(回答)
	//}
}
