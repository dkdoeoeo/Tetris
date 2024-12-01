<script setup>
import {onMounted, ref, watch} from 'vue';

const props = defineProps({
  Board: {
    type: Array, //[20][10]
    Required: true,
  }, 
  UIHeight: {
    type: Number,
    Required: true
  },
  NextBlockType: {
    type: String,
    Required: true
  },
  Score: {
    type:Number,
    Required: true
  }
})

// current theme settings
let cellBgColor = "linear-gradient(145deg, #4c9ed9, #2a6f91)"

// board
// const boardRef = useTemplateRef('TetrisBoardRef')
let boardHeight = ref(0)
let cellSideLength = ref(0)
let cellPadding = ref(0)

// next block preview
let nextBlockSideLength = ref(0)
let nextBlockColor = ""
let nextBlockDisplay = ref([
        [0, 0, 0, 0],
        [0, 0, 0, 0],
        [0, 0, 0, 0],
        [0, 0, 0, 0]
      ])



const resizeBoard = () => { 
  // set height to a factor of BoardHeight
  let heightFactor = 0.8
  boardHeight.value = props.UIHeight * heightFactor

  // //  totalWidth = boardHeight.value * 0.5 + boardHeight.value * 0.3
  // let totalWidth = boardHeight.value * 0.8
  // let totalWidthBound = window.innerWidth * 0.3

  // if (totalWidth > totalWidthBound) {
  //   let boardWidthRatio = 3 / 8
  //   boardHeight.value = totalWidthBound * boardWidthRatio * 2
  // }
  cellSideLength.value = boardHeight.value * 0.05
  cellPadding.value = cellSideLength.value * 0.1

  nextBlockSideLength.value = boardHeight.value * 0.3
}

const calNextBlockDisplay = () => {
  let nextBlockType = props.NextBlockType
  console.log(nextBlockType)
  switch (nextBlockType) {
      case "I":
        nextBlockColor = 'lightblue'
        nextBlockDisplay.value = [
          [0, 0, 0, 0],
          [0, 0, 0, 0],
          [1, 1, 1, 1],
          [0, 0, 0, 0]
        ]
        break
      case "L":
        nextBlockColor = 'orange'
        nextBlockDisplay.value = [
          [0, 0, 0, 0],
          [0, 1, 0, 0],
          [0, 1, 0, 0],
          [0, 1, 1, 0]
        ]
        break
      case "J":
        nextBlockColor = 'darkblue'
        nextBlockDisplay.value = [
          [0, 0, 0, 0],
          [0, 0, 1, 0],
          [0, 0, 1, 0],
          [0, 1, 1, 0]
        ]
        break
      case "O":
        nextBlockColor = 'yellow'
        nextBlockDisplay.value = [
          [0, 0, 0, 0],
          [0, 1, 1, 0],
          [0, 1, 1, 0],
          [0, 0, 0, 0]
        ]
        break
      case "Z":
        nextBlockColor = 'red'
        nextBlockDisplay.value = [
          [0, 0, 0],
          [1, 1, 0],
          [0, 1, 1],
          [0, 0, 0],
        ]
        break
      case "S":
        nextBlockColor = 'green'
        nextBlockDisplay.value = [
          [0, 0, 0],
          [0, 1, 1],
          [1, 1, 0],
          [0, 0, 0],
        ]
        break
      case "T":
        nextBlockColor = 'purple'
        nextBlockDisplay.value = [
          [0, 0, 0],
          [0, 1, 0],
          [1, 1, 1],
          [0, 0, 0],
        ]
        break
    }
}

// ================================ Board ==============================
const getBoardCellStyle = (rowIndex, colIndex) => {
    let cellVal = props.Board[rowIndex][colIndex]
    let color = "gray"

    switch (cellVal) {
      case 0:
        color = '';
        break
      case 1:
        color = 'lightblue';
        break
      case 2:
        color = 'yellow';
        break
      case 3:
        color = 'purple';
        break
      case 4:
        color = 'red';
        break
      case 5:
        color = 'green';
        break
      case 6:
        color = 'darkblue';
        break
      case 7:
        color = 'orange';
        break
      case 8:
        color = 'gray';
        break
    }
    

    return {
      width: `${cellSideLength.value}px`,
      height: `${cellSideLength.value}px`,
      top: `0px`,
      left: `${cellSideLength.value * colIndex}px`,
      border: `${cellPadding.value}px solid`,
      borderColor: `lightgray gray gray lightgray`,
      backgroundClip: `content-box`,
      backgroundColor:`${color}`,
      backgroundImage: `${color} ${cellBgColor}`,

      // Theme 0 ${color} linear-gradient(145deg, #4c9ed9, #2a6f91)
    };
   
}

const getBoardRowStyle = (rowIndex) => {
  return {
      width: `${cellSideLength.value * 10}px`,
      height: `${cellSideLength.value}px`,
      top: `${cellSideLength.value * rowIndex}px`,
      left: `0px`,
    };
}
// ================================ Board ==============================



// ================================ Next Block (Preview) ==============================
const getNextBlockCellStyle = (rowIndex, colIndex) => {

  if (nextBlockDisplay.value[rowIndex][colIndex] == 1) {
    return {
      width: `${cellSideLength.value}px`,
      height: `${cellSideLength.value}px`,
      border: `${cellPadding.value}px solid`,
      borderColor: `lightgray gray gray lightgray`,
      backgroundClip: `content-box`,
      backgroundColor:`${nextBlockColor}`,
    }
  }
  else {
    return {
      width: `${cellSideLength.value}px`,
      height: `${cellSideLength.value}px`,
      backgroundColor:`transparent`,
    }
  }
}
// ================================ Next Block (Preview) ==============================


const resizeCallback = () => {
  resizeBoard()
}

// watch(props.UIHeight, ()=>{resizeBoard()})
watch(() => props.NextBlockType, calNextBlockDisplay)

onMounted(() => {
  resizeBoard()
  calNextBlockDisplay()
  window.addEventListener("resize", resizeCallback)
})

</script>

<template>
  <div id="TetrisUI" class="relative flex flex-row justify-center gap-3 items-start min-h-min min-w-min" :style="{height: `${UIHeight}`,}">
    <div class="flex flex-col gap-7">
      <div class="relative flex flex-col items-center justify-between border-black rounded-3xl border-4 min-h-min p-2" :style="{width: `${nextBlockSideLength}px`}">
        <h2 class="font-jersey text-5xl whitespace-nowrap">Next Block:</h2>
        <div id="NextBlock" class="flex flex-col justify-center items-center">
          <div v-for="(row, previewRowIndex) in nextBlockDisplay" :key="`preview_row_${previewRowIndex}`" class="flex flex-row justify-center gap-0" :style="{width: `${cellSideLength * nextBlockDisplay.length}px`, height: `${cellSideLength}px`}">
            <span v-for="(cellVal, previewColIndex) in row" :key="`preview_cell_${previewRowIndex * nextBlockDisplay.length + previewColIndex}`" class="" :style="getNextBlockCellStyle(previewRowIndex, previewColIndex)">
            </span>
          </div>
        </div>
      </div>
      <div class="relative flex flex-col items-center justify-between border-black rounded-3xl border-4 min-h-min p-2" :style="{width: `${nextBlockSideLength}px`, height: `${nextBlockSideLength * 0.3}px`}">
        <h2 class="font-jersey text-5xl whitespace-nowrap">Score:</h2>
        <h2 class="font-jersey text-5xl whitespace-nowrap">{{ Score }}</h2>

      </div>
    </div>
    <div id="board" ref="TetrisBoardRef" class="flex gap-0 flex-col relative aspect-[1/2]" :style="{height: `${boardHeight}px`}">
      <div v-for="(row, rowIndex) in Board" :key="`row_${rowIndex}`" 
      class="absolute"
      :style="getBoardRowStyle(rowIndex)">
        <span v-for="(cellVal, colIndex) in row" :key="`cell_${rowIndex * 10 + colIndex}`" 
        class="absolute min-w-2 min-h-2" 
        :style = "getBoardCellStyle(rowIndex, colIndex)">
        </span>
      </div>
    </div>
  </div>
  
</template>

<style scoped>



</style>
