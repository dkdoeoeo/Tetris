<script>
export default {
  data() {
    return {
      ws: null, // WebSocket 連線
      gameState: null // 遊戲狀態
    };
  },
  mounted() {
    // 建立 WebSocket 連線
    this.ws = new WebSocket("ws://localhost:8080/game/1");

    this.ws.onopen = () => {
    console.log("WebSocket 連線成功");
  };

    // 接收伺服器回傳的遊戲狀態
    this.ws.onmessage = (event) => {
      const data = JSON.parse(event.data);
      this.updateGameState(data); // 更新遊戲狀態
      console.log("收到後端的訊息:", event.data);
    };

    // 鍵盤事件偵聽（用戶輸入）
    window.addEventListener("keydown", this.handleKeyPress);
  },
  methods: {
    handleKeyPress(event) {
      console.log("按鍵:", event.key);
      let action = null;
      switch (event.key) {
        case "ArrowLeft":
          action = "moveLeft";
          break;
        case "ArrowRight":
          action = "moveRight";
          break;
        case "ArrowDown":
          action = "softDrop";
          break;
        case "ArrowUp":
          action = "rotate";
          break;
        default:
          return; // 忽略無關按鍵
      }

      // 傳送操作輸入給伺服器
      if (action && this.ws) {
        this.ws.send(JSON.stringify({ type: "playerAction", action: action }));
      }
    },
    updateGameState(data) {
      // 更新 Vue 的 data 中的遊戲狀態
      this.gameState = data;
      // 此處可以根據 gameState 進行畫面的重繪
    }
  },
  beforeDestroy() {
    // 清除 WebSocket 連線與事件
    if (this.ws) {
      this.ws.close();
    }
    window.removeEventListener("keydown", this.handleKeyPress);
  }
};
</script>