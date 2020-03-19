package rules

import (
	"fmt"
	"sokolov-bot/method"
)

//复读机，每当有人复读消息，就跟着复读一次，

func Repeat(ch chan string, group string, url string) {
	var messageSlice []string = []string{"", "", "", ""}

	for message := range ch {
		for i := 0; i < 4; i++ {
			if i == 3 {
				_, err := method.PostJSON(
					fmt.Sprintf(jsonString, group, message),
					url,
				)
				if err != nil {
					fmt.Println(err)
				}
				messageSlice = []string{"", "", "", ""}
				break
			}
			if messageSlice[i] == message {
				continue
			} else {
				messageSlice[i] = message
				break
			}
		}
	}

}
