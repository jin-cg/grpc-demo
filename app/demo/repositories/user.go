package repositories

import (
	"context"
	"github.com/jin-cg/grpc-demo/app/model"
	"github.com/jin-cg/grpc-demo/app/pkg/grpc"
)

func (*Repository) FindOneUser(ctx context.Context, name string) (*model.User, error) {
	db, err := grpc.GetTransaction(ctx)
	if err != nil {
		return nil, err
	}
	db.First((&model.User{Name:name})
	return &model.User{
		ID:   1,
		Name: name,
	}
}
