# aichat_to_dingding

目前接入了chatgpt 和 claude 支持 钉钉群聊@ 和 私聊

# 使用方式

build.sh 构建应用

把 dist 目录的文件放到服务器

配置环境变量

```api
dd_secret=钉钉企业机器人密钥 这个是chatgpt
dd_secret2=钉钉企业机器人密钥2 这个是claude
secret_id= https://www.aigcaas.cn/ 申请的id
secret_key= 上面网站申请的key
claude_cookie=https://slack.com/ 创建工作组以后抓包获取
claude_token=不会的去网上看教程
claude_botid=机器人id
claude_roomid=房间id
look_url=http://xxx:6699/ 用于查看代码有代码高亮的
#PORT=6689 运行端口 调试时候打开 如果是部署 在 docker-compose.yml 里面配置
```

运行 `docker-compose up -d` 启动服务

钉钉开放平台配置

创建好应用以后,进入应用功能,消息推送,

填入 消息接收地址 为 http://xxx:6699/aichat/claude 这个就是claude机器人

填入 消息接收地址 为 http://xxx:6699/aichat/aigcaas 这个就是chatgpt机器人

# 免责声明

本代码仅供学习交流使用，不得用于商业用途，如有侵权，请联系作者删除.
