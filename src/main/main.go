package main

import (
	"fmt"
)

func determinant(matrix [][]float64) float64 { //Only works for 2x2 right now
	size := [2]int{len(matrix), len(matrix[0])}
	if size[0] != size[1] {
		return 0
	} else if size[0] > 2 {
		var det float64
		var j int
		for j = 0; j < size[0]; j++ {
			if j%2 == 0 {

			} else if j%2 == 1 {

			}
		}
	} else if size[0] == 2 {
		return (matrix[0][0] * matrix[1][1]) - (matrix[0][1] * matrix[1][0])
	} else {
		return matrix[0][0]
	}

}

func matrixTransform2d(matrix [2][2]float64, vector [2]float64) [2]float64 {
	var result [2]float64                                               //define result vector as a list with two items
	result[0] = (vector[0] * matrix[0][0]) + (vector[1] * matrix[1][0]) //x component
	result[1] = (vector[0] * matrix[0][1]) + (vector[1] * matrix[1][1]) //y component
	return result
}

func constantTimesMatrix2d(constant float64, matrix [2][2]float64) [2][2]float64 {
	matrix[0][0] *= constant
	matrix[0][1] *= constant
	matrix[1][0] *= constant
	matrix[1][1] *= constant
	return matrix
}

func inverseMatrix2d(matrix [][]float64) [2][2]float64 {
	var inverseDeterminant float64 = 1 / determinant(matrix)
	var newMatrix [2][2]float64
	newMatrix[0][0] = matrix[1][1]
	newMatrix[0][1] = -matrix[0][1]
	newMatrix[1][0] = -matrix[1][0]
	newMatrix[1][1] = matrix[0][0]
	newMatrix = constantTimesMatrix2d(inverseDeterminant, newMatrix)
	return newMatrix
}

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
	var inverseFirstMatrix [2][2]float64 = inverseMatrix2d(firstMatrix)
	var mbTerms [2]float64
	mbTerms = matrixTransform2d(inverseFirstMatrix, equalVector)
	return mbTerms
}

func main() {
	myMatrix := [][]float64{{1, 2}, {2, 1}}
	myMatrixDeterminant := determinant(myMatrix)
	fmt.Println(myMatrixDeterminant)
}
