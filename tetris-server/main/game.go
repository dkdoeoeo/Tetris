package main

// 更新遊戲以及相關操作函式
func update_game_status(Player int, Move string, curGameState GameState) GameState {
	//重設Eliminate_rows
	curGameState.Player1_Eliminate_rows = [20]int{}
	curGameState.Player2_Eliminate_rows = [20]int{}

	switch Move {
	case "left":
		curGameState = Move_Left(Player, curGameState)
	case "right":
		curGameState = Move_Right(Player, curGameState)
	case "rotate":
		curGameState = Rotate(Player, curGameState)
	case "soft_drop":
		curGameState = Soft_Drop(Player, curGameState)
	case "hard_drop":
		curGameState = Hard_Drop(Player, curGameState)
	case "hold":
		curGameState = Hold(Player, curGameState)
	}
	return curGameState
}

// 執行方塊下落一格
func Move_Down(Player int, curGameState GameState) GameState {
	if Player == 1 {
		tryChangeBlock := curGameState.Player1_cur_block
		tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
		tmpboard := removeBoardWithBlock(curGameState.Player1_Block_Board, curGameState.Player1_cur_block)
		flag := Check_collision(tryChangeBlock, tmpboard)

		if flag {
			tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
			curGameState.Player1_cur_block.pos.y += 1
			curGameState.Player1_Block_Board = tmpboard
		} else {
			curGameState = Check_row_full(Player, curGameState)
			curGameState = Eliminate_rows(Player, curGameState)
			curGameState.Player1_This_Round_Hold_flag = false
			curGameState = switchNextBlock(Player, curGameState)
		}
	} else {
		tryChangeBlock := curGameState.Player2_cur_block
		tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
		tmpboard := removeBoardWithBlock(curGameState.Player2_Block_Board, curGameState.Player2_cur_block)
		flag := Check_collision(tryChangeBlock, tmpboard)

		if flag {
			tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
			curGameState.Player2_cur_block.pos.y += 1
			curGameState.Player2_Block_Board = tmpboard
		} else {
			curGameState = Check_row_full(Player, curGameState)
			curGameState = Eliminate_rows(Player, curGameState)
			curGameState.Player2_This_Round_Hold_flag = false
			curGameState = switchNextBlock(Player, curGameState)
		}
	}
	return curGameState
}

// 執行方塊左移一格
func Move_Left(Player int, curGameState GameState) GameState {
	if Player == 1 {
		tryChangeBlock := curGameState.Player1_cur_block
		tryChangeBlock.pos.x = tryChangeBlock.pos.x - 1
		tmpboard := removeBoardWithBlock(curGameState.Player1_Block_Board, curGameState.Player1_cur_block)
		flag := Check_collision(tryChangeBlock, tmpboard)

		if flag {
			tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
			curGameState.Player1_cur_block.pos.x -= 1
			curGameState.Player1_Block_Board = tmpboard
		}
	} else {
		tryChangeBlock := curGameState.Player2_cur_block
		tryChangeBlock.pos.x = tryChangeBlock.pos.x - 1
		tmpboard := removeBoardWithBlock(curGameState.Player2_Block_Board, curGameState.Player2_cur_block)
		flag := Check_collision(tryChangeBlock, tmpboard)

		if flag {
			tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
			curGameState.Player2_cur_block.pos.x -= 1
			curGameState.Player2_Block_Board = tmpboard
		}
	}
	return curGameState
}

// 執行方塊右移一格
func Move_Right(Player int, curGameState GameState) GameState {
	if Player == 1 {
		tryChangeBlock := curGameState.Player1_cur_block
		tryChangeBlock.pos.x = tryChangeBlock.pos.x + 1
		tmpboard := removeBoardWithBlock(curGameState.Player1_Block_Board, curGameState.Player1_cur_block)
		flag := Check_collision(tryChangeBlock, tmpboard)

		if flag {
			tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
			curGameState.Player1_cur_block.pos.x += 1
			curGameState.Player1_Block_Board = tmpboard
		}
	} else {
		tryChangeBlock := curGameState.Player2_cur_block
		tryChangeBlock.pos.x = tryChangeBlock.pos.x + 1
		tmpboard := removeBoardWithBlock(curGameState.Player2_Block_Board, curGameState.Player2_cur_block)
		flag := Check_collision(tryChangeBlock, tmpboard)

		if flag {
			tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
			curGameState.Player2_cur_block.pos.x += 1
			curGameState.Player2_Block_Board = tmpboard
		}
	}
	return curGameState
}

// 執行方塊旋轉
func Rotate(Player int, curGameState GameState) GameState {
	if Player == 1 {
		tryChangeBlock := curGameState.Player1_cur_block
		tryChangeBlock = Rotate90(tryChangeBlock)
		tmpboard := removeBoardWithBlock(curGameState.Player1_Block_Board, curGameState.Player1_cur_block)

		switch tryRotate(tryChangeBlock, tmpboard) {
		case failRotate:
			return curGameState
		case origin:
		case left:
			tryChangeBlock.pos.x = tryChangeBlock.pos.x - 1
		case right:
			tryChangeBlock.pos.x = tryChangeBlock.pos.x + 1
		case up:
			tryChangeBlock.pos.y = tryChangeBlock.pos.y - 1
		case left2:
			tryChangeBlock.pos.x = tryChangeBlock.pos.x - 2
		case right2:
			tryChangeBlock.pos.x = tryChangeBlock.pos.x + 2
		case down:
			tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
		}

		tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
		curGameState.Player1_cur_block = tryChangeBlock
		curGameState.Player1_Block_Board = tmpboard
	} else {
		tryChangeBlock := curGameState.Player2_cur_block
		tryChangeBlock = Rotate90(tryChangeBlock)
		tmpboard := removeBoardWithBlock(curGameState.Player2_Block_Board, curGameState.Player2_cur_block)

		switch tryRotate(tryChangeBlock, tmpboard) {
		case failRotate:
			return curGameState
		case origin:
		case left:
			tryChangeBlock.pos.x = tryChangeBlock.pos.x - 1
		case right:
			tryChangeBlock.pos.x = tryChangeBlock.pos.x + 1
		case up:
			tryChangeBlock.pos.y = tryChangeBlock.pos.y - 1
		case left2:
			tryChangeBlock.pos.x = tryChangeBlock.pos.x - 2
		case right2:
			tryChangeBlock.pos.x = tryChangeBlock.pos.x + 2
		case down:
			tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
		}

		tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
		curGameState.Player2_cur_block = tryChangeBlock
		curGameState.Player2_Block_Board = tmpboard
	}
	return curGameState
}

// 執行方塊快速下降
func Soft_Drop(Player int, curGameState GameState) GameState {
	return Move_Down(Player, curGameState)
}

// 執行方塊瞬間下降
func Hard_Drop(Player int, curGameState GameState) GameState {
	flag := true
	for flag {
		if Player == 1 {
			tryChangeBlock := curGameState.Player1_cur_block
			tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
			tmpboard := removeBoardWithBlock(curGameState.Player1_Block_Board, curGameState.Player1_cur_block)
			flag := Check_collision(tryChangeBlock, tmpboard)

			if flag {
				tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
				curGameState.Player1_cur_block.pos.y += 1
				curGameState.Player1_Block_Board = tmpboard
			} else {
				curGameState = Check_row_full(Player, curGameState)
				curGameState = Eliminate_rows(Player, curGameState)
				curGameState.Player1_This_Round_Hold_flag = false
				curGameState = switchNextBlock(Player, curGameState)
				break
			}
		} else {
			tryChangeBlock := curGameState.Player2_cur_block
			tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
			tmpboard := removeBoardWithBlock(curGameState.Player2_Block_Board, curGameState.Player2_cur_block)
			flag := Check_collision(tryChangeBlock, tmpboard)

			if flag {
				curGameState.Player2_cur_block.pos.y += 1
				curGameState.Player2_Block_Board = tmpboard
			} else {
				curGameState = Check_row_full(Player, curGameState)
				curGameState = Eliminate_rows(Player, curGameState)
				curGameState.Player2_This_Round_Hold_flag = false
				curGameState = switchNextBlock(Player, curGameState)

				break
			}
		}
	}
	return curGameState
}

// 執行暫存方塊
func Hold(Player int, curGameState GameState) GameState {
	if Player == 1 {
		if curGameState.Player1_This_Round_Hold_flag == false {
			curGameState.Player1_This_Round_Hold_flag = true
			//擦除目前方塊
			curGameState.Player1_Block_Board = removeBoardWithBlock(curGameState.Player1_Block_Board, curGameState.Player1_cur_block)

			//更新遊戲狀態
			oldBlockType := curGameState.Player1_cur_block_type
			var newBlockType string

			//如果有Hold的方塊
			if curGameState.Player1_Hold_Block_type != "" {
				newBlockType = curGameState.Player1_Hold_Block_type
			} else { //拿下個方塊
				newBlockType = curGameState.Player1_Next_Block
				curGameState.Player1_Next_Block = generateRandomBlockType()
			}
			curGameState.Player1_cur_block = generateRandomBlock(newBlockType)
			//檢查是否衝突
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
			curGameState.Player1_cur_block_type = newBlockType
			curGameState.Player1_Hold_Block_type = oldBlockType
		}
	} else {
		if curGameState.Player2_This_Round_Hold_flag == false {
			curGameState.Player2_This_Round_Hold_flag = true
			//擦除目前方塊
			curGameState.Player2_Block_Board = removeBoardWithBlock(curGameState.Player2_Block_Board, curGameState.Player2_cur_block)

			//更新遊戲狀態
			oldBlockType := curGameState.Player2_cur_block_type
			var newBlockType string

			//如果有Hold的方塊
			if curGameState.Player2_Hold_Block_type != "" {
				newBlockType = curGameState.Player2_Hold_Block_type
			} else { //拿下個方塊
				newBlockType = curGameState.Player2_Next_Block
				curGameState.Player2_Next_Block = generateRandomBlockType()
			}
			curGameState.Player2_cur_block = generateRandomBlock(newBlockType)
			//檢查是否衝突
			for _, offset := range curGameState.Player2_cur_block.Offsets {
				x := curGameState.Player2_cur_block.pos.x + offset[0]
				y := curGameState.Player2_cur_block.pos.y + offset[1]

				if x >= 0 && x < 10 && y >= 0 && y < 20 {
					if curGameState.Player2_Block_Board[y][x] != 0 { //位置已經有方塊
						curGameState.ifGameOver = 1 //玩家一獲勝
						return curGameState         //處理玩家落敗
					}
				}
			}
			//填上新方塊位置
			curGameState.Player2_Block_Board = fillBoardWithBlock(curGameState.Player2_Block_Board, curGameState.Player2_cur_block)
			curGameState.Player2_cur_block_type = newBlockType
			curGameState.Player2_Hold_Block_type = oldBlockType
		}
	}
	return curGameState
}
