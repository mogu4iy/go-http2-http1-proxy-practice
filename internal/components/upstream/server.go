package upstream

type Server interface {
	getWeight() int
	GetAddress() string
}

type server struct {
	addr   string
	weight int
}

func (s *server) getWeight() int {
	return s.weight
}

func (s *server) GetAddress() string {
	return s.addr
}
