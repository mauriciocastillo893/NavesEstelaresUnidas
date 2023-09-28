package scenes

import (
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"modules/models"
	"os"
	"time"

	"fyne.io/fyne/v2"
	"fyne.io/fyne/v2/canvas"
	"fyne.io/fyne/v2/container"
	"fyne.io/fyne/v2/layout"
	"fyne.io/fyne/v2/widget"
)

type Space struct {
	canvasWidth  float32
	canvasHeight float32
	fps          int
	then         int64
	margin       int
}

func NewSpace(canvasWidth float32, canvasHeight float32, fps int, then int64, margin int) *Space {
	return &Space{canvasWidth: canvasWidth, canvasHeight: canvasHeight, fps: fps, then: then, margin: margin}
}

func ActualizarEscena(sprite *image.RGBA, playerSprites image.Image, naveModelo *models.Nave, spriteSize image.Point, scene *Space, mainContainer *fyne.Container, windowWidth float32, windowHeight float32) {
	fpsInterval := int64(1000 / scene.Fps())
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
			// naveImg = canvas.NewRasterFromImage(sprite)

			if naveModelo.XMov() != 0 || naveModelo.YMov() != 0 {
				naveModelo.SetX(naveModelo.X() + naveModelo.XMov())
				naveModelo.SetY(naveModelo.Y() + naveModelo.YMov())
				naveModelo.SetFrameX((naveModelo.FrameX() + 1) % naveModelo.CyclesX())
				naveModelo.SetXMov(0)
				naveModelo.SetYMov(0)
			} else {
				naveModelo.SetFrameX(0)
			}
			mainContainer.Refresh()
		}
	}
}

func CreateScene(myApp fyne.App, w fyne.Window) {
	now := time.Now().UnixMilli()
	scene := *NewSpace(564,
		314,
		60,
		now,
		10)

	// fmt.Println(fpsInterval)
	naveModelo := models.NewCharacter(200, 332, 40, 72, 0, 0, 4, 3, 0, 1, 2, 14, 0, 0)

	background, playerSprites, nil := UploadImage()
	fmt.Println(nil)

	imgBackground := canvas.NewImageFromImage(background)
	imgBackground.FillMode = canvas.ImageFillOriginal

	sprite := image.NewRGBA(background.Bounds())

	naveImg := canvas.NewRasterFromImage(sprite)
	spriteSize := image.Pt(naveModelo.Width(), naveModelo.Height())

	// Crear un widget de etiqueta para mostrar la hora actual
	horaLabel := widget.NewLabel("Hora actual: ")

	go GetHour(horaLabel)

	// Variable para el contador de gasolina
	contadorGasolina := 10000
	contadorLabel := widget.NewLabel(fmt.Sprintf("Cargador de Litio: %d", contadorGasolina))

	// Crear un botón para salir del juego
	salirButton := widget.NewButton("Salir del juego", func() {
		myApp.Quit()
	})

	// Crear un contenedor horizontal para la hora, el botón de salir y el contador
	horaContainer := container.NewHBox(
		layout.NewSpacer(),
		horaLabel,
		layout.NewSpacer(),
		salirButton,
		layout.NewSpacer(),
		contadorLabel,
		layout.NewSpacer(),
	)

	// Crear un contenedor principal para el fondo y la hora
	emptyObject := canvas.NewText("", color.Black)
	emptyObject.Hide()

	mainContainer := container.New(layout.NewBorderLayout(horaContainer, emptyObject, emptyObject, emptyObject), imgBackground, horaContainer, naveImg)

	w.SetContent(mainContainer)

	// Ajusta el tamaño de la ventana según las dimensiones del fondo
	windowWidth := float32(w.Canvas().Size().Width)
	windowHeight := float32(w.Canvas().Size().Height)

	w.Resize(fyne.NewSize(windowWidth, windowHeight))

	// Canal para comunicarse con la goroutine de recarga de litio
	recargaLitio := make(chan bool)
	stopRecargar := make(chan bool)

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

		contadorGasolina--
		contadorLabel.SetText(fmt.Sprintf("Cargador de Litio: %d", contadorGasolina))

		// Informar a la goroutine de recarga de litio que una tecla se presionó
		recargaLitio <- true

		// Reiniciar el temporizador de recarga de litio
		stopRecargar <- true
	})

	// Goroutine para recargar el litio
	go models.RecargarLitio(contadorLabel, recargaLitio, stopRecargar, &contadorGasolina)
	// Goroutine para actualizar la escena
	go ActualizarEscena(sprite, playerSprites, naveModelo, spriteSize, &scene, mainContainer, windowWidth, windowHeight)

	// Coloca la nave en el centro de la ventana
	naveModelo.SetX(int(windowWidth) / 2)
	naveModelo.SetY(int(windowHeight) / 2)
}

// Goroutine para actualizar la hora
func GetHour(horaLabel *widget.Label) {
	for {
		horaLabel.SetText("Hora actual: " + time.Now().Format("15:04:05"))
		time.Sleep(time.Second) // Actualiza cada segundo
	}
}

func Load(filePath string) (image.Image, error) {
	imgFile, err := os.Open(filePath)
	if err != nil {
		return nil, err
	}
	defer imgFile.Close()
	imgData, _, err := image.Decode(imgFile)
	if err != nil {
		return nil, err
	}
	return imgData, nil
}

func UploadImage() (image.Image, image.Image, error) {
	background, err := Load("./assets/map/USS Discovery Estacion3-4.png")
	if err != nil {
		return nil, nil, fmt.Errorf("error al cargar el fondo: %v", err)
	}
	playerSprites, err := Load("./assets/sprites/SpriteNave2.png")
	if err != nil {
		return nil, nil, fmt.Errorf("error al cargar los sprites del jugador: %v", err)
	}
	return background, playerSprites, nil
}




func (s *Space) CanvasWidth() float32 {
	return s.canvasWidth
}

func (s *Space) SetCanvasWidth(canvasWidth float32) {
	s.canvasWidth = canvasWidth
}

func (s *Space) CanvasHeight() float32 {
	return s.canvasHeight
}

func (s *Space) SetCanvasHeight(canvasHeight float32) {
	s.canvasHeight = canvasHeight
}

func (s *Space) Fps() int {
	return s.fps
}

func (s *Space) SetFps(fps int) {
	s.fps = fps
}

func (s *Space) Then() int64 {
	return s.then
}

func (s *Space) SetThen(then int64) {
	s.then = then
}

func (s *Space) Margin() int {
	return s.margin
}

func (s *Space) SetMargin(margin int) {
	s.margin = margin
}
