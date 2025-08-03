package main

import (
	"fmt"
	"math"
	"github.com/snugml/go"  
)

func FloatMatrixToIntSlice(matrix [][]float64) []int {
	ints := make([]int, len(matrix))
	for i, row := range matrix {
		if len(row) > 0 {
			ints[i] = int(math.Round(row[0])) // o simplemente int(row[0]) para truncar
		} else {
			ints[i] = 0 // o algún valor por defecto
		}
	}
	return ints
}

func IntSliceToStringSlice(ints []int) []string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return strs
}


func main() {
	// xor sample
	X := [][]float64{
		{0, 0},
		{0, 1},
		{1, 0},
		{1, 1},
	}
	y := [][]float64{{0}, {1}, {1}, {0}}
	
	// Instancia del modelo
	model := ml.NewMLPClassifier(2, 4, 5, 0.1)

	// Entrenamiento del modelo
	model.Fit(X, y, 10000)

	// Predicciones del modelo
	var yPredict []int

	for _, f := range X {
		output := model.Predict(f)
		if len(output) == 1 {
			// Binaria: redondear
			if output[0] >= 0.5 {
				yPredict = append(yPredict, 1)
			} else {
				yPredict = append(yPredict, 0)
			}
		} else {
			// Multiclase: índice de la mayor probabilidad
			maxIdx := 0
			for i := 1; i < len(output); i++ {
				if output[i] > output[maxIdx] {
					maxIdx = i
				}
			}
			yPredict = append(yPredict, maxIdx)
		}
	}


	accuracy := ml.AccuracyScore(
		IntSliceToStringSlice(FloatMatrixToIntSlice(y)),
		IntSliceToStringSlice(yPredict),
	)

	fmt.Println("\nPredicted labels:")
	fmt.Println(yPredict)

	fmt.Printf("\nAccuracy score: %.4f\n", accuracy)


}