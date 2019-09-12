package main

import (
	"log"
	"net/http"
	"net/rpc"
)

type Rect struct {

}

// 声明一个参数的结构体, 字段首字母要大写
type Params struct {
	// 长 和 宽
	Witch int
	Height int
}

// 定义求矩形面积的方法
func (r *Rect) Area (p Params, result *int) error {
	*result = p.Witch * p.Height
	return nil
}

// 定义求矩形周长的方法
func (r *Rect) Perimeter (p Params, result *int) error {
	*result = (p.Height + p.Witch) * 2
	return nil
}

func main()  {

	var (
		err error
	)

	// 注册服务
	rect := new(Rect)

	if err = rpc.Register(rect); err != nil {
		log.Fatal(err)
	}

	// 把服务处理绑定到http协议上
	rpc.HandleHTTP()

	// 监听服务
	if err = http.ListenAndServe("127.0.0.1:8090", nil); err != nil {
		log.Fatal(err)
	}

}






