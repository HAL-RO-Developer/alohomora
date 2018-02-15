package websocket

import (
	"net/http"

	"github.com/trevex/golem"
)

var conns []golem.Connection

type joinModel struct {
	ID int `json:"id"`
}

func GetHandle() func(http.ResponseWriter, *http.Request) {
	return createRouter().Handler()
}

func createRouter() *golem.Router {
	router := golem.NewRouter()
	router.OnConnect(connectHndle)

	return router
}

// connection接続時の処理はここに書く
func connectHndle(conn *golem.Connection, http *http.Request) {
	conns = append(conns, *conn)
}

func SendAll(msg string) {
	for _, conn := range conns {
		conn.Emit("alohomora", struct {
			Msg string `json:"msg"`
		}{Msg: msg})
	}
}
