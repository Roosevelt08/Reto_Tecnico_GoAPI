<template>
  <div>
    <h1>Matriz Original y Rotada</h1>
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
        <li v-for="(row, index) in rotatedMatrix" :key="'rotated-'+index">
          {{ row }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      originalMatrix: [],
      rotatedMatrix: []
    };
  },
  methods: {
    async fetchMatrix() {
      try {
        const response = await axios.post('http://localhost:8080/rotate', { matrix: [[1, 2], [3, 4]] });
        this.originalMatrix = [[1, 2], [3, 4]];
        this.rotatedMatrix = response.data;
      } catch (error) {
        console.error('Error fetching matrix:', error);
      }
    }
  }
};
</script>
