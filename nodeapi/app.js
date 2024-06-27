const express = require('express');
const app = express();
const bodyParser = require('body-parser');

app.use(bodyParser.json());

app.get('/', (req, res) => {
  res.send('Hola desde Node.js Roosevelt!');
});

app.post('/stats', (req, res) => {
  const rotatedMatrix = req.body.rotatedMatrix;
  
  if (!Array.isArray(rotatedMatrix) || rotatedMatrix.length === 0) {
    return res.status(400).json({ error: 'matriz no válido' });
  }

  const flatMatrix = rotatedMatrix.flat();
  const max = Math.max(...flatMatrix);
  const min = Math.min(...flatMatrix);
  const suma = flatMatrix.reduce((acc, num) => acc + num, 0);
  const promedio = suma / flatMatrix.length;
  const diagonal = rotatedMatrix.every((row, i) => row.every((value, j) => (i === j) || (value === 0)));

  const stats = {
    max,
    min,
    promedio,
    suma,
    diagonal
  };

  res.json(stats);
});

// Cambia la forma de iniciar el servidor para pruebas
if (!module.parent) {  // Solo si el módulo no es requerido por otro módulo
  const PORT = process.env.PORT || 3000;
  app.listen(PORT, () => {
    console.log(`Servidor corriendo en el puerto ${PORT}`);
  });
}

module.exports = app; // Exporta para pruebas y otros usos
