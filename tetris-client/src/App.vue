<template>
  <div>
    <!-- 顯示遊戲畫面 -->
    <h1>Tetris Game</h1>
    <p v-if="connected">Successfully connected to the server!</p>
    <p v-else>Waiting for connection...</p>
  </div>
</template>

<script>

export default {
  data() {
    return {
      connected: false,  // 用來顯示連線狀態
      player1Move: null,
      socket: null,  // WebSocket 客戶端
    };
  },
  mounted() {
    // 當組件加載時，建立 WebSocket 連線
    this.socket = new WebSocket("ws://localhost:8081/game");

    // 註冊 WebSocket 事件
    this.socket.onopen = () => {
      this.connected = true;  // 當 WebSocket 連線成功時設為 true
      console.log("Connected to server!");
    };

    this.socket.onclose = () => {
      this.connected = false;  // 當 WebSocket 斷開時設為 false
      console.log("Disconnected from server!");
    };

    this.socket.onerror = (error) => {
      console.log("WebSocket Error:", error);
    };

    this.socket.onmessage = (event) => {
      // 監聽來自伺服器的消息
      const gameState = JSON.parse(event.data); // 假設伺服器發送 JSON 資料
      console.log("Received game state:", gameState);
    };

    // 監聽鍵盤事件
    window.addEventListener("keydown", this.handleKeyDown);
  },
  beforeDestroy() {
    // 在組件銷毀前移除監聽器並關閉連線
    window.removeEventListener("keydown", this.handleKeyDown);
    if (this.socket) {
      this.socket.close();  // 斷開 WebSocket 連線
    }
  },
  methods: {
    handleKeyDown(event) {
      // 監聽鍵盤按鍵事件
      switch (event.key) {
        case 'ArrowLeft':
          this.moveLeft();
          break;
        case 'ArrowRight':
          this.moveRight();
          break;
        case 'ArrowDown':
          this.accelerateFall();
          break;
        case ' ':
          this.rotateBlock();
          break;
        case 'p':
          this.pauseGame();
          break;
        default:
          break;
      }
    },
    moveLeft() {
      // 觸發左移操作
      console.log('Move left');
      this.sendMoveToServer('left');
    },
    moveRight() {
      // 觸發右移操作
      console.log('Move right');
      this.sendMoveToServer('right');
    },
    accelerateFall() {
      // 觸發加速下落
      console.log('Accelerate fall');
      this.sendMoveToServer('down');
    },
    rotateBlock() {
      // 觸發旋轉操作
      console.log('Rotate block');
      this.sendMoveToServer('rotate');
    },
    pauseGame() {
      // 觸發暫停遊戲操作
      console.log('Pause game');
      this.sendMoveToServer('pause');
    },
    sendMoveToServer(moveType) {
       // 使用 WebSocket 發送消息到伺服器
       const message = JSON.stringify({
        player: 1,  // 或者依照玩家的ID設置
        move: moveType,
      });
      this.socket.send(message);  // 發送消息
    }
  },
};
</script>

<style scoped>
/* 這裡可以加上樣式 */
</style>
