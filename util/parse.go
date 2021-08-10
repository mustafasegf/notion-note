package util

import (
	"strings"
)

func CheckIfCommand(body string) (valid bool) {
	if body[0:1] == "/" {
		valid = true
	}
	return
}

func GetCommand(body string) (command string) {
	lines := strings.SplitN(body, " ", 2)
	command = lines[0][1:]
	return
}

func ParseText(body string) (title, content string) {
	lines := strings.SplitN(body[6:], "\n", 2)
	title = lines[0]
	content = lines[1]
	return
}
