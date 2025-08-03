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

func createPlot(X, Y []float64, linePts plotter.XYs) {
	filename := "plot.png"
	p := plot.New()
	p.Title.Text = "Scatter & Prediction (Polynomial)"
	p.X.Label.Text = "X"
	p.Y.Label.Text = "Y"

	// Scatter con datos reales
	data := make(plotter.XYs, len(X))
	for i := range X {
		data[i].X, data[i].Y = X[i], Y[i]
	}

	scatter, err := plotter.NewScatter(data)
	if err != nil {
		log.Fatalf("Error creating scatter: %v", err)
	}

	line, err := plotter.NewLine(linePts)
	if err != nil {
		log.Fatalf("Error creating line: %v", err)
	}

	scatter.GlyphStyle.Color = color.RGBA{R: 255, A: 255} // rojo
	line.LineStyle.Color = color.RGBA{B: 255, A: 255}     // azul

	p.Add(scatter, line)
	p.Legend.Add("Real", scatter)
	p.Legend.Add("Predicted", line)

	if err := p.Save(6*vg.Inch, 4*vg.Inch, filename); err != nil {
		log.Fatalf("Error saving plot: %v", err)
	}
}

func main() {
	// Datos de ejemplo (X e Y)
	X := []float64{0, 1, 2, 3, 4, 5, 6, 7, 8, 9}
	y := []float64{1, 4, 1, 5, 3, 7, 2, 7, 4, 9}
	degree := 3

	// Instancia de LinearRegression
	model := ml.PolynomialRegression{Degree:degree}

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

	// generar curva
	const steps = 100
	minX, maxX := X[0], X[len(X)-1]
	step := (maxX - minX) / float64(steps-1)
	linePts := make(plotter.XYs, steps)

	for i := 0; i < steps; i++ {
		x := minX + float64(i)*step
		y := model.Predict([]float64{x})[0]
		linePts[i].X = x
		linePts[i].Y = y
	}

	createPlot(X, y, linePts)


}