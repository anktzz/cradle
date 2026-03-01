package gates

type AndGate struct {
	A, B *Wire
	Out  *Wire
}

func NewAndGate(a, b, out *Wire) *AndGate {
	return &AndGate{A: a, B: b, Out: out}
}

func (a *AndGate) Evaluate() {
	a.Out.Set(a.A.Get() && a.B.Get())
}
