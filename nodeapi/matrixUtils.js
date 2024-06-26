
function rotateMatrix(matrix) {
    const rows = matrix.length;
    const cols = matrix[0].length;
    const rotated = [];

    for (let col = 0; col < cols; col++) {
        const newRow = [];
        for (let row = rows - 1; row >= 0; row--) {
            newRow.push(matrix[row][col]);
        }
        rotated.push(newRow);
    }

    return rotated;
}

module.exports = { rotateMatrix };
