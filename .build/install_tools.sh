#!/bin/bash
source $(dirname "$0")/function.sh

function install_protoc() {
    local version="3.15.2"

    local installedVersion="0.0.0"
    if which protoc; then
        installedVersion=$(/usr/local/bin/protoc --version | grep -oE "\b[0-9][0-9.]+")
    fi
    
    if verlt $installedVersion $version; then
        local os=$(echo $(uname -s) | tr '[:upper:]' '[:lower:]')
        local arch=$(uname -m)
        local pbRelease="https://github.com/protocolbuffers/protobuf/releases"
        local fileName="protoc-${version}-${os}-${arch}.zip"
        local downloadURL="$pbRelease/download/v${version}/${fileName}"    
        curl -LO ${downloadURL}
        unzip -o -d /usr/local/ ${fileName}
        rm -vf $fileName
        
        chmod a+rx /usr/local/bin/protoc
    fi
}

function install_protoc_gen_go() {
    go get -u google.golang.org/protobuf/cmd/protoc-gen-go 
    go get -u github.com/golang/mock/mockgen
    go get -u golang.org/x/tools/cmd/goimports
}

function install_protoc-gen-go-grpc() {
	go get -u google.golang.org/grpc/cmd/protoc-gen-go-grpc
}

function install_protoc-gen-go-grpc-http() {
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-grpc-gateway
    go get -u github.com/grpc-ecosystem/grpc-gateway/protoc-gen-openapiv2
}

function install_protoc-gen-grpc-web() {
    local version="1.2.1"

    local binName="protoc-gen-grpc-web"
    local os=$(echo $(uname -s) | tr '[:upper:]' '[:lower:]')
    local arch=$(uname -m)
    local pbRelease="https://github.com/grpc/grpc-web/releases/"
    local fileName="${binName}-${version}-${os}-${arch}"
    local savePath=$(dirname $(which protoc))/${binName} 
    if [ "${os}" = "windows" ]; then        
        fileName="${filename}.exe"
        savePath="${savePath}.exe"
    fi
    local downloadURL="$pbRelease/download/${version}/${fileName}"    
    echo $downloadURL
    curl -L --output "${savePath}" ${downloadURL}
    if [ "${os}" != "windows" ]; then
        chmod a+rx ${savePath}
    fi
}


function waiting_for() {
    yellow_println "waiting all tools installing..."
    
    while [ $(ps aux | grep $(basename ${0}) | grep -v grep | wc -l) -gt 2 ]; do
        sleep 1
        yellow_print "."          
    done

    green_println "all tools installed."
}

################# 执行主入口 ###################

pushd ~
	install_protoc &
	install_protoc_gen_go &
	install_protoc-gen-go-grpc &
    install_protoc-gen-go-grpc-http &
    install_protoc-gen-grpc-web &

    waiting_for
popd