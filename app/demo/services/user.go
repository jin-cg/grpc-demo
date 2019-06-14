package services

import "github.com/jin-cg/grpc-demo/app/model"

func (s *Service) GetUser(name string) *model.User {
	return s.dao.FindOneUser(name)
}
