<script setup>
import {onMounted, ref} from 'vue';

const props = defineProps({
  Board: {
    type: Array, //[20][10]
    Required: true,
  }, 
  UIHeight: {
    type: Number,
    Required: true
  }

})

// board
// const boardRef = useTemplateRef('TetrisBoardRef')
let boardHeight = ref(0)
let cellSideLength = ref(0)
let cellPadding = ref(0)

// next block preview
let nextBlockSideLength = ref(0)



const resizeBoard = () => { 
  // set height to a factor of BoardHeight
  let heightFactor = 0.8
  boardHeight.value = props.UIHeight * heightFactor
  cellSideLength.value = boardHeight.value * 0.05
  cellPadding.value = cellSideLength.value * 0.1
  
  nextBlockSideLength.value = boardHeight.value * 0.3
}

const getCellStyle = (rowIndex, colIndex) => {
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
      backgroundImage: `${color} linear-gradient(145deg, #4c9ed9, #2a6f91)`,

      // Theme 0 ${color} linear-gradient(145deg, #4c9ed9, #2a6f91)
    };
   
}

const getRowStyle = (rowIndex) => {
  return {
      width: `${cellSideLength.value * 10}px`,
      height: `${cellSideLength.value}px`,
      top: `${cellSideLength.value * rowIndex}px`,
      left: `0px`,
    };
}




const resizeCallback = () => {
  resizeBoard()
}

// watch(props.UIHeight, ()=>{resizeBoard()})

onMounted(() => {
  resizeBoard()
  window.addEventListener("resize", resizeCallback)
})

</script>

<template>
  <div id="TetrisUI" class="flex flex-row justify-center gap-3 items-start min-h-min min-w-min">
    <div id="NextBlock" class="flex flex-col justify-between border-black rounded-3xl border-4" :style="{height: `${nextBlockSideLength}px`, width: `${nextBlockSideLength}px`}">
      <h2>Next Block</h2>
    </div>
    <div id="board" ref="TetrisBoardRef" class="flex gap-0 flex-col relative aspect-[1/2]" :style="{height: `${boardHeight}px`}">
      <div v-for="(row, rowIndex) in Board" :key="`row_${rowIndex}`" 
      class="absolute"
      :style="getRowStyle(rowIndex)">
        <span v-for="(cellVal, colIndex) in row" :key="`cell_${rowIndex * 10 + colIndex}`" 
        class="absolute min-w-2 min-h-2" 
        :style = "getCellStyle(rowIndex, colIndex)">
        </span>
      </div>
    </div>
  </div>
  
</template>

<style scoped>
</style>
