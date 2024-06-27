<template>
  <div>
    <h2>Matriz Original y Rotada</h2>
    <button @click="fetchMatrix">Cargar Matriz</button>
    <div v-if="originalMatrix.length">
      <h2>Matriz Original</h2>
      <ul>
        <li v-for="(row, index) in originalMatrix" :key="'orig-'+index">
          {{ row }}
        </li>
      </ul>
    </div>
    <div v-if="rotatedMatrix.length">
      <h2>Matriz Rotada</h2>
      <ul>
        <li v-for="(row, index) in rotatedMatrix" :key="'rot-'+index">
          {{ row }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    originalMatrix: Array,
    rotatedMatrix: Array
  }
}
</script>

<style scoped>
ul {
  list-style: none;
  padding: 0;
}
li {
  margin-bottom: 5px;
}
</style>
