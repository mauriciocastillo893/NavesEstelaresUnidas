package main

import (
	"modules/scenes"
	"modules/models"
	"fmt"
	"image"
	"image/draw"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/app"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
)

func load(filePath string) (image.Image, error) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		return nil, err // Devuelve el error si no puedes abrir el archivo
	}
	defer imgFile.Close()

	imgData, _, err := image.Decode(imgFile) // Cambia png.Decode a image.Decode
	if err != nil {
		return nil, err // Devuelve el error si no puedes decodificar la imagen
	}

	return imgData, nil // Devuelve la imagen y nil (sin error) si todo está bien
}

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Naves Estelares Unidas")
	w.SetFixedSize(true)

	background, err := load("./assets/map/USS Discovery Estacion3-4.png")
	if err != nil {
		fmt.Println("Error al cargar el fondo:", err)
		// Manejar el error de carga de fondo, por ejemplo, mostrar un mensaje de error
		return
	}
	
	playerSprites, err := load("./assets/sprites/SpriteNave2.png")
	if err != nil {
		fmt.Println("Error al cargar los sprites del jugador:", err)
		// Manejar el error de carga de sprites, por ejemplo, mostrar un mensaje de error
		return
	}

	now := time.Now().UnixMilli()
	scene := scenes.NewSpace(564,
		314,
		60,
		now,
		10)

	fpsInterval := int64(1000 / scene.Fps())
	fmt.Println(fpsInterval)
	naveModelo := models.NewCharacter(200, 332, 40, 72, 0, 0, 4, 3, 0, 1, 2, 14, 0, 0)

	imgBackground := canvas.NewImageFromImage(background)
	imgBackground.FillMode = canvas.ImageFillOriginal

	sprite := image.NewRGBA(background.Bounds())

	naveImg := canvas.NewRasterFromImage(sprite)
	spriteSize := image.Pt(naveModelo.Width(), naveModelo.Height())

	c := container.New(layout.NewStackLayout(), imgBackground, naveImg)
	w.SetContent(c)
	// Ajusta el tamaño de la ventana según las dimensiones del fondo
	windowWidth := float32(w.Canvas().Size().Width)
	windowHeight := float32(w.Canvas().Size().Height)

	w.Resize(fyne.NewSize(windowWidth, windowHeight))
	// WHAT KEY USER PRESSES
	w.Canvas().SetOnTypedKey(func(k *fyne.KeyEvent) {
		switch k.Name {
		case fyne.KeyDown:
			if naveModelo.Y() < int(windowHeight)-naveModelo.Height()-scene.Margin() {
				naveModelo.SetYMov(naveModelo.Speed())
			}
			naveModelo.SetFrameY(naveModelo.DownY())
		case fyne.KeyUp:
			if naveModelo.Y() > scene.Margin() {
				naveModelo.SetYMov(-naveModelo.Speed())
			}
			naveModelo.SetFrameY(naveModelo.UpY())
		case fyne.KeyLeft:
			if naveModelo.X() > scene.Margin() {
				naveModelo.SetXMov(-naveModelo.Speed())
			}
			naveModelo.SetFrameY(naveModelo.LeftY())
		case fyne.KeyRight:
			if naveModelo.X() < int(windowWidth)-naveModelo.Width()-scene.Margin() {
				naveModelo.SetXMov(naveModelo.Speed())
			}
			naveModelo.SetFrameY(naveModelo.RightY())
		}
	})

	
	
	go func() {
		for {
			time.Sleep(time.Millisecond)
			thisTime := time.Now().UnixMilli()
			elapsed := thisTime - scene.Then()

			if elapsed > fpsInterval {
				scene.SetThen(thisTime)
				spriteNave := image.Pt(naveModelo.Width()*naveModelo.FrameX(), naveModelo.Height()*naveModelo.FrameY())
				size := image.Rectangle{spriteNave, spriteNave.Add(spriteSize)}
				dataPoint := image.Pt(naveModelo.X(), naveModelo.Y())
				rectangule := image.Rectangle{dataPoint, dataPoint.Add(spriteSize)}

				draw.Draw(sprite, sprite.Bounds(), image.Transparent, image.Point{}, draw.Src)
				draw.Draw(sprite, rectangule, playerSprites, size.Min, draw.Src)
				naveImg = canvas.NewRasterFromImage(sprite)

				if naveModelo.XMov() != 0 || naveModelo.YMov() != 0 {
					naveModelo.SetX(naveModelo.X() + naveModelo.XMov())
					naveModelo.SetY(naveModelo.Y() + naveModelo.YMov())
					naveModelo.SetFrameX((naveModelo.FrameX() + 1) % naveModelo.CyclesX())
					naveModelo.SetXMov(0)
					naveModelo.SetYMov(0)
				} else {
					naveModelo.SetFrameX(0)
				}
				c.Refresh()
			}
		}
	}()

	// Coloca la nave en el centro de la ventana
	naveModelo.SetX(int(windowWidth) / 2)
	naveModelo.SetY(int(windowHeight) / 2)
	w.CenterOnScreen()
	w.ShowAndRun()
}
