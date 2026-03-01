package gates

type NotGate struct {
	In  *Wire
	Out *Wire
}

func NewNotGate(in, out *Wire) *NotGate {
	return &NotGate{In: in, Out: out}
}

func (n *NotGate) Evaluate() {
	n.Out.Set(!n.In.Get())
}
