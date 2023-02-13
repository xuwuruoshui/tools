package router

import (
	"context"
	"end/api/handler"
	"end/bootstrap"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"net/http"
)

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

func Ws(r *handler.RouterWrapper, app *bootstrap.App) {

	r.GET("/ws", func(c *gin.Context) any {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return handler.ApiResp(handler.UNKNOWN,nil)
		}
		defer conn.Close()

		ctx, cancel := context.WithCancel(context.Background())
		var readChan = make(chan []byte, 100)
		var writeChan = make(chan []byte, 100)
		go ReadLoop(conn, ctx, cancel, readChan)
		go WriteLoop(conn, ctx, cancel, writeChan)
		for {
			select {
			case msg := <-readChan:
				writeChan <- msg
			case <-ctx.Done():
				goto End
			}
		}
	End:
		return handler.ApiResp(handler.OK,nil)
	})
}

func ReadLoop(conn *websocket.Conn, ctx context.Context, cancel context.CancelFunc, readChan chan []byte) {

	for {
		select {
		case <-ctx.Done():
			goto End
		default:
			// Read message from browser
			_, message, err := conn.ReadMessage()
			if err == websocket.ErrCloseSent {
				conn.Close()
				cancel()
				goto End
			} else if err != nil {
				fmt.Println("receive msg error:", err)
			}
			readChan <- message
		}
	}
End:
}

func WriteLoop(conn *websocket.Conn, ctx context.Context, cancel context.CancelFunc, writeChan chan []byte) {
	for {
		select {
		case tmp := <-writeChan:
			bytes := []byte("服务端消息")

			tmp = append(bytes, tmp...)
			err := conn.WriteMessage(websocket.TextMessage, tmp)
			if err == websocket.ErrCloseSent {
				conn.Close()
				cancel()
				goto End
			} else if err != nil {
				fmt.Println("send msg error:", err)
			}
		case <-ctx.Done():
			goto End
		}

	}
End:
}
