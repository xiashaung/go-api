package lib

import (
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var WsConnections = make(map[int]*websocket.Conn)

func WsConnection(c *gin.Context,user_id int)  {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}
	go func() {
		defer ws.Close()
		var uid = user_id;
		WsConnections[uid] = ws
		//ws.WriteMessage(websocket.TextMessage,(byte)user_id)
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			err = ws.WriteMessage(mt, message)
			if err != nil {
				log.Println("write:", err)
				break
			}
		}
		delete(WsConnections,uid)
	}()
}