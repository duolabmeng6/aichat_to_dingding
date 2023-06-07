package claude

import (
	"aichat/aichat"
	"html"
)

type E机器人连续聊天 struct {
	系统提示 string
	id   string
	Name string
}

func New机器人连续聊天() *E机器人连续聊天 {
	return &E机器人连续聊天{}
}
func (c *E机器人连续聊天) Clone() aichat.E机器人连续聊天接口 {
	newChatbot := *c
	return &newChatbot
}

func (c *E机器人连续聊天) GetName() string {
	return "claude"
}
func (c *E机器人连续聊天) SetName(Name string) {
	c.Name = Name
}

func (c *E机器人连续聊天) E设定聊天内容(聊天内容 string) {
	c.系统提示 = 聊天内容
}

func (c *E机器人连续聊天) E清空对话() {
	c.id = ""
	c.E设定聊天内容("")
}

func (c *E机器人连续聊天) _获取机器人回答(内容 string) string {

	//message := "oldId:" + c.id
	//c.id = ecore.E文本取随机字母(3, 1)
	//message += " newId:" + c.id
	//
	//return message
	var message string
	c.id, message = 提问(c.id, 内容)
	return message
}

func (c *E机器人连续聊天) E发送消息(问题 string) string {
	if c.系统提示 != "" {
		问题 = c.系统提示 + 问题
		c.系统提示 = ""
	}

	机器人回答 := c._获取机器人回答(问题)
	机器人回答 = html.UnescapeString(机器人回答)
	//ecore.E调试输出(机器人回答)
	return 机器人回答
}
