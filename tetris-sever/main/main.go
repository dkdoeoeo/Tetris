package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

// 定義遊戲狀態
type GameState struct {
	Player1Block string
	Player2Block string
	Player1Score int
	Player2Score int
}

// 記錄房間的遊戲狀態
var rooms = make(map[string]GameState)

// 升級 HTTP 請求到 WebSocket 的工具
var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool {
		return true
	},
}

// 隨機生成方塊
func generateRandomBlock() string {
	blocks := []string{"I", "O", "T", "S", "Z", "L", "J"}
	return blocks[rand.Intn(len(blocks))]
}

// 處理每個客戶端的 WebSocket 連線
func handleConnection(conn *websocket.Conn, roomID string) {
	// 初始化房間狀態
	if _, exists := rooms[roomID]; !exists {
		rooms[roomID] = GameState{
			Player1Block: generateRandomBlock(),
			Player2Block: generateRandomBlock(),
		}
	}

	// 發送初始遊戲狀態
	conn.WriteJSON(rooms[roomID])

	for {
		// 接收玩家操作
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// 根據操作更新遊戲狀態
		gameState := rooms[roomID]
		fmt.Println(string(msg))
		if string(msg) == "move_left" {
			gameState.Player1Score += 10 // 假設為移動加分
		}

		// 更新遊戲狀態
		rooms[roomID] = gameState

		// 傳送更新後的狀態
		conn.WriteJSON(gameState)
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()

	// Gin 路由來處理 WebSocket 連接
	r.GET("/game", func(c *gin.Context) {
		//roomID := c.Param("roomID")
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Error upgrading connection:", err)
			return
		}
		defer conn.Close()

		// 處理此 WebSocket 連線
		handleConnection(conn, "1")
	})

	r.Run(":8080")
}
