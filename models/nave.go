package models

import (
	"fmt"
	"time"

	"fyne.io/fyne/v2/widget"
)

type Nave struct {
	x       int
	y       int
	width   int
	height  int
	frameX  int
	frameY  int
	cyclesX int
	upY     int
	downY   int
	leftY   int
	rightY  int
	speed   int
	xMov    int
	yMov    int
}

func NewCharacter(x int, y int, width int, height int, frameX int, frameY int, cyclesX int, upY int, downY int, leftY int, rightY int, speed int, xMov int, yMov int) *Nave{
	return &Nave{x: x, y: y, width: width, height: height, frameX: frameX, frameY: frameY, cyclesX: cyclesX, upY: upY, downY: downY, leftY: leftY, rightY: rightY, speed: speed, xMov: xMov, yMov: yMov}
}

func RecargarLitio(contadorLabel *widget.Label, recargaLitio chan bool, stopRecargar chan bool, contadorGasolina *int) {
	for {
		select {
		case <-recargaLitio:
			// Se presionó una tecla, no recargues
		case <-time.After(10 * time.Second):
			// Pasaron 10 segundos sin presionar una tecla, recargar
			if *contadorGasolina < 10000 {
				*contadorGasolina++
				contadorLabel.SetText(fmt.Sprintf("Cargador de Litio: %d", *contadorGasolina))
			}
		case <-stopRecargar:
			// Reiniciar el temporizador de recarga de litio
			go func() {
				<-time.After(10 * time.Second)
				if *contadorGasolina < 10000 {
					*contadorGasolina++
					contadorLabel.SetText(fmt.Sprintf("Cargador de Litio: %d", *contadorGasolina))
				}
			}()
		}
	}
}

func (n *Nave) X() int {
	return n.x
}

func (n *Nave) SetX(x int) {
	n.x = x
}

func (n *Nave) Y() int {
	return n.y
}

func (n *Nave) SetY(y int) {
	n.y = y
}

func (n *Nave) Width() int {
	return n.width
}

func (n *Nave) SetWidth(width int) {
	n.width = width
}

func (n *Nave) Height() int {
	return n.height
}

func (n *Nave) SetHeight(height int) {
	n.height = height
}

func (n *Nave) FrameX() int {
	return n.frameX
}

func (n *Nave) SetFrameX(frameX int) {
	n.frameX = frameX
}

func (n *Nave) FrameY() int {
	return n.frameY
}

func (n *Nave) SetFrameY(frameY int) {
	n.frameY = frameY
}

func (n *Nave) CyclesX() int {
	return n.cyclesX
}

func (n *Nave) SetCyclesX(cyclesX int) {
	n.cyclesX = cyclesX
}

func (n *Nave) UpY() int {
	return n.upY
}

func (n *Nave) SetUpY(upY int) {
	n.upY = upY
}

func (n *Nave) DownY() int {
	return n.downY
}

func (n *Nave) SetDownY(downY int) {
	n.downY = downY
}

func (n *Nave) LeftY() int {
	return n.leftY
}

func (n *Nave) SetLeftY(leftY int) {
	n.leftY = leftY
}

func (n *Nave) RightY() int {
	return n.rightY
}

func (n *Nave) SetRightY(rightY int) {
	n.rightY = rightY
}

func (n *Nave) Speed() int {
	return n.speed
}

func (n *Nave) SetSpeed(speed int) {
	n.speed = speed
}

func (n *Nave) XMov() int {
	return n.xMov
}

func (n *Nave) SetXMov(xMov int) {
	n.xMov = xMov
}

func (n *Nave) YMov() int {
	return n.yMov
}

func (n *Nave) SetYMov(yMov int) {
	n.yMov = yMov
}