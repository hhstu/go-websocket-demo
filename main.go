package main

import (

	"fmt"
	"golang.org/x/net/websocket"
	"net/http"
	"os"
	"time"
)
func main() {
	http.Handle("/lc", websocket.Handler(upper))
	if err := http.ListenAndServe(":9999", nil); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	for   {
		time.Sleep(1 *time.Second)
	}
}

func upper(ws *websocket.Conn) {

	var message string
	for {
		// 接收数据

		err := websocket.Message.Receive(ws, &message)

		if err != nil {

			fmt.Println("连接异常, ",err)
			break
		}else {
			fmt.Println("收到信息，", message)

			for  {
				err := websocket.Message.Send(ws, "666")
				if err != nil{

					fmt.Println("发送出错: " + err.Error())
					break
				}else {
					fmt.Println("send  success: ")
				}
				time.Sleep(1 *time.Second)
			}


		}
	}
}

