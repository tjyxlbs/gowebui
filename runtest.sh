#!/bin/bash

function initTestCase() {
    # 证书测试
    go test -v ./case/cms/site_test.go ./case/cms/cms.go

    # 证书链测试
    go test -v ./case/cms/ca_test.go ./case/cms/cms.go

    # 服务测试
    go test -v ./case/server/appserver_test.go  
}

# 清理测试数据
function clear() {
    # 1. 删除服务
    go test -v ./case/server/appserver_test.go -run TestApp0
    # 2. 删除证书链接
    go test -v ./case/cms/ca_test.go ./case/cms/cms.go -run TestCa0
    # 3. 删除证书
    go test -v ./case/cms/site_test.go ./case/cms/cms.go -run TestC0
}

function help() {
    echo "-i 初始化测试，证书->证书链->服务"
    echo "-c 清理数据， 服务->证书链->证书"
}

case $1 in
    "-i")
    	initTestCase
     ;;
    "-c")
    	clear
     ;;
     *)
     	help
      ;;
esac

