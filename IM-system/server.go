package main

import (
	"fmt"
	"io"
	"net"
	"sync"
	"time"
)

type Server struct {
	Ip   string
	Port string

	// 在线用户列表
	OlineMap map[string]*User
	mapLock  sync.RWMutex

	// 信息广播的通道
	Message chan string
}

// 根据一个服务器创建一个 服务
func Newserver(ip, port string) *Server {
	return &Server{
		Ip:       ip,
		Port:     port,
		OlineMap: make(map[string]*User),
		Message:  make(chan string),
	}
}

//根据 连接 connection 进行操作方法

func (this *Server) Handler(conn net.Conn) {
	// fmt.Println("连接操作成功")
	user := NewUser(conn)
	// 用户上线
	user.Online()

	//接受客户端发送的信息

	isLive := make(chan bool)

	go func() {
		buf := make([]byte, 4096)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				// 用户下线
				user.Offline()
				return
			}

			if err != nil && err != io.EOF {
				fmt.Println("Conn Read err ", err)
				return
			}
			// 去除用户信息后面的 '\n'
			msg := string(buf[:n-1])

			// 将得到的信息进行广播
			// 将用户发过来的信息 进行广播
			// this.BoardCast(user, msg)
			user.DoMessage(msg)
			isLive <- true
		}
	}()

	// handel 阻塞
	// 处理长时间不发信息, 将用户踢出
	for {
		select {

		case <-isLive:

		case <-time.After(30 * time.Second):
			// 已经超时
			user.SendMsg("你被踢出")

			//关闭连接和通道
			close(user.C)
			user.conn.Close()
			return

		}
	}
}

// 启动服务器
func (this *Server) Start() {
	// socket listen
	listener, err := net.Listen("tcp", fmt.Sprintf("%s:%s", this.Ip, this.Port))
	if err != nil {
		fmt.Println("listen err ", err)
		return
	}
	defer listener.Close()

	go this.MessageListener()

	// accpet
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("listen accpet err ", err)
		}
		// 单独开协成处理当前的 连接  不阻塞下一个服务器的链接
		go this.Handler(conn)
	}
}

// 广播信息, 传递信息
func (this *Server) BoardCast(user *User, msg string) {
	sendMessage := "[" + user.UserAddress + "]" + user.UserName + ":" + msg

	this.Message <- sendMessage
}

// 监听message 信息
func (this *Server) MessageListener() {

	for {
		msg := <-this.Message

		for _, val := range this.OlineMap {
			val.C <- msg
		}
	}

}
