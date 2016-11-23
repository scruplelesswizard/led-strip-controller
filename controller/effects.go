package controller

// Off sets all LED's on the strip to off
func (s *Strip) Off() {
	s.SetColor(s.Color.Off())
}

// func (s *Strip) Fade(color RGB, duration time.Duration) {
//
// 	// calculate step duration and # of steps
// 	stepDuration := time.Duration(100) * time.Millisecond
// 	steps := duration / stepDuration
// 	maxDiff := s.Color.maxDifference(color)
//
// 	if steps > maxDiff {
// 		steps = maxDiff
// 		stepDuration = duration / steps
// 	}
//
// 	// calculate differences
//
// 	// TODO: Fades should be done in the HSV space for more appeasing gradients
//
// 	rDiff, gDiff, bDiff := s.Color.difference(color)
//
// 	// calculate change per steps
//
// 	// This will cause fractional steps to be discarded, meaning there
// 	// is the possibility that changes won't occur, or that the final color
// 	// won't be the requested color
//
// 	rStep := rDiff / steps
// 	gStep := gDiff / steps
// 	bStep := bDiff / steps
//
// 	for step := 0; step < steps; step++ {
// 		s.SetColor(s.Color.add(rStep, gStep, bStep))
// 		time.Sleep(stepDuration)
// 	}
//
// 	// HACK: to ensure the final color is the requested color
//
// 	s.SetColor(color)
//
// }
//
// func (s *Strip) FadeBetween(a, b RGB, duration time.Duration) {
//
// 	s.Fade(a, duration/2)
// 	// HACK: This will block. Consider using channel to break when required
// 	for {
// 		s.Fade(a, duration/2)
// 		s.Fade(b, duration/2)
// 	}
//
// }
//
// func (s *Strip) FadeOut(duration time.Duration) {
// 	s.Fade(s.Color.Off(), duration)
// }
//
// func (s *Strip) FlashBetween(c []RGB, d time.Duration) {
//
// 	// HACK: This will block. Consider using channel to break when required
// 	for {
// 		for _, color := range c {
// 			s.SetColor(color)
// 			time.Sleep(d)
// 		}
// 	}
// }
//
// func (s *Strip) Flash(c RGB, d time.Duration) {
// 	s.FlashBetween([]RGB{c, c.Off()}, d)
// }
