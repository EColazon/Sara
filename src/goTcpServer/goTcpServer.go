package goTcpServer

import (
	"fmt"
	"net"
	"strings"
	"time"
)

type Client struct {
	C chan string //用户发送数据的管道
	Name string   //用户名
	Addr string   //网络地址
}

//保存在线用户 cliAddr ---> Client
var onlineMap map[string]Client

var message = make(chan string)


func GoTcpServer() {
	//监听
	listener, err := net.Listen("tcp", ":32730")
	if err != nil {
		fmt.Println("---> net.Listen err : ", err)
		return
	}
	defer listener.Close()

	//新开一个协程转发消息,只要有消息到来，遍历map,给map每个成员发送此消息
	go Manager()

	//主协程,循环阻塞等待用户连接
	for {
		conn, err := listener.Accept()
		if err != nil {
			fmt.Println("---> listener.Accept err : ", err)
			continue
		}
		go HandleConn(conn) //处理用户连接
	}
}

func Manager() {
	//给map分配空间
	onlineMap = make(map[string]Client)

	for {
		msg := <- message //没有消息,会阻塞

		//遍历map,给map每个成员发送此消息
		for _, cli := range onlineMap {
			cli.C <- msg
		}
	}
}

func HandleConn(conn net.Conn) {
	//处理用户连接
	defer conn.Close()

	//获取客户端的网络地址
	cliAddr := conn.RemoteAddr().String()

	//创建一个结构体,默认用户名和网络地址一致
	cli := Client{make(chan string), cliAddr, cliAddr}
	//将结构体添加到map
	onlineMap[cliAddr] = cli

	//新开一个协程专门给当前客户端发送消息
	go WriteMsgToClient(cli, conn)
	
	//广播某个在线
	//message <- "[" + cli.Addr + "]" + cli.Name + ": login"
	message <- MakeMsg(cli, "login")

	//提示我是谁
	cli.C <- MakeMsg(cli, "I am here.")

	isQuit := make(chan bool) //对方是否主动掉线
	hasData := make(chan bool) //对方是否有数据发送

	//新建一个协程,接收用户发过来的数据
	go func() {
		buf := make([]byte, 2048)
		for {
			n, err := conn.Read(buf)
			if n == 0 {
				//对方断开或者出问题
				isQuit <- true
				fmt.Println("---> conn.Read err: ", err)
				return
			}

			msg := string(buf[:n-1]) //win-1,一个换行符
			if len(msg) == 3 && msg == "who" {
				//遍历map,给当前用户发送所有成员
				conn.Write([]byte("user list: \n"))
				for _, tmp := range onlineMap {
					msg = tmp.Addr + ":" + tmp.Name + "\n"
					conn.Write([]byte(msg))
				}
			} else if len(msg) == 8 && msg[:6] == "rename" {
				//rename|mike
				name := strings.Split(msg, "|")[1]
				cli.Name = name
				onlineMap[cliAddr] = cli
				conn.Write([]byte("rename ok\n"))
			} else {
				//转发此内容
				message <- MakeMsg(cli, msg)
			}
			hasData <- true //代表有数据
		}
	}()

	for {
		//通过select检测channel的流动
		select {
		case <- isQuit:
			delete(onlineMap, cliAddr) //当前用户从map移除
			message <- MakeMsg(cli, "login out.")
			return
		case <- hasData:

		case <- time.After(30*time.Second):
			delete(onlineMap, cliAddr) //当前用户从map移除
			message <- MakeMsg(cli, "timeout leave out.") //广播谁下线了
			return
		}
	}
}

func WriteMsgToClient(cli Client, conn net.Conn) {
	for msg := range cli.C { //给当前客户端发送消息
		conn.Write([]byte(msg + "\n"))
	}
}

func MakeMsg(cli Client, msg string) (buf string) {
	buf = "[" + cli.Addr + "]" + cli.Name + ": " + msg
	return
}


