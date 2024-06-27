<template>
  <div>
    <h2>Ingresar Matriz</h2>
    <textarea v-model="inputMatrix" rows="5" cols="30"></textarea>
    <br />
    <button @click="rotateMatrix">Rotar Matriz</button>
  </div>
</template>

<script>
export default {
  data() {
    return {
      inputMatrix: '{"matrix": [[1, 2], [3, 4]]}',
    };
  },
  methods: {
    async rotateMatrix() {
      try {
        const response = await fetch("http://localhost:8080/rotate", {
          method: "POST",
          headers: {
            "Content-Type": "application/json",
          },
          body: this.inputMatrix,
        });
        const result = await response.json();
        this.$emit("matrix-rotated", result.rotatedMatrix, result.statistics);
      } catch (error) {
        console.error("Error rotando la matriz:", error);
      }
    },
  },
};
</script>

<style scoped>
textarea {
  width: 100%;
}
</style>