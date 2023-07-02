package stack

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
	Size  *dimension.D2
	Dummy bool
}

type Component struct {
	Config *Config
	buffer [][]rune
	lines  []string
	index  int
}

func New(conf *Config) *Component {
	c := &Component{
		Config: conf,
		buffer: utils.InitRuneMatrix(conf.Size.X, conf.Size.Y, ' '),
		lines:  make([]string, conf.Size.Y),
	}
	if conf.Dummy {
		go func() {
			for range time.NewTicker(time.Millisecond * 250).C {
				n := rand.Intn(26)
				c.Append(fmt.Sprintf("%s_%d", "\x1b[1;31mHello\x1b[0m my name is eco", n))
			}
		}()
	}
	return c
}

func (c *Component) Append(line string) {
	if c.index >= c.Config.Size.Y {
		c.lines = c.lines[1:]
		c.lines = append(c.lines, line)
	} else {
		c.lines[c.index] = line
		c.index++
	}
	c.renderList()
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

func (c *Component) renderList() {
	for i := range c.lines {
		c.renderLine(i)
	}
}

func (c *Component) renderLine(index int) {
	line := c.lines[index]
	r := component.ResizeLine(line, c.Config.Size.X)
	c.buffer[index] = []rune(r)
}
