package main

import (
	"net"
	"strings"
)

type User struct {
	UserName    string
	UserAddress string
	C           chan string
	// 为什么要单独一个 表示链接
	conn   net.Conn
	Server *Server
}

func NewUser(conn net.Conn) *User {

	userAdr := conn.RemoteAddr().String()

	user := &User{
		UserName:    userAdr,
		UserAddress: userAdr,
		C:           make(chan string),
		conn:        conn,
	}
	// 监听新创建的user
	go user.Listener()

	return user
}

// 监听用户信息  一旦有信息就发送给对端客户端

func (this *User) Listener() {
	for {
		msg := <-this.C

		this.conn.Write([]byte(msg + "\n"))
	}
}

// 用户上线
func (this *User) Online() {

	this.Server.mapLock.Lock()
	this.Server.OlineMap[this.UserName] = this
	this.Server.mapLock.Unlock()

	this.Server.BoardCast(this, "已上线")
}

// 用户下线
func (this *User) Offline() {
	this.Server.mapLock.Lock()
	delete(this.Server.OlineMap, this.UserName)
	this.Server.mapLock.Unlock()

	this.Server.BoardCast(this, "已下线")
}

// 发送信息
func (this *User) SendMsg(msg string) {
	this.conn.Write([]byte(msg))
}

// 处理用户信息

func (this *User) DoMessage(msg string) {

	if msg == "who" {
		// 加锁
		this.Server.mapLock.Lock()
		// 查询当前在线有哪些用户
		for _, user := range this.Server.OlineMap {
			onlineMsg := "[" + user.UserAddress + "]" + user.UserName + "在线 ...ing"
			this.SendMsg(onlineMsg)
		}
		this.Server.mapLock.Unlock()
	} else if len(msg) > 7 && msg[:7] == "rename|" {
		// 修改名称
		newName := strings.Split(msg, "|")[1]

		_, ok := this.Server.OlineMap[newName]
		if ok {
			// 存在这个名字
			this.SendMsg("当前用户名被使用 \n")
		} else {
			// 修改用户名
			this.Server.mapLock.Lock()
			delete(this.Server.OlineMap, this.UserName)
			this.Server.OlineMap[newName] = this
			this.Server.mapLock.Unlock()
			this.UserName = newName
			this.SendMsg("您已经跟新用户名" + newName + "\n")

		}
	} else if len(msg) > 4 && msg[:3] == "to|" {
		// 给user 发信息
		toUserName := strings.Split(msg, "|")[1]

		if toUserName == "" {
			this.SendMsg("发送信息格式不正确 请使用 \"to|张三|你好啊\" 格式")
			return
		}

		toUser, ok := this.Server.OlineMap[toUserName]
		if !ok {
			this.SendMsg("当前用户不存在")
			return
		}
		content := strings.Split(msg, "|")[2]

		if content == "" {
			this.SendMsg("发送信息不能为空，请重新编辑")
			return
		}
		toUser.SendMsg(this.UserName + "对您说" + content)

	} else {
		this.Server.BoardCast(this, msg)
	}
}

func (this *User) GetInfo() string {
	return "[" + this.UserAddress + "]" + this.UserName
}
