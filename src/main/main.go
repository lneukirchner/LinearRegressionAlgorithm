package main

func matrixTransform2d(matrix [2][2]float64, vector [2]float64) [2]float64 {
	var result [2]float64                                               //define result vector as a list with two items
	result[0] = (vector[0] * matrix[0][0]) + (vector[1] * matrix[1][0]) //x component
	result[1] = (vector[0] * matrix[0][1]) + (vector[1] * matrix[1][1]) //y component
	return result
}

func determinant2d(matrix [2][2]float64) float64 {
	return matrix[0][0]*matrix[1][1] - matrix[0][1]*matrix[1][0] //ad-bc formula
}

func constantTimesMatrix2d(constant float64, matrix [2][2]float64) [2][2]float64 {
	matrix[0][0] *= constant
	matrix[0][1] *= constant
	matrix[1][0] *= constant
	matrix[1][1] *= constant
	return matrix
}

func inverseMatrix2d(matrix [2][2]float64) [2][2]float64 {
	var inverseDeterminant float64 = 1 / determinant2d(matrix)
	var newMatrix [2][2]float64
	newMatrix[0][0] = matrix[1][1]
	newMatrix[0][1] = -matrix[0][1]
	newMatrix[1][0] = -matrix[1][0]
	newMatrix[1][1] = matrix[0][0]
	newMatrix = constantTimesMatrix2d(inverseDeterminant, newMatrix)
	return newMatrix

}

func main() {

}
