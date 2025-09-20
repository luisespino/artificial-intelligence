package main

import (
	"fmt"
	"log"
	"github.com/snugml/go"
)

func main() {
	// Datos de entrenamiento
	X := [][]float64{
		{1, 2},
		{1, 3},
		{2, 2},
		{3, 3},
		{3, 4},
	}
	y := []interface{}{0, 0, 1, 1, 0}

	// Crear el modelo
	model := ml.GaussianNB{}

	// Ajustar modelo a los datos
	if err := model.Fit(X, y); err != nil {
		log.Fatal(err)
	}

	// Realizar predicciones
	predictions := model.Predict(X)

	// Mostrar resultados
	fmt.Println("Predicciones:", predictions)
}
