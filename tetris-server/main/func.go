package main

import (
	"fmt"
	"math/rand"
)

// 這邊開始是一些會用到的函數
func Rotate90(block TetrisBlock) TetrisBlock {
	if block.boardType == typeO {
		return block
	}
	var rotatedBlock = block
	for i, offset := range block.Offsets {
		dx, dy := offset[0], offset[1]
		// 90 度順時針旋轉: (dx, dy) -> (-dy, dx)
		rotatedBlock.Offsets[i] = [2]int{-dy, dx}
	}
	return rotatedBlock
}

// 隨機生成方塊類型
func generateRandomBlockType() string {
	blocks := []string{"I", "O", "T", "S", "Z", "L", "J"}
	return blocks[rand.Intn(len(blocks))]
}

// 根據方塊類型生成初始方塊
func generateRandomBlock(BlockType string) TetrisBlock {
	switch BlockType {
	case "I":
		return IBlock
	case "O":
		return OBlock
	case "T":
		return TBlock
	case "S":
		return SBlock
	case "Z":
		return ZBlock
	case "L":
		return LBlock
	case "J":
		return JBlock
	}
	fmt.Println("generateRandomBlock Error")
	return IBlock
}

func tryRotate(tryChangeBlock TetrisBlock, tmpboard [20][10]int) int {
	if Check_collision(tryChangeBlock, tmpboard) {
		return origin
	}

	tryChangeBlock.pos.x = tryChangeBlock.pos.x - 1
	if Check_collision(tryChangeBlock, tmpboard) {
		return left
	}

	tryChangeBlock.pos.x = tryChangeBlock.pos.x + 2
	if Check_collision(tryChangeBlock, tmpboard) {
		return right
	}

	tryChangeBlock.pos.x = tryChangeBlock.pos.x - 1
	tryChangeBlock.pos.y = tryChangeBlock.pos.y - 1
	if Check_collision(tryChangeBlock, tmpboard) {
		return up
	}

	tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
	tryChangeBlock.pos.x = tryChangeBlock.pos.x - 2
	if Check_collision(tryChangeBlock, tmpboard) {
		return left2
	}

	tryChangeBlock.pos.x = tryChangeBlock.pos.x + 4
	if Check_collision(tryChangeBlock, tmpboard) {
		return right2
	}

	tryChangeBlock.pos.x = tryChangeBlock.pos.x - 4
	tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
	if Check_collision(tryChangeBlock, tmpboard) {
		return down
	}

	return failRotate
}

func fillBoardWithBlock(curBoard [20][10]int, newBlock TetrisBlock) [20][10]int {
	for _, offset := range newBlock.Offsets {
		x := newBlock.pos.x + offset[0]
		y := newBlock.pos.y + offset[1]

		if x >= 0 && x < 10 && y >= 0 && y < 20 {
			curBoard[y][x] = newBlock.boardType // 在新的板上對應的位置設置為對應的值
		} else {
			fmt.Println("fillBoardWithNewBlock Error")
			return curBoard
		}
	}
	return curBoard
}

func removeBoardWithBlock(curBoard [20][10]int, curBlock TetrisBlock) [20][10]int {
	for _, offset := range curBlock.Offsets {
		x := curBlock.pos.x + offset[0]
		y := curBlock.pos.y + offset[1]

		if x >= 0 && x < 10 && y >= 0 && y < 20 {
			curBoard[y][x] = 0 // 在新的板上對應的位置設置為 0
		} else {
			fmt.Println("removeBoardWithBlock Error")
			return curBoard
		}
	}
	return curBoard
}

// 檢查方塊衝突,true:無衝突,false:衝突
func Check_collision(tryChangeBlock TetrisBlock, tmpboard [20][10]int) bool {

	for _, offset := range tryChangeBlock.Offsets {
		x := tryChangeBlock.pos.x + offset[0]
		y := tryChangeBlock.pos.y + offset[1]

		if x >= 0 && x < 10 && y >= 0 && y < 20 {
			if tmpboard[y][x] != 0 { //位置已經有方塊
				return false
			}
		} else {
			return false //超出範圍
		}
	}

	return true
}

// 將next block切換成當前方塊並生成新的next block
func switchNextBlock(Player int, curGameState GameState) GameState {
	//切換當前方塊
	curGameState.Player_cur_block_type[Player-1] = curGameState.Player_Next_Block[Player-1]
	curGameState.Player_cur_block[Player-1] = generateRandomBlock(curGameState.Player_Next_Block[Player-1])
	//檢查新的方塊是否可以加入
	for _, offset := range curGameState.Player_cur_block[Player-1].Offsets {
		x := curGameState.Player_cur_block[Player-1].pos.x + offset[0]
		y := curGameState.Player_cur_block[Player-1].pos.y + offset[1]

		if x >= 0 && x < 10 && y >= 0 && y < 20 {
			if curGameState.Player_Block_Board[Player-1][y][x] != 0 { //位置已經有方塊
				curGameState.ifGameOver = 3 - Player //另一位玩家獲勝
				return curGameState                  //處理玩家落敗
			}
		}
	}
	//填上新方塊位置
	curGameState.Player_Block_Board[Player-1] = fillBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])
	//生成新的Next_Block
	curGameState.Player_Next_Block[Player-1] = generateRandomBlockType()

	return curGameState
}

// 檢查row有沒有填滿，紀錄在curGameState.Player_Eliminate_rows
func Check_row_full(Player int, curGameState GameState) GameState {
	for row := 0; row < len(curGameState.Player_Block_Board[Player-1]); row++ {
		full := true
		for col := 0; col < len(curGameState.Player_Block_Board[Player-1][row]); col++ {
			if curGameState.Player_Block_Board[Player-1][row][col] == 0 || curGameState.Player_Block_Board[Player-1][row][col] == 8 {
				full = false
				break // 如果有任何元素是 0或8，這一列不滿
			}
		}
		if full {
			curGameState.Player_Eliminate_rows[Player-1][row] = 1 // 如果該列滿了，設為 1
		} else {
			curGameState.Player_Eliminate_rows[Player-1][row] = 0 // 否則設為 0
		}
	}

	return curGameState
}

// 消除行
func Eliminate_rows(Player int, curGameState GameState) GameState {
	var full_row_num = 0
	for row := 0; row < len(curGameState.Player_Eliminate_rows[Player-1]); row++ {
		if curGameState.Player_Eliminate_rows[Player-1][row] != 0 { //該列滿了
			moveBlocksDown(&curGameState.Player_Block_Board[Player-1], row) //清空
			full_row_num++
		}
	}

	switch full_row_num {
	case 1:
		curGameState.PlayerScore[Player-1] += 40
	case 2:
		curGameState.PlayerScore[Player-1] += 100
		curGameState.Player_garbage_line[2-Player] += 1
	case 3:
		curGameState.PlayerScore[Player-1] += 300
		curGameState.Player_garbage_line[2-Player] += 2
	case 4:
		curGameState.PlayerScore[Player-1] += 1200
		curGameState.Player_garbage_line[2-Player] += 4
	}
	return curGameState
}

func generateGarbageLine(Player int, curGameState GameState) GameState {
	for i := 0; i < curGameState.Player_garbage_line[Player-1]; i++ {
		moveBlocksUp(&curGameState.Player_Block_Board[Player-1])
		for j := 0; j < 10; j++ {
			curGameState.Player_Block_Board[Player-1][19][j] = typeGarbage
		}
		curGameState.Player_Block_Board[Player-1][19][rand.Intn(10)] = 0
	}
	curGameState.Player_garbage_line[Player-1] = 0
	return curGameState
}

func moveBlocksDown(a *[20][10]int, row int) {
	// 從下往上遍歷，將 row 之上的每一列的方塊下移一列
	for r := row; r > 0; r-- {
		for c := 0; c < len(a[r]); c++ {
			a[r][c] = a[r-1][c] // 將上方的方塊下移一列
		}
	}
}

func moveBlocksUp(a *[20][10]int) {
	// 從下往上遍歷，將 row 之上的每一列的方塊下移一列
	// for r := row; r > 0; r-- {
	// 	for c := 0; c < len(a[r]); c++ {
	// 		a[r][c] = a[r-1][c] // 將上方的方塊下移一列
	// 	}
	// }
	for r := 0; r < Row-1; r++ {
		for c := 0; c < len(a[r]); c++ {
			a[r][c] = a[r+1][c]
		}
	}
	for c := 0; c < len(a[Row-1]); c++ {
		a[Row-1][c] = 0
	}
}

func generatePreviewBlock(Player int, curGameState GameState) GameState {
	//移除預覽方塊
	clearPreviewBlock(curGameState.Player_Block_Board[Player-1])
	//暫時移除當前方塊
	tmpboard := removeBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])

	PreviewBlock := curGameState.Player_cur_block[Player-1]
	PreviewBlock.boardType = 8
	flag := true

	for flag {
		flag := Check_collision(PreviewBlock, tmpboard)

		if flag {
			tmpboard = fillBoardWithBlock(tmpboard, PreviewBlock)
			PreviewBlock.pos.y += 1
		}
	}
	//填上當前方塊
	tmpboard = fillBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])
	return curGameState
}

func clearPreviewBlock(Player_Block_Board [20][10]int) [20][10]int {
	for r := 0; r < Row; r++ {
		for c := 0; c < Col; c++ {
			if Player_Block_Board[r][c] == 8 {
				Player_Block_Board[r][c] = 0
			}
		}
	}
	return Player_Block_Board
}
