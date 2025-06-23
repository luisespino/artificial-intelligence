package main

import (
	"fmt"

	"github.com/snugml/go"
)

func IntSliceToStringSlice(ints []int) []string {
	strs := make([]string, len(ints))
	for i, v := range ints {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return strs
}

func main() {
	outlook := []string{"sunny", "sunny", "overcast", "rain", "rain", "rain", "overcast",
		"sunny", "sunny", "rain", "sunny", "overcast", "overcast", "rain"}
	temperature := []string{"hot", "hot", "hot", "mild", "cool", "cool", "cool",
		"mild", "cool", "mild", "mild", "mild", "hot", "mild"}
	humidity := []string{"high", "high", "high", "high", "normal", "normal", "normal",
		"high", "normal", "normal", "normal", "high", "normal", "high"}
	windy := []string{"false", "true", "false", "false", "false", "true", "true",
		"false", "false", "false", "true", "true", "false", "true"}
	label := []string{"N", "N", "P", "P", "P", "N", "P", "N", "P", "P", "P", "P", "P", "N"}

	// Crear encoder
	encoder := ml.LabelEncoder{}

	encOut, err := encoder.FitTransform(outlook)
	if err != nil {
		panic(err)
	}
	encTem, err := encoder.FitTransform(temperature)
	if err != nil {
		panic(err)
	}
	encHum, err := encoder.FitTransform(humidity)
	if err != nil {
		panic(err)
	}
	encWin, err := encoder.FitTransform(windy)
	if err != nil {
		panic(err)
	}
	encLab, err := encoder.FitTransform(label)
	if err != nil {
		panic(err)
	}

	// Combinar características en [][]int
	features := make([][]int, len(outlook))
	for i := range outlook {
		features[i] = []int{encOut[i], encTem[i], encHum[i], encWin[i]}
	}

	// Instanciar DecisionTreeClassifier con MaxDepth=5
	model := ml.DecisionTreeClassifier{MaxDepth: 5}

	err = model.Fit(features, encLab)
	if err != nil {
		panic(err)
	}

	encYPredict, err := model.Predict(features)
	if err != nil {
		panic(err)
	}

	yPredict, err := encoder.InverseTransform(encYPredict)
	if err != nil {
		panic(err)
	}

	accuracy := ml.AccuracyScore(
		IntSliceToStringSlice(encLab),
		IntSliceToStringSlice(encYPredict),
	)

	fmt.Println("Encoded features:")
	for _, f := range features {
		fmt.Println(f)
	}

	fmt.Println("\nPredicted labels:")
	fmt.Println(yPredict)

	fmt.Printf("\nAccuracy score: %.4f\n", accuracy)

	fmt.Println("\nDescriptive tree:")
	fmt.Println(model.PrintTree())

	fmt.Println("\nGain track:")
	fmt.Println(model.Gain)
}
