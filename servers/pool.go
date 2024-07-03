package servers

type Pool struct {
	healthy   map[string]bool
	unhealthy map[string]bool
	counter   int
}

func (p Pool) GetNextServerAddr() string {
	servers := p.GetHealthyServerURLs()
	if len(servers) == 0 {
		return ""
	}
	nextIdx := p.counter % len(servers)
	return servers[nextIdx]
}

func (p Pool) Suspend(s string) {
	delete(p.healthy, s)
	p.unhealthy[s] = true
}

func (p Pool) Activate(s string) {
	delete(p.unhealthy, s)
	p.healthy[s] = true
}

func (p Pool) GetHealthyServerURLs() []string {
	servers := make([]string, len(p.healthy))
	i := 0
	for server := range p.healthy {
		servers[i] = server
		i++
	}
	return servers
}

func (p Pool) GetUnhealthyServerURLs() []string {
	servers := make([]string, len(p.unhealthy))
	i := 0
	for server := range p.unhealthy {
		servers[i] = server
		i++
	}
	return servers
}

func NewPool(healthy []string) *Pool {
	servers := make(map[string]bool, len(healthy))
	for _, s := range healthy {
		servers[s] = true
	}
	return &Pool{healthy: servers, unhealthy: map[string]bool{}}
}
