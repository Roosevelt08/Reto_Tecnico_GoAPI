
function rotateMatrix(matrix) {
    const rows = matrix.length; // Obtiene el número de filas de la matriz
    const cols = matrix[0].length; // Obtiene el número de columnas de la matriz (asume que todas las filas tienen la misma longitud)
    const rotated = []; // Crea un array vacío que almacenará la nueva matriz rotada
    
    // Bucle que recorre cada columna de la matriz original
    for (let col = 0; col < cols; col++) {
        const newRow = []; // Crea un nuevo array para la nueva fila en la matriz rotada
        for (let row = rows - 1; row >= 0; row--) {
            newRow.push(matrix[row][col]);
        }
        rotated.push(newRow);// Agrega la nueva fila al array de la matriz rotada
    }

    return rotated;
}

module.exports = { rotateMatrix };
