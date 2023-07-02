package basic

import (
	"fmt"
	"math/rand"
	"term/internal/component"
	"term/internal/models/dimension"
	"term/internal/utils"
	"time"
)

const (
	ResetAllModes string = "\x1b[0m"
)

type Config struct {
	Title string
	Size  *dimension.D2
}

type Component struct {
	Config *Config
	buffer [][]rune
	lines  []string
}

func New(conf *Config) *Component {
	c := &Component{
		Config: conf,
		buffer: utils.InitRuneMatrix(conf.Size.X, conf.Size.Y, ' '),
		lines:  make([]string, conf.Size.Y),
	}
	go func() {
		n := rand.Intn(26)
		count := 0
		for range time.NewTicker(time.Millisecond * 500).C {
			c.buffer = utils.InitRuneMatrix(c.Config.Size.X, c.Config.Size.Y, rune(n+97))
			c.Insert(count, "\x1b[1;31mHello\x1b[0m my name is eco")
			count = (count + 1) % (c.Config.Size.Y)
		}
	}()
	return c
}

func (c *Component) Insert(index int, line string) error {
	if index >= (c.Config.Size.Y) {
		return fmt.Errorf("index out of range. index: %d, size: %d", index, c.Config.Size.Y)
	}
	c.lines[index] = line
	c.renderLine(index)
	return nil
}

func (c *Component) GetTitle() []rune {
	title := component.ResizeLine(c.Config.Title, c.Config.Size.X)
	return title
}

func (c *Component) Clear(index int) {
	c.buffer[index] = []rune(fmt.Sprintf("%-*s", c.Config.Size.X, " "))
}

func (c *Component) GetSize() *dimension.D2 {
	return c.Config.Size
}

func (c *Component) GetBuffer() [][]rune {
	return c.buffer
}

func (c *Component) renderLine(index int) {
	line := c.lines[index]
	r := component.ResizeLine(line, c.Config.Size.X)
	c.buffer[index] = []rune(r)
}
