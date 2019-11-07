package clientpkg

func ClientConn() {
	conn := Login("chen", "123456")
	defer conn.Close()

	for {
		PushMsg(EditMsg(), conn)
	}

}
