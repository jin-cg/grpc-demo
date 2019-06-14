package controllers

import (
	"context"
	"fmt"
	"github.com/jin-cg/grpc-demo/app/demo/proto"
)

func (c *Controller) Hei(ctx context.Context, req *proto.Request) (*proto.Response, error) {
	user := c.s.GetUser(req.Name)
	return &proto.Response{
		Content: fmt.Sprintf("%d-%s", user.ID, user.Name),
	}, nil
}
