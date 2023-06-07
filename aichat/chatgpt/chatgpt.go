package chatgpt

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"github.com/duolabmeng6/goefun/ehttp"
	"github.com/duolabmeng6/goefun/etool"
	"math/rand"
	"strconv"
	"time"
)

func getSha256(content string) string {
	hasher := sha256.New()
	hasher.Write([]byte(content))
	return hex.EncodeToString(hasher.Sum(nil))
}

type H map[string]interface{}

func 连续聊天(content []H, secret_id string, secret_key string) (string, error) {

	application_name := "chatgpt_chat"
	apiName := "chat_com"
	eh := ehttp.NewHttp()
	url := "https://api.aigcaas.cn/product/" + application_name + "/api/" + apiName

	data := H{
		"messages": content,
	}

	nonce := strconv.Itoa(rand.Intn(10000) + 1)
	timestamp := strconv.FormatInt(time.Now().Unix(), 10)
	token := getSha256(fmt.Sprintf("%s%s%s", timestamp, secret_key, nonce))

	headers := H{
		"SecretID":     secret_id,
		"Nonce":        nonce,
		"Token":        token,
		"Timestamp":    timestamp,
		"Content-Type": "application/json",
	}

	ret, _ := eh.Post(url, data, headers)
	回答 := etool.Json解析文本(ret, "choices.0.message.content")

	return 回答, nil
}
