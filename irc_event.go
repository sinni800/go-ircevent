package irc

import (
	"strings"
)

func (e *IRCEvent) IsChannelMessage() bool {
	if strings.HasPrefix(e.Arguments[0], "#") {
		return true
	}
	return false
}

func (e *IRCEvent) GetMessageSource() string {
	if len(e.Arguments) == 0 {
		return e.Message()
	}
	if strings.HasPrefix(e.Arguments[0], "#") {
		return e.Arguments[0]
	}
	return e.Nick
}
