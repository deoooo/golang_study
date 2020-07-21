package balance

import "fmt"

type Mgr struct {
	allVolume map[string]Balancer
}

var mgr = Mgr{allVolume: make(map[string]Balancer)}

func (m *Mgr) registerBalancer(name string, b Balancer) {
	m.allVolume[name] = b
}

func RegisterBalancer(name string, b Balancer) {
	mgr.registerBalancer(name, b)
}

func SelectByName(name string, volumes []*Volume) (v *Volume, err error) {
	b, ok := mgr.allVolume[name]
	if !ok {
		err = fmt.Errorf("not found %s balance", name)
		return
	}

	fmt.Printf("use %s volume\n", name)
	v, err = b.Select(volumes)
	return
}
