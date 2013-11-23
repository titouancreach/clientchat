package parser

import (
	"regexp"
	"strings"
)

func ParseMessage(msg string) string {
	pattern, _ := regexp.Compile("^<img.*[0-9]{2}:[0-9]{2} ")
	newLine, _ := regexp.Compile("<br />")
	pseudo, _ := regexp.Compile("&gt;")

	msg = pattern.ReplaceAllString(msg, "")
	msg = newLine.ReplaceAllString(msg, "\n")
	msg = pseudo.ReplaceAllString(msg, "> ")
	return msg
}

func isHook(r rune) bool {
	return r == ']' || r == '['
}

func IsHL(msg string, pseudo string) (bool, string) {
	pattern, _ := regexp.Compile(pseudo + ">")

	if pattern.MatchString(msg) {
		sender := strings.FieldsFunc(msg, isHook)
		return true, sender[0]
	}
	return false, ""
}

func IsNoticed(msg string) (bool, string) {
	pattern, _ := regexp.Compile("^Notice")

	if pattern.MatchString(msg) {
		sender := strings.FieldsFunc(msg, isHook)
		return true, sender[1]
	}
	return false, ""
}

func IsJoined(msg string) (bool, string) {
	pattern, _ := regexp.Compile("[0-9]{2}:[0-9]{2}.*(a rejoint)")

	if pattern.MatchString(msg) {
		stripHour, _ := regexp.Compile(`^\d{2}:\d{2} `)
		stripMessage, _ := regexp.Compile(` a rejoint.*`)

		msg = stripHour.ReplaceAllString(msg, "")
		msg = stripMessage.ReplaceAllString(msg, "")

		msg = strings.Replace(msg, "\n", "", -1)

		return true, msg
	}

	return false, ""
}
