package main

import (
	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
)

var wsMap map[string]*websocket.Conn

func main() {
	http.Handle("/lc", websocket.Handler(upper))
	wsMap = make(map[string]*websocket.Conn)
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func upper(ws *websocket.Conn) {

	var message string
	for {
		// 接收数据
		err := websocket.Message.Receive(ws, &message)
		username := ws.Request().Header.Get("username")
		wsMap[username] = ws
		if err != nil {
			delete(wsMap, username)
			fmt.Println("连接异常, ", err)
			break
		} else {
			fmt.Println("收到信息，", message)
			fmt.Println(len(wsMap))
			for k, v := range wsMap {
				if k != username {
					err := websocket.Message.Send(v, "收到来自 "+username+" 的消息: "+message)
					if err != nil {
						delete(wsMap, username)
						_ = ws.Close()
						fmt.Println("发送出错: " + err.Error())
						break
					} else {
						fmt.Println("send  success: ")
					}
				}
			}

		}
	}
}
