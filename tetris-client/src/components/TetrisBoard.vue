<script setup>
const props = defineProps({
  Board: {
    type: Array, //[20][10]
    Required: true,
    validator: (value) => {
      return Array.isArray(value) 
      && value.length === 20 
      && value.every(row => Array.isArray(row) 
          && row.length === 10 
          && row.every(cell => typeof cell === 'number')
        );
    }
  }
})
</script>

<template>
  <div id="board" class="">
    <div v-for="(row, row_index) in Board" :key="row_index" class="row">
      <span v-for="(cell, cell_index) in row" :key="cell_index" class="cell">
        {{ cell }}
      </span>
    </div>
  </div>
</template>

<style scoped>
#board {
  display: grid;
  width: 400px;
  height: 400px;
  background-color: white;
}

.cell {
  width: 10px;
  height: 10px;
  background-color: white;
}
</style>
