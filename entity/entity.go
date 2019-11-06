package entity

import "sync"

//传递出去的信息结构体存储
type Msg struct {
	To   string
	From string
	Body string
	Time string
}

//存储读到的msg
type Msgs struct {
	Str string
	Mu  sync.Mutex
}
