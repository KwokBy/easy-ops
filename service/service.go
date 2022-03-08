package service

import "context"

type IDemoService interface {
	GetLongDemo(ctx context.Context) (string, error)
}

type IUserService interface {
}

type IHostService interface {
}

type IMirrorService interface {
}

type ITaskService interface {
}
