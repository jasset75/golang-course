package utils

import (
	"log"

	"gioui.org/app"
	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
	"gonum.org/v1/plot/plotutil"
)

func Draw() {
	go func() {
		// Create gioui window
		w := new(app.Window)
		if err := draw(w); err != nil {
			log.Fatal(err)
		}
	}()
	app.Main()
}

func draw(w *app.Window) error {
	// Tu número complejo
	z := complex(2, 3)

	// Create plot
	p := plot.New()
	p.Title.Text = "Complex Number"
	p.X.Label.Text = "Real Part"
	p.Y.Label.Text = "Imaginary Part"

	// Crear coordenate point
	pts := make(plotter.XYs, 1)
	pts[0].X = real(z)
	pts[0].Y = imag(z)

	// Add points to plot
	err := plotutil.AddScatters(p, "Original", pts)
	if err != nil {
		return err
	}

	// Show graphic
	return p.Draw(draw2DContext(w))
}

func draw2DContext(w *app.Window) *plot.Context {
	// Se implementa para mostrar el gráfico en la ventana.
	// Gioui renderiza gráficos en una ventana utilizando su propio canvas.
	return nil
}
