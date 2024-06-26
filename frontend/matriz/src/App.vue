<template>
  <div id="app">
    <MatrixDisplay />
  </div>
</template>

<script>
import MatrixDisplay from "./components/MatrixDisplay.vue";

export default {
  name: "App",
  components: {
    MatrixDisplay,
  },
};
</script>
