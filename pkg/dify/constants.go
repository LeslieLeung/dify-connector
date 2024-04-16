package dify

const (
	apiChatMessages           = "/v1/chat-messages"
	apiChatMessagesStop       = "/v1/chat-messages/{task_id}/stop"
	apiCompletionMessages     = "/v1/completion-messages"
	apiCompletionMessagesStop = "/v1/completion-messages/{task_id}/stop"
	apiWorkflowRun            = "/v1/workflow/run"
	apiWorkflowRunStop        = "/v1/workflow/run/{task_id}/stop"
	apiMessages               = "/v1/messages"
	apiMessagesFeedbacks      = "/v1/messages/{message_id}/feedbacks"
	apiConversations          = "/v1/conversations"
	apiParameters             = "/v1/parameters"
	apiSuggested              = "/v1/messages/{message_id}/suggested"
)

const (
	ResponseModeBlocking  = "blocking"
	ResponseModeStreaming = "streaming"
)

const (
	RatingLike    = "like"
	RatingDislike = "dislike"
	RatingNull    = "null"
)

const (
	AppTypeTextGenerator = iota
	AppTypeChatApp
	AppTypeWorkflow
)
