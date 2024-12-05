<script setup>
import { nextTick, onMounted, onUnmounted, ref, useTemplateRef, watch } from 'vue';
import TetrisBoard from '@/components/TetrisUI.vue';
import { Icon } from '@iconify/vue';

let socket;
let connected = ref(false)
let opponentFound = ref(false)
// game starts after 3 secs of finding a opponent
let gameStartCountDown = ref(-1) // -1 = not found, 0 = game start, > 0 = counting down
let gameStartCountDownInterval = null

// identity indicators
let playerId;
let playerName;
// TODO: 之後改成可以改
const player1Color = "rgba(185, 28, 28, 1)"
const player2Color = "rgba(202, 138, 4, 1)"


let blockBoardPlayer1 = ref(Array.from({ length: 20 }, 
                                    () => Array(10).fill(0)));
let blockBoardPlayer2 = ref(Array.from({ length: 20 }, 
                                    () => Array(10).fill(0)));

let nextBlockTypePlayer1 = ref("")
let nextBlockTypePlayer2 = ref("")

let holdBlockTypePlayer1 = ref("")
let holdBlockTypePlayer2 = ref("")

let scorePlayer1 = ref(0)
let scorePlayer2 = ref(0)

let redrawTetrisUICounter = ref(0)
let uiHeight = ref(0)
let TetrisUIContainerRef = useTemplateRef("TetrisUIContainerRef")

// export default {
//   data() {
//     return {
//       connected: false,  // 用來顯示連線狀態
//       player1Move: null,
//       socket: null,  // WebSocket 客戶端
//       playerId: null,
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

      if (gameStartCountDown.value != 0) return

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

  

onMounted(() => {
    // window.addEventListener("keydown", handleKeyDown);


    // 當組件加載時，建立 WebSocket 連線
    socket = new WebSocket("ws://localhost:8080/game");
    // 註冊 WebSocket 事件
    socket.onopen = () => {
      connected.value = true;  // 當 WebSocket 連線成功時設為 true
      console.log("Connected to server!");
    };

    socket.onclose = () => {
      connected.value = false;  // 當 WebSocket 斷開時設為 false
      opponentFound.value = false
      gameStartCountDown.value = -1
      console.log("Disconnected from server!");
    };

    socket.onerror = (error) => {
      console.log("WebSocket Error:", error);
    };

     // 接收後端訊息
    socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
    
      if (data.player_id) {
        opponentFound.value = true

        // game start after 3 sec
        gameStartCountDown.value = 3;
        gameStartCountDownInterval = setInterval(() => {
          gameStartCountDown.value -= 1;
          if (gameStartCountDown.value <= 0) { 
            nextTick(() => { handleResize() })
            clearInterval(gameStartCountDownInterval);
           }
        }, 1000)

        playerId = data.player_id;  // 設置玩家ID
        console.log('玩家ID:', playerId);
        playerName = playerId == 1 ? "Player 1" : "Player 2" 
      }
      
      if(data.Player_Block_Board) {
        blockBoardPlayer1.value = data.Player_Block_Board[0]
        blockBoardPlayer2.value = data.Player_Block_Board[1]
      }

      if(data.Player_Next_Block) {
        nextBlockTypePlayer1.value = data.Player_Next_Block[0]
        nextBlockTypePlayer2.value = data.Player_Next_Block[1]
      }
      
      
      if(data.Player_Hold_Block_type) {
        holdBlockTypePlayer1.value = data.Player_Hold_Block_type[0]
        holdBlockTypePlayer2.value = data.Player_Hold_Block_type[1]
      }


      if(data.PlayerScore) {
        scorePlayer1.value = data.PlayerScore[0]
        scorePlayer2.value = data.PlayerScore[1]
      }

      redrawTetrisUICounter.value += 1
    }
    
  })


// 監聽鍵盤事件
onUnmounted(() => {
    // 在組件銷毀前移除監聽器並關閉連線
    window.removeEventListener("keydown", handleKeyDown);
    window.removeEventListener("resize", handleResize);
    if (socket) {
      socket.close();  // 斷開 WebSocket 連線
    }
  })



const handleResize = () => {
  if(!TetrisUIContainerRef.value) return
  const tetrisUIHeightFactor = 0.8
  uiHeight.value = TetrisUIContainerRef.value.clientHeight * tetrisUIHeightFactor
}

//check resize if TetrisBoardRerenders
watch(redrawTetrisUICounter, handleResize)

window.addEventListener("keydown", handleKeyDown);
window.addEventListener("resize", handleResize);

</script>

<template>    
  <!-- <h1>Tetris Game</h1>
  <p v-if="connected">Successfully connected to the server!</p>
  <p v-else>Waiting for connection...</p> -->

  <div class="flex flex-col justify-center items-center h-screen w-screen">
    <div v-show="!opponentFound">
      <RouterLink to="/" class="absolute top-0 left-0 text-gray-600 hover:text-black hover:scale-110 ease-linear duration-[80ms]"><Icon icon="line-md:arrow-small-left" width="92" height="92"/></RouterLink>
      <div class="flex flex-row items-end">
        <h1 class="font-jersey text-7xl">Looking for your next opponent</h1>
        <Icon icon="svg-spinners:3-dots-bounce" width="48" height="48" />
      </div>
    </div>

    <div v-show="gameStartCountDown > 0"  class="fixed flex flex-col justify-center items-center h-screen w-screen top-0 left-0 bg-white z-10">
        <h1 class="font-jersey text-7xl text-center">Game will start in: </h1>
        <div class="flex items-center justify-center h-36 w-36">
          <Transition name="gameCountdown">
          <!-- 3 -->
          <Icon v-if="gameStartCountDown == 3" icon="formkit:number" width="144" height="144" class="absolute" /> 
          <!-- 2 -->
          <Icon v-else-if="gameStartCountDown == 2" icon="tabler:circle-number-2" width="144" height="144" class=" absolute"/> 
          <!-- 1 -->
          <Icon v-else icon="mdi:number-1-box-multiple-outline" width="144" height="144" class="absolute"/>
        </Transition>
        </div>
        <!-- <h1 class="font-jersey text-9xl text-center">{{gameStartCountDown}}</h1> -->
    </div>

    <Transition name="gameSceneContainerTransition">
      <div id="gameSceneContainer" v-show="gameStartCountDown == 1 || gameStartCountDown == 0" class="h-full w-full">
        <div id="Header" class="absolute top-0 left-0 w-full z-10">
          <div class="flex flex-col justify-center">
            <div class="w-full">
              <RouterLink to="/" class="absolute left-0 text-gray-600 hover:text-black hover:scale-110 ease-linear duration-[80ms]"><Icon icon="line-md:arrow-small-left" width="92" height="92"/></RouterLink>
            </div>

            <span class="inline-flex flex-row justify-center items-end">
              <h2 class="font-jersey text-7xl mr-8 h-full">
                You Are 
              </h2>
              <h2 class="font-jersey text-7xl h-full" :style="{color: `${playerId == 1 ? player1Color : player2Color}`}">
                {{ playerName }}
              </h2>
            </span>
          </div>
        </div>
        
          <div  id="TetrisUIContainer" ref="TetrisUIContainerRef" class="relative flex flex-row justify-around items-center h-full w-full z-0">
            <div class="flex flex-col justify-center items-center h-full w-full">
              <h2 class="font-jersey text-7xl text-red-700">Player 1</h2>
              <TetrisBoard :Board=blockBoardPlayer1 :key="`Player1_${redrawTetrisUICounter}`" :UIHeight="uiHeight" :NextBlockType="nextBlockTypePlayer1" :HoldBlockType="holdBlockTypePlayer1" :Score="scorePlayer1"></TetrisBoard>
            </div>
            <div class="flex flex-col justify-center items-center h-full w-full">
              <h2 class="font-jersey text-7xl text-yellow-600">Player 2</h2>
              <TetrisBoard :Board=blockBoardPlayer2 :key="`Player2_${redrawTetrisUICounter}`" :UIHeight="uiHeight" :NextBlockType="nextBlockTypePlayer2" :HoldBlockType="holdBlockTypePlayer2" :Score="scorePlayer2"></TetrisBoard>   
            </div>
          </div>
        </div>
    </Transition>
  </div>
    
  </template>

<style lang="scss" scoped>

.gameCountdown-enter-active,
.gameCountdown-leave-active {
  position: absolute;
  transition: opacity 0.5s ease, transform 0.5s ease;
}

.gameCountdown-enter-from,
.gameCountdown-leave-to {
  opacity: 0;
  transform: scale(0.9);
}

.gameCountdown-enter-to,
.gameCountdown-leave-from {
  opacity: 1;
  transform: scale(1);
}

.gameSceneContainerTransition-enter-active,
.gameSceneContainerTransition-leave-active {
  position: absolute;
  transition: opacity 2s ease, transform 2s ease;
}

.gameSceneContainerTransition-enter-from,
.gameSceneContainerTransition-leave-to {
  opacity: 0;
  transform: scale(0.9);
}

.gameSceneContainerTransition-enter-to,
.gameSceneContainerTransition-leave-from {
  opacity: 1;
  transform: scale(1);
}


</style>
