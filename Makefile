GO_PATH=`go env GOPATH`

gen-proto:
	protoc --proto_path=./app/demo/proto/ --go_out=plugins=grpc:./app/demo/proto/ --go_out=${GO_PATH}/src ./app/demo/proto/*.proto
	goimports -w ${GO_PATH}/src/github.com/jin-cg/grpc-demo/app/demo/proto/
get-vendor:
	cd app; dep ensure
