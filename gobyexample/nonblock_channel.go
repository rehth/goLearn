package main

import "fmt"

func main() {
	msg := make(chan string)
	signals := make(chan bool)
	// 使用带一个 default 子句的 select 来实现非阻塞接收
	select {
	case info := <-msg:
		fmt.Println("收到消息：", info)
	default:
		fmt.Println("没有收到消息")
	}

	info := "hello"
	// 使用带一个 default 子句的 select 来实现非阻塞发送
	select {
	case msg <- info:
		fmt.Println("发送消息", info)
	default:
		fmt.Println("没有发送消息", len(msg))
	}

	// 多个 case 配合实现多路的非堵塞的选择器
	select {
	case info := <-msg:
		fmt.Println("收到消息", info)
	case sig := <-signals:
		fmt.Println("收到信号", sig)
	default:
		fmt.Println("没有任何行为")
	}

}
