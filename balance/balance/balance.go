package balance

type Balancer interface {
	Select([]*Volume) (*Volume, error)
}
