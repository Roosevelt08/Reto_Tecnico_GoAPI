<template>
  <div v-if="statistics">
    <h2>Estadísticas</h2>
    <ul>
      <li>Valor Máximo: {{ statistics.maxValue }}</li>
      <li>Valor Mínimo: {{ statistics.minValue }}</li>
      <li>Promedio: {{ statistics.average }}</li>
      <li>Suma Total: {{ statistics.totalSum }}</li>
      <li>Matriz Diagonal: {{ statistics.isDiagonal ? 'Sí' : 'No' }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    statistics: Object
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
