const express = require('express');
const app = express();

app.use(express.json()); 

app.post('/stats', (req, res) => {
  const { rotatedMatrix } = req.body;
  if (!rotatedMatrix || !Array.isArray(rotatedMatrix) || !rotatedMatrix.length) {
      return res.status(400).json({ error: 'Entrada inválida' });
  }
  const stats = calculateStats(rotatedMatrix);
  res.json(stats);
});

function calculateStats(matrix) {
  return {
      promedio: calculatePromedio(matrix),
      max: findMax(matrix),
      min: findMin(matrix)
  };
}

function calculatePromedio(matrix) {
  let total = 0;
  let count = 0;
  matrix.forEach(row => {
      row.forEach(value => {
          total += value;
          count++;
      });
  });
  return total / count;
}

function findMax(matrix) {
  let max = -Infinity;
  matrix.forEach(row => {
      row.forEach(value => {
          if (value > max) max = value;
      });
  });
  return max;
}

function findMin(matrix) {
  let min = Infinity;
  matrix.forEach(row => {
      row.forEach(value => {
          if (value < min) min = value;
      });
  });
  return min;
}

// Cambia la forma de iniciar el servidor para pruebas
if (!module.parent) {  // Solo si el módulo no es requerido por otro módulo
  const PORT = process.env.PORT || 3000;
  app.listen(PORT, () => {
    console.log(`Servidor corriendo en el puerto ${PORT}`);
  });
}

module.exports = app; // Exporta para pruebas y otros usos
