const express = require('express');
const app = express();
const bodyParser = require('body-parser'); // analiza el cuerpo de las peticiones HTTP

app.use(bodyParser.json());
// Definir una ruta GET en la raíz del servidor
app.get('/', (req, res) => {
  res.send('Hola desde Node.js Roosevelt!');
});

app.post('/stats', (req, res) => {
  const rotatedMatrix = req.body.rotatedMatrix;
  
  if (!Array.isArray(rotatedMatrix) || rotatedMatrix.length === 0) {
    return res.status(400).json({ error: 'matriz no válido' });
  }

  const flatMatrix = rotatedMatrix.flat();
  const max = Math.max(...flatMatrix); // Calcular el máximo valor en la matriz aplanada
  const min = Math.min(...flatMatrix); // Calcular el mínimo valor en la matriz aplanada
  const suma = flatMatrix.reduce((acc, num) => acc + num, 0);// Calcular la suma de todos los valores en la matriz aplanada
  const promedio = suma / flatMatrix.length; // Calcular el promedio de los valores en la matriz aplanada
  const diagonal = rotatedMatrix.every((row, i) => row.every((value, j) => (i === j) || (value === 0))); // Determinar si la matriz es diagonal

  // Crear un objeto con las estadísticas calculadas
  const stats = {
    max,
    min,
    promedio,
    suma,
    diagonal
  };
  // Devolver las estadísticas como una respuesta JSON
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
