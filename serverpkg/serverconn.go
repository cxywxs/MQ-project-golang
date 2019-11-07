package serverpkg

import (
	"bufio"
	"fmt"
	"io"
	"net"
	"time"
)

// 心跳计时，根据GravelChannel判断Client是否在设定时间内发来信息
func HeartBeating(conn net.Conn, readerChannel chan byte, timeout int) {
	select {
	case _ = <-readerChannel:
		fmt.Println(conn.RemoteAddr().String(), "get message, keeping heartbeating...")
		conn.SetDeadline(time.Now().Add(time.Duration(timeout) * time.Second))
		break
	case <-time.After(time.Second * 5):
		fmt.Println("It's really weird to get Nothing!!!")
		conn.Close()
	}
}

//启动server
func ServerConn() {
	listen, err := net.Listen("tcp", ":6010")
	if err != nil {
		fmt.Println("conn start err:", err)
	}
	defer listen.Close()
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("connected err:", err)
		}
		go ConnHandle(conn)
	}
}

func ConnHandle(conn net.Conn) {

	Deamon(conn)
}

func Deamon(conn net.Conn) {
	c := make(chan byte)
	go HeartBeating(conn, c, 20)
	for {

		r := bufio.NewReader(conn)
		line, err := r.ReadString('\n') //将r的内容也就是conn的数据按照换行符进行读取。
		MsgHandle(line)
		c <- byte(1)
		if err != io.EOF {
			//SaveMsg(line)
			MsgHandle(line)
			err = nil
		} else {
			break
		}
	}

}
