package services

import (
	"github.com/jin-cg/grpc-demo/app/demo/repositories"
	"github.com/jinzhu/gorm"
)

type Service struct {
	dao *repositories.Repository
}

func NewService(db *gorm.DB) *Service {
	return &Service{
		dao: repositories.NewRepository(db),
	}
}
