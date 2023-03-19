package telegram_tree

type nodesList []*node

func (n nodesList) setupCallBacks(callback string) error {
	for i := range n {
		newCallback, err := incrementCallback(callback, n[i].payload, i)
		if err != nil {
			return err
		}
		n[i].callback = newCallback
	}
	return nil
}

func (n nodesList) toInterface() []Node {
	var resp = make([]Node, len(n))
	for i := range n {
		resp[i] = n[i].toInterface()
	}
	return resp
}
