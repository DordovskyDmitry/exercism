package react

const testVersion = 5

type ConcreteInputCell struct {
	value int
	r *ReactSystem
}

func (in *ConcreteInputCell) Value() int {
	return in.value
}

func (in *ConcreteInputCell) SetValue(value int) {
	if in.value != value {
		in.value = value
		updateCells(in.r.dependants[in])
	}
}

type ConcreteComputeCell struct {
	value int
	constructor (func() int)
	r *ReactSystem
	callbacks [](*func(int))
}

func (c *ConcreteComputeCell) Value() int {
	return c.value
}

func updateComputeCellValues(c *ConcreteComputeCell) {
	oldValue := c.value
	c.value = c.constructor()

	if c.value != oldValue {
		updateCells(c.r.dependants[c])
		for _, callback := range c.callbacks {
			(*callback)(c.value)
		}
	}
}

func updateCells(cells []*ConcreteComputeCell) {
	for _, cell := range cells {
		updateComputeCellValues(cell)
	}
}


type ConcreteCanceler struct {
	f (*func(int))
}

func (canceler ConcreteCanceler) Cancel() {
	*canceler.f = func(int){}
}

func (c *ConcreteComputeCell) AddCallback(f func(int)) Canceler {
	c.callbacks = append(c.callbacks, &f)
	return &ConcreteCanceler{&f}
}

type ReactSystem struct {
	dependants map[Cell]([]*ConcreteComputeCell)
}

func (r *ReactSystem) CreateInput(v int) InputCell {
	return &ConcreteInputCell{value: v, r: r}
}

func (r *ReactSystem) CreateCompute1(in Cell, f (func(int) int)) ComputeCell {
	c := createCompute(r, func() int { return f(in.Value()) })
	r.dependants[in] = append(r.dependants[in], &c)
	return &c
}

func (r *ReactSystem) CreateCompute2(in1 Cell, in2 Cell, f func(int, int) int) ComputeCell {
	c := createCompute(r, func() int { return f(in1.Value(), in2.Value()) })
	r.dependants[in1] = append(r.dependants[in1], &c)
	r.dependants[in2] = append(r.dependants[in2], &c)
	return &c
}

func createCompute(r *ReactSystem, f func() int) ConcreteComputeCell {
	return ConcreteComputeCell{constructor: f, value: f(), r: r}
}

func New() Reactor {
	return &ReactSystem{dependants: make(map[Cell]([](*ConcreteComputeCell)), 0)}
}
