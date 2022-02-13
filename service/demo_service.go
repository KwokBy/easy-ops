package service

import (
	"context"
	"github.com/KwokBy/easy-ops/repo"
)

type demoService struct {
	demoRepo repo.IDemoRepo
}

func NewDemoService(demoRepo repo.IDemoRepo) IDemoService {
	return &demoService{
		demoRepo: demoRepo,
	}
}

func (d *demoService) GetLongDemo(ctx context.Context) (string,error) {
	demos,err:= d.demoRepo.GetDemos(ctx)
	if err!= nil {
		return "",err
	}
	for _,demo := range demos {
		if len(demo.Name) > 10 {
			return demo.Name,nil
		}
	}
	return "" ,nil
}