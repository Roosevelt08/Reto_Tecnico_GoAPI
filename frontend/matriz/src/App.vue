<template>
  <div id="app">
    <header>
      <h1>Reto Técnico - Matriz y Estadísticas</h1>
    </header>
    <MatrixInput @matrix-rotated="updateMatrix" />
    <MatrixDisplay :originalMatrix="originalMatrix" :rotatedMatrix="rotatedMatrix" />
    <StatisticsDisplay :statistics="statistics" />
  </div>
</template>

<script>
import MatrixInput from './components/MatrixInput.vue'
import MatrixDisplay from './components/MatrixDisplay.vue'
import StatisticsDisplay from './components/StatisticsDisplay.vue'

export default {
  components: {
    MatrixInput,
    MatrixDisplay,
    StatisticsDisplay
  },
  data() {
    return {
      originalMatrix: [],
      rotatedMatrix: [],
      statistics: {}
    };
  },
  methods: {
    updateMatrix(rotatedMatrix, statistics) {
      this.rotatedMatrix = rotatedMatrix;
      this.statistics = statistics;
    }
  }
}
</script>

<style>
body {
  font-family: Arial, sans-serif;
}

header {
  background-color: #333;
  color: white;
  padding: 1rem;
  text-align: center;
}

h2 {
  margin-top: 1rem;
}

textarea {
  width: 100%;
}
</style>