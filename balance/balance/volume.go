package balance

import "fmt"

type Volume struct {
	host string
	port int
}

func NewVolume(host string, port int) *Volume {
	return &Volume{
		host: host,
		port: port,
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
