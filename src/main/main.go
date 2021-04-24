package main

import (
	"fmt"
)

func determinant(matrix [][]float64) float64 { //Only works for 2x2 right now
	size := [2]int{len(matrix), len(matrix[0])}
	var det float64
	if size[0] != size[1] {
		return 0
	} else if size[0] > 2 {
		var j int
		for j = 0; j < size[0]; j++ {
			if j%2 == 0 { //If j is in an odd column, which is even in the index.
				if j == 0 {
					det += matrix[j][0] * determinant(matrix[1:][1:])
				} else if j == size[0]-1 {
					det += matrix[j][0] * determinant(matrix[:j][1:size[1]])
				} else {
					//var submatrixLeft [][]float64
					//var submatrixRight [][]float64
					//submatrixLeft = matrix[:j][1:size[1]]
					//submatrixRight = matrix[j+1:][1:size[1]]
					submatrix := append(matrix[:j][1:size[1]], matrix[j+1:][1:]...) //simplify code by replacing variables with their declarations
					det += matrix[j][0] * determinant(submatrix)
				}
				fmt.Println(det)
			} else if j%2 == 1 { //If we're on an even column
				if j == size[0]-1 {
					det -= matrix[j][0] * determinant(matrix[:j][1:size[1]])
				} else {
					submatrix := append(matrix[:j][1:size[1]], matrix[j+1:][1:]...)
					det -= matrix[j][0] * determinant(submatrix)
				}
				fmt.Println(det)

			}
		}
		return det
	} else if size[0] == 2 {
		return (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0])
	} else {
		return matrix[0][0]
	}
}

func dotProduct(v1, v2 []float64) float64 { //WORKS
	if len(v1) != len(v2) {
		return 0
	} else {
		var elements int = len(v1)
		var i int
		var total float64
		for i = 0; i < elements; i++ {
			total += v1[i] * v2[i]
		}
		return total
	}
}

func matrixTransform(matrix [][]float64, vector []float64) []float64 {
	if len(matrix) != len(vector) {
		return []float64{}
	} else {
		var i int
		var outputVector []float64
		for i = 0; i < len(matrix[0]); i++ { //Go through each row
			var j int
			var intermediateVector []float64
			for j = 0; j < len(matrix); j++ { //Go through each column to make a vector for the dot product.
				intermediateVector = append(intermediateVector, matrix[j][i])
			}
			outputVector = append(outputVector, dotProduct(vector, intermediateVector))
		}
		return outputVector
	}
}

func constantTimesMatrix(constant float64, matrix [][]float64) [][]float64 { //WORKS
	var columnCount int = len(matrix)
	var rowCount int = len(matrix[0])
	var i, j int
	for i = 0; i < columnCount; i++ {
		for j = 0; j < rowCount; j++ {
			matrix[i][j] *= constant
		}
	}
	return matrix
}

func inverseMatrix(matrix [][]float64) [][]float64 { //Only works for 2x2
	var inverseDeterminant float64 = 1 / determinant(matrix)
	var newMatrix [][]float64
	newMatrix[0][0] = matrix[1][1]
	newMatrix[0][1] = -matrix[0][1]
	newMatrix[1][0] = -matrix[1][0]
	newMatrix[1][1] = matrix[0][0]
	newMatrix = constantTimesMatrix(inverseDeterminant, newMatrix)
	return newMatrix
}

/*
func linearRegression2d(pointsList [][2]float64) [2]float64 {
	var xSquaredTotal, xTotal, xyTotal, yTotal, oneTotal float64
	var j int
	for j = 0; j < len(pointsList); j++ {
		xSquaredTotal += pointsList[j][0] * pointsList[j][0]
		xTotal += pointsList[j][0]
		xyTotal += pointsList[j][0] * pointsList[j][1]
		yTotal += pointsList[j][1]
		oneTotal += 1
	}
	firstMatrix := [][]float64{{xSquaredTotal, xTotal}, {xTotal, oneTotal}}
	equalVector := [2]float64{xyTotal, yTotal}
	var inverseFirstMatrix [][]float64 = inverseMatrix(firstMatrix)
	var mbTerms [2]float64
	mbTerms = matrixTransform(inverseFirstMatrix, equalVector)
	return mbTerms
}
*/
func main() {
	matrix := [][]float64{{5, 6, 4}, {8, 9, 7}, {-4, -5, -2}}
	vector := []float64{2, -3, 1}
	fmt.Println(matrixTransform(matrix, vector))
}
