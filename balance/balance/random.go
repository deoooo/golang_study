package balance

import (
	"errors"
	"math/rand"
)

func init() {
	RegisterBalancer("random", &Random{})
}

type Random struct {
}

func (r Random) Select(volumes Volumes) (v *Volume, err error) {
	lens := len(volumes)
	if lens == 0 {
		err = errors.New("No Volume")
	}
	index := rand.Intn(lens)
	v = volumes[index]
	return
}
