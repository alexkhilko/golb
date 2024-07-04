package servers

type Pool struct {
	Healthy   *NodeList
	Unhealthy *NodeList
}

func (p Pool) GetNextServerAddr() string {
	n := p.Healthy.Pop()
	if n == nil {
		return ""
	}
	p.Healthy.Push(n.Val)
	return n.Val
}

func (p Pool) Suspend(n *Node) {
	p.Healthy.Remove(n)
	p.Unhealthy.Push(n.Val)
}

func (p Pool) Activate(n *Node) {
	p.Unhealthy.Remove(n)
	p.Healthy.Push(n.Val)
}

func NewPool(servers []string) *Pool {
	h := NewNodeList()
	for _, s := range servers {
		h.Push(s)
	}
	return &Pool{Healthy: h, Unhealthy: NewNodeList()}
}
