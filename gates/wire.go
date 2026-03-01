package gates

type Wire struct {
	State bool
}

func (w *Wire) Set(value bool) {
	w.State = value
}

func (w *Wire) Get() bool {
	return w.State
}
