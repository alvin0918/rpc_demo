package main

import (
	"fmt"
	"log"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 声明算数运算结构体
type Arith struct{}

type Param struct {
	B int
}

func main() {

	var (
		cli *rpc.Client
		err error
		res int
	)

	if cli, err = jsonrpc.Dial("tcp", "127.0.0.1:8091"); err != nil {
		log.Fatal(err)
	}

	if err = cli.Call("Arith.A", Param{B: 123}, &res); err != nil {
		log.Fatal(err)
	}

	fmt.Println("结果是：", res)

}
