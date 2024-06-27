<template>
  <div v-if="statistics">
    <h2>Estadísticas</h2>
    <ul>
      <li>Valor Máximo: {{ statistics.max }}</li>
      <li>Valor Mínimo: {{ statistics.min }}</li>
      <li>Promedio: {{ statistics.average }}</li>
      <li>Suma Total: {{ statistics.sum }}</li>
      <li>Matriz Diagonal: {{ statistics.diagonal ? 'Sí' : 'No' }}</li>
    </ul>
  </div>
</template>

<script>
export default {
  props: {
    statistics: {
      type: Object,
      required: true
    }
  }
}
</script>