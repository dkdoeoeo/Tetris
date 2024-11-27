<script setup>
import {useTemplateRef, onMounted, ref} from 'vue';

const props = defineProps({
  Board: {
    type: Array, //[20][10]
    Required: true,
  }, 

})

const boardRef = useTemplateRef('TetrisBoardRef')
let cellSideLength = ref(0)

const resizeBoard = () => { 
  if (!boardRef.value) return
  console.log("Resizing Tetris Board")
  // maintain ratio height : width = 2 : 1
  let height = boardRef.value.clientHeight
  let width = boardRef.value.clientHeight / 2
  boardRef.value.height = height;
  boardRef.value.width = width
  cellSideLength.value = height / 20
}

const getCellStyle = (rowIndex, colIndex) => {
    let cellVal = props.Board[rowIndex][colIndex]
    let color = "gray"

    switch (cellVal) {
      case 0:
        color = 'black';
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
    
    console.log()

    return {
      width: `${cellSideLength.value}px`,
      height: `${cellSideLength.value}px`,
      top: `0px`,
      left: `${cellSideLength.value * colIndex}px`,
      backgroundColor: `${color}`,
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

onMounted(() => {
  resizeBoard()
  window.addEventListener("resize", resizeCallback)

})
</script>

<template>
  <div id="board" ref="TetrisBoardRef" class="flex gap-0 flex-col relative h-3/5">
    <div v-for="(row, rowIndex) in Board" :key="`row_${rowIndex}`" 
    class="absolute"
    :style="getRowStyle(rowIndex)">
      <span v-for="(cellVal, colIndex) in row" :key="`cell_${rowIndex * 10 + colIndex}`" 
      class="absolute min-w-4 min-h-4"  
      :style = "getCellStyle(rowIndex, colIndex)">
      </span>
    </div>
  </div>
</template>

<style scoped>
</style>
