package schema

// 用戶訊息
type User struct {
	id   string      // key, ip:port
	name string      // 預設值同 id,之後可以改
	msg  chan string // 當前用戶的訊息管道
}

// 建立user函數
func NewUser(addr string) *User {
	return &User{
		id:   addr,
		name: addr,
		msg:  make(chan string, 10),
	}
}

// 獲取user id
func (u *User) GetUserId() string {
	return u.id
}

// 獲取user name
func (u *User) GetUserName() string {
	return u.name
}

// 獲取user msg
func (u *User) GetUserMsg() chan string {
	return u.msg
}

// 變更 user msg
func (u *User) SetUserMsg(msg string) {
	u.msg <- msg
}

// 變更 user name
func (u *User) SetUserName(name string) {
	u.name = name
}
