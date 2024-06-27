<template>
  <div>
    <h2>Matriz Original y Rotada</h2>
    <button @click="fetchMatrix">Cargar Matriz</button>
    <div v-if="originalMatrix.length > 0">
      <h3>Matriz Original</h3>
      <ul>
        <li v-for="(row, index) in originalMatrix" :key="'orig-' + index">
          {{ row }}
        </li>
      </ul>
    </div>
    <div v-if="rotatedMatrix.length > 0">
      <h3>Matriz Rotada</h3>
      <ul>
        <li v-for="(row, index) in rotatedMatrix" :key="'rotated-' + index">
          {{ row }}
        </li>
      </ul>
    </div>
  </div>
</template>

<script>
export default {
  props: {
    originalMatrix: {
      type: Array,
      required: true
    },
    rotatedMatrix: {
      type: Array,
      required: true
    }
  },
  methods: {
    fetchMatrix() {
      // Esta función puede ser implementada para obtener matrices desde una fuente si es necesario
    }
  }
}
</script>