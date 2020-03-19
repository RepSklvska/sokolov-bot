package rules

//看见别人发特定的话，就直接复读，尝试带领更多人复读

func RepeatLeader(ch chan string, group string, url string) {

	for message := range ch {
		_ = message
	}
}
