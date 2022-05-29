package main

import (
	"fmt"
	"gonum.org/v1/gonum/mat"
	"inteligencia-artificial-siglo-21/models"
	"inteligencia-artificial-siglo-21/utils"
	"log"
)

func main() {
	rows, columns := 10, 10
	centeredMat, err := utils.LoadFileAsMatrix("data/centered.txt", rows, columns)
	if err != nil {
		log.Fatalf("Error loading input file. %v", err)
		return
	}
	fmt.Println("First pattern loaded")
	utils.PrettyPrintMatrixImage(centeredMat)

	uncenteredMat, err := utils.LoadFileAsMatrix("data/uncentered.txt", rows, columns)
	if err != nil {
		log.Fatalf("Error loading input file. %v", err)
		return
	}
	fmt.Println("Second pattern loaded")
	utils.PrettyPrintMatrixImage(uncenteredMat)

	fmt.Println("Train Hopfield model with Hebb method")
	patterns := []*mat.Dense{centeredMat, uncenteredMat}
	model := models.NewHopfieldModel()
	err = model.TrainWithHebbMethod(patterns)
	if err != nil {
		log.Fatalf("Error training the model using Hebb Method. %v", err)
		return
	}

	weightX, weightY := model.WeightDims()
	fmt.Printf("Size of parameters %d x %d\n", weightX, weightY)

	fmt.Println("First pattern with 20% noise probability")
	centeredWithNoise := mat.DenseCopyOf(centeredMat)
	utils.AddNoiseToMatrix(centeredWithNoise, 0.1)
	utils.PrettyPrintMatrixImage(centeredWithNoise)
	fmt.Println("Result of passing to Hopfield Model")
	result := model.Pass(centeredWithNoise)
	utils.PrettyPrintMatrixImage(result)

	fmt.Println("Second pattern with 20% noise probability")
	uncenteredWithNoise := mat.DenseCopyOf(uncenteredMat)
	utils.AddNoiseToMatrix(uncenteredWithNoise, 0.1)
	utils.PrettyPrintMatrixImage(uncenteredWithNoise)
	fmt.Println("Result of passing to Hopfield Model")
	result = model.Pass(uncenteredWithNoise)
	utils.PrettyPrintMatrixImage(result)
}
