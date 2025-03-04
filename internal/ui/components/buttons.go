package components

import (
	"image/color"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/widget"
)

type CustomButton struct {
	widget.Button
	bgColor color.Color
	bgRect  *canvas.Rectangle
	label   *widget.Label
}

func (b *CustomButton) CreateRenderer() fyne.WidgetRenderer {
	b.bgRect = canvas.NewRectangle(b.bgColor)
	b.bgRect.CornerRadius = 6
	b.label = widget.NewLabel(b.Text)
	b.label.TextStyle.Bold = true
	b.label.Alignment = fyne.TextAlignCenter
	return widget.NewSimpleRenderer(container.NewStack(b.bgRect, container.NewCenter(b.label)))
}

func (b *CustomButton) Tapped(*fyne.PointEvent) {
	if b.OnTapped != nil {
		b.animateTap()
		b.OnTapped()
	}
}

func (b *CustomButton) animateTap() {
	originalSize := b.Size()
	smallerSize := fyne.NewSize(originalSize.Width*0.9, originalSize.Height*0.9)
	b.Resize(smallerSize)
	b.Refresh()
	// Restore after 100ms
	time.AfterFunc(100*time.Millisecond, func() {
		b.Resize(originalSize)
		b.Refresh()
	})
}

func NewPrimaryButton(text string, tapped func()) *CustomButton {
	btn := &CustomButton{
		bgColor: color.RGBA{R: 0, G: 150, B: 255, A: 255}, // blue
		label:   widget.NewLabel(text),
	}
	btn.ExtendBaseWidget(btn)
	btn.OnTapped = tapped
	btn.Text = text
	return btn
}

func NewDangerButton(text string, tapped func()) *CustomButton {
	btn := &CustomButton{
		bgColor: color.RGBA{R: 255, G: 0, B: 0, A: 255}, // red
		label:   widget.NewLabel(text),
	}
	btn.ExtendBaseWidget(btn)
	btn.OnTapped = tapped
	btn.Text = text
	return btn
}
