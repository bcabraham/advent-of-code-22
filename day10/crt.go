package day10

import (
	"fmt"
	"strings"
)

const (
	CRT_HEIGHT   = 6
	CRT_WIDTH    = 40
	SPRITE_WIDTH = 3
)

type CRT struct {
	pixels [][]string
	buffer []string
}

func (c *CRT) newBuffer() []string {
	b := make([]string, CRT_WIDTH)

	return b
}

func NewCRT() CRT {
	p := make([][]string, CRT_HEIGHT)

	b := make([]string, CRT_WIDTH)
	crt := CRT{p, b}

	return crt
}

func getCRTCoords(cycle int) (int, int) {
	return cycle / CRT_WIDTH, cycle % CRT_WIDTH
}

// if the sprite's horizontal position puts its pixels where the CRT is currently drawing,
// then those pixels will be drawn.
func (c *CRT) spriteIsVisible(cycle, register int) bool {
	_, xPos := getCRTCoords(cycle)

	return xPos-1 <= register && register <= xPos+1
}

func (c *CRT) Draw(cycle, register int) {
	y, x := getCRTCoords(cycle)
	fmt.Println(fmt.Sprintf("Coords: %d, %d", x, y))

	if c.spriteIsVisible(cycle, register) {
		c.buffer[x] = "#"
	} else {
		c.buffer[x] = "."
	}

	// c.PrintBuffer()

	if x == CRT_WIDTH-1 {
		c.Flush(y)
	}
}

func (c *CRT) Flush(row int) {
	c.pixels[row] = c.buffer

	c.buffer = c.newBuffer()

	fmt.Println(c)
}

func (c *CRT) PrintBuffer() {
	fmt.Println("CRT Buffer:", strings.Join(c.buffer, ""))
}

func (c *CRT) String() string {
	output := []string{}

	for _, l := range c.pixels {
		output = append(output, strings.Join(l, ""))
	}

	return strings.Join(output, "\n")
}
