package main

import (
	"flag"
	"fmt"
	"io"
	"net"
	"os"
)

type Client struct {
	ServerIp   string
	ServerPort int
	conn       net.Conn
	Name       string
	flag       int
}

func NewClient(ServerIp string, ServerPort int) *Client {
	client := &Client{
		ServerIp:   ServerIp,
		ServerPort: ServerPort,
		flag:       999,
	}
	conn, err := net.Dial("tcp", fmt.Sprintf("%s:%d", ServerIp, ServerPort))

	if err != nil {
		fmt.Println("net Dial err:", err)
	}

	client.conn = conn

	return client
}

var ServerIp string
var ServerPort int

// ./client -ip 127.0.0.1 -port 8888
func init() {
	flag.StringVar(&ServerIp, "ip", "127.0.0.1", "设置服务器IP地址(默认是 127.0.0.1)")
	flag.IntVar(&ServerPort, "port", 8888, "设置服务器端口(默认是 8888)")

}

func (client *Client) menu() bool {

	var flag int

	fmt.Println("1.公聊模式")
	fmt.Println("2.私聊模式")
	fmt.Println("3.更新用户名")
	fmt.Println("0.退出")

	fmt.Scanln(&flag)

	if flag >= 0 && flag <= 3 {
		client.flag = flag
		return true
	} else {
		fmt.Println("请输入合法的数字")
		return false
	}
}
func (client *Client) UpdateName() bool {
	fmt.Println(">>>>请输入用户名")
	fmt.Scanln(&client.Name)

	sendMsg := "rename|" + client.Name + "\n"

	_, err := client.conn.Write([]byte(sendMsg))

	if err != nil {
		fmt.Println("conn Write err: ", err)
		return false
	}
	return true
}

// 公聊模式
func (client *Client) PublicChat() {
	//提示用户输入消息
	var chatMsg string
	fmt.Println(">>>>请输入聊天内容, exit退出.")
	fmt.Scanln(&chatMsg)

	for chatMsg != "exit" {
		//发给服务器
		//消息不为空则发送
		if len(chatMsg) != 0 {
			sendMsg := chatMsg + "\n"
			_, err := client.conn.Write([]byte(sendMsg))
			if err != nil {
				fmt.Println("conn Write err:", err)
				break
			}
		}

		chatMsg = ""
		fmt.Println(">>>>请输入聊天内容, exit退出.")
		fmt.Scanln(&chatMsg)
	}
}

// 处理服务端传递的信息

func (client *Client) DealResponse() {
	// 一旦client.conn有数据，就直接诶copy到stdout标准输出上，永久阻塞监听
	io.Copy(os.Stdout, client.conn)
}

func (client *Client) Run() {
	for client.flag != 0 {
		for client.menu() == false {
		}

		switch client.flag {
		case 1:
			fmt.Println("公聊模式")
			client.PublicChat()
			break
		case 2:
			fmt.Println("私聊模式")
			break
		case 3:
			fmt.Println("跟新用户名")
			client.UpdateName()
			break
		}
	}

}

// 查询在线用户
func (client *Client) SelectUsers() {
	sendMsg := "who\n"
	_, err := client.conn.Write([]byte(sendMsg))
	if err != nil {
		fmt.Println("conn Write err:", err)
		return
	}
}
func (client *Client) PrivateChat() {
	var remoteName string
	var chatMsg string

	client.SelectUsers()
	fmt.Println(">>>>请输入聊天对象[用户名], exit退出:")
	fmt.Scanln(&remoteName)

	for remoteName != "exit" {
		fmt.Println(">>>>请输入消息内容, exit退出:")
		fmt.Scanln(&chatMsg)

		for chatMsg != "exit" {
			//消息不为空则发送
			if len(chatMsg) != 0 {
				sendMsg := "to|" + remoteName + "|" + chatMsg + "\n\n"
				_, err := client.conn.Write([]byte(sendMsg))
				if err != nil {
					fmt.Println("conn Write err:", err)
					break
				}
			}

			chatMsg = ""
			fmt.Println(">>>>请输入消息内容, exit退出:")
			fmt.Scanln(&chatMsg)
		}

		client.SelectUsers()
		fmt.Println(">>>>请输入聊天对象[用户名], exit退出:")
		fmt.Scanln(&remoteName)
	}
}

func main() {

	flag.Parse()

	client := NewClient(ServerIp, ServerPort)

	if client == nil {
		fmt.Println(">>>>>>>>服务器连接失败<<<<<<<<")
	}

	fmt.Println(">>>>>>>>服务器链接成功<<<<<<<<")

	go client.DealResponse()

	client.Run()
	select {}
}
