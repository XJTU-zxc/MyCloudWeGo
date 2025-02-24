.PHONY: gen-demo-proto
gen-demo-proto:
	@cd demo/demo_proto && cwgo server -I ../../idl --type RPC --module github.com/XJTU-zxc/MyCloudWeGo/demo/demo_proto --service demo_proto --idl ../../idl/echo.proto

.PHONY: gen-demo-thrift
gen-demo-thrift:
	@cd demo/demo_thrift && cwgo server -I ../../idl --type RPC --module github.com/XJTU-zxc/MyCloudWeGo/demo/demo_thrift --service demo_thrift --idl ../../idl/echo.thrift

.PHONY: gen_frontend
gen_frontend
	@cd app/frontend && cwgo server -I ../../idl --type HTTP --module github.com/XJTU-zxc/MyCloudWeGo/app/frontend --service frontend  --idl ../../idl/frontend/home.proto