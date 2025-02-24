package main

import (
	"context"
	"errors"
	"fmt"
	"log"

	"github.com/XJTU-zxc/MyCloudWeGo/demo/demo_proto/kitex_gen/pbapi"
	"github.com/XJTU-zxc/MyCloudWeGo/demo/demo_proto/kitex_gen/pbapi/echoservice"
	"github.com/XJTU-zxc/MyCloudWeGo/demo/demo_proto/middleware"
	"github.com/bytedance/gopkg/cloud/metainfo"
	"github.com/cloudwego/kitex/client"
	"github.com/cloudwego/kitex/pkg/kerrors"
	"github.com/cloudwego/kitex/pkg/transmeta"
	"github.com/cloudwego/kitex/transport"
	consul "github.com/kitex-contrib/registry-consul"
)

func main() {
	r, err := consul.NewConsulResolver("127.0.0.1:8500")
	if err != nil {
		log.Fatal(err)
	}
	c, err := echoservice.NewClient("demo_proto", client.WithResolver(r),
		client.WithMetaHandler(transmeta.ClientHTTP2Handler),
		client.WithTransportProtocol(transport.GRPC),
		client.WithMiddleware(middleware.Middleware),
	)
	if err != nil {
		log.Fatal(err)
	}
	var bizErr *kerrors.GRPCBizStatusError
	ctx := metainfo.WithPersistentValue(context.Background(), "CLIENT_NAME", "demo_proto")
	res, err := c.Echo(ctx, &pbapi.Request{Message: "hello"})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
	res, err = c.Echo(ctx, &pbapi.Request{Message: "error"})
	if err != nil {
		if ok := errors.As(err, &bizErr); ok {
			fmt.Printf("\n%#v\n", bizErr)
		}
		log.Fatal(err)
	}
	fmt.Printf("%v", res)
}
