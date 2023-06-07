package global

import (
	"fmt"
	"github.com/duolabmeng6/goefun/ecore"
)

var G struct {
	Dd_secret     string //问答4
	Dd_secret2    string //问答2
	Secret_id     string
	Secret_key    string
	Look_url      string
	PORT          string
	Claude_cookie string
	Claude_token  string
	Claude_botid  string
	Claude_roomid string
}

func E加载环境变量(配置文件路径 string) {
	fmt.Println("加载环境变量", 配置文件路径)
	if !ecore.E加载环境变量_从文件(配置文件路径) {
		panic("加载环境变量失败")
	}
	G.Dd_secret = ecore.E读环境变量("dd_secret")
	G.Dd_secret2 = ecore.E读环境变量("dd_secret2")
	G.Secret_id = ecore.E读环境变量("secret_id")
	G.Secret_key = ecore.E读环境变量("secret_key")
	G.Claude_cookie = ecore.E读环境变量("claude_cookie")
	G.Claude_token = ecore.E读环境变量("claude_token")
	G.Claude_botid = ecore.E读环境变量("claude_botid")
	G.Claude_roomid = ecore.E读环境变量("claude_roomid")

	G.Look_url = ecore.E读环境变量("look_url")
	G.PORT = ecore.E读环境变量("PORT")

}

type E钉钉消息 struct {
	AtUsers                   []AtUsers `json:"atUsers"`
	ChatbotCorpID             string    `json:"chatbotCorpId"`
	ChatbotUserID             string    `json:"chatbotUserId"`
	ConversationID            string    `json:"conversationId"`
	ConversationType          string    `json:"conversationType"`
	CreateAt                  int64     `json:"createAt"`
	IsAdmin                   bool      `json:"isAdmin"`
	MsgID                     string    `json:"msgId"`
	Msgtype                   string    `json:"msgtype"`
	RobotCode                 string    `json:"robotCode"`
	SenderCorpID              string    `json:"senderCorpId"`
	SenderID                  string    `json:"senderId"`
	SenderNick                string    `json:"senderNick"`
	SenderStaffID             string    `json:"senderStaffId"`
	SessionWebhook            string    `json:"sessionWebhook"`
	SessionWebhookExpiredTime int64     `json:"sessionWebhookExpiredTime"`
	Text                      Text      `json:"text"`
}
type AtUsers struct {
	DingtalkID string `json:"dingtalkId"`
}
type Text struct {
	Content string `json:"content"`
}
