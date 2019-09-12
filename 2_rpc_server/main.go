package main

import (
	"errors"
	"log"
	"net/http"
	"net/rpc"
)

// 声明算数运算结构体
type Arith struct {
}

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

// 乘积运算
func (this *Arith) Multiply(req ArithRequest, res *ArithResponse) error {

	res.Pro = req.A * req.B

	return nil
}

// 商 和 余数
func (this *Arith) Divide(req ArithRequest, res *ArithResponse) error {

	if req.B == 0 {
		return errors.New("除数不能为零！")
	}

	// 除
	res.Que = req.A / req.B

	// 余数
	res.Rem = req.A % req.B

	return nil
}

func main() {

	var (
		re  *Arith
		err error
	)

	// 注册服务
	re = new(Arith)

	if err = rpc.Register(re); err != nil {
		log.Fatal(err)
	}

	// 采用Http作为Rpc的载体
	rpc.HandleHTTP()

	// 监听服务
	if err = http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}

}
