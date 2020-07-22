package balance

import "fmt"

type Volume struct {
	host   string
	port   int
	weight int
}

type Volumes []*Volume

func (volumes Volumes) getMaxWeight() int {
	max := 0
	for _, v := range volumes {
		if max < v.weight {
			max = v.weight
		}
	}
	return max
}

func NewVolume(host string, port int, weight int) *Volume {
	return &Volume{
		host:   host,
		port:   port,
		weight: weight,
	}
}

func (v *Volume) GetHost() string {
	return v.host
}

func (v *Volume) GetPort() int {
	return v.port
}

func (v *Volume) String() string {
	return fmt.Sprintf("%s:%d", v.GetHost(), v.GetPort())
}
