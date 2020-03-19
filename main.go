package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io/ioutil"
	"net/url"
	"sokolov-bot/method"
	"sokolov-bot/rules"
	"strings"
)

const (
	URL = "http://localhost:9999"
)

func main() {
	go rules.EightOClock("572808546", URL)
	messageChannelXiaoDian := make(chan string)
	go rules.Repeat(messageChannelXiaoDian, "572808546", URL)

	router := gin.Default()

	router.POST("/", func(context *gin.Context) {
		data, err := ioutil.ReadAll(context.Request.Body)
		if err != nil {
			fmt.Println(err)
		}
		jsonString, err := url.QueryUnescape(string(data))
		if err != nil {
			fmt.Println(err)
		}
		jsonString, err = method.UnescapeUnicode(jsonString)
		if err != nil {
			fmt.Println(err)
		}

		if strings.Contains(jsonString, `"group":572808546`) &&
			!strings.Contains(jsonString, `,"imageInfo":`) &&
			strings.Contains(jsonString, `"type":2`) &&
			strings.Contains(jsonString, `"subtype":1`) {
			start := strings.Index(jsonString, `"msg":`) + 7
			end := strings.Index(jsonString, `,"font"`) - 1
			message := jsonString[start:end]
			messageChannelXiaoDian <- message
			fmt.Println(message)
		}
		fmt.Println("Received a post:")
		fmt.Println(url.QueryUnescape(string(data)))
	})

	router.Run(":3002")
}
