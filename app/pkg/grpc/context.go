package grpc

import (
	"context"
	"errors"
	"github.com/jinzhu/gorm"
)

const (
	Transaction = iota + 1
)

func GetTransaction(ctx context.Context) (*gorm.DB, error) {
	value := ctx.Value(Transaction)
	tx, ok := value.(*gorm.DB)
	if !ok {
		return nil, errors.New("error!!!")
	}
	return tx, nil
}
