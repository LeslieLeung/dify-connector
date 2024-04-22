package controller

import (
	"github.com/gin-gonic/gin"
	"github.com/leslieleung/dify-connector/internal/api/typedef"
	"github.com/leslieleung/dify-connector/internal/database"
	"github.com/leslieleung/dify-connector/internal/misc"
	"net/http"
)

func Model(c *gin.Context) {
	resp := typedef.ModelResponse{Object: "list"}

	list, err := database.GetEnabledApps(misc.ToContext(c))
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	modelList := make([]typedef.Model, 0)
	for _, app := range list {
		modelList = append(modelList, typedef.Model{
			ID:      app.Name,
			Object:  "model",
			Created: app.CreatedAt.Unix(),
			OwnedBy: "dify",
		})
	}
	resp.Data = modelList
	c.JSON(http.StatusOK, resp)
}
