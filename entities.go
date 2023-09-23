package telegram_tree

const (
	CallBackClose   = "*"
	CallBackIgnore  = "#"
	CallBackSkip    = "@"
	callbackDivider = ">"
)

type Meta interface {
	GetCallback() string
	SetupCallback(in string)
	SetIsMiddle(middle bool)
	GetIsMiddle() bool
}

type metaRealization struct {
	callback string
	middle   bool
}

func (m *metaRealization) GetCallback() string     { return m.callback }
func (m *metaRealization) SetupCallback(in string) { m.callback = in }
func (m *metaRealization) SetIsMiddle(middle bool) { m.middle = middle }
func (m *metaRealization) GetIsMiddle() bool       { return m.middle }

func newMeta(in string) Meta {
	return &metaRealization{
		callback: in,
	}
}

type nodesList []Node

func (n nodesList) setupCallBacks(callback string) error {
	for i := range n {
		newCallback, err := incrementCallback(callback, n[i].getPayload(), i)
		if err != nil {
			return err
		}
		n[i].setCallback(newCallback)
	}
	return nil
}

type Payload interface {
	GetKey() string
	GetValue() string
}

type payload struct {
	key   string
	value string
}

func (p *payload) GetKey() string   { return p.key }
func (p *payload) GetValue() string { return p.value }

func NewPayload(key, value string) Payload {
	return &payload{
		key:   key,
		value: value,
	}
}
