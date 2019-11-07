package clientpkg

import (
	"bufio"
	"encoding/json"
	"example.com/m/v2/entity"
	"fmt"
	"io"
	"net"
	"os"
	"time"
)

//发送信息操作
func PushMsg(msg entity.Msg, conn net.Conn) error {
	json, err := json.Marshal(msg)
	if err != nil {
		fmt.Println("Make json err", err)
	}

	_, err = conn.Write(json)
	if err != nil {
		conn = Login("chen", "123456")
		_, err = conn.Write(json)
	}
	return err
}

//读取信息操作
func PullMsg() {

}

//读取信息的守护进程
//危险点：有可能会因为断开连接而导致接收不到read
func Deamon(conn net.Conn) {

	for {

		r := bufio.NewReader(conn)
		line, err := r.ReadString('\n') //将r的内容也就是conn的数据按照换行符进行读取。

		if err != io.EOF {
			SaveMsg(line)
			err = nil
		} else {
			break
		}
	}

}

//界面输入模块
func ScanShell() {
	fmt.Println("-1：关闭；0：登陆；1:显示数据；2：发送数据；")
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	conn := Login("chen", "123456")
	go Deamon(conn)
	for {
		input, _ := f.ReadString('\n')
		if len(input) == 1 {
			continue
		}
		if input == "exit\n" {
			break
		}
		switch input {
		case "-1\n":
			os.Exit(0)
		case "1\n":
			ReadMsg()
		case "2\n":
			PushMsg(EditMsg(), conn)

		}

	}
}

//需要在主程序下建立一个结构体保存数据
var MkMsgs entity.Msgs

//save信息模块，主要用于保存信息，以后可以改进为.log或数据库保存
func SaveMsg(msg string) {
	defer MkMsgs.Mu.Unlock()
	MkMsgs.Mu.Lock()
	MkMsgs.Str = MkMsgs.Str + msg
}

//read信息模块，负责将数据按行打印出来,同时将打印过的
//需要将Str加一层锁，以免读取和存储混乱
func ReadMsg() {
	defer MkMsgs.Mu.Unlock()
	MkMsgs.Mu.Lock()
	if MkMsgs.Str != "" {
		fmt.Println(MkMsgs.Str)
	}
	MkMsgs.Str = ""
}

//登陆模块
func Login(user, password string) net.Conn {
	conn, err := net.Dial("tcp", "127.0.0.1:6010")
	if err != nil {
		fmt.Println("Conn err", err)

	}
	return conn
}

//编辑信息
func EditMsg() entity.Msg {
	var msg entity.Msg
	fmt.Println("输入to")
	f := bufio.NewReader(os.Stdin) //读取输入的内容
	to, _ := f.ReadString('\n')
	fmt.Println("输入from")
	from, _ := f.ReadString('\n')
	fmt.Println("输入body")
	body, _ := f.ReadString('\n')
	msg.To = to
	msg.From = from
	msg.Body = body
	msg.Time = time.Now().Format("2006-1114 22:36")
	return msg

}

//连接判断机制
