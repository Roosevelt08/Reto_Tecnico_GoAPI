<template>
  <div id="app">
    <header>
      <h1>Reto Técnico - Matriz y Estadísticas</h1>
    </header>
    <MatrixInput @matrix-rotated="handleMatrixRotated"/>
    <MatrixDisplay :originalMatrix="originalMatrix" :rotatedMatrix="rotatedMatrix"/>
    <StatisticsDisplay :statistics="statistics"/>
  </div>
</template>

<script>
import MatrixInput from './components/MatrixInput.vue'
import MatrixDisplay from './components/MatrixDisplay.vue'
import StatisticsDisplay from './components/StatisticsDisplay.vue'

export default {
  name: 'App',
  components: {
    MatrixInput,
    MatrixDisplay,
    StatisticsDisplay
  },
  data() {
    return {
      originalMatrix: [],
      rotatedMatrix: [],
      statistics: null
    }
  },
  methods: {
    handleMatrixRotated(data) {
      this.originalMatrix = data.originalMatrix;
      this.rotatedMatrix = data.rotatedMatrix;
      this.statistics = data.statistics;
    }
  }
}
</script>

<style>
header {
  text-align: center;
  padding: 20px;
  background: #333;
  color: white;
}
</style>
