package day9

type Dimension struct {
	MinX int
	MaxX int
	MinY int
	MaxY int
}

func (d *Dimension) SetMinMax(k Knot) {
	if k.X < d.MinX {
		d.MinX = k.X
	}
	if k.X > d.MaxX {
		d.MaxX = k.X
	}
	if k.Y < d.MinY {
		d.MinY = k.Y
	}
	if k.Y > d.MaxY {
		d.MaxY = k.Y
	}
}

func getDimensions(moves []Move) Dimension {
	k := Knot{"K", 0, 0}
	d := Dimension{}

	for _, move := range moves {
		k.MoveMany(move)
		d.SetMinMax(k)
	}

	return d
}

func (d *Dimension) Height() int {
	return d.MaxY - d.MinY + 1
}

func (d *Dimension) Width() int {
	return d.MaxX - d.MinX + 1
}

func (d *Dimension) XOffset() int {
	if d.MinX < 0 {
		return d.MinX
	}

	return 0
}

func (d *Dimension) YOffset() int {
	if d.MinY < 0 {
		return d.MinY
	}

	return 0
}
