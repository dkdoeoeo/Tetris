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
    this.socket = new WebSocket("ws://140.118.178.119:8080/game");

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

     // 接收後端訊息
     this.socket.onmessage = (event) => {
      const data = JSON.parse(event.data);
      if (data.player_id) {
        this.playerId = data.player_id;  // 設置玩家ID
        console.log('玩家ID:', this.playerId);
      }
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
        case "ArrowUp":
          this.rotateBlock();
          break;
        case ' ':
          this.hard_drop();
          break;
        case 'Shift':
          this.hold();
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
      this.sendMoveToServer('soft_drop');
    },
    rotateBlock() {
      // 觸發旋轉操作
      console.log('Rotate block');
      this.sendMoveToServer('rotate');
    },
    hard_drop() {
      // 觸發加速下降操作
      console.log('hard_drop');
      this.sendMoveToServer('hard_drop');
    },
    hold() {
      // 觸發加速下降操作
      console.log('hold');
      this.sendMoveToServer('hold');
    },
    sendMoveToServer(moveType) {
       // 使用 WebSocket 發送消息到伺服器
       const message = JSON.stringify({
        player: this.playerId,  // 依照玩家的ID設置
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
