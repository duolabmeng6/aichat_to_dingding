package chatgpt

import (
	"aichat/aichat"
	"github.com/duolabmeng6/goefun/ecore"
)

type E机器人连续聊天 struct {
	问答列表       []H
	问答轮数       int
	系统提示       H
	Secret_id  string
	Secret_key string
	Name       string
}

func New机器人连续聊天(Secret_id, Secret_key string) *E机器人连续聊天 {
	return &E机器人连续聊天{
		问答轮数:       5,
		Secret_id:  Secret_id,
		Secret_key: Secret_key,
	}
}
func (c *E机器人连续聊天) Clone() aichat.E机器人连续聊天接口 {
	newChatbot := *c
	return &newChatbot
}
func (c *E机器人连续聊天) GetName() string {
	return "chatgpt"
}

func (c *E机器人连续聊天) SetName(Name string) {
	c.Name = Name
}

func (c *E机器人连续聊天) E设定聊天内容(聊天内容 string) {
	c.问答列表 = []H{}
	c.系统提示 = H{"role": "system", "content": 聊天内容}
	//c.问答列表 = append(c.问答列表, H{"role": "system", "content": 聊天内容})
}

func (c *E机器人连续聊天) E清空对话() {
	c.E设定聊天内容("你是AI.请直接回答.使用markdown格式回答")
}

func (c *E机器人连续聊天) _获取机器人回答(内容 []H) string {
	//message := "I am chatgpt"
	//return message
	message, _ := 连续聊天(内容, c.Secret_id, c.Secret_key)
	return message
}

func (c *E机器人连续聊天) E发送消息(问题 string) string {
	c.问答列表 = append(c.问答列表, H{"role": "user", "content": 问题})

	// 将 c.系统提示 加入到 c.问答列表 的第一个
	临时问答列表 := append([]H{c.系统提示}, c.问答列表...)

	机器人回答 := c._获取机器人回答(临时问答列表)
	c.问答列表 = append(c.问答列表, H{"role": "assistant", "content": 机器人回答})

	//ecore.E调试输出(临时问答列表, 机器人回答)
	if len(c.问答列表) > c.问答轮数*2 {
		c.问答列表 = c.问答列表[2:]
	}
	ecore.E调试输出(c.问答列表)
	ecore.E调试输出(len(c.问答列表))
	return 机器人回答
}
