package lib

import (
	"encoding/json"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"log"
	"net/http"
	"reflect"
)

var upgrader = websocket.Upgrader{
	// 解决跨域问题
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

var WsConnections = make(map[string]*websocket.Conn)

func WsConnection(c *gin.Context, userId string) {
	ws, err := upgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		log.Print("upgrade:", err)
		return
	}

	go func() {
		var uid = userId
		WsConnections[uid] = ws
		log.Println("上线啦!", uid)
		log.Println("在线人数:", len(WsConnections))

		defer ws.Close()
		//ws.WriteMessage(websocket.TextMessage,(byte)user_id)
		for {
			mt, message, err := ws.ReadMessage()
			if err != nil {
				log.Println("read:", err)
				break
			}
			log.Printf("recv: %s", message)
			var decode map[string]interface{}
			if err := json.Unmarshal(message, &decode); err != nil {
				log.Printf("decode err: %s", err.Error())
				break
			}
			sendTo := decode["send_to"]
			log.Println(decode)
			log.Println(sendTo)
			if sendTo != nil {
				sendTo := sendTo.(string)
				var sendMsg = make(map[string]interface{})
				sendMsg["from"] = uid
				sendMsg["message"] = decode["message"].(string)
				sendMsg["name"] = decode["name"].(string)
				msg, _ := json.Marshal(sendMsg)
				if WsConnections[sendTo] != nil {
					err = WsConnections[sendTo].WriteMessage(mt, msg)
					if err != nil {
						log.Printf("WriteMessage err: %s", err.Error())
					}
				}

			}
		}
		delete(WsConnections, uid)
	}()
}

/**
判断是否在线
*/
func isWsOnline(uid string) bool {
	return reflect.TypeOf(WsConnections[uid]) != nil
}
