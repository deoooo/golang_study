package balance

import (
	"errors"
)

var i int = -1
var currentWeight int

type weights struct {
}

func init() {
	mgr.registerBalancer("weights", &weights{})
}

func (w weights) Select(volumes Volumes) (*Volume, error) {
	if len(volumes) < 1 {
		return nil, errors.New("No Volume")
	}

	for {
		i = (i + 1) % len(volumes)
		if i == 0 {
			currentWeight--
			if currentWeight <= 0 {
				currentWeight = volumes.getMaxWeight()
			}
		}

		if volumes[i].weight >= currentWeight {
			return volumes[i], nil
		}

	}

}
