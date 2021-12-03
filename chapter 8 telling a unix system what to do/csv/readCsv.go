package main

import (
	"encoding/csv"
	"fmt"
	"os"
	"strconv"

	"github.com/Arafatk/glot"
)

func main() {
	if len(os.Args) != 2 {
		fmt.Println("Need a file!")
	}
	//wether the file exists or nor
	_, err := os.Stat(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	f, err := os.Open(os.Args[1])
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()

	r := csv.NewReader(f)
	r.FieldsPerRecord = -1
	records, err := r.ReadAll()
	if err != nil {
		fmt.Println(err.Error())
		return
	}
	xp := []float64{}
	xy := []float64{}
	for _, coord := range records {
		x, _ := strconv.ParseFloat(coord[0], 64)
		y, _ := strconv.ParseFloat(coord[1], 64)
		xp = append(xp, x)
		xy = append(xy, y)
	}

	points := [][]float64{}
	points = append(points, xp)
	points = append(points, xy)

	dimensions, persist, debug := 2, true, false
	plot, _ := glot.NewPlot(dimensions, persist, debug)
	plot.SetTitle("CSV Data alv")
	plot.SetXLabel("X")
	plot.SetYLabel("Y")
	plot.AddPointGroup("Circle:", "circle", points)
	plot.SavePlot("output.jpg")
}
