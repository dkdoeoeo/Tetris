package main

import "math/rand"

type Pos struct {
	x int
	y int
}

// 定義方塊類型，包含中心座標和偏移量
type TetrisBlock struct {
	pos     Pos       //方塊中心座標
	Offsets [4][2]int // 方塊範圍的相對偏移座標 (共4個格子)
}

var (
	IBlock = TetrisBlock{
		pos: Pos{x: 5, y: 0},
		Offsets: [4][2]int{
			{-1, 0}, // 左邊一格
			{0, 0},  // 中心
			{1, 0},  // 右邊一格
			{2, 0},  // 最右邊
		},
	}

	LBlock = TetrisBlock{
		pos: Pos{x: 6, y: 1},
		Offsets: [4][2]int{
			{0, -1}, // 上方一格
			{0, 0},  // 中心
			{-1, 0}, // 左邊一格
			{-2, 0}, // 最左邊
		},
	}

	JBlock = TetrisBlock{
		pos: Pos{x: 4, y: 1},
		Offsets: [4][2]int{
			{0, -1}, // 上方一格
			{0, 0},  // 中心
			{1, 0},  // 右邊一格
			{2, 0},  // 最右邊
		},
	}

	OBlock = TetrisBlock{
		pos: Pos{x: 5, y: 1},
		Offsets: [4][2]int{
			{0, 0},  // 中心
			{1, 0},  // 右邊一格
			{0, -1}, // 上方一格
			{1, -1}, // 右上角
		},
	}

	ZBlock = TetrisBlock{
		pos: Pos{x: 5, y: 1},
		Offsets: [4][2]int{
			{0, 0},   // 中心
			{1, 0},   // 右邊一格
			{0, -1},  // 上方一格
			{-1, -1}, // 左上角
		},
	}

	SBlock = TetrisBlock{
		pos: Pos{x: 5, y: 1},
		Offsets: [4][2]int{
			{0, 0},  // 中心
			{-1, 0}, // 左邊一格
			{0, -1}, // 上方一格
			{1, -1}, // 右上角
		},
	}

	TBlock = TetrisBlock{
		pos: Pos{x: 5, y: 0},
		Offsets: [4][2]int{
			{0, 0},  // 中心
			{-1, 0}, // 左邊一格
			{1, 0},  // 右邊一格
			{0, 1},  // 下方一格
		},
	}
)

func (block *TetrisBlock) Rotate90() {
	for i, offset := range block.Offsets {
		dx, dy := offset[0], offset[1]
		// 90 度順時針旋轉: (dx, dy) -> (-dy, dx)
		block.Offsets[i] = [2]int{-dy, dx}
	}
}

// 隨機生成方塊
func generateRandomBlock() string {
	blocks := []string{"I", "O", "T", "S", "Z", "L", "J"}
	return blocks[rand.Intn(len(blocks))]
}

func update_game_status(Player int, Move string, curGameState GameState) GameState {
	if Player == 1 {
		curGameState.Player1_input = Move
	} else if Player == 2 {
		curGameState.Player2_input = Move
	}
	return curGameState
}

//執行方塊下落一格
func Move_Down(Player int, curGameState GameState) GameState {
	return curGameState
}

//執行方塊左移一格
func Move_Left(Player int, curGameState GameState) GameState {
	return curGameState
}

//執行方塊右移一格
func Move_Right(Player int, curGameState GameState) GameState {
	return curGameState
}

//執行方塊旋轉
func Rotate(Player int, curGameState GameState) GameState {
	return curGameState
}

//執行方塊快速下降
func Soft_Drop(Player int, curGameState GameState) GameState {
	return curGameState
}

//執行方塊瞬間下降
func Hard_Drop(Player int, curGameState GameState) GameState {
	return curGameState
}

//執行暫存方塊
func Hold(Player int, curGameState GameState) GameState {
	return curGameState
}

//檢查方塊是否衝突
func Check_collision(Player int, block_type string, block_pos Pos, curGameState GameState) bool {
	return true
}

//檢查row有沒有填滿
func Check_row_full(curGameState GameState) bool {
	return true
}

//消除row
func Eliminate_rows(curGameState GameState) bool {
	return true
}
