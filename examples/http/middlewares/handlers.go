package main

import (
	"context"

	"github.com/go-kratos/kratos/examples/helloworld/helloworld"
	"github.com/go-kratos/kratos/v2/transport"
	"github.com/go-kratos/kratos/v2/transport/http"
	"github.com/go-kratos/kratos/v2/transport/http/binding"
)

func sayHelloHandler(ctx http.Context) error {
	var in helloworld.HelloRequest
	if err := ctx.Bind(&in); err != nil {
		return err
	}

	// binding /hello/{name} to in.Name
	if err := binding.BindVars(ctx.Vars(), &in); err != nil {
		return err
	}

	transport.SetOperation(ctx, "/helloworld.Greeter/SayHello")
	h := ctx.Middleware(func(ctx context.Context, req interface{}) (interface{}, error) {
		return &helloworld.HelloReply{Message: "test:" + req.(*helloworld.HelloRequest).Name}, nil
	})
	return ctx.Returns(h(ctx, &in))
}