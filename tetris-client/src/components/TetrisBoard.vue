<script setup>
import { onMounted, useTemplateRef} from 'vue';

const props = defineProps({
  Board: {
    type: Array, //[20][10]
    Required: true,
  }
})

const canvasRef = useTemplateRef('TetrisCanvasRef')

const getColor = (row_index, col_index) => {
    let cellVal = props.Board[row_index][col_index]
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


// watch(props.Board, (newBoard, oldBoard) => {drawTetrisBoard()}, {deep: true})

const resizeCallback = () => {
  recalCellRelaPos()
  drawTetrisBoard()
}
window.addEventListener("resize", resizeCallback)

onMounted(() => {
  recalCellRelaPos()
  drawTetrisBoard()
  canvasRef.value.focus()
})



</script>

<template>
  <div id="board">
    <canvas id="TetrisCanvas" ref="TetrisCanvasRef"></canvas>
  </div>
</template>

<style scoped>
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
