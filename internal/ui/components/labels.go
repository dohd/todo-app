package components

import (
	"image/color"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CustomLabel struct {
	widget.BaseWidget
	Text        string
	TextColor   color.Color
	BgColor     color.Color
	Bold        bool
	Italic      bool
	Underline   bool
	textCanvas  *canvas.Text
	bgRectangle *canvas.Rectangle
}

func (c *CustomLabel) CreateRenderer() fyne.WidgetRenderer {
	c.bgRectangle = canvas.NewRectangle(c.BgColor)
	c.textCanvas = canvas.NewText(c.Text, c.TextColor)
	c.textCanvas.Text = c.Text
	c.textCanvas.Alignment = fyne.TextAlignCenter
	if c.Bold {
		c.textCanvas.TextStyle.Bold = true
	}
	if c.Italic {
		c.textCanvas.TextStyle.Italic = true
	}
	if c.Underline {
		c.textCanvas.TextStyle.Underline = true
	}
	return widget.NewSimpleRenderer(container.NewStack(c.bgRectangle, c.textCanvas))
}

func (c *CustomLabel) SetText(text string) *CustomLabel {
	c.Text = text
	c.Refresh()
	return c
}

func NewLabel(text string) *CustomLabel {
	lbl := &CustomLabel{
		Text:      text,
		TextColor: color.RGBA{R: 0, G: 150, B: 255, A: 255}, // blue
		BgColor:   color.Transparent,
	}
	lbl.ExtendBaseWidget(lbl)
	return lbl
}
