package irc

import (
	"strings"
)

func (e *Event) IsChannelMessage() bool {
	if strings.HasPrefix(e.Arguments[0], "#") {
		return true
	}
	return false
}

func (e *Event) GetMessageSource() string {
	if len(e.Arguments) == 0 {
		return e.Message()
	}
	if strings.HasPrefix(e.Arguments[0], "#") {
		return e.Arguments[0]
	}
	return e.Nick
}
