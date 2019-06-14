package interceptor

import (
	"context"
	pkgrpc "github.com/jin-cg/grpc-demo/app/pkg/grpc"
	"github.com/jin-cg/grpc-demo/app/pkg/logger"
	"github.com/jinzhu/gorm"
	"google.golang.org/grpc"
)

func DBInterceptor(db *gorm.DB) grpc.UnaryServerInterceptor {
	return func(ctx context.Context, req interface{}, info *grpc.UnaryServerInfo, handler grpc.UnaryHandler) (interface{}, error) {
		tx := db.BeginTx(ctx, nil)
		ctx = context.WithValue(ctx, pkgrpc.Transaction, tx)
		res, err := handler(ctx, req)
		if err != nil {
			if tx.Rollback(); tx.Error != nil {
				logger.Errorf(ctx, "transaction rollback failed, err:%+v", tx.Error)
			}
			return nil, err
		}
		if tx.Commit(); tx.Error != nil {
			logger.Errorf(ctx, "transaction commit failed, err:%+v", tx.Error)
			return nil, tx.Error
		}
		return res, nil
	}
}
