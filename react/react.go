package react

type canceler struct {
	c  *cell
	id int
}

func (c *canceler) Cancel() {
	delete(c.c.callbacks, c.id)
}

type cell struct {
	value          int
	reactor        *reactor
	update         func() int
	nextCallbackId func() int
	callbacks      map[int]func(int)
}

func (c *cell) Value() int {
	return c.value
}

func (c *cell) SetValue(value int) {
	container := Container{}
	c.setValue(value, &container)

	if *container.currentValue != *container.previousValue {
		for _, callback := range container.callbacks {
			callback(*container.currentValue)
		}
	}
}

type Container struct {
	previousValue *int
	currentValue  *int
	callbacks     map[int]func(int)
}

func (c *Container) setCurrentValue(value int) {
	c.currentValue = &value
}

func (c *Container) setPreviousValue(value int) {
	c.previousValue = &value
}

func (c *cell) setValue(value int, container *Container) {
	if !c.reactor.nodes.HasEntriesFor(c) {
		if container.previousValue == nil {
			container.setPreviousValue(c.value)
		}
		container.setCurrentValue(value)
		container.callbacks = c.callbacks
	}

	c.value = value

	for _, child := range c.reactor.nodes[c] {
		child.setValue(child.update(), container)
	}
}

func (c *cell) AddCallback(callback func(int)) Canceler {
	id := c.nextCallbackId()
	c.callbacks[id] = callback
	return &canceler{
		c:  c,
		id: id,
	}
}

type NodesMap map[Cell][]*cell

func (n NodesMap) HasEntriesFor(c Cell) bool {
	return n != nil && len(n[c]) != 0
}

type reactor struct {
	nodes NodesMap
}

func New() Reactor {
	return &reactor{
		nodes: map[Cell][]*cell{},
	}
}

func (r *reactor) CreateInput(initial int) InputCell {
	return &cell{
		value:   initial,
		reactor: r,
	}
}

func (r *reactor) CreateCompute1(dep Cell, compute func(int) int) ComputeCell {
	c := &cell{
		reactor: r,
		value:   compute(dep.Value()),
		update: func() int {
			return compute(dep.Value())
		},
		callbacks:      map[int]func(int){},
		nextCallbackId: CallBackIdGenerator(),
	}
	r.nodes[dep] = append(r.nodes[dep], c)

	return c
}

func (r *reactor) CreateCompute2(dep1, dep2 Cell, compute func(int, int) int) ComputeCell {
	c := &cell{
		reactor: r,
		value:   compute(dep1.Value(), dep2.Value()),
		update: func() int {
			return compute(dep1.Value(), dep2.Value())
		},
		callbacks:      map[int]func(int){},
		nextCallbackId: CallBackIdGenerator(),
	}
	r.nodes[dep1] = append(r.nodes[dep1], c)
	r.nodes[dep2] = append(r.nodes[dep2], c)

	return c
}

func CallBackIdGenerator() func() int {
	id := 0
	return func() int {
		id++
		return id
	}
}
