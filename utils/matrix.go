package utils

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"math/rand"
)

func flipValue(v float64) float64 {
	return -v
}

func AddNoiseToMatrix(matrix *mat.Dense, prob float64) {
	matrix.Apply(func(i, j int, v float64) float64 {
		// Distribucion uniforme entre 0 y 1
		rng := rand.Float64()
		if prob < rng {
			return v
		}
		return flipValue(v)
	}, matrix)
}

func PrettyPrintMatrixImage(matrix *mat.Dense) {
	rows, columns := matrix.Dims()
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			el := matrix.At(i, j)
			if el == -1 {
				fmt.Printf("%d ", 0)
			} else if el == 1 {
				fmt.Printf("%d ", 1)
			}
		}
		fmt.Println()
	}
}

func PrettyPrintMatrix(matrix *mat.Dense) {
	rows, columns := matrix.Dims()
	for i := 0; i < rows; i++ {
		for j := 0; j < columns; j++ {
			el := matrix.At(i, j)
			fmt.Printf("%.2f ", el)
		}
		fmt.Println()
	}
}
