package utils

import (
	"bufio"
	"fmt"
	"gonum.org/v1/gonum/mat"
	"os"
)

func LoadFileAsMatrix(filename string, rows, columns int) (*mat.Dense, error) {
	data := make([]float64, rows*columns)

	file, err := os.Open(filename)

	if err != nil {
		return nil, fmt.Errorf("error opening file %s. %w", filename, err)
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for i := 0; scanner.Scan(); i++ {
		line := scanner.Text()
		for j, char := range line {
			num := float64(char - 48)
			if num == 0 {
				data[i*columns+j] = -1
			} else if num == 1 {
				data[i*columns+j] = 1
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading file %s. %w", filename, err)
	}

	matrix := mat.NewDense(rows, columns, data)

	return matrix, nil
}
