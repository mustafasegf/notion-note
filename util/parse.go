package util

import (
	"strings"
)

func ParseText(body string) (title, content string) {
	if body[0:6] == "/note " {
		lines := strings.SplitN(body, "\n", 2)
		title = lines[0][6:]
		content = strings.ReplaceAll(lines[1], "\n", "")
	}
	return
}
