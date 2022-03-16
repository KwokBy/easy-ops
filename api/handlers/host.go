package handlers

import "github.com/gin-gonic/gin"

type HostHandler struct {
}

func NewHostHandler() HostHandler {
	return HostHandler{}
}

// GetHosts
func (h *HostHandler) GetHosts(c *gin.Context) {

}
