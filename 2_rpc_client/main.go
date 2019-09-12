package main

import (
	"fmt"
	"log"
	"net/rpc"
)

// 声明参数接受结构体
type ArithRequest struct {
	A int
	B int
}

// 声明返回客户端参数结构体
type ArithResponse struct {
	// 乘积
	Pro int
	// 商
	Que int
	// 余数
	Rem int
}

func main() {

	var (
		cl     *rpc.Client
		err    error
		result ArithResponse
	)

	// 连接远程的RPC
	if cl, err = rpc.DialHTTP("tcp", "127.0.0.1:8090"); err != nil {
		log.Fatal(err)
	}

	if err = cl.Call("Arith.Multiply", ArithRequest{100, 6}, &result); err != nil {
		log.Fatal(err)
	}

	if err = cl.Call("Arith.Divide", ArithRequest{100, 6}, &result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("计算结果：", result.Pro)
	fmt.Println("计算结果：", result.Que)
	fmt.Println("计算结果：", result.Rem)

}
