package scenes

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
