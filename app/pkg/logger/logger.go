package logger

import (
	"context"
	"log"
)

func Errorf(ctx context.Context, format string, args ...interface{}) {
	//todo ctx識別子をログに反映する
	log.Printf(format, args...)
}
