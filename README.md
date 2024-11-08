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
