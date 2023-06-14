package day10

import "fmt"

const (
	SIGNAL_START = 20
	SIGNAL_CYCLE = 40
)

type Processor struct {
	clock           int
	register        int
	bufferIn        Queue
	bufferOut       Queue
	currentOp       Operation
	cmdCache        Queue
	signals         []int
	signalsOnlyFlag bool
	crt             CRT
}

func NewProcessor(input []string, signalsOnly bool) Processor {
	clock := 0
	register := 1
	bufferIn := NewQueue(input)
	bufferOut := NewQueue([]string{})
	currentOp := Operation{}
	cmdCache := NewQueue([]string{})
	signals := []int{}
	crt := NewCRT()

	return Processor{clock, register, bufferIn, bufferOut, currentOp, cmdCache, signals, signalsOnly, crt}
}

func (p *Processor) Cycle() bool {
	if p.bufferIn.IsEmpty() && p.cmdCache.IsEmpty() {
		return false
	}

	p.crt.Draw(p.clock, p.register)
	p.clock += 1
	p.Display()

	if !p.cmdCache.IsEmpty() {
		p.currentOp = p.cmdCache.Pop()
		switch p.currentOp.Cmd {
		case "addx":
			p.register += p.currentOp.Arg
		}

		return true
	}

	p.currentOp = p.bufferIn.Pop()

	if p.currentOp.Cmd != "noop" {
		p.cmdCache.Push(p.currentOp)
	}

	return true
}

// Check signalStrength at the 20th cycle and every 40 cycles after that
func shouldCheckSignalStrength(cycle int, signalStart int, signalCycle int) bool {
	return cycle == signalStart || cycle > signalStart && (cycle-signalStart)%signalCycle == 0
}

func (p *Processor) signalStrength() (int, bool) {
	if shouldCheckSignalStrength(p.clock, SIGNAL_START, SIGNAL_CYCLE) {
		s := p.clock * p.register
		p.signals = append(p.signals, s)
		return s, true
	}

	return -1, false
}

// the X register controls the horizontal position of a sprite. Specifically, the sprite is 3
// pixels wide, and the X register sets the horizontal position of the middle of that sprite.
// This CRT screen draws the top row of pixels left-to-right, then the row below that,
// and so on. The left-most pixel in each row is in position 0, and the right-most pixel
// in each row is in position 39.
// You should be able to determine whether the sprite is visible the instant each pixel is
// drawn. If the sprite is positioned such that one of its three pixels is the pixel
// currently being drawn, the screen produces a lit pixel (#); otherwise, the screen leaves
// the pixel dark (.).

func (p *Processor) Display() {
	signal, showSignal := p.signalStrength()

	if !p.signalsOnlyFlag || showSignal {
		fmt.Println(fmt.Sprintf("Clock: %5d, Register: %5d, Op: %s, Signal: %d", p.clock, p.register, &p.currentOp, signal))
	}
}

func (p *Processor) GetSignals() []int {
	s := p.signals

	return s
}

func (p *Processor) SumSignals() int {
	sum := 0

	for _, s := range p.signals {
		sum += s
	}

	return sum
}
