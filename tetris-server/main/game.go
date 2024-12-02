package main

// 更新遊戲以及相關操作函式
func update_game_status(Player int, Move string, curGameState GameState) GameState {
	//重設Eliminate_rows
	curGameState.Player_Eliminate_rows[0] = [20]int{}
	curGameState.Player_Eliminate_rows[1] = [20]int{}

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
	curGameState = generatePreviewBlock(Player, curGameState)
	return curGameState
}

// 執行方塊下落一格
func Move_Down(Player int, curGameState GameState) GameState {
	//先嘗試下移
	tryChangeBlock := curGameState.Player_cur_block[Player-1]
	tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
	tmpboard := removeBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])
	//檢查下移是否衝突
	flag := Check_collision(tryChangeBlock, tmpboard)

	//如果可以下移
	if flag {
		//真的下移
		tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
		curGameState.Player_cur_block[Player-1].pos.y += 1
		curGameState.Player_Block_Board[Player-1] = tmpboard
	} else {
		//處理消排
		curGameState = Check_row_full(Player, curGameState)
		curGameState = Eliminate_rows(Player, curGameState)
		curGameState.Player_This_Round_Hold_flag[Player-1] = false
		curGameState = generateGarbageLine(Player, curGameState)
		curGameState = switchNextBlock(Player, curGameState)
	}

	return curGameState
}

// 執行方塊左移一格
func Move_Left(Player int, curGameState GameState) GameState {
	tryChangeBlock := curGameState.Player_cur_block[Player-1]
	tryChangeBlock.pos.x = tryChangeBlock.pos.x - 1
	tmpboard := removeBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])
	flag := Check_collision(tryChangeBlock, tmpboard)

	if flag {
		tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
		curGameState.Player_cur_block[Player-1].pos.x -= 1
		curGameState.Player_Block_Board[Player-1] = tmpboard
	}
	return curGameState
}

// 執行方塊右移一格
func Move_Right(Player int, curGameState GameState) GameState {
	tryChangeBlock := curGameState.Player_cur_block[Player-1]
	tryChangeBlock.pos.x = tryChangeBlock.pos.x + 1
	tmpboard := removeBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])
	flag := Check_collision(tryChangeBlock, tmpboard)

	if flag {
		tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
		curGameState.Player_cur_block[Player-1].pos.x += 1
		curGameState.Player_Block_Board[Player-1] = tmpboard
	}
	return curGameState
}

// 執行方塊旋轉
func Rotate(Player int, curGameState GameState) GameState {
	tryChangeBlock := curGameState.Player_cur_block[Player-1]
	tryChangeBlock = Rotate90(tryChangeBlock)
	tmpboard := removeBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])

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
	curGameState.Player_cur_block[Player-1] = tryChangeBlock
	curGameState.Player_Block_Board[Player-1] = tmpboard
	return curGameState
}

// 執行方塊快速下降
func Soft_Drop(Player int, curGameState GameState) GameState {
	return Move_Down(Player, curGameState)
}

// 執行方塊瞬間下降
func Hard_Drop(Player int, curGameState GameState) GameState {
	flag := true
	//不斷下移一格
	for flag {
		tryChangeBlock := curGameState.Player_cur_block[Player-1]
		tryChangeBlock.pos.y = tryChangeBlock.pos.y + 1
		tmpboard := removeBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])
		flag := Check_collision(tryChangeBlock, tmpboard)

		if flag {
			tmpboard = fillBoardWithBlock(tmpboard, tryChangeBlock)
			curGameState.Player_cur_block[Player-1].pos.y += 1
			curGameState.Player_Block_Board[Player-1] = tmpboard
		} else {
			curGameState = Check_row_full(Player, curGameState)
			curGameState = Eliminate_rows(Player, curGameState)
			curGameState.Player_This_Round_Hold_flag[Player-1] = false
			curGameState = generateGarbageLine(Player, curGameState)
			curGameState = switchNextBlock(Player, curGameState)
			break
		}
	}
	return curGameState
}

// 執行暫存方塊
func Hold(Player int, curGameState GameState) GameState {
	if !curGameState.Player_This_Round_Hold_flag[Player-1] {
		curGameState.Player_This_Round_Hold_flag[Player-1] = true
		//擦除目前方塊
		curGameState.Player_Block_Board[Player-1] = removeBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])

		//更新遊戲狀態
		oldBlockType := curGameState.Player_cur_block_type[Player-1]
		var newBlockType string

		//如果有Hold的方塊
		if curGameState.Player_Hold_Block_type[Player-1] != "" {
			newBlockType = curGameState.Player_Hold_Block_type[Player-1]
		} else { //拿下個方塊
			newBlockType = curGameState.Player_Next_Block[Player-1]
			curGameState.Player_Next_Block[Player-1] = generateRandomBlockType()
		}
		curGameState.Player_cur_block[Player-1] = generateRandomBlock(newBlockType)
		//檢查是否衝突
		for _, offset := range curGameState.Player_cur_block[Player-1].Offsets {
			x := curGameState.Player_cur_block[Player-1].pos.x + offset[0]
			y := curGameState.Player_cur_block[Player-1].pos.y + offset[1]

			if x >= 0 && x < 10 && y >= 0 && y < 20 {
				if curGameState.Player_Block_Board[Player-1][y][x] != 0 { //位置已經有方塊
					curGameState.ifGameOver = 2 //玩家二獲勝
					return curGameState         //處理玩家落敗
				}
			}
		}
		//填上新方塊位置
		curGameState.Player_Block_Board[Player-1] = fillBoardWithBlock(curGameState.Player_Block_Board[Player-1], curGameState.Player_cur_block[Player-1])
		curGameState.Player_cur_block_type[Player-1] = newBlockType
		curGameState.Player_Hold_Block_type[Player-1] = oldBlockType
	}
	return curGameState
}
