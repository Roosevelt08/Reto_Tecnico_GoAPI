const request = require('supertest');
const app = require('./app');

// Define un grupo de pruebas usando Jest
describe('API tests', () => {
  // Primer test para la ruta GET en la raíz
  test('GET / devuelve el mensaje de saludo', async () => {
    const response = await request(app).get('/');
    expect(response.statusCode).toBe(200);
    expect(response.text).toContain('Hola desde Node.js Roosevelt!');
  });

  // Segundo test para la ruta POST que calcula estadísticas de matrices
  test('devuelve estadísticas para la matriz rotada', async () => {
    const rotatedMatrix = [
      [3, 1],
      [4, 2]
    ];

    const response = await request(app)
      .post('/stats')
      .send({ rotatedMatrix });

    expect(response.statusCode).toBe(200); // Verifica que el código de estado de la respuesta sea 200
    expect(response.body).toEqual({  // Verifica que el cuerpo de la respuesta coincida con las estadísticas esperadas
      max: 4,
      min: 1,
      average: 2.5,
      sum: 10,
      isDiagonal: false
    });
  });

  // Tercer test para verificar la funcionalidad de detección de matrices diagonales
  test('devuelve la verificación de matriz diagonal', async () => {
    const rotatedMatrix = [
      [1, 0],
      [0, 2]
    ];

    // Realiza una solicitud POST a la ruta '/stats', enviando la matriz diagonal en el cuerpo
    const response = await request(app)
      .post('/stats')
      .send({ rotatedMatrix });

    expect(response.statusCode).toBe(200);
    expect(response.body.isDiagonal).toBe(true);// Verifica que la propiedad 'isDiagonal' del cuerpo de la respuesta sea verdadera
  });
});

