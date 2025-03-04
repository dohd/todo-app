package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/widget"
)

type TappableRectangle struct {
	widget.BaseWidget
	Rect     *canvas.Rectangle
	OnTapped func()
}

func NewTappableRectangle(color color.Color, onTap func()) *TappableRectangle {
	t := &TappableRectangle{
		Rect:     canvas.NewRectangle(color),
		OnTapped: onTap,
	}
	t.ExtendBaseWidget(t)
	return t
}

func (t *TappableRectangle) Tapped(_ *fyne.PointEvent) {
	if t.OnTapped != nil {
		t.OnTapped()
	}
}

func (t *TappableRectangle) CreateRenderer() fyne.WidgetRenderer {
	return widget.NewSimpleRenderer(t.Rect)
}
