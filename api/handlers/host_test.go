package handlers

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func TestHostHandler_GetHosts(t *testing.T) {
	type args struct {
		c *gin.Context
	}
	tests := []struct {
		name string
		h    *HostHandler
		args args
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.h.GetHosts(tt.args.c)
		})
	}
}
