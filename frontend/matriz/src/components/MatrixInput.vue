<template>
  <div>
    <h2>Ingresar Matriz</h2>
    <textarea v-model="matrixInput" rows="5"></textarea>
    <button @click="rotateMatrix">Rotar Matriz</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      matrixInput: '{"matrix": [[1, 2], [3, 4]]}',
    };
  },
  methods: {
    async rotateMatrix() {
      try {
        const response = await fetch('http://localhost:8080/rotate', {
          method: 'POST',
          headers: {
            'Content-Type': 'application/json',
          },
          body: this.matrixInput,
        });

        const data = await response.json();

        if (response.ok) {
          this.$emit('matrix-rotated', data.rotatedMatrix, data.statistics);
        } else {
          console.error('Error rotando la matriz:', data.error);
        }
      } catch (error) {
        console.error('Error rotando la matriz:', error);
      }
    },
  },
};
</script>