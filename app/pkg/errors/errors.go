package errors

import (
	"context"
	"github.com/pkg/errors"
)

func New(ctx context.Context, msg string) error {
	//todo ctx中から識別子を取得する
	id := 1
	return errors.Errorf("%s id:%d", msg, id)
}
