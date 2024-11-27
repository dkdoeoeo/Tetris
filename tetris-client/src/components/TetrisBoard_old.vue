<script setup>
import { onUpdated, ref, watch} from 'vue';

const props = defineProps({
  Board: {
    type: Array, //[20][10]
    Required: true,
  }
})

const getCellColor = (row_index, col_index) => {
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

</script>

<template>
  <div id="board">
    <div v-for="(row, row_index) in Board" :key="row_index" class="row">
      <span v-for="(cellVal, col_index) in row" :key="row_index * 10 + col_index" 
        class="cell"  :style = "{backgroundColor: getCellColor(row_index, col_index)}">
        {{getCellColor(cellVal)}}
      </span>
    </div>
  </div>
</template>

<style scoped>
#board {
  display: flex;
  flex-direction: column;
  justify-content: space-between;
  column-gap: 0;
  width: 400px;
  height: auto;
  border: 2px solid darkgray;
}
.row {
  display: inline-flex;
  flex-direction: row;
  justify-content: space-between;
  gap: 0;
}

.cell {
  margin: 0;
  padding: 0;
  width: 30px;
  height: 30px;
  background-color: white;
}
</style>
