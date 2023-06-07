package claude

import (
	"aichat/global"
	"errors"
	"fmt"
	"html"
	"strings"
	"time"

	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/ehttp"
	"github.com/duolabmeng6/goefun/etool"
)

type H map[string]interface{}

var headers = H{}
var token = ""
var 机器人id = ""
var 房间id = ""

func Claude加载配置() {
	token = global.G.Claude_token
	机器人id = global.G.Claude_botid
	房间id = global.G.Claude_roomid
	headers = H{
		"Cookie": global.G.Claude_cookie,
	}
}

func 提问(m_ts string, 提问内容 string) (string, string) {
	提问内容 = html.EscapeString(提问内容)

	data := H{
		"token":                      token,
		"channel":                    房间id,
		"ts":                         "1681546073.xxxxx5",
		"type":                       "message",
		"unfurl":                     "[]",
		"blocks":                     `[{"type":"rich_text","elements":[{"type":"rich_text_section","elements":[{"type":"user","user_id":"` + 机器人id + `"},{"type":"text","text":" ` + 提问内容 + `"}]}]}]`,
		"include_channel_perm_error": "true",
		"_x_reason":                  "webapp_message_send",
		"_x_mode":                    "online",
		"_x_sonic":                   "true",
	}
	if m_ts != "" {
		data["thread_ts"] = m_ts
	}
	eh := ehttp.NewHttp()

	b, _ := eh.PostFile("https://app.slack.com/api/chat.postMessage", data, headers)

	jsondata := string(b)
	//println(jsondata)

	new_m_ts := etool.Json解析文本(jsondata, "message.thread_ts")
	if new_m_ts == "" {
		new_m_ts = etool.Json解析文本(jsondata, "message.ts")
	}
	if new_m_ts == "" {
		fmt.Println("提问失败", jsondata)
		return "", "获取回答失败"
	}
	for i := 0; i < 90; i++ {
		fmt.Println("等待回答", i, new_m_ts)
		time.Sleep(time.Second * 2)
		ret, _ := 获取回答内容(new_m_ts)
		if ret != "" {
			return new_m_ts, ret
		}
	}
	return new_m_ts, "a"
}

func 获取回答内容(m_thread_ts string) (string, error) {

	url_file := H{
		"token":     token,
		"channel":   房间id,
		"ts":        m_thread_ts,
		"inclusive": "ture",
		"limit":     "28",
		"oldest":    m_thread_ts,
		"_x_reason": "history-api/fetchReplies",
		"_x_mode":   "online",
		"_x_sonic":  "true",
	}

	eh := ehttp.NewHttp()
	jsondata, _ := eh.Post("https://yierco.slack.com/api/conversations.replies", url_file, headers)
	//fmt.Println(jsondata)
	if strings.Contains(jsondata, "_Typing") {
		return "", errors.New("正在输入")
	}
	数组长度 := etool.Json解析文本(jsondata, "messages.#")
	数字1 := ecore.E到文本(ecore.E到整数(数组长度) - 2)
	数字2 := ecore.E到文本(ecore.E到整数(数组长度) - 1)

	user1 := etool.Json解析文本(jsondata, "messages."+数字1+".user")
	user2 := etool.Json解析文本(jsondata, "messages."+数字2+".user")
	text1 := etool.Json解析文本(jsondata, "messages."+数字1+".text")
	text2 := etool.Json解析文本(jsondata, "messages."+数字2+".text")
	//println("数字1", 数字1, "数字2", 数字2)
	//println("user1", user1)
	//println("user2", user2)
	//println("text1", text1)
	//println("text2", text2)
	回答内容 := ""
	if user1 == user2 {
		回答内容 = text1
	} else {
		回答内容 = text2
	}

	if strings.Contains(回答内容, "<@") {
		return "", errors.New("正在输入")
	}

	return 回答内容, nil

}
