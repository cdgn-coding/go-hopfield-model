package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"gonum.org/v1/gonum/stat"
	"inteligencia-artificial-siglo-21/models"
	"inteligencia-artificial-siglo-21/utils"
)

func makeExperiment(model *models.HopfieldModel, pattern *mat.Dense, noiseProb float64) (int, bool) {
	centeredWithNoise := mat.DenseCopyOf(pattern)
	utils.AddNoiseToMatrix(centeredWithNoise, noiseProb)

	previousResult := model.Pass(centeredWithNoise)
	var passes int
	for passes = 1; ; passes++ {
		currentResult := model.Pass(previousResult)
		if mat.Equal(previousResult, currentResult) {
			break
		}
		if passes > 1000 {
			break
		}
	}

	correctResult := mat.Equal(previousResult, pattern)

	return passes, correctResult
}

func boolToFloat(v bool) float64 {
	if v {
		return 1.0
	}
	return 0.
}

func makeKExperiments(k int, model *models.HopfieldModel, pattern *mat.Dense, noiseProb float64) (float64, float64) {
	numberOfPasses := make([]float64, k)
	successRecord := make([]float64, k)
	for i := 0; i < k; i++ {
		passes, success := makeExperiment(model, pattern, noiseProb)
		numberOfPasses[i] = float64(passes)
		successRecord[i] = boolToFloat(success)
	}
	return stat.Mean(numberOfPasses, nil), stat.Mean(successRecord, nil)
}

func main() {
	rows, columns := 10, 10
	centeredMat, _ := utils.LoadFileAsMatrix("data/centered.txt", rows, columns)
	uncenteredMat, _ := utils.LoadFileAsMatrix("data/uncentered.txt", rows, columns)
	patterns := []*mat.Dense{centeredMat, uncenteredMat}
	model := models.NewHopfieldModel()
	_ = model.TrainWithHebbMethod(patterns)

	meanPasses, successRatio := makeKExperiments(100, model, centeredMat, 0.2)
	fmt.Printf("%.0f%% Noise. Stabilization with %.0f passes. Success ratio of %.2f\n", 20., meanPasses, successRatio)

	meanPasses, successRatio = makeKExperiments(100, model, centeredMat, 0.3)
	fmt.Printf("%.0f%% Noise. Stabilization with %.0f passes. Success ratio of %.2f\n", 30., meanPasses, successRatio)

	meanPasses, successRatio = makeKExperiments(100, model, centeredMat, 0.4)
	fmt.Printf("%.0f%% Noise. Stabilization with %.0f passes. Success ratio of %.2f\n", 40., meanPasses, successRatio)

	//meanPasses, successRatio = makeKExperiments(100, model, centeredMat, 0.2)
	//fmt.Printf("%.0f Noise. Stabilization with %.0f passes. Success ratio of %.2f\n", 20., meanPasses, successRatio)
}
