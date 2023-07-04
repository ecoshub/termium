package screen

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"time"
)

func (s *Screen) calculateFPS() {
	if !s.ShowFPS {
		return
	}
	delta := time.Since(s.lastRender)
	fps := 1.0 / float64(delta.Milliseconds()) * 1000
	fpsString := fmt.Sprintf("fps:%-4d", int(fps))
	runes := []rune(fpsString)
	copy(s.buffer[0][s.sizeX-1-8:s.sizeX-1], runes[:8])
}

func ListenInterrupt() {
	chanInterrupt := make(chan os.Signal, 1)
	signal.Notify(chanInterrupt, os.Interrupt)
	for {
		select {
		case <-chanInterrupt:
			log.Print("Interrupted. Exiting...")
			os.Exit(0)
		}
	}
}
