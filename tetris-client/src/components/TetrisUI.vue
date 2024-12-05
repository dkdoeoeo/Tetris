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
  HoldBlockType: {
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
let cellBgColorDarker = "linear-gradient(145deg, #387ca7, #1f556d)"

// board
// const boardRef = useTemplateRef('TetrisBoardRef')
let boardHeight = ref(0)
let boardWidth = ref(0)
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

// hold block
let holdBlockSideLength = ref(0)
let holdBlockColor = ""
let holdBlockDisplay = ref([
        [0, 0, 0, 0],
        [0, 0, 0, 0],
        [0, 0, 0, 0],
        [0, 0, 0, 0]
      ])

const resizeBoard = () => { 
  // set height to a factor of BoardHeight
  let heightFactor = 0.8
  boardHeight.value = props.UIHeight * heightFactor
  boardWidth.value = boardHeight.value * 0.5
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

const calDisplays = () => {
  let blockTypes = [props.NextBlockType, props.holdBlockDisplay]
  let blockDisplay;
  let blockColor;

  for (let index = 0; index < 2; index++) {
    let blockType = blockTypes[index];
    switch (blockType) {
      case "I":
        blockColor = 'lightblue'
        blockDisplay = [
          [0, 0, 0, 0],
          [0, 0, 0, 0],
          [1, 1, 1, 1],
          [0, 0, 0, 0]
        ]
        break
      case "L":
        blockColor = 'darkblue'
        blockDisplay = [
          [0, 0, 0, 0],
          [0, 1, 0, 0],
          [0, 1, 0, 0],
          [0, 1, 1, 0]
        ]
        break
      case "J":
        blockColor = 'orange'
        blockDisplay = [
          [0, 0, 0, 0],
          [0, 0, 1, 0],
          [0, 0, 1, 0],
          [0, 1, 1, 0]
        ]
        break
      case "O":
        blockColor = 'yellow'
        blockDisplay = [
          [0, 0, 0, 0],
          [0, 1, 1, 0],
          [0, 1, 1, 0],
          [0, 0, 0, 0]
        ]
        break
      case "Z":
        blockColor = 'green'
        blockDisplay = [
          [0, 0, 0],
          [1, 1, 0],
          [0, 1, 1],
          [0, 0, 0],
        ]
        break
      case "S":
        blockColor = 'red'
        blockDisplay = [
          [0, 0, 0],
          [0, 1, 1],
          [1, 1, 0],
          [0, 0, 0],
        ]
        break
      case "T":
        blockColor = 'purple'
        blockDisplay = [
          [0, 0, 0],
          [0, 1, 0],
          [1, 1, 1],
          [0, 0, 0],
        ]
        break
      //none holding
      default:
        blockColor = 'transparent'
        blockDisplay = [
          [0, 0, 0],
          [0, 0, 0],
          [0, 0, 0],
          [0, 0, 0],
        ]
        break
    }
    
    if (index == 0) { //NextBlockType
      nextBlockColor = blockColor
      nextBlockDisplay.value = blockDisplay
    }
    else { //HoldBlockType
      holdBlockColor = blockColor
      holdBlockDisplay.value = blockDisplay
    }
  }
  
}

// ================================ Board ==============================
const getBoardCellStyle = (rowIndex, colIndex) => {
    let cellVal = props.Board[rowIndex][colIndex]
  

    let color = ""
  
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
      borderColor: cellVal == 9 ?  `black` : `lightgray gray gray lightgray`,
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



// ================================ Next Block ==============================
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
// ================================ Next Block ==============================

// ================================ Hold Block ==============================
const getHoldBlockCellStyle = (rowIndex, colIndex) => {

if (holdBlockDisplay.value[rowIndex][colIndex] == 1) {
  return {
    width: `${cellSideLength.value}px`,
    height: `${cellSideLength.value}px`,
    border: `${cellPadding.value}px solid`,
    borderColor: `lightgray gray gray lightgray`,
    backgroundClip: `content-box`,
    backgroundColor:`${holdBlockColor}`,
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
// ================================ Hold Block ==============================

const resizeCallback = () => {
  resizeBoard()
}

// watch(props.UIHeight, ()=>{resizeBoard()})
watch(() => props.NextBlockType, calDisplays)

onMounted(() => {
  resizeBoard()
  calDisplays()
  window.addEventListener("resize", resizeCallback)
})

</script>

<template>
  <div id="TetrisUI" class="relative flex flex-row justify-center gap-3 items-start min-h-min min-w-min " :style="{height: `${UIHeight}`,}">
    <div class="flex flex-col gap-7">
      <div class="relative flex flex-col items-center justify-between border-black rounded-3xl border-4 min-h-min p-2" :style="{width: `${nextBlockSideLength}px`}">
        <h2 class="font-jersey text-5xl whitespace-nowrap">Next Block:</h2>
        <div id="NextBlock" class="flex flex-col justify-center items-center">
          <div v-for="(row, nextBlockRowIndex) in nextBlockDisplay" :key="`next_row_${nextBlockRowIndex}`" class="flex flex-row justify-center gap-0" :style="{width: `${cellSideLength * nextBlockDisplay.length}px`, height: `${cellSideLength}px`}">
            <span v-for="(cellVal, nextBlockColIndex) in row" :key="`next_cell_${nextBlockRowIndex * nextBlockDisplay.length + nextBlockColIndex}`" class="" :style="getNextBlockCellStyle(nextBlockRowIndex, nextBlockColIndex)">
            </span>
          </div>
        </div>
      </div>
      <div class="relative flex flex-col items-center justify-between border-black rounded-3xl border-4 min-h-min p-2" :style="{width: `${nextBlockSideLength}px`, height: `${nextBlockSideLength * 0.3}px`}">
        <h2 class="font-jersey text-5xl whitespace-nowrap">Score:</h2>
        <h2 class="font-jersey text-5xl whitespace-nowrap">{{ Score }}</h2>
      </div>
    </div>
    
    <div id="board" ref="TetrisBoardRef" class="flex gap-0 flex-col relative aspect-[1/2] border-[10px] border-solid rounded-3xl border-black" :style="{height: `${boardHeight + 20}px`, width: `${boardWidth + 20}px`}">
      <div v-for="(row, rowIndex) in Board" :key="`row_${rowIndex}`" 
      class="absolute"
      :style="getBoardRowStyle(rowIndex)">
        <span v-for="(cellVal, colIndex) in row" :key="`cell_${rowIndex * 10 + colIndex}`" 
        class="absolute min-w-2 min-h-2" 
        :style = "getBoardCellStyle(rowIndex, colIndex)">
        </span>
      </div>
    </div>
    <div class="flex flex-col gap-7">
      <div class="relative flex flex-col items-center justify-between border-black rounded-3xl border-4 min-h-min p-2" :style="{width: `${nextBlockSideLength}px`}">
        <h2 class="font-jersey text-5xl whitespace-nowrap">Holding:</h2>
        <div id="NextBlock" class="flex flex-col justify-center items-center">
          <div v-for="(row, holdBlockRowIndex) in nextBlockDisplay" :key="`hold_row_${holdBlockRowIndex}`" class="flex flex-row justify-center gap-0" :style="{width: `${cellSideLength * nextBlockDisplay.length}px`, height: `${cellSideLength}px`}">
            <span v-for="(cellVal, holdBlockColIndex) in row" :key="`hold_cell_${holdBlockRowIndex * nextBlockDisplay.length + holdBlockColIndex}`" class="" :style="getHoldBlockCellStyle(holdBlockRowIndex, holdBlockColIndex)">
            </span>
          </div>
        </div>
      </div>
    </div>
  </div>
  
</template>

<style scoped>



</style>
