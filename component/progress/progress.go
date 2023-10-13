package progress

import (
	"errors"
	"math"
	"strings"

	"github.com/ecoshub/termium/component/config"
	"github.com/ecoshub/termium/component/panel"
	"github.com/ecoshub/termium/component/style"
)

var (
	DefaultBarStyle []rune = []rune{'[', '=', '>', '-', ']'}
)

type Config struct {
	Width        int
	BarStyle     []rune
	ContentStyle *style.Style
}

type ProgressBar struct {
	config *Config
	*panel.Base
}

func New(conf *Config) (*ProgressBar, error) {
	if conf.BarStyle == nil {
		conf.BarStyle = DefaultBarStyle
	}
	if conf.ContentStyle == nil {
		conf.ContentStyle = &style.Style{}
	}
	if len(conf.BarStyle) != 5 {
		return nil, errors.New("bar style must contain 5 char. example [=>-]")
	}
	if conf.Width < 5 {
		return nil, errors.New("progress bar width must greater than 5")
	}
	return &ProgressBar{
		config: conf,
		Base: panel.NewBasicPanel(&config.Config{
			Width:        conf.Width,
			Height:       1,
			ContentStyle: conf.ContentStyle,
		}),
	}, nil
}

func (pb *ProgressBar) Update(percent float64) {
	pb.Write(0, pb.draw(pb.Config.Width, percent))
}

func (pb *ProgressBar) draw(width int, percent float64) string {
	if percent >= 1 {
		return pb.fullBar(width)
	}
	if percent == 0 {
		return pb.emptyBar(width)
	}
	return pb.percentBar(width, percent)
}

func (pb *ProgressBar) percentBar(width int, percent float64) string {
	barWidth := width - 2
	doneCount := int(math.Floor(percent * float64(barWidth)))
	bar := strings.Builder{}
	bar.WriteRune(pb.config.BarStyle[0])
	for i := 0; i < barWidth; i++ {
		if i < doneCount {
			bar.WriteRune(pb.config.BarStyle[1])
		} else if i == doneCount {
			bar.WriteRune(pb.config.BarStyle[2])
		} else {
			bar.WriteRune(pb.config.BarStyle[3])
		}
	}
	bar.WriteRune(pb.config.BarStyle[4])
	return bar.String()
}

func (pb *ProgressBar) fullBar(width int) string {
	return fillBar(width, pb.config.BarStyle[0], pb.config.BarStyle[1], pb.config.BarStyle[4])
}

func (pb *ProgressBar) emptyBar(width int) string {
	return fillBar(width, pb.config.BarStyle[0], pb.config.BarStyle[3], pb.config.BarStyle[4])
}

func fillBar(width int, start, char, end rune) string {
	s := strings.Builder{}
	s.WriteRune(start)
	innerSize := width - 2
	for i := 0; i < innerSize; i++ {
		s.WriteRune(char)
	}
	s.WriteRune(end)
	return s.String()
}
