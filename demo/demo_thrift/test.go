package main

import (
	"context"
	"fmt"
	demo_thrift "github.com/XJTU-zxc/MyCloudWeGo/demo/demo_thrift/kitex_gen/api"
	"github.com/apache/thrift/lib/go/thrift"
	"os"
)

func main() {
	transportFactory := thrift.NewTFramedTransportFactory(thrift.NewTTransportFactory())
	protocolFactory := thrift.NewTBinaryProtocolFactoryDefault()

	// 连接到服务
	transport, err := thrift.NewTSocket("127.0.0.1:8888")
	if err != nil {
		fmt.Println("Error opening socket:", err)
		os.Exit(1)
	}
	defer transport.Close()

	useTransport, err := transportFactory.GetTransport(transport)
	if err != nil {
		fmt.Println("Error getting transport:", err)
		os.Exit(1)
	}
	client := demo_thrift.NewEchoClientFactory(useTransport, protocolFactory)
	if err := transport.Open(); err != nil {
		fmt.Println("Error opening connection:", err)
		os.Exit(1)
	}
	defer transport.Close()

	// 创建请求
	req := demo_thrift.Request{Message: "Hello World 12345"}
	ctx := context.Background()
	resp, err := client.Echo(ctx, &req)
	if err != nil {
		fmt.Println("Error calling Echo:", err)
		os.Exit(1)
	}
	fmt.Println("Response:", resp.Message)
}
