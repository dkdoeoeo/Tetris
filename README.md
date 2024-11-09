前置準備:

    git設置:
        
	下載git:https://git-scm.com，創個帳號

	給我你github的username，我邀你進git協作。(去信箱接受邀請)

	看你習慣放哪，執行"git clone https://github.com/dkdoeoeo/Tetris.git"
    
    環境設置:

	我自己是用visual studio code，去官網下載

	開啟visual studio code後，選open folder，開啟你剛載的專案資料夾

	前端:

	    點開"./tetris-client/src/App.vue"，也會自動問你要不要載vue

	    下載node.js後重開電腦。https://nodejs.org/zh-tw

	    嘗試執行"node -v"、"npm -v"。
	    
    	    如果報錯已停用指令碼執行，以系統管理員身分執行PowerShell，執行:"Set-ExecutionPolicy RemoteSigned"後輸入y

	    執行"npm install"

	    到這邊應該可以執行"npm run dev"來讓前端跑起來了
	
	GO:

   	    點開"./tetris-sever/main/main.go"會自動問你要不要載GO，載就對了。
	
	    然後跟你說找不到路徑，要先去GO官網載。https://go.dev/dl/

	    重啟visual studio code按F5執行main.go十，再載一次應該就可以GO了

	有其他問題再問我或是gpt


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
