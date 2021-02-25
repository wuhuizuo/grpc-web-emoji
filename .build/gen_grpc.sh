#!/bin/bash

source $(dirname "$0")/function.sh

# param $1 pb file path
function gen_pb() {
    local pbFile=$1

	# 基本的
	local protoOpts=""
	protoOpts="${protoOpts} -I ."

	# 普通的message
	protoOpts="${protoOpts} --go_out ./"
	protoOpts="${protoOpts} --go_opt paths=source_relative"

	# grpc service
	protoOpts="${protoOpts} --go-grpc_out ./"
	protoOpts="${protoOpts} --go-grpc_opt paths=source_relative"

	# grpc gateway
	protoOpts="${protoOpts} --grpc-gateway_out ./"
	protoOpts="${protoOpts} --grpc-gateway_opt logtostderr=true"	
	protoOpts="${protoOpts} --grpc-gateway_opt paths=source_relative" 
	protoOpts="${protoOpts} --grpc-gateway_opt generate_unbound_methods=true"

	# open api v2
	protoOpts="${protoOpts} --openapiv2_out ./"
	protoOpts="${protoOpts} --openapiv2_opt logtostderr=true"

	# javascript
	protoOpts="${protoOpts} --js_out ./"
	protoOpts="${protoOpts} --js_opt import_style=commonjs"
	protoOpts="${protoOpts} --grpc-web_out ./"
	protoOpts="${protoOpts} --grpc-web_opt import_style=commonjs"
	protoOpts="${protoOpts} --grpc-web_opt mode=grpcwebtext"

    pushd $(dirname ${pbFile})
		protoc ${protoOpts} $(basename ${pbFile})
	popd
}

############### 执行入口 ##################
PATH="$PATH:$(go env GOPATH)/bin" gen_pb $@