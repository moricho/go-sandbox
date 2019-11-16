package greeter

import (
	"context"
	"errors"
	"fmt"

	pb "github.com/moricho/go-sandbox/microservice-sample/greeter/proto"
)

type GreeterService struct{}

var ctxKeyUserName = "userNames"

func (gs *GreeterService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloReply, error) {
	if req.Name == "" {
		return nil, errors.New("error: empty input")
	}
	msg := fmt.Sprintf("Hello, %s!", req.Name)
	// user := ctx.Value(ctxKeyUserName).(string)

	return &pb.HelloReply{
		Message: msg,
		// User:    user,
	}, nil
}
