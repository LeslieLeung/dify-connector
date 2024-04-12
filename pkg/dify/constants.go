package dify

const (
	apiChatMessages       = "/v1/chat-messages"
	apiCompletionMessages = "/v1/completion-messages"
	apiMessages           = "/v1/messages"
	apiConversations      = "/v1/conversations"
	apiParameters         = "/v1/parameters"
)

const (
	OutputModeBlocking  = "blocking"
	OutputModeStreaming = "streaming"
)

const (
	AppTypeTextGenerator = iota
	AppTypeChatApp
	AppTypeWorkflow
)
