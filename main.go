package main

import (
	"aichat/aichat"
	"aichat/aichat/chatgpt"
	"aichat/aichat/claude"
	"aichat/global"
	"aichat/tool"
	"github.com/duolabmeng6/goefun/ecore"
	"github.com/duolabmeng6/goefun/egin"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	global.E加载环境变量("./.env")
	claude.Claude加载配置()

	// 创建一个Gin路由器
	r := gin.Default()
	r.GET("/", func(c *gin.Context) {
		c.String(200, "aichat to dingding")
	})

	// 定义一个GET请求的路由处理函数
	bot_chatgpt := aichat.New机器人连续聊天(chatgpt.New机器人连续聊天(global.G.Secret_id, global.G.Secret_key))
	bot_claude := aichat.New机器人连续聊天(claude.New机器人连续聊天())

	r.Group("/").
		Use(tool.DDChatMiddleware(bot_chatgpt, global.G.Dd_secret)).
		POST("/aichat/aigcaas")

	r.Group("/").
		Use(tool.DDChatMiddleware(bot_claude, global.G.Dd_secret2)).
		POST("/aichat/claude")

	r.GET("/login", func(c *gin.Context) {
		//输出一个登录页面的html 最简单的
		html := `
<html>
	<head>
		<title>登录</title>
		<meta name="viewport" content="width=device-width, initial-scale=1.0">
		<style>
			body {
				font-size: 16px;
				margin: 0;
				padding: 20px;
			}
			
			form {
				max-width: 400px;
				margin: 0 auto;
			}
			
			input[type="text"],
			input[type="submit"] {
				display: block;
				width: 100%;
				margin-bottom: 10px;
				padding: 10px;
				font-size: 16px;
				border-radius: 5px;
			}
			
			input[type="submit"] {
				background-color: #4CAF50;
				color: white;
				cursor: pointer;
			}
		</style>
	</head>
	<body>
		<form action="/login" method="get">
			<input type="text" name="key" placeholder="查看密码" />
			<input type="hidden" name="url" value="` + c.Query("url") + `" />
			<input type="submit" value="登录" />
		</form>
	</body>
</html>
`
		url := c.Query("url")
		key := c.Query("key")
		if key == "a" {
			//设置cookie
			c.SetCookie("key", key, 60*60*24*30, "/", "", false, true)
			c.Redirect(302, url)
		} else {
			//html头信息
			c.Header("Content-Type", "text/html; charset=utf-8")
			c.String(200, html)
		}
	})

	r.GET("/look", func(c *gin.Context) {
		//检查cookie是否有登录信息
		cookie, err := c.Cookie("key")
		if err != nil {
			c.Redirect(302, "/login?url="+c.Request.URL.String())
			return
		}
		println(cookie)
		// 打印所有请求的头信息
		//ecore.E调试输出(c.Request.Header)

		data := egin.I(c, "data", "")
		内容 := ecore.E读入文本("./回答数据/" + data + ".txt")
		//内容 = html.EscapeString(内容)

		内容 = tool.MarkdownToHtml(内容)
		c.Header("Content-Type", "text/html; charset=utf-8")
		c.String(http.StatusOK, 内容)
	})

	// 启动Gin服务
	r.Run("0.0.0.0:" + global.G.PORT)
}
