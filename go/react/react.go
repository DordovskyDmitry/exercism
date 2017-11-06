package react

const testVersion = 5

type ConcreteInputCell struct {
	value int
	r     *ReactSystem
}

func (in *ConcreteInputCell) Value() int {
	return in.value
}

func (in *ConcreteInputCell) SetValue(value int) {
	if in.value != value {
		in.value = value
		markCells(in.r.dependants[in])
		updateCellsFromRoot(in)
	}
}

type ConcreteComputeCell struct {
	value       int
	constructor (func() int)
	r           *ReactSystem
	callbacks   [](*func(int))
}

func (c *ConcreteComputeCell) Value() int {
	return c.value
}

func markCells(cells []*ConcreteComputeCell) {
	for _, c := range cells {
		c.r.marked[c] += 1
		markCells(c.r.dependants[c])
	}
}

func updateCellsFromRoot(in *ConcreteInputCell) {
	updateQueue := in.r.dependants[in]
	for len(updateQueue) > 0 {
		elem := updateQueue[0]
		if in.r.marked[elem] == 0 {
			updateQueue = updateQueue[1:]
			continue
		}
		if in.r.marked[elem] == 1 {
			updateValue(elem)
		}
		in.r.marked[elem] -= 1
		updateQueue = append(updateQueue[1:], elem.r.dependants[elem]...)
	}
}

func updateValue(c *ConcreteComputeCell) {
	newValue := c.constructor()
	if newValue != c.value {
		for _, callback := range c.callbacks {
			(*callback)(newValue)
		}
		c.value = newValue
	}
}

type ConcreteCanceler struct {
	f (*func(int))
}

func (canceler ConcreteCanceler) Cancel() {
	*canceler.f = func(int) {}
}

func (c *ConcreteComputeCell) AddCallback(f func(int)) Canceler {
	c.callbacks = append(c.callbacks, &f)
	return &ConcreteCanceler{&f}
}

type ReactSystem struct {
	dependants map[Cell]([]*ConcreteComputeCell)
	marked     map[*ConcreteComputeCell]int
}

func (r *ReactSystem) CreateInput(v int) InputCell {
	return &ConcreteInputCell{value: v, r: r}
}

func (r *ReactSystem) CreateCompute1(in Cell, f func(int) int) ComputeCell {
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
	val := f()
	return ConcreteComputeCell{constructor: f, value: val, r: r}
}

func New() Reactor {
	return &ReactSystem{dependants: make(map[Cell]([](*ConcreteComputeCell)), 0), marked: make(map[*ConcreteComputeCell]int, 0)}
}
