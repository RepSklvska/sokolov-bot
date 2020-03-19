package rules

import (
	"fmt"
	"sokolov-bot/method"
	"time"
)

//每到八点钟，就提醒群员八点了

func EightOClock(group string, url string) {
	for {
		switch time.Now().Format("15:04:05") {
		case "08:00:00":
			_, err := method.PostJSON(
				fmt.Sprintf(jsonString, group, "八点了！"),
				url,
			)
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Minute * 59)
		case "20:00:00":
			_, err := method.PostJSON(
				fmt.Sprintf(jsonString, group, "八点了！"),
				url,
			)
			if err != nil {
				fmt.Println(err)
			}
			time.Sleep(time.Minute * 59)
		}
	}
}
