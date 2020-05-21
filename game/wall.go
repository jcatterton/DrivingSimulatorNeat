package game

import (
	"github.com/faiface/pixel"
	"github.com/faiface/pixel/imdraw"
	"golang.org/x/image/colornames"
)

type Wall struct {
	x1 float64
	y1 float64
	x2 float64
	y2 float64
}

func (w *Wall) Draw(t pixel.Target, imd imdraw.IMDraw) {
	imd.Color = colornames.Black
	imd.Push(pixel.V(w.x1, w.y1))
	imd.Push(pixel.V(w.x2, w.y2))
	imd.Line(2)
}

func (w *Wall) Line() pixel.Line {
	return pixel.L(pixel.V(w.x1, w.y1), pixel.V(w.x2, w.y2))
}
