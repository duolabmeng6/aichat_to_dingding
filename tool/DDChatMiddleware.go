package tool

import (
	"aichat/aichat"
	. "aichat/global"
	"fmt"
	"github.com/duolabmeng6/goefun/dingding"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"sync"
)

var Global_bot sync.Map

// DDChatMiddleware 钉钉消息统一处理中间件
func DDChatMiddleware(chatbot *aichat.E机器人连续聊天, ddsecret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		var data map[string]interface{}
		if err := c.ShouldBindJSON(&data); err != nil {
			log.Println("Error parsing JSON:", err)
			c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid JSON"})
			return
		}
		//ecore.E写到文件("test.json", []byte(ecore.E到文本(data)))
		var msg E钉钉消息
		ecore.E到结构体(data, &msg)

		收到消息 := ecore.E删首尾空(msg.Text.Content)
		发送人名称 := msg.SenderNick
		fmt.Println("发送人", 发送人名称, "消息", 收到消息)

		var chatbotNew *aichat.E机器人连续聊天
		v, _ := Global_bot.Load(chatbot.GetName() + msg.SenderID)
		if v == nil {
			chatbotNew = chatbot.Clone()
			Global_bot.Store(chatbot.GetName()+msg.SenderID, chatbotNew)
		} else {
			chatbotNew = v.(*aichat.E机器人连续聊天)
		}

		if ecore.E取文本左边(收到消息, 2) == "清空" {
			chatbotNew.E清空对话()
			c.JSON(200, gin.H{"msgtype": "text", "text": gin.H{"content": "已清空"}})
			return
		}

		if ecore.E取文本左边(收到消息, 3) == "自定义" {
			chatbotNew.E清空对话()
			预设 := ecore.StrCut(收到消息, "自定义$")
			chatbotNew.E设定聊天内容(预设)
			c.JSON(200, gin.H{"msgtype": "text", "text": gin.H{"content": "已设定," + 预设}})
			return
		}

		if ecore.E取文本左边(收到消息, 2) == "帮助" {
			c.JSON(200, gin.H{"msgtype": "text", "text": gin.H{"content": `有以下指令:
清空 - 清空对话
自定义预设 - 自定义你是一名golang程序员
查看密码为 a
`}})
			return
		}

		机器人回答 := chatbotNew.E发送消息(收到消息)
		fmt.Println("发送人", 发送人名称, "消息", 收到消息, "机器人回答", 机器人回答)

		//机器人回答 := chatbot.GetName()
		回答链接 := ""
		if ecore.E判断文本(机器人回答, "```", "`") || ecore.E判断文本(收到消息, "代码") {
			ecore.E写到文件("./回答数据/"+ecore.E到文本(msg.CreateAt)+".txt", []byte(ecore.E到文本(机器人回答)))
			回答链接 = fmt.Sprintf("[查看原文本](%slook?data=%s)", G.Look_url, ecore.E到文本(msg.CreateAt))
		}

		新回答 := fmt.Sprintf("%s\n\n%s", 机器人回答, 回答链接)
		//fmt.Println("回答", 新回答)

		钉钉机器人 := dingding.New钉钉机器人(msg.SessionWebhook, ddsecret)
		钉钉机器人.E发送markdown消息("回答", 新回答, 2, msg.SenderStaffID)

		c.JSON(200, gin.H{"msgtype": "empty"})
		c.Next()
	}
}
