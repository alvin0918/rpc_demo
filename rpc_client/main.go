package main

import (
	"fmt"
	"log"
	"net/rpc"
)

type Params struct {
	Witch int
	Height int
}

// 调用服务
func main()  {
	var (
		err error
		rp *rpc.Client
		result int
	)
	// 1、连接RPC
	if rp, err = rpc.DialHTTP("tcp", "127.0.0.1:8090"); err != nil {
		log.Fatal(err)
	}

	//  初始化变量
	result = 0

	// 2、调用远程的方法
	if err = rp.Call("Rect.Area", Params{50, 100}, &result); err != nil {
		log.Fatal(err)
	}

	// 求面积（获取返回结果）
	fmt.Println("面积是：" , result)

	if err = rp.Call("Rect.Perimeter", Params{50, 100}, &result); err != nil {
		log.Fatal(err)
	}

	fmt.Println("周长是：", result)

}
