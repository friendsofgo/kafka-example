package kafkaexample

type Message struct {
	Username string `json:"username"`
	Message  string `json:"message"`
}

const SystemID = "system"

// NewSystemMessage when we want to publish global message to all users in the room
func NewSystemMessage(message string) Message {
	return Message{
		Username: SystemID,
		Message:  message,
	}
}

// NewMessage prepare a message for publish into a stream
func NewMessage(username, message string) Message {
	return Message{username, message}
}
