11/9更新
前置需要下載的東東:

  明天晚上我在另一台新電腦測試

前端:

    主要任務:
    
        1.讀取玩家鍵盤輸入，傳到後端伺服器。

        2.接收後端傳來的遊戲狀態，更新網頁顯示

    訊息傳送格式:

	ex:左移
 	{
	    "player" : 1
	    "move" : "left"
	}

    待做事項:

	根據傳入的[20][10] int畫出兩個玩家的畫面。
	
	顯示分數格

	顯示垃圾行計量條
後端:

    主要任務:
    
        1.接收前端訊息，計算遊戲狀態並回傳

    訊息傳送格式:

	ex:GameState
	{
	    Player1Block         [20][10]int
	    Player2Block         [20][10]int
	    Player1Score         int
	    Player2Score         int
    	    Player1_garbage_line int
	    Player2_garbage_line int
	    Player1_input        string
	    Player2_input        string
	}

	ex:Player
	{
	    PlayerID int `json:"player_id"`
	}

    待做事項:

	func update_game_status(Player int, Move string, curGameState GameState)
	
	    輸入哪個玩家以及操作，更新curGameState

	func moveBlockDown(curGameState GameState) GameState
	    
	    實現方塊下落

定義:

    GameState [20][10] int 內數字意義

    0:無方塊

    1:I型方塊(淺藍色)

    2:O型方塊(黃色)

    3:T型方塊(紫色)

    4:S型方塊(紅色)

    5:Z型方塊(綠色)

    6:L型方塊(深藍色)

    7:J型方塊(橘色)
