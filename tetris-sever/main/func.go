package main

import (
	"fmt"
	"math/rand"
)

// 這邊開始是一些會用到的函數
func Rotate90(block TetrisBlock) TetrisBlock {
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

func fillBoardWithBlock(curBoard [20][10]int, newBlock TetrisBlock) [20][10]int {
	for _, offset := range newBlock.Offsets {
		x := newBlock.pos.x + offset[0]
		y := newBlock.pos.y + offset[1]

		if x >= 0 && x < 10 && y >= 0 && y < 20 {
			curBoard[y][x] = 1 // 在新的板上對應的位置設置為 1
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

// 檢查方塊是否衝突,如果沒有就把tryChangeBlock填入並回傳新的tmpboard
func Check_collision(Player int, tryChangeBlock TetrisBlock, curGameState GameState) (bool, *[20][10]int) {
	var tmpboard [20][10]int

	if Player == 1 {
		tmpboard = removeBoardWithBlock(curGameState.Player1_Block_Board, curGameState.Player1_cur_block) //暫時移除當前方塊
	} else {
		tmpboard = removeBoardWithBlock(curGameState.Player2_Block_Board, curGameState.Player2_cur_block) //暫時移除當前方塊
	}

	for _, offset := range tryChangeBlock.Offsets {
		x := tryChangeBlock.pos.x + offset[0]
		y := tryChangeBlock.pos.y + offset[1]

		if x >= 0 && x < 10 && y >= 0 && y < 20 {
			if tmpboard[y][x] != 0 { //位置已經有方塊
				return false, nil
			}
		} else {
			return false, nil //超出範圍
		}
	}

	tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)

	return true, &tmpboard
}

// 將next block切換成當前方塊並生成新的next block
func switchNextBlock(Player int, curGameState GameState) GameState {
	if Player == 1 {
		//切換當前方塊
		curGameState.Player1_cur_block_type = curGameState.Player1_Next_Block
		curGameState.Player1_cur_block = generateRandomBlock(curGameState.Player1_Next_Block)
		//檢查新的方塊是否可以加入
		for _, offset := range curGameState.Player1_cur_block.Offsets {
			x := curGameState.Player1_cur_block.pos.x + offset[0]
			y := curGameState.Player1_cur_block.pos.y + offset[1]

			if x >= 0 && x < 10 && y >= 0 && y < 20 {
				if curGameState.Player1_Block_Board[y][x] != 0 { //位置已經有方塊
					curGameState.ifGameOver = 2 //玩家二獲勝
					return curGameState         //處理玩家落敗
				}
			}
		}
		//填上新方塊位置
		curGameState.Player1_Block_Board = fillBoardWithBlock(curGameState.Player1_Block_Board, curGameState.Player1_cur_block)
		//生成新的Next_Block
		curGameState.Player1_Next_Block = generateRandomBlockType()
	} else {
		//切換當前方塊
		curGameState.Player2_cur_block_type = curGameState.Player2_Next_Block
		curGameState.Player2_cur_block = generateRandomBlock(curGameState.Player2_Next_Block)
		//檢查新的方塊是否可以加入
		for _, offset := range curGameState.Player2_cur_block.Offsets {
			x := curGameState.Player2_cur_block.pos.x + offset[0]
			y := curGameState.Player2_cur_block.pos.y + offset[1]

			if x >= 0 && x < 10 && y >= 0 && y < 20 {
				if curGameState.Player2_Block_Board[y][x] != 0 { //位置已經有方塊
					curGameState.ifGameOver = 1 //玩家一獲勝
					return curGameState         //玩家落敗
				}
			}
		}
		//填上新方塊位置
		curGameState.Player2_Block_Board = fillBoardWithBlock(curGameState.Player2_Block_Board, curGameState.Player2_cur_block)
		//生成新的Next_Block
		curGameState.Player2_Next_Block = generateRandomBlockType()
	}

	return curGameState
}

// 檢查row有沒有填滿，紀錄在curGameState.Player_Eliminate_rows
func Check_row_full(Player int, curGameState GameState) GameState {
	if Player == 1 {
		for row := 0; row < len(curGameState.Player1_Block_Board); row++ {
			full := true
			for col := 0; col < len(curGameState.Player1_Block_Board[row]); col++ {
				if curGameState.Player1_Block_Board[row][col] == 0 {
					full = false
					break // 如果有任何元素是 0，這一列不滿
				}
			}
			if full {
				curGameState.Player1_Eliminate_rows[row] = 1 // 如果該列滿了，設為 1
			} else {
				curGameState.Player1_Eliminate_rows[row] = 0 // 否則設為 0
			}
		}
	} else {
		for row := 0; row < len(curGameState.Player2_Block_Board); row++ {
			full := true
			for col := 0; col < len(curGameState.Player2_Block_Board[row]); col++ {
				if curGameState.Player2_Block_Board[row][col] == 0 {
					full = false
					break // 如果有任何元素是 0，這一列不滿
				}
			}
			if full {
				curGameState.Player2_Eliminate_rows[row] = 1 // 如果該列滿了，設為 1
			} else {
				curGameState.Player2_Eliminate_rows[row] = 0 // 否則設為 0
			}
		}
	}
	return curGameState
}

// 消除行
func Eliminate_rows(Player int, curGameState GameState) GameState {
	if Player == 1 {
		for row := 0; row < len(curGameState.Player1_Eliminate_rows); row++ {
			if curGameState.Player1_Eliminate_rows[row] != 0 { //該列滿了
				moveBlocksDown(&curGameState.Player1_Block_Board, row) //清空
			}
		}
	} else {
		for row := 0; row < len(curGameState.Player2_Eliminate_rows); row++ {
			if curGameState.Player2_Eliminate_rows[row] != 0 { //該列滿了
				moveBlocksDown(&curGameState.Player2_Block_Board, row) //清空
			}
		}
	}
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
