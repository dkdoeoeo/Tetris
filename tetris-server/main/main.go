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
	Player int    `json:"player"`
	Move   string `json:"move"`
}

type Player struct {
	PlayerID int `json:"player_id"`
}

// 定義遊戲狀態
type GameState struct {
	Player1_Block_Board          [20][10]int
	Player2_Block_Board          [20][10]int
	Player1Score                 int
	Player2Score                 int
	Player1_garbage_line         int
	Player2_garbage_line         int
	Player1_cur_block_type       string
	Player2_cur_block_type       string
	Player1_cur_block            TetrisBlock
	Player2_cur_block            TetrisBlock
	Player1_Hold_Block_type      string
	Player2_Hold_Block_type      string
	Player1_Next_Block           string
	Player2_Next_Block           string
	Player1_Eliminate_rows       [20]int
	Player2_Eliminate_rows       [20]int
	Player1_This_Round_Hold_flag bool
	Player2_This_Round_Hold_flag bool
	ifGameOver                   int //0遊戲繼續、1玩家一獲勝、2玩家二獲勝
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

// 將GameState需要的東西初始化
func initGameState() GameState {
	player1BlockType := generateRandomBlockType()
	player2BlockType := generateRandomBlockType()
	player1Block := generateRandomBlock(player1BlockType)
	player2Block := generateRandomBlock(player2BlockType)

	newGameState := GameState{
		Player1_cur_block_type: player1BlockType,
		Player2_cur_block_type: player2BlockType,
		Player1_cur_block:      player1Block,
		Player2_cur_block:      player2Block,
		Player1_Next_Block:     generateRandomBlockType(),
		Player2_Next_Block:     generateRandomBlockType(),
	}
	newGameState.Player1_Block_Board = fillBoardWithBlock(newGameState.Player1_Block_Board, player1Block)
	newGameState.Player2_Block_Board = fillBoardWithBlock(newGameState.Player2_Block_Board, player2Block)
	return newGameState
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
		rooms[roomID] = initGameState()

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
		go updateGameLoop(player1Conn, player2Conn, roomID)
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

		var input PlayerMove
		err = json.Unmarshal(msg, &input)
		if err != nil {
			fmt.Println("Error unmarshalling message:", err)
			break
		}

		//updateGameState 函數來處理遊戲邏輯
		rooms[roomID] = update_game_status(input.Player, input.Move, rooms[roomID])

		if rooms[roomID].ifGameOver != 0 {
			//處理勝負
		}

		// 傳送更新後雙方遊戲盤面、分數、垃圾行數量
		conn.WriteJSON(rooms[roomID].Player1_Block_Board)
		conn.WriteJSON(rooms[roomID].Player2_Block_Board)
		conn.WriteJSON(rooms[roomID].Player1Score)
		conn.WriteJSON(rooms[roomID].Player2Score)
		conn.WriteJSON(rooms[roomID].Player1_garbage_line)
		conn.WriteJSON(rooms[roomID].Player2_garbage_line)
		conn.WriteJSON(rooms[roomID].Player1_Hold_Block_type)
		conn.WriteJSON(rooms[roomID].Player2_Hold_Block_type)
		conn.WriteJSON(rooms[roomID].Player1_Next_Block)
		conn.WriteJSON(rooms[roomID].Player2_Next_Block)
		conn.WriteJSON(rooms[roomID].Player1_Eliminate_rows)
		conn.WriteJSON(rooms[roomID].Player2_Eliminate_rows)
		//test
		printInfo(rooms[roomID])
	}
}

// test
func printInfo(curGameState GameState) {
	fmt.Print("\033[H\033[2J")
	fmt.Print("Player1Score: ")
	fmt.Print(curGameState.Player1Score)
	fmt.Println()
	fmt.Print("Player1_Hold_Block_type: ")
	fmt.Print(curGameState.Player1_Hold_Block_type)
	fmt.Println()
	fmt.Print("Player1_Next_Block: ")
	fmt.Print(curGameState.Player1_Next_Block)
	fmt.Println()
	fmt.Print("Player1_This_Round_Hold_flag: ")
	fmt.Print(curGameState.Player1_This_Round_Hold_flag)
	fmt.Println()
	fmt.Print("Player1_cur_block_type: ")
	fmt.Print(curGameState.Player1_cur_block_type)
	fmt.Println()
	fmt.Print("Player1_garbage_line: ")
	fmt.Print(curGameState.Player1_garbage_line)
	fmt.Println()
	fmt.Print("ifGameOver: ")
	fmt.Print(curGameState.ifGameOver)
	fmt.Println()

	for i := 0; i < len(curGameState.Player1_Block_Board); i++ {
		for j := 0; j < len(curGameState.Player1_Block_Board[i]); j++ {
			// 輸出每個格子
			fmt.Print(curGameState.Player1_Block_Board[i][j], " ")
		}
		// 每列跳行
		fmt.Println()
	}
}

// 定期更新遊戲狀態的函數
func updateGameLoop(Player1_conn *websocket.Conn, Player2_conn *websocket.Conn, roomID string) {
	ticker := time.NewTicker(1 * time.Second) // 設置每秒更新一次
	defer ticker.Stop()

	for range ticker.C {
		curGameState := rooms[roomID]

		// 讓方塊向下移動
		//curGameState = Move_Down(1, curGameState)
		//curGameState = Move_Down(2, curGameState)

		// 更新遊戲狀態
		rooms[roomID] = curGameState

		if rooms[roomID].ifGameOver != 0 {
			//處理勝負
		}

		// 向房間內的玩家廣播更新後的遊戲狀態
		Player1_conn.WriteJSON(rooms[roomID])
		Player2_conn.WriteJSON(rooms[roomID])
	}
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

	r.Run(":8080")
}