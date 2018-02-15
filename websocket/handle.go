package websocket

import (
	"net/http"

	"github.com/trevex/golem"
)

var conn *golem.Connection

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
func connectHndle(c *golem.Connection, http *http.Request) {
	conn = c
}

func SendAll(msg string) {
	conn.Emit("alohomora", struct {
		Msg string `json:"msg"`
	}{Msg: msg})

}
