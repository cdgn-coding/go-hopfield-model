package models

import (
	"errors"
	"gonum.org/v1/gonum/mat"
	"inteligencia-artificial-siglo-21/utils"
)

var NoPatternsGiven = errors.New("no patterns were given to the training method")

type HopfieldModel struct {
	weights *mat.Dense
}

func NewHopfieldModel() *HopfieldModel {
	return &HopfieldModel{}
}

// eye returns a new identity matrix of size n×n.
func eye(n int) *mat.Dense {
	d := make([]float64, n*n)
	for i := 0; i < n*n; i += n + 1 {
		d[i] = 1
	}
	return mat.NewDense(n, n, d)
}

func (h *HopfieldModel) PrintWeights() {
	utils.PrettyPrintMatrix(h.weights)
}

func (h *HopfieldModel) TrainWithHebbMethod(patterns []*mat.Dense) error {
	if len(patterns) == 0 {
		return NoPatternsGiven
	}

	rows, columns := patterns[0].Caps()
	weights := mat.NewDense(rows*columns, rows*columns, nil)

	for k := 0; k < len(patterns); k++ {
		weightsK := mat.NewDense(rows*columns, rows*columns, nil)
		patternK := mat.NewDense(1, rows*columns, patterns[k].RawMatrix().Data)
		patternKTransposed := mat.Transpose{Matrix: patternK}
		weightsK.Mul(patternKTransposed, patternK)
		weightsK.Sub(weightsK, eye(rows*columns))
		weights.Add(weights, weightsK)
	}

	h.weights = weights

	return nil
}

func (h *HopfieldModel) Pass(image *mat.Dense) *mat.Dense {
	rows, columns := image.Dims()
	pattern := mat.NewDense(1, rows*columns, image.RawMatrix().Data)
	result := mat.DenseCopyOf(pattern)
	result.Mul(result, h.weights)
	result.Apply(func(i, j int, v float64) float64 {
		if v >= 0 {
			return 1.0
		}
		return -1.0
	}, result)
	imageResult := mat.NewDense(rows, columns, result.RawMatrix().Data)
	return imageResult
}
