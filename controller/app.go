package controller

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"

	"github.com/HAL-RO-Developer/alohomora/websocket"
	"github.com/gin-gonic/gin"
	uuid "github.com/satori/go.uuid"
)

var device_id = ""
var NowUUid = ""

func SetToken(c *gin.Context) {
	token := c.Query("token")
	if token != "BAC92613-0C14-4F78-9C1E-93D431AF6870" {
		c.JSON(400, "token err")
		return
	}
	NowUUid = uuid.NewV4().String()
	fmt.Println(NowUUid)
	c.JSON(200, gin.H{
		"uuid": NowUUid,
	})
}

func App(c *gin.Context) {
	id := c.Param("uid")
	if NowUUid != id {
		c.Redirect(301, "https://konojunya.com")
		return
	}
	c.HTML(200, "index.html", nil)
}

func Open(c *gin.Context) {
	uid := c.Query("uuid")
	if NowUUid == uid {
		c.JSON(401, "uuid err")
		return
	}
	go websocket.SendAll("open!")

	API("1")
}
func Close(c *gin.Context) {
	uid := c.Query("uuid")
	if NowUUid == uid {
		c.JSON(401, "uuid err")
		return
	}
	go websocket.SendAll("close!")
	API("1")
}

func SetDevice(c *gin.Context) {
	id := c.Query("id")
	device_id = id
	c.JSON(200, gin.H{
		"device_id": device_id,
	})
}

func API(arg string) {
	url := "https://hal-iot.net/api/function"

	payload := strings.NewReader("{\n    \"device_id\":\"" + device_id + "\",\n    \"port\":[\n        {\n            \"port_no\": 1,\n            \"function\": 84,\n            \"args\":[" + arg + ", 5000]\n        }\n    ]\n}")

	req, _ := http.NewRequest("POST", url, payload)

	req.Header.Add("Content-Type", "application/json")

	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	fmt.Println(res)
	fmt.Println(string(body))

}
