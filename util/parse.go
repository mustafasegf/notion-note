package util

import (
	"strings"
)

var (
	Help = `Create new note
/note [title]
[content]

Append to last note
/add [content]`
)

func CheckIfCommand(body string) (valid bool) {
	if body[0:1] == "/" {
		valid = true
	}
	return
}

func GetCommand(body string) (command string) {
	lines := strings.FieldsFunc(body, func(r rune) bool { return r == ' ' || r == '\n' })
	command = lines[0][1:]
	return
}

func ParseText(body string) (title, content string) {
	lines := strings.SplitN(body[6:], "\n", 2)
	title = lines[0]
	content = lines[1]
	return
}

func ParseTextAdd(body string) (content string) {
	content = body[5:]
	return
}
