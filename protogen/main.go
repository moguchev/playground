package main

import (
	"context"

	"github.com/moguchev/playground/protogen/bar"
	"github.com/moguchev/playground/protogen/foo"
	"google.golang.org/grpc"
)

type server struct {
	client bar.BarClient
}

func (s *server) GetFoo(ctx context.Context, in *foo.GetFooRequest, opts ...grpc.CallOption) (*foo.GetFooResponse, error) {
	return s.client.GetFooBar(ctx, in)
}

func NewServer(bc bar.BarClient) foo.FooServer {
	return &server{bc}
}

func main() {

}
