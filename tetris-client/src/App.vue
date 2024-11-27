<script setup>
import { onMounted, onUnmounted, ref, watch, useTemplateRef} from 'vue';
// import TetrisBoard from '@/components/TetrisBoard.vue';
let connected = ref(false)
let socket;
let playerId;

let Player1_Block_Board = ref(Array.from({ length: 20 }, 
                                    () => Array(10).fill(0)));
let Player2_Block_Board = ref(Array.from({ length: 20 }, 
                                    () => Array(10).fill(0)));

// export default {
//   data() {
//     return {
//       connected: false,  // 用來顯示連線狀態
//       player1Move: null,
//       socket: null,  // WebSocket 客戶端
//       playerId: null,
//       Player1_Block_Board: Array.from({ length: 20 }, () => Array(10).fill(0)),
//       Player2_Block_Board: Array.from({ length: 20 }, () => Array(10).fill(0)),
//       player1Score: 0,  // 玩家1分數
//       player2Score: 0,  // 玩家2分數
//       Player1_garbage_line: null,
//       Player2_garbage_line: null,
//       Player1_cur_block_type: null,
//       Player2_cur_block_type: null,
//       Player1_cur_block: Array.from({ length: 4 }, () => Array(2).fill(0)),
//       Player2_cur_block: Array.from({ length: 4 }, () => Array(2).fill(0)),
//       Player1_Hold_Block_type: null,
//       Player2_Hold_Block_type: null,
//       Player1_Next_Block: null,
//       Player2_Next_Block: null,
//       Player1_Eliminate_rows: Array.from({ length: 20 }),
//       Player2_Eliminate_rows: Array.from({ length: 20 }),
//       Player1_This_Round_Hold_flag: null,
//       Player2_This_Round_Hold_flag: null,`
//     };
//   },

onMounted(() => {
  
    // Test
    recalCellRelaPos()
    drawTetrisBoard()

    // 當組件加載時，建立 WebSocket 連線
    socket = new WebSocket("ws://localhost:8080/game");
    // 註冊 WebSocket 事件
    socket.onopen = () => {
      connected.value = true;  // 當 WebSocket 連線成功時設為 true
      console.log("Connected to server!");
    };

    socket.onclose = () => {
      connected.value = false;  // 當 WebSocket 斷開時設為 false
      console.log("Disconnected from server!");
    };

    socket.onerror = (error) => {
      console.log("WebSocket Error:", error);
    };

     // 接收後端訊息
     socket.onmessage = (event) => {
       const data = JSON.parse(event.data);
       
       if (data.player_id) {
         playerId = data.player_id;  // 設置玩家ID
         console.log('玩家ID:', playerId);
        }
        
        if(data.Player_Block_Board) {
          // make sure vue notice the change
          Player1_Block_Board.value = data.Player_Block_Board[0]
          Player2_Block_Board.value = data.Player_Block_Board[1]
          drawTetrisBoard()
        }
      }

  })


const moveLeft = () => {
      // 觸發左移操作
      console.log('Move left');
      sendMoveToServer('left');
    }
const moveRight= () => {
      // 觸發右移操作
      console.log('Move right');
      sendMoveToServer('right');
    }

const accelerateFall = () => {
      // 觸發加速下落
      console.log('Accelerate fall');
      sendMoveToServer('soft_drop');
    }

const rotateBlock = () => {
      // 觸發旋轉操作
      console.log('Rotate block');
      sendMoveToServer('rotate');
    }
const hard_drop = () => {
      // 觸發加速下降操作
      console.log('hard_drop');
      sendMoveToServer('hard_drop');
    }

const hold = () => {
      // 觸發加速下降操作
      console.log('hold');
      sendMoveToServer('hold');
    }
const sendMoveToServer = (moveType) => {
       // 使用 WebSocket 發送消息到伺服器
       const message = JSON.stringify({
        player: playerId,  // 依照玩家的ID設置
        move: moveType,
      });

      socket.send(message);  // 發送消息
    }

const handleKeyDown = (event) => {
      // 監聽鍵盤按鍵事件
      switch (event.key) {
        case 'ArrowLeft':
          moveLeft();
          break;
        case 'ArrowRight':
          moveRight();
          break;
        case 'ArrowDown':
          accelerateFall();
          break;
        case "ArrowUp":
          rotateBlock();
          break;
        case ' ':
          hard_drop();
          break;
        case 'Shift':
          hold();
          break;
        default:
          break;
      }
    }


// test
const canvasRef = useTemplateRef('TetrisCanvasRef')

const getColor = (row_index, col_index) => {
    let cellVal = Player1_Block_Board.value[row_index][col_index]
    switch (cellVal) {
      case 0:
        return 'black';
      case 1:
        return 'lightblue';
      case 2:
        return 'yellow';
      case 3:
        return 'purple';
      case 4:
        return 'red';
      case 5:
        return 'green';
      case 6:
        return 'darkblue';
      case 7:
        return 'orange';
      case 8:
        return 'gray';
    }
}

let cellRelaPos = Array.from({ length: 20 }, () => Array(10).fill([0, 0]))
let cellWidth = 0
let cellHeight = 0

const recalCellRelaPos = () => { 
  let width = canvasRef.value.clientWidth
  let height = canvasRef.value.clientHeight
  canvasRef.value.width = width;
  canvasRef.value.height = height;

  let cellWid = width / 10 
  let cellHei = height / 20 
  cellWidth = cellWid
  cellHeight = cellHei

  for (let row = 0; row < cellRelaPos.length; row++) {
    for (let col = 0; col < cellRelaPos[row].length; col++) {
      cellRelaPos[row][col] = [cellHei * row, cellWid * col]
    }
  }
  
};

const drawTetrisBoard = () => {
    if (!canvasRef.value) return; // Ensure canvasRef is not null
    console.log("Redrawing")
    let ctx = canvasRef.value.getContext("2d");
    for (let index = 0; index < 100; index++) {
      ctx.clearRect(0, 0, canvasRef.value.width, canvasRef.value.height);
      for (let row = 0; row < cellRelaPos.length; row++) {
      for (let col = 0; col < cellRelaPos[row].length; col++) {
          let [y, x] = cellRelaPos[row][col]
          ctx.fillStyle = getColor(row, col)
          ctx.fillRect(x, y, cellWidth, cellHeight)
        }
      }
    }
    
    

  }


// watch(Player1_Block_Board, (newBoard, oldBoard) => {drawTetrisBoard()}, {deep: true, immediate: true})

const resizeCallback = () => {
  recalCellRelaPos()
  drawTetrisBoard()
}

window.addEventListener("resize", resizeCallback)
// test


// 監聽鍵盤事件
onUnmounted(() => {
    // 在組件銷毀前移除監聽器並關閉連線
    window.removeEventListener("keydown", handleKeyDown);
    if (socket) {
      socket.close();  // 斷開 WebSocket 連線
    }

  })

// watch([() => Player1_Block_Board, () => Player2_Block_Board], ([newP1Board, newP2Board]) => {
//   Player1_Block_Board = newP1Board;
//   Player2_Block_Board = newP2Board;
//   console.log("parent watch detected change")
// })


// window.addEventListener("keydown", handleKeyDown);
window.addEventListener("keydown", handleKeyDown);

</script>

<template>
  <div id = "AppContainer">
    <h1>Tetris Game</h1>
    <p v-if="connected">Successfully connected to the server!</p>
    <p v-else>Waiting for connection...</p>
    <div id="TetrisBoardContainer">
      <!-- <TetrisBoard :Board=Player1_Block_Board></TetrisBoard>
      <TetrisBoard :Board=Player2_Block_Board></TetrisBoard> -->

      <!-- test -->
      <div id="board">
        <canvas id="TetrisCanvas" ref="TetrisCanvasRef"></canvas>
      </div>
    
    
    </div>  
  </div>
</template>

<style scoped>
#AppContainer {
  display: flex;
  flex-direction: column;
  justify-content:space-between;
  width: 100%;
  height: auto;
}

#TetrisBoardContainer{
  width: 100%;
  display: flex;
  flex-direction: row;
  justify-content: space-evenly;

}


/* test */
#board {
  display: inline-flex;
  flex-direction: column;
  justify-content: space-between;
  column-gap: 0;
  width: 400px;
  height: 800px;
  border: 2px solid darkgray;
  padding: 5px;
}

#TetrisCanvas {
  width: 100%;
  height: 100%;
}
</style>
