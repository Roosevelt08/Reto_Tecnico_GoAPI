const request = require('supertest');
const app = require('./app');

describe('API tests', () => {
  test('GET / devuelve el mensaje de saludo', async () => {
    const response = await request(app).get('/');
    expect(response.statusCode).toBe(200);
    expect(response.text).toContain('Hola desde Node.js Roosevelt!');
  });

  test('devuelve estadísticas para la matriz rotada', async () => {
    const rotatedMatrix = [
      [3, 1],
      [4, 2]
    ];

    const response = await request(app)
      .post('/stats')
      .send({ rotatedMatrix });

    expect(response.statusCode).toBe(200);
    expect(response.body).toEqual({
      max: 4,
      min: 1,
      average: 2.5,
      sum: 10,
      isDiagonal: false
    });
  });

  test('devuelve la verificación de matriz diagonal', async () => {
    const rotatedMatrix = [
      [1, 0],
      [0, 2]
    ];

    const response = await request(app)
      .post('/stats')
      .send({ rotatedMatrix });

    expect(response.statusCode).toBe(200);
    expect(response.body.isDiagonal).toBe(true);
  });
});

