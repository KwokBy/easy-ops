package service

import "context"

type IDemoService interface {
	GetLongDemo(ctx context.Context) (string,error)
}
