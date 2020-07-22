package balance

type Balancer interface {
	Select(Volumes) (*Volume, error)
}
