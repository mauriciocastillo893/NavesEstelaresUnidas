package main

import (
	"modules/scenes"
	"fyne.io/fyne/v2/app"
)

func main() {
	myApp := app.New()
	w := myApp.NewWindow("Naves Estelares Unidas")
	w.SetFixedSize(true)

	// Cargar las imagenes de la escena y la nave
	scenes.UploadImage()
	// Crear y cargar la escena a partir de la app y ventana mandadas como parametro
	scenes.CreateScene(myApp, w)
	// Centrar la aplicación al centro
	w.CenterOnScreen()
	// Correrla (la aplicación)
	w.ShowAndRun()
}
