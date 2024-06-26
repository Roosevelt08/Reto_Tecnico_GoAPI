const request = require('supertest');
const { rotateMatrix } = require('./matrixUtils'); 
const app = require('./app'); // Asegúrate de que esta es la ruta correcta hacia tu aplicación

describe('API tests', () => {
  test('GET / obtiene', async () => {
    const response = await request(app).get('/');
    expect(response.statusCode).toBe(200);
    expect(response.text).toContain('Hola desde Node.js Roosevelt!');
  });

  test('devolvera la rotación de matriz correcta', () => {
    const matrix = [[1, 2], [3, 4]];
    const rotatedMatrix = rotateMatrix(matrix); // Asegúrate de que rotateMatrix está correctamente definida e importada
    expect(rotatedMatrix).toEqual([[3, 1], [4, 2]]);
  });
});
