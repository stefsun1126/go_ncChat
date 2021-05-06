package main

import (
	"fmt"
	"net"

	schema "./schema"
)

// 全局map紀錄所有用戶
var allUser = make(map[string]*schema.User)

// 全局msg管道
var allMsg = make(chan string)

func main() {

	// 創建服務器
	server, err := net.Listen("tcp", ":8080")
	if err != nil {
		fmt.Printf("net.Listen() err : %v\n", err)
	}
	fmt.Println("伺服器啟動!")

	// 監聽全局訊息管道
	go listenAllMsg()

	// 接收多個連接
	for {
		conn, err := server.Accept()
		fmt.Println("接收新連線!")
		if err != nil {
			fmt.Printf("server.Accept() err : %v\n", err)
			return
		}
		go handleConnection(conn)
	}
}

// 處理連接
func handleConnection(conn net.Conn) {
	// 建立用戶資料
	// 用戶連接訊息只用處理一次 所以寫在無限迴圈外
	addr := conn.RemoteAddr().String()
	currentUser := schema.NewUser(addr)

	// 將此用戶加入所有用戶map
	allUser[currentUser.GetUserId()] = currentUser

	// 監聽將用戶的管道的訊息寫到終端
	go writeToWindows(currentUser, conn)

	// 寫入登入訊息
	loginInfo := fmt.Sprintf("[%v:%v] login !!!\n", currentUser.GetUserId(), currentUser.GetUserName())
	allMsg <- loginInfo

	// 一個連接會有多次發送
	for {

		// 用來接收獲取到的訊息
		buffer := make([]byte, 2048)
		// 讀取連接
		len, err := conn.Read(buffer)
		if err != nil {
			fmt.Printf("conn.Read() err : %v\n", err)
			return
		}
		// 打印 -1是因為 nc傳過來會有個換行
		msgInfo := fmt.Sprintf("[%v:%v] %v \n", currentUser.GetUserId(), currentUser.GetUserName(), string(buffer[:len-1]))
		allMsg <- msgInfo
	}
}

// 監聽全局訊息管道
func listenAllMsg() {
	for {
		// 將管道訊息讀出
		info := <-allMsg

		// 遍歷所有用戶 寫進用戶管道訊息
		for _, user := range allUser {
			// 這裡如果用無緩衝管道會阻塞在這，因為無緩衝管道需要讀寫同時
			user.SetUserMsg(info)
		}
	}
}

// 將用戶的管道的訊息寫到終端
func writeToWindows(user *schema.User, conn net.Conn) {
	for msg := range user.GetUserMsg() {
		conn.Write([]byte(msg))
	}
}
