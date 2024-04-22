package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/leslieleung/dify-connector/internal/api/typedef"
	"github.com/leslieleung/dify-connector/internal/database"
	db "github.com/leslieleung/dify-connector/internal/database/typedef"
	"github.com/leslieleung/dify-connector/internal/misc"
	dify2 "github.com/leslieleung/dify-connector/pkg/dify"
	"net/http"
	"time"
)

func ChatCompletion(c *gin.Context) {
	var req typedef.ChatCompletionRequest
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	// validate model
	enabledApps, err := database.GetEnabledApps(misc.ToContext(c))
	if err != nil {
		c.JSON(500, gin.H{"error": err.Error()})
		return
	}
	var dify *db.DifyApp
	for _, app := range enabledApps {
		if app.Name == req.Model {
			dify = app
			break
		}
	}
	if dify == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "model not found"})
		return
	}

	if dify.Type == dify2.AppTypeChatApp && !req.Stream {
		c.JSON(http.StatusBadRequest, gin.H{"error": "stream is required for chat app"})
		return
	}

	if req.User == "" {
		req.User = uuid.NewString()
	}

	switch dify.Type {
	case dify2.AppTypeTextGenerator:
		handleTextGenerator(c, req, dify)
	case dify2.AppTypeChatApp:
		handleChatApp(c, req, dify)
	default:
		c.JSON(http.StatusBadRequest, gin.H{"error": "unsupported model type"})
		return
	}
}

func handleTextGenerator(c *gin.Context, req typedef.ChatCompletionRequest, dify *db.DifyApp) {
	difyApp := dify2.New(dify.BaseURL, dify.APIKey)
	difyApp.SetDebug()
	if req.Stream {
		// TODO
		c.AbortWithStatusJSON(http.StatusNotImplemented, gin.H{"error": "streaming is not implemented"})
	}
	resp, err := difyApp.CompletionMessage(dify2.CompletionMessageRequest{
		Inputs: map[string]interface{}{
			"query": req.Messages[0].Content,
		},
		ResponseMode: dify2.ResponseModeBlocking,
		User:         req.User,
	})
	if err != nil {
		return
	}
	res := typedef.ChatCompletionResponse{
		ID:      uuid.NewString(), // TODO
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   req.Model,
		Choices: []typedef.ChatCompletionChoice{
			{
				Index: 0,
				Message: typedef.ChatCompletionMessage{
					Role:    "assistant",
					Content: resp.Answer,
				},
				FinishReason: typedef.FinishReasonStop,
			},
		},
	}
	c.AbortWithStatusJSON(http.StatusOK, res)
}

func streamHandler(c *gin.Context) {
	// TODO
}

func handleChatApp(c *gin.Context, req typedef.ChatCompletionRequest, dify *db.DifyApp) {
	difyApp := dify2.New(dify.BaseURL, dify.APIKey)
	resp, err := difyApp.ChatMessageStream(dify2.ChatMessageRequest{
		Inputs: map[string]interface{}{
			"query": req.Messages[0].Content,
		},
		User: req.User,
	})
	if err != nil {
		return
	}
	res, err := resp.Wait()
	if err != nil {
		return
	}

	r := typedef.ChatCompletionResponse{
		ID:      uuid.NewString(), // TODO
		Object:  "chat.completion",
		Created: time.Now().Unix(),
		Model:   req.Model,
		Choices: []typedef.ChatCompletionChoice{
			{
				Index: 0,
				Message: typedef.ChatCompletionMessage{
					Role:    "assistant",
					Content: res,
				},
				FinishReason: typedef.FinishReasonStop,
			},
		},
	}
	c.AbortWithStatusJSON(http.StatusOK, r)
}
