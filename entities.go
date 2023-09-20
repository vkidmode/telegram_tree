package telegram_tree

const (
	CallBackClose   = "*"
	CallBackIgnore  = "#"
	CallBackSkip    = "@"
	callbackDivider = ">"
)

type meta interface {
	GetCallback() Callback
}

type metaRealization struct {
	callback Callback
}

func (m *metaRealization) GetCallback() Callback {
	return m.callback
}

func newMeta(in Callback) meta {
	return &metaRealization{
		callback: in,
	}
}

type nodesList []Node

func (n nodesList) setupCallBacks(callback Callback) error {
	for i := range n {
		newCallback, err := incrementCallback(callback, n[i].GetPayload(), i)
		if err != nil {
			return err
		}
		n[i].setCallback(newCallback)
	}
	return nil
}

type payload struct {
	key   string
	value string
}

func (p *payload) GetKey() string {
	return p.key
}

func (p *payload) GetValue() string {
	return p.value
}

type Payload interface {
	GetKey() string
	GetValue() string
}

func NewPayload(key, value string) Payload {
	return &payload{
		key:   key,
		value: value,
	}
}
