package main

import (
	"fmt"
	"log"
	"net"
	"net/rpc"
	"net/rpc/jsonrpc"
)

// 声明算数运算结构体
type Arith struct{}

type Param struct {
	B int
}

func (r *Arith) A(p Param, res *int) error {

	*res = p.B

	return nil
}

func main() {

	var (
		err  error
		lis  net.Listener
		conn net.Conn
	)

	// 注册RPC服务
	if err = rpc.Register(new(Arith)); err != nil {
		log.Fatal(err)
	}

	// 把服务绑定到HTTP上
	//rpc.HandleHTTP()

	//if err = http.ListenAndServe(":8091", nil); err != nil {
	//	log.Fatal(err)
	//}

	if lis, err = net.Listen("tcp", "127.0.0.1:8091"); err != nil {
		log.Fatal(err)
	}

	for {
		if conn, err = lis.Accept(); err != nil {
			log.Fatal(err)
		}

		go func(conn net.Conn) {
			fmt.Println("new Client!")
			jsonrpc.ServeConn(conn)
		}(conn)
	}
}
