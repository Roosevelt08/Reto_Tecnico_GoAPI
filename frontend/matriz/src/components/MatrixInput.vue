<template>
  <div>
    <h2>Ingresar Matriz</h2>
    <textarea v-model="matrixInput" rows="10" cols="50"></textarea>
    <button @click="rotateMatrix">Rotar Matriz</button>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  data() {
    return {
      matrixInput: '{"matrix": [[1, 2], [3, 4]]}',
      rotatedMatrix: null
    };
  },
  methods: {
    async rotateMatrix() {
      try {
        const response = await axios.post('http://localhost:8080/rotate', JSON.parse(this.matrixInput));
        this.rotatedMatrix = response.data;
        console.log(this.rotatedMatrix);
      } catch (error) {
        console.error("Error rotando la matriz:", error);
      }
    }
  }
};
</script>