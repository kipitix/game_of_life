package cell

type cellImpl struct {
	state State
}

func NewCell() Cell {
	return &cellImpl{}
}

var _ Cell = (*cellImpl)(nil)

func (c *cellImpl) State() State {
	return c.state
}

func (c *cellImpl) SetState(state State) {
	c.state = state
}
