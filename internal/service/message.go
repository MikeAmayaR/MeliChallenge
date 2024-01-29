package service

func GetMessage(messages ...[]string) (msg string) {
	if len(messages) == 0 {
		return ""
	}

	maxLength := 0
	for _, m := range messages {
		if maxLength < len(m) {
			maxLength = len(m)
		}
	}

	completeMsg := make([]string, maxLength)
	for i := 0; i < maxLength; i++ {
		for _, m := range messages {
			if i < len(m) && m[i] != "" && completeMsg[i] == "" {
				completeMsg[i] = m[i]
				break
			}
		}
	}

	msg = ""
	for _, word := range completeMsg {
		if word != "" {
			msg += word + " "
		}
	}

	if len(msg) > 0 {
		msg = msg[:len(msg)-1]
	}

	return msg
}
