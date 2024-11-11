package main

import "math/rand"

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

func Move_Down(Player int, curGameState GameState) GameState {
	return curGameState
}

func Move_Left(Player int, curGameState GameState) GameState {
	return curGameState
}

func Move_Right(Player int, curGameState GameState) GameState {
	return curGameState
}

func Rotate(Player int, curGameState GameState) GameState {
	return curGameState
}

func Soft_Drop(Player int, curGameState GameState) GameState {
	return curGameState
}

func Hard_Drop(Player int, curGameState GameState) GameState {
	return curGameState
}

func Hold(Player int, curGameState GameState) GameState {
	return curGameState
}

func Check_collision(Player int, move string, curGameState GameState) bool {
	return true
}

func Check_row_full(curGameState GameState) bool {
	return true
}

func Eliminate_rows(curGameState GameState) bool {
	return true
}
