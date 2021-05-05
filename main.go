package main

import (
	"fmt"
	"net"
)

func main() {

	// 創建服務器
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("net.Listen() err : %v\n", err)
	}
	fmt.Println("伺服器啟動!")

	// 接收多個連接
	for {
		conn, err := server.Accept()
		fmt.Println("接收新連線!")
		if err != nil {
			fmt.Printf("server.Accept() err : %v\n", err)
		}
		go handleConnection(conn)
	}
}

// 處理連接
func handleConnection(conn net.Conn) {
	// 一個連接會有多次發送
	for {
		// 用來接收獲取到的訊息
		buffer := make([]byte, 2048)
		// 讀取連接
		len, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("conn.Read() err : %v\n", err)
		}
		// 打印 -1是因為 nc傳過來會有個換行
		fmt.Printf("獲取到的訊息是: %v\n", string(buffer[:len-1]))
	}
}