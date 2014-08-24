package controllers

import (
	"container/list"
	"github.com/astaxie/beego"
	"github.com/gorilla/websocket"
	"net/http"
)

type WebSocketController struct {
	beego.Controller
}

type Subscriber struct {
	Conn *websocket.Conn // Only for WebSocket users; otherwise nil.
}

var (
	message_chan      = make(chan string, 10)
	subscribe_chan    = make(chan Subscriber, 10)
	subscribes_list   = list.New()
	unsubcribe_chan   = make(chan Subscriber, 10)
	unsubscribes_list = list.New()
)

func (this *WebSocketController) Get() {
	ws, err := websocket.Upgrade(this.Ctx.ResponseWriter, this.Ctx.Request, nil, 1024, 1024)
	if _, ok := err.(websocket.HandshakeError); ok {
		http.Error(this.Ctx.ResponseWriter, "Not a websocket handshake", 400)
		return
	} else if err != nil {
		beego.Error("Cannot setup WebSocket connection:", err)
		return
	}
	subscribe_chan <- Subscriber{Conn: ws}

	for {
		_, p, err := ws.ReadMessage()
		if err != nil {
			return
		}
		lineInfo := string(p)
		message_chan <- lineInfo
		// publish <- newEvent(models.EVENT_MESSAGE, uname, string(p))
		beego.Debug("socket receive:", lineInfo)
	}

	// TODO: 从request的body中获取行号信息
	// this.Ctx.Request.Body
}

func publish() {
	for {
		select {
		case message := <-message_chan:
			for sub := subscribes_list.Front(); sub != nil; sub = sub.Next() {
				beego.Debug("helper:", subscribes_list.Len())

				// beego.Debug("haha")
				ws := sub.Value.(Subscriber).Conn
				if ws != nil {
					if ws.WriteMessage(websocket.TextMessage, []byte(message)) != nil {
						// User disconnected.
						unsubcribe_chan <- sub.Value.(Subscriber)
					}
				}
			}
		case subcribe := <-subscribe_chan:
			subscribes_list.PushBack(subcribe)
		case unsubcribe := <-unsubcribe_chan:
			// beego.Debug("unsub", unsubcribe)
			for sub := subscribes_list.Front(); sub != nil; sub = sub.Next() {
				if sub.Value.(Subscriber).Conn == unsubcribe.Conn {
					subscribes_list.Remove(sub)
					// Clone connection.
					ws := sub.Value.(Subscriber).Conn
					if ws != nil {
						ws.Close()
						// beego.Error("WebSocket closed:", unsubcribe.Conn)
					}
				}
			}

		}
	}
}

func init() {
	go publish()
}

// func (this *WebSocketController) Publish() {
// 	ws.WriteMessage(websocket.TextMessage)

// }
