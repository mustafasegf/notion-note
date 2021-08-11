package util

import (
	"strings"
)

var (
	Help = `Create new note
/note [title]
[content]

Append to last note
/add [content]

Set notion token
/token [token]

Set notion database id
/page [database_id]`
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

func ParseTextOne(body string) (content string) {
	lines := strings.FieldsFunc(body, func(r rune) bool { return r == ' ' || r == '\n' })
	if len(lines) > 0 {
		content = lines[1]
	}
	return
}
