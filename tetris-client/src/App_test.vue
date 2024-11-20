<template>
  <div>
    <!-- 顯示遊戲畫面 -->
    <h1>Tetris Game</h1>
    <p v-if="connected">Successfully connected to the server!</p>
    <p v-else>Waiting for connection...</p>
  </div>
  <div class="game">
    <div class="game-div">
      <div class="game-min">
        <div class="row" v-for="(row, i) in frame" :key="i">
          <p class="element" v-for="(col, j) in row" :key="j" :style="{ background: col.bg }"></p>
        </div>
      </div>
      <div class="right-div">
        <div class="ass">
          <div class="row" v-for="(row, i) in secondaryScreen" :key="i">
            <p class="element" v-for="(col, j) in row" :key="j" :style="{ background: col.bg }"></p>
          </div>
        </div>
        <div>
          <p>得分</p>
          <p class="data">{{ score }}</p>
        </div>
        <div>
          <p>等級</p>
          <p class="data">{{ level }}</p>
        </div>
        <div>
          <p>消除</p>
          <p class="data">{{ times }}</p>
        </div>
        <div class="play" @click="stopGame()">
          開始
        </div>
      </div>
    </div>
    <div class="control">
      <div class="change bt" @click="change1()">变化</div>
      <div class="zy">
        <div class="left bt" @click="moveLeft()">向左</div>
        <div class="right bt" @click="moveRight()">向右</div>
      </div>
      <div class="down bt" @click="moveDown()">向下</div>
    </div>
  </div>
</template>

<script>

export default {
  data() {
    return {
      connected: false,  // 用來顯示連線狀態
      player1Move: null,
      socket: null,  // WebSocket 客戶端
      row: 20, //行
      col: 10, //列
      frame: [],//界面
      secondaryScreen: [],//副屏幕
      bg: '#eee',
      block: [],//方块集合
      now: { b: 0, c: 0 },//当前方块以及其旋转角度
      next: { b: 0, c: 0 },//下一个方块以及其旋转角度
      nowBlock: [],//当前形状数据
      nextBlock: [],//下一个形状数据
      xz: 0,//当前旋转角度
      timer: '',//自动下落
      speed: 800,//速度
      score: 0,//得分
      level: 1,//等级
      times: 0,//消除次数
      stop: true,//是否停止
      removeRow: [],//消除的行记录

    };
  },
  mounted() {
    // 當組件加載時，建立 WebSocket 連線
    this.socket = new WebSocket("ws://140.118.214.40:8080/game");

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
      if (data.Player1Block) {
        this.Player1Block = data.Player1Block;
        console.log('\nPlayer1Block:', this.Player1Block);
        this.fillblock();
      }
      else if (data.Player2Block) {
        this.Player2Block = data.Player2Block;
        console.log('\nPlayer2Block:', this.Player2Block);
        this.fillblock();
      }
    };
    // 監聽鍵盤事件
    window.addEventListener("keydown", this.handleKeyDown);
    this.gameFrame();
    // fillblock();
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
    },
    //游戏框架
    gameFrame() {
      //主屏幕
      for (let i = 0; i < this.row; i++) {
        let a = [];
        for (let j = 0; j < this.col; j++) {
          let b = {
            data: 0,
            bg: this.bg,
          };
          a.push(b);
        }
        this.frame.push(a);
      }
      //副屏
      for (let i = 0; i < 4; i++) {
        let a = [];
        for (let j = 0; j < 4; j++) {
          let b = {
            data: 0,
            bg: this.bg,
          };
          a.push(b);
        }
        this.secondaryScreen.push(a);
      }
      //模拟格子被占用
      // this.frame[4][4].bg = "#00aaee";
      // this.frame[4][4].data = 1;
    },

    fillblock() {
      for (let i = 0; i < 20; i++) {
        for (let j = 0; j < 10; j++) {
          if (this.Player1Block[i][j] == 1 || this.Player2Block[i][j] == 1) {
            this.frame[i][j].bg = "#66B3FF";
          }
          else if (this.Player1Block[i][j] == 2 || this.Player2Block[i][j] == 2) {
            this.frame[i][j].bg = "#FFFF6F";
          }
          else if (this.Player1Block[i][j] == 3 || this.Player2Block[i][j] == 3) {
            this.frame[i][j].bg = "#BE77FF";
          }
          else if (this.Player1Block[i][j] == 4 || this.Player2Block[i][j] == 4) {
            this.frame[i][j].bg = "#FF2D2D";
          }
          else if (this.Player1Block[i][j] == 5 || this.Player2Block[i][j] == 5) {
            this.frame[i][j].bg = "#79FF79";
          }
          else if (this.Player1Block[i][j] == 6 || this.Player2Block[i][j] == 6) {
            this.frame[i][j].bg = "#0000E3";
          }
          else if (this.Player1Block[i][j] == 7 || this.Player2Block[i][j] == 7) {
            this.frame[i][j].bg = "#FF8040";
          }
          else if (this.Player1Block[i][j] == 8 || this.Player2Block[i][j] == 8) {
            this.frame[i][j].bg = "#5B5B5B";  //灰色
          }
          else {
            this.frame[i][j].bg = "#eee";
          }
        }
      }
    },


    //渲染方块
    getBlock(e) {
      this.block = blockMod(color[e]);
      console.log(this.block);
    },

    //下一个形状
    async getNext() {
      //随机获取形状
      // this.next.b = Math.floor(Math.random() * this.block.length)
      // this.next.c = Math.floor(Math.random() * 4);
    },
    //渲染当前形状
    init() {
      //获取到下一个形状数据
      this.now = JSON.parse(JSON.stringify(this.next));
      this.xz = this.now.c

      //当前形状数据
      this.nowBlock = JSON.parse(JSON.stringify(this.block[this.now.b]));

      //渲染形状数据
      this.renderBlock(this.nowBlock, this.frame, 1);
      //旋转
      if (this.now.c > 0) {
        for (let i = 0; i < this.now.c; i++) {
          this.change(this.nowBlock, this.frame, this.now, i);
        }
      }

      this.getNext().then(() => {
        if (this.nextBlock.site) {
          this.renderBlock(this.nextBlock, this.secondaryScreen, 0);
        }
        //下一个形状
        this.nextBlock = JSON.parse(JSON.stringify(this.block[this.next.b]));
        //渲染形状数据
        this.renderBlock(this.nextBlock, this.secondaryScreen, 1);
        //旋转
        if (this.next.c > 0) {
          for (let i = 0; i < this.next.c; i++) {
            this.change(this.nextBlock, this.secondaryScreen, this.next, i);
          }
        }

      }
      )
    },

    //渲染形状
    //b:方块，d：位置，n:0擦除，1生成，2确定落到最下层
    // renderBlock(b, d, n) {
    //   let c = b.site;
    //   if (n == 0) {
    //     //擦除
    //     for (let i = 0; i < c.length; i += 2) {
    //       d[c[i]][c[i + 1]].bg = this.bg;
    //     }
    //   } else if (n == 1) {
    //     //生成
    //     for (let i = 0; i < c.length; i += 2) {
    //       d[c[i]][c[i + 1]].bg = b.color;
    //     }
    //   } else if (n == 2) {
    //     //确定位置
    //     for (let i = 0; i < c.length; i += 2) {
    //       d[c[i]][c[i + 1]].data = 1;
    //     }
    //   }

    // },
    //旋转b:当前方块、d:要渲染的位置、z:渲染的对象现在还是下一个、xz:当前选转角度
    //旋转
    // change(b, d, z, xz) {
    //   this.renderBlock(b, d, 0);
    //   //记录第一块位置
    //   const x = b.site[0];
    //   const y = b.site[1];
    //   for (let i = 0; i < b.site.length; i += 2) {
    //     let a = b.site[i];
    //     b.site[i] = b.site[i + 1] - y + x + transition[z.b][xz].x;
    //     b.site[i + 1] = -(a - x) + y + transition[z.b][xz].y;
    //   }
    //   xz++;
    //   if (xz == 4) {
    //     xz = 0;
    //   }
    //   this.renderBlock(b, d, 1);
    // },

    //自动下落
    // autMoveDown() {

    //   this.timer = setInterval(() => {
    //     this.moveDown();

    //   }, this.speed);
    // },
    //开始与暂停
    stopGame() {
      this.stop = !this.stop;
      if (this.stop) {
        clearInterval(this.timer);
      } else {
        this.autMoveDown();
      }
    },

    // //向下移动
    // moveDown() {
    //   if (this.canMove(3)) {
    //     //先清理
    //     this.renderBlock(this.nowBlock, this.frame, 0);
    //     for (let i = 0; i < this.nowBlock.site.length; i += 2) {
    //       //向下移动一位
    //       this.nowBlock.site[i]++;
    //     }
    //     //再渲染形状数据
    //     this.renderBlock(this.nowBlock, this.frame, 1);
    //   } else {
    //     //已经不能下落了
    //     // clearInterval(this.timer);

    //     this.renderBlock(this.nowBlock, this.frame, 2);

    //     //判断是否可以消除
    //     this.removeBlock();
    //   }
    // },

    // //向左移动
    // moveLeft() {
    //   if (this.canMove(2)) {
    //     //先清理
    //     this.renderBlock(this.nowBlock, this.frame, 0);
    //     for (let i = 0; i < this.nowBlock.site.length; i += 2) {
    //       //向下移动一位
    //       this.nowBlock.site[i + 1]--;
    //     }
    //     //再渲染形状数据
    //     this.renderBlock(this.nowBlock, this.frame, 1);
    //   }
    // },

    // //向右移动
    // moveRight() {
    //   if (this.canMove(1)) {
    //     //先清理
    //     this.renderBlock(this.nowBlock, this.frame, 0);
    //     for (let i = 0; i < this.nowBlock.site.length; i += 2) {
    //       //向下移动一位
    //       this.nowBlock.site[i + 1]++;
    //     }
    //     //再渲染形状数据
    //     this.renderBlock(this.nowBlock, this.frame, 1);
    //   }
    // },

    //下落时旋转
    //旋转b:当前方块、xz:当前选转角度
    change1() {
      const b = JSON.parse(JSON.stringify(this.nowBlock));
      //记录第一块位置
      const x = b.site[0];
      const y = b.site[1];

      let n = true;
      for (let i = 0; i < b.site.length; i += 2) {
        let a = b.site[i];
        b.site[i] = b.site[i + 1] - y + x + transition[this.now.b][this.xz].x;
        b.site[i + 1] = -(a - x) + y + transition[this.now.b][this.xz].y;

        //判断旋转后该点是否合理
        if (b.site[i + 1] < 0 || b.site[i + 1] >= this.col || b.site[i] >= this.row || this.frame[b.site[i]][b.site[i + 1]].data > 0) {
          n = false;
        }
      }
      //符合条件 
      if (n) {
        this.renderBlock(this.nowBlock, this.frame, 0);
        this.xz++;
        if (this.xz == 4) {
          this.xz = 0;
        }
        this.nowBlock = b;
        this.renderBlock(this.nowBlock, this.frame, 1);
      }
    },

    //是否可移动判断
    //预判能否移动或变化，e:1.右移，2.左移，3.下移，4.变化
    canMove(e) {
      let d = 0;
      let c = this.nowBlock.site;
      switch (e) {
        case 1:
          for (let i = 0; i < c.length; i += 2) {
            if (c[i + 1] >= this.col - 1) {
              return false;
            }
            //判断下一个位置是否被占用
            d += this.frame[c[i]][c[i + 1] + 1].data;
          }
          if (d > 0) {
            return false;
          }
          break;
        case 2:
          for (let i = 0; i < c.length; i += 2) {
            if (c[i + 1] <= 0) {
              return false;
            }
            //判断下一个位置是否被占用
            d += this.frame[c[i]][c[i + 1] - 1].data;
          }
          if (d > 0) {
            return false;
          }
          break;
        case 3:
          for (let i = 0; i < c.length; i += 2) {
            if (c[i] >= this.row - 1) {
              return false;
            }
            //判断下一个位置是否被占用
            d += this.frame[c[i] + 1][c[i + 1]].data;
          }
          if (d > 0) {
            return false;
          }
          break;
      }
      return true;
    },
    //判断是否可以消除
    removeBlock() {
      //遍历整个界面
      for (let i = 0; i < this.row; i++) {
        let a = 0;
        for (let j = 0; j < this.col; j++) {
          if (this.frame[i][j].data == 1) {
            a++
          }
        }
        if (a == this.col) {
          //说明这一行已经占满可以消除
          this.removeRow.push(i);
        }
      }
      //获取是否可以消除行
      if (this.removeRow.length > 0) {
        let l = this.removeRow;
        for (let i = 0; i < l.length; i++) {
          let j = 0;
          let timer = setInterval(() => {
            this.frame[l[i]][j] = { bg: this.bg, data: 0, }
            j++;
            if (j == this.col) {
              clearInterval(timer);
            }
          }, 20)
        }
        setTimeout(() => {
          //上面方块下移,从下往上判断
          for (let i = this.row - 1; i >= 0; i--) {
            let a = 0;
            for (let j = 0; j < l.length; j++) {
              if (l[j] > i) {
                a++;
              }
            }
            if (a > 0) {
              for (let k = 0; k < this.col; k++) {
                if (this.frame[i][k].data == 1) {
                  //先向下移动
                  this.frame[i + a][k] = this.frame[i][k];
                  //再清楚当前
                  this.frame[i][k] = { data: 0, bg: this.bg };
                }
              }
            }
          }
          //清除行记录
          this.removeRow = [];
          //生成下一个
          this.init();
        }, 20 * this.col)
        //数据处理
        //消除次数+1
        this.times++;
        //等级向下取整+1
        let lev = Math.floor(this.times / 10) + 1;
        if (lev > this.level) {
          this.level = lev;
          //速度
          if (this.level < 21) {
            //20级内做减法
            this.speed = 800 - (this.level - 1) * 40;
          } else {
            this.speed = 30;
          }
          //清除当前下落
          clearInterval(this.timer);
          //加速
          this.autMoveDown();
        }
        this.level = this.times;
        //分数消除一行100，两行300，三行600，四行1000
        this.score += ((l.length) * (l.length + 1) / 2) * 100 * this.level;

      } else {
        //生成下一个
        this.init();
      }
    }
  }
};
</script>

<style lang="less" scoped>
.game {
  .game-div {
    display: flex;

    .game-min {

      .row {
        display: flex;
        padding-top: 1px;

        .element {
          width: 20px;
          height: 20px;
          padding: 0;
          margin: 0 1px 0 0;
        }
      }
    }

    .right-div {
      padding-left: 12px;

      .ass {
        .row {
          display: flex;
          padding-top: 1px;

          .element {
            width: 16px;
            height: 16px;
            padding: 0;
            margin: 0 1px 0 0;
          }
        }
      }

      .data {
        font-weight: 700;
        font-size: 16px;
        color: #00aaee;
      }

      .play {
        width: 56px;
        height: 32px;
        background: #bbb;
        text-align: center;
        line-height: 32px;
      }
    }
  }

  .control {
    padding-top: 12px;

    .bt {
      height: 36px;
      background: #aaa;
      width: 94px;
      margin: 6px;
      text-align: center;
      line-height: 36px;
      color: #fff;
    }

    .zy {
      display: flex;
    }

    .change,
    .down {
      width: 200px;
    }
  }
}
</style>