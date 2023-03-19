package telegram_tree

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
