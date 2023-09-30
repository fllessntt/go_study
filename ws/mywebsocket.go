package ws

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/gorilla/websocket"
	"io"
	"log"
	"net/http"
)

func handleHttp(w http.ResponseWriter, r *http.Request) {
	res, _ := fmt.Fprintf(w, "hello World!")
	log.Println(res)
}

var wsUpgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
	CheckOrigin: func(r *http.Request) bool {
		return r.Header.Get("access_token") != ""
	},
}

func handleWebSocket(w http.ResponseWriter, r *http.Request) {
	// 升级HTTP连接为WebSocket连接
	// WebSocket 连接底层并不是 HTTP 连接，但是 WebSocket 连接的建立需要依赖于 HTTP 连接。
	conn, err := wsUpgrader.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	log.Printf("client %v%v connect success", r.Host, r.URL)
	//defer后接匿名函数, 宿主函数返回之前执行
	//关闭连接, 忽略错误
	defer func(conn *websocket.Conn) { _ = conn.Close() }(conn)
	listenWsMessage(conn)
}

// 监听收到的WebSocket消息
func listenWsMessage(conn *websocket.Conn) {
	for {
		// 读取消息
		messageType, p, err := conn.ReadMessage()
		if err != nil {
			log.Println(err)
			return
		}
		log.Println("Received message:", string(p))

		// 发送消息
		err = conn.WriteMessage(messageType, []byte("Hello, world!"))
		if err != nil {
			log.Println(err)
			return
		}
	}
}

func Run() {
	http.HandleFunc("/ws", handleWebSocket)
	http.HandleFunc("/", handleHttp)
	log.Println("Server will start on 0.0.0.0:8080/ws")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func TestHttpGet() {
	client := &http.Client{}
	req, err := http.NewRequest("GET", "https://api.vvhan.com/api/moyu?type=json", nil)
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer func(Body io.ReadCloser) { _ = Body.Close() }(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}

func GuessByPost(word string) {
	client := &http.Client{}
	headers := map[string]string{
		"origin":       "https://lab.magiconch.com",
		"referer":      "https://lab.magiconch.com/nbnhhsh/",
		"user-agent":   "Mozilla/5.0 (Windows NT 10.0; WOW64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/80.0.3987.163 Safari/537.36",
		"Content-Type": "application/json",
	}
	data := map[string]any{"text": word}
	jsonBytes, err := json.Marshal(data)
	if err != nil {
		fmt.Println(err)
		return
	}
	buf := bytes.NewBuffer(jsonBytes)
	req, err := http.NewRequest("POST", "https://lab.magiconch.com/api/nbnhhsh/guess", buf)
	for key, val := range headers {
		req.Header.Set(key, val)
	}
	if err != nil {
		fmt.Println(err)
		return
	}
	resp, err := client.Do(req)
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {
			fmt.Println(err)
			return
		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
