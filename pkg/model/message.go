package model

var MESSAGES []string = []string{
	MessageNone,
}

var (
	MessageNone string = ""
)

func GetMessages() []string {
	return MESSAGES
}
