package router

import (
	"fmt"
	"image"
	"image/color"
	"image/png"
	"net/http"
	"strconv"

	"github.com/llgcode/draw2d"
	"github.com/llgcode/draw2d/draw2dimg"
)

func BreakventPoint(w http.ResponseWriter, r *http.Request) {

	xPoint, _ := strconv.Atoi(r.FormValue("x"))
	yPoint, _ := strconv.Atoi(r.FormValue("y"))
	maxNum, _ := strconv.Atoi(r.FormValue("max"))
	xstep, _ := strconv.Atoi(r.FormValue("xstep"))
	ystep, _ := strconv.Atoi(r.FormValue("ystep"))
	yTag := r.FormValue("ytag")
	if yTag == "" {
		yTag = "Y Axis"
	}
	xTag := r.FormValue("xtag")
	if xTag == "" {
		xTag = "X Axis"
	}

	//-------------------- testing -------------------
	numFC, _ := strconv.Atoi(r.FormValue("fc"))
	hideTag := r.FormValue("hide")
	if hideTag == "" {
		hideTag = "false"
	}
	//-------------------- testing -------------------

	mult := 40
	maxPoint := mult
	pad := 40

	for i := 0; i <= maxNum; i++ {
		maxPoint += mult
	}

	dest := image.NewRGBA(image.Rect(0, 0, maxPoint+pad*2, maxPoint+pad*2))
	gc := draw2dimg.NewGraphicContext(dest)
	gcLine := draw2dimg.NewGraphicContext(dest)
	text := draw2dimg.NewGraphicContext(dest)

	gc.SetStrokeColor(color.Black)
	gc.SetLineWidth(5)

	gcLine.SetStrokeColor(color.RGBA{0, 0, 255, 255})
	gcLine.SetLineDash([]float64{5, 5, 5, 5}, 0)
	gcLine.SetLineCap(draw2d.ButtCap)
	gcLine.SetLineJoin(draw2d.RoundJoin)
	gcLine.SetLineWidth(3)

	draw2d.SetFontFolder("font")
	text.SetFontData(draw2d.FontData{Name: "luxi", Family: draw2d.FontFamilyMono})
	text.SetFillColor(image.Black)
	text.SetFontSize(12)

	gc.BeginPath()
	gcLine.BeginPath()

	// straigh line from top to bottom
	gc.MoveTo(float64(pad), float64(pad))
	gc.LineTo(float64(pad), float64(maxPoint))
	if hideTag == "false" {
		text.FillStringAt(yTag, 10, 10)
	}

	// straight line from bottom to right
	gc.MoveTo(float64(pad), float64(maxPoint))
	gc.LineTo(float64(maxPoint), float64(maxPoint))
	if hideTag == "false" {
		text.FillStringAt(xTag, float64(maxPoint), float64(maxPoint))
	}

	lineY := float64(pad)
	line2Y := float64(pad)
	line2YFC := float64(pad)

	//-------------------- testing -------------------
	lineFC := 0.0
	//-------------------- testing -------------------

	numY := 0
	for i := maxPoint; i > pad; i -= mult {
		if hideTag == "false" {
			text.FillStringAt(fmt.Sprintf("%d - ", numY), 0, float64(i))
		}
		if inBetween(yPoint, numY-ystep, numY) {
			lineY = float64(i) - 5
			if hideTag == "true" {
				text.FillStringAt(fmt.Sprintf("%d-", yPoint), 0, float64(i))
			}
		}

		if inBetween(yPoint*2, numY-ystep, numY) {
			line2Y = float64(i) - 5
		}

		//-------------------- testing -------------------

		if inBetween(numFC, numY-ystep, numY) {
			lineFC = float64(i)
			if numFC <= numY {
				lineFC += 5
			} else if numFC > numY-ystep {
				lineFC -= 5
			}

			if hideTag == "true" {
				text.FillStringAt(fmt.Sprintf("%d-", numFC), 0, float64(i))
			}
		}

		if inBetween(yPoint*2-numFC, numY-ystep, numY) {
			line2YFC = float64(i) - 5
		}

		//-------------------- testing -------------------

		numY += ystep
	}

	lineX := float64(pad)
	line2X := float64(pad)
	line2XFC := float64(pad)

	numX := 0
	for i := pad; i < maxPoint; i += mult {
		if hideTag == "false" {
			text.FillStringAt("| ", float64(i), float64(maxPoint+20))
			text.FillStringAt(fmt.Sprintf("%d", numX), float64(i), float64(maxPoint+20*2))
		}
		if xPoint == numX {
			lineX = float64(i) + 5
			if hideTag == "true" {
				text.FillStringAt("| ", float64(i), float64(maxPoint+20))
				text.FillStringAt(fmt.Sprintf("%d", xPoint), float64(i), float64(maxPoint+20*2))
			}
		}

		if xPoint*2 == numX {
			line2X = float64(i) + 5
		}

		if xPoint*2 >= numX {
			line2XFC = float64(i) + 5
		}

		numX += xstep
	}

	// make line from mid-left to mid-right
	gcLine.MoveTo(float64(pad), lineY)
	gcLine.LineTo(lineX, lineY)

	// make line from mid-bottom to mid-top
	// gcLine.MoveTo(lineX, float64(maxPoint))
	gcLine.LineTo(lineX, float64(maxPoint))

	//-------------------- testing -------------------

	// garis laba

	gcLineGreen := draw2dimg.NewGraphicContext(dest)
	gcLineGreen.SetStrokeColor(color.RGBA{0, 128, 0, 255})
	gcLineGreen.SetLineWidth(3)
	gcLineGreen.MoveTo(float64(pad), float64(maxPoint))
	gcLineGreen.LineTo(line2X, line2Y)
	gcLineGreen.FillStroke()

	// garis total cost
	gcLineRed := draw2dimg.NewGraphicContext(dest)
	gcLineRed.SetStrokeColor(color.RGBA{255, 0, 0, 255})
	gcLineRed.SetLineWidth(3)
	gcLineRed.MoveTo(float64(pad), lineFC)
	gcLineRed.LineTo(line2XFC, line2YFC)
	gcLineRed.FillStroke()

	// garis fix cost
	gcLinePurple := draw2dimg.NewGraphicContext(dest)
	gcLinePurple.SetStrokeColor(color.RGBA{128, 0, 128, 255})
	gcLinePurple.SetLineWidth(3)
	gcLinePurple.MoveTo(float64(pad), lineFC)
	gcLinePurple.LineTo(float64(maxPoint), lineFC)
	gcLinePurple.Close()
	gcLinePurple.FillStroke()

	//-------------------- testing -------------------

	gc.Close()
	text.Close()

	gcLine.FillStroke()
	gc.FillStroke()

	err := png.Encode(w, dest)
	if err != nil {
		fmt.Fprintln(w, err.Error())
	}
}

func inBetween(i, min, max int) bool {
	return (i > min) && (i <= max)
}
