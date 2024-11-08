package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
)

type PlayerMove struct {
	Player int    `json:"player`
	Move   string `json:"move"`
}

type Player struct {
	PlayerID int `json:"player_id"`
}

// 定義遊戲狀態
type GameState struct {
	Player1Block         [20][10]int
	Player2Block         [20][10]int
	Player1Score         int
	Player2Score         int
	Player1_garbage_line int
	Player2_garbage_line int
	Player1_input        string
	Player2_input        string
}

var (
	rooms          = make(map[string]GameState) // 房間狀態
	waitingPlayers = make([]*websocket.Conn, 0) // 等待匹配的玩家隊列
	mu             sync.Mutex                   // 保護共享資源
	roomIDCounter  = 1                          // 房間編號
)

func createRoom() string {
	mu.Lock()
	defer mu.Unlock()

	// 創建一個新的房間並返回房間ID
	roomID := fmt.Sprintf("room-%d", roomIDCounter)
	roomIDCounter++
	return roomID
}

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
func handleConnection(conn *websocket.Conn) {
	// 先將玩家加入等待隊列
	mu.Lock()
	waitingPlayers = append(waitingPlayers, conn)
	mu.Unlock()

	// 如果有兩位玩家，開始匹配
	if len(waitingPlayers) >= 2 {
		player1Conn := waitingPlayers[0]
		player2Conn := waitingPlayers[1]

		// 創建一個新的遊戲房間
		roomID := createRoom()

		// 初始化遊戲狀態
		rooms[roomID] = GameState{}

		// 分配玩家編號
		player1 := Player{PlayerID: 1}
		player2 := Player{PlayerID: 2}

		// 發送玩家編號給兩位玩家
		player1Conn.WriteJSON(player1)
		player2Conn.WriteJSON(player2)

		// 發送初始遊戲狀態給兩位玩家
		player1Conn.WriteJSON(rooms[roomID])
		player2Conn.WriteJSON(rooms[roomID])

		// 清除等待匹配的隊列
		waitingPlayers = waitingPlayers[2:]
		go listenForPlayerInput(player1Conn, roomID)
		go listenForPlayerInput(player2Conn, roomID)
		go updateGameLoop(player1Conn, roomID)
		go updateGameLoop(player2Conn, roomID)
	} else {
		// 如果還沒有對手，告知玩家等候
		conn.WriteJSON(map[string]string{"message": "Waiting for another player"})
	}

}

func listenForPlayerInput(conn *websocket.Conn, roomID string) {
	defer conn.Close()
	for {
		// 接收玩家操作消息
		_, msg, err := conn.ReadMessage()
		if err != nil {
			fmt.Println("Error reading message:", err)
			break
		}

		// 根據玩家操作更新遊戲狀態
		// 假設 msg 是一個包含操作的 JSON 字串
		// 例如：{ "player": 1, "move": "left" }
		var input PlayerMove
		err = json.Unmarshal(msg, &input)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
			break
		}

		// 假設有一個 updateGameState 函數來處理遊戲邏輯
		rooms[roomID] = update_game_status(input.Player, input.Move, rooms[roomID])

		// 傳送更新後的遊戲狀態
		conn.WriteJSON(rooms[roomID])
	}
}

func update_game_status(Player int, Move string, curGameState GameState) GameState {
	if Player == 1 {
		curGameState.Player1_input = Move
	} else if Player == 2 {
		curGameState.Player2_input = Move
	}
	return curGameState
}

// 定期更新遊戲狀態的函數
func updateGameLoop(conn *websocket.Conn, roomID string) {
	ticker := time.NewTicker(1 * time.Second) // 設置每秒更新一次
	defer ticker.Stop()

	for range ticker.C {
		curGameState := rooms[roomID]

		// 自動讓方塊向下移動（可以根據遊戲規則更新）
		curGameState = moveBlockDown(curGameState)
		curGameState = moveBlockDown(curGameState)

		// 更新遊戲狀態
		rooms[roomID] = curGameState

		// 向房間內的玩家廣播更新後的遊戲狀態
		conn.WriteJSON(rooms[roomID])
	}
}

func moveBlockDown(curGameState GameState) GameState {
	return curGameState
}

func main() {
	rand.Seed(time.Now().UnixNano())
	r := gin.Default()

	// Gin 路由來處理 WebSocket 連接
	r.GET("/game", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			fmt.Println("Error upgrading connection:", err)
			return
		}

		// 處理此 WebSocket 連線
		handleConnection(conn)
	})

	r.Run(":8081")
}
