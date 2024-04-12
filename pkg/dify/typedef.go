package dify

// credit: https://github.com/langgenius/dify-sdk-go

type ErrorResponse struct {
	Code    string `json:"code"`
	Message string `json:"message"`
	Params  string `json:"params"`
}

type CompletionMessageRequest struct {
	Inputs         map[string]interface{} `json:"inputs"`
	ResponseMode   string                 `json:"response_mode"`
	ConversationID string                 `json:"conversation_id,omitempty"`
	User           string                 `json:"user"`
}

type ChatCompletionMessageResponse struct {
	ID     string `json:"id"`
	Answer string `json:"answer"`
}

type ChatMessageRequest struct {
	Inputs         map[string]interface{} `json:"inputs"`
	Query          string                 `json:"query"`
	ResponseMode   string                 `json:"response_mode"`
	ConversationID string                 `json:"conversation_id,omitempty"`
	User           string                 `json:"user"`
}

type ChatMessageResponse struct {
	ID             string `json:"id"`
	Answer         string `json:"answer"`
	ConversationID string `json:"conversation_id"`
	CreatedAt      int    `json:"created_at"`
}

type ChatMessageStreamResponse struct {
	Event          string `json:"event"`
	TaskID         string `json:"task_id"`
	ID             string `json:"id"`
	Answer         string `json:"answer"`
	CreatedAt      int64  `json:"created_at"`
	ConversationID string `json:"conversation_id"`
}

type ChatMessageStreamChannelResponse struct {
	ChatMessageStreamResponse
	Err error `json:"-"`
}

type MessagesFeedbacksRequest struct {
	MessageID string `json:"message_id,omitempty"`
	Rating    string `json:"rating,omitempty"`
	User      string `json:"user"`
}

type MessagesFeedbacksResponse struct {
	HasMore bool                            `json:"has_more"`
	Data    []MessagesFeedbacksDataResponse `json:"data"`
}

type MessagesFeedbacksDataResponse struct {
	ID             string `json:"id"`
	Username       string `json:"username"`
	PhoneNumber    string `json:"phone_number"`
	AvatarURL      string `json:"avatar_url"`
	DisplayName    string `json:"display_name"`
	ConversationID string `json:"conversation_id"`
	LastActiveAt   int64  `json:"last_active_at"`
	CreatedAt      int64  `json:"created_at"`
}

type MessagesRequest struct {
	ConversationID string `json:"conversation_id"`
	FirstID        string `json:"first_id,omitempty"`
	Limit          int    `json:"limit"`
	User           string `json:"user"`
}

type MessagesResponse struct {
	Limit   int                    `json:"limit"`
	HasMore bool                   `json:"has_more"`
	Data    []MessagesDataResponse `json:"data"`
}

type MessagesDataResponse struct {
	ID             string                 `json:"id"`
	ConversationID string                 `json:"conversation_id"`
	Inputs         map[string]interface{} `json:"inputs"`
	Query          string                 `json:"query"`
	Answer         string                 `json:"answer"`
	Feedback       interface{}            `json:"feedback"`
	CreatedAt      int64                  `json:"created_at"`
}

type ConversationsRequest struct {
	LastID string `json:"last_id,omitempty"`
	Limit  int    `json:"limit"`
	User   string `json:"user"`
}

type ConversationsResponse struct {
	Limit   int                         `json:"limit"`
	HasMore bool                        `json:"has_more"`
	Data    []ConversationsDataResponse `json:"data"`
}

type ConversationsDataResponse struct {
	ID        string            `json:"id"`
	Name      string            `json:"name"`
	Inputs    map[string]string `json:"inputs"`
	Status    string            `json:"status"`
	CreatedAt int64             `json:"created_at"`
}

type ConversationsRenamingRequest struct {
	ConversationID string `json:"conversation_id,omitempty"`
	Name           string `json:"name"`
	User           string `json:"user"`
}

type ConversationsRenamingResponse struct {
	Result string `json:"result"`
}

type ParametersRequest struct {
	User string `json:"user"`
}

type ParametersResponse struct {
	OpeningStatement              string        `json:"opening_statement"`
	SuggestedQuestions            []interface{} `json:"suggested_questions"`
	SuggestedQuestionsAfterAnswer struct {
		Enabled bool `json:"enabled"`
	} `json:"suggested_questions_after_answer"`
	MoreLikeThis struct {
		Enabled bool `json:"enabled"`
	} `json:"more_like_this"`
	UserInputForm []map[string]interface{} `json:"user_input_form"`
}
