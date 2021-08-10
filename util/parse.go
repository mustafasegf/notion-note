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
	if body[0:6] == "/note " {
		lines := strings.SplitN(body, "\n", 2)
		title = lines[0][6:]
		content = strings.ReplaceAll(lines[1], "\n", "")
	}
	return
}
