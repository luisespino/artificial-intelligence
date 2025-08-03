package main

import (
	"fmt"
	"log"
	"image/color"
	"github.com/snugml/go"  // Importa el paquete ml donde está la lógica de LinearRegression

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/vg"
)

func createPlot(X, Y, Ypredict []float64) {
	filename := "plot.png"
	p := plot.New()
	p.Title.Text = "Scatter & Prediction"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	data := make(plotter.XYs, len(X))
	pred := make(plotter.XYs, len(X))
	for i := range X {
		data[i].X, data[i].Y = X[i], Y[i]
		pred[i].X, pred[i].Y = X[i], Ypredict[i]
	}

	scatter, err := plotter.NewScatter(data)
	if err != nil {
		log.Fatalf("Scatter create error: %v", err)
	}
	line, err := plotter.NewLine(pred)
	if err != nil {
		log.Fatalf("Line create error: %v", err)
	}

	scatter.GlyphStyle.Color = color.RGBA{R: 255, G: 0, B: 0, A: 255} // rojo
	line.LineStyle.Color = color.RGBA{B: 255, A: 255}                 // azul

	p.Add(scatter, line)
	p.Legend.Add("Train", scatter)
	p.Legend.Add("Predicted", line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		log.Fatalf("Saving file error: %v", err)
	}
}

func main() {
	// Datos de ejemplo (X e Y)
	X := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := []float64{1, 4, 1, 5, 3, 7, 2, 7, 4, 9}

	// Instancia de LinearRegression
	model := ml.LinearRegression{}

	// Entrenamiento del modelo
	model.Fit(X, y)

	// Predicciones del modelo
	yPredict := model.Predict(X)

	// Calcular el MSE y R^2
	mse := model.MSE(y, yPredict)
	r2 := model.R2(y, yPredict)

	// Imprimir los resultados
	fmt.Println("X:", X)
	fmt.Println("y:", y)
	fmt.Println("yPredict:", yPredict)
	fmt.Printf("MSE: %.4f\n", mse)
	fmt.Printf("R2: %.4f\n", r2)

	createPlot(X, y, yPredict)
}