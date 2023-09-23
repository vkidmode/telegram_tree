package telegram_tree

type TelegramOpt func(v *telegram)

type Telegram interface {
	GetMessage() string
	GetTabTxt() string
	GetHideBar() bool
	GetSwitchInlineQueryCurrentChat() *string
	GetEnablePreview() bool
	DeleteMessage() bool
	GetResendMsg() bool
	setDefaultMessage(string)
}

type telegram struct {
	message                      string
	tabTxt                       string
	hideBar                      bool
	enablePreview                bool
	deleteMsg                    bool
	resendMsg                    bool
	switchInlineQueryCurrentChat *string
}

func (t *telegram) GetMessage() string     { return t.message }
func (t *telegram) GetTabTxt() string      { return t.tabTxt }
func (t *telegram) GetHideBar() bool       { return t.hideBar }
func (t *telegram) GetEnablePreview() bool { return t.enablePreview }
func (t *telegram) DeleteMessage() bool    { return t.deleteMsg }
func (t *telegram) GetResendMsg() bool     { return t.resendMsg }
func (t *telegram) GetSwitchInlineQueryCurrentChat() *string {
	return t.switchInlineQueryCurrentChat
}
func (t *telegram) setDefaultMessage(in string) { t.message = in }

func NewTelegram(options ...TelegramOpt) Telegram {
	tg := telegram{}
	for _, opt := range options {
		opt(&tg)
	}
	return &tg
}

func WithSwitchInline(inline string) TelegramOpt {
	return func(v *telegram) {
		v.switchInlineQueryCurrentChat = &inline
	}
}

func WithHideBar() TelegramOpt {
	return func(v *telegram) {
		v.hideBar = true
	}
}

func WithMessage(msg string) TelegramOpt {
	return func(v *telegram) {
		v.message = msg
	}
}

func WithTabTxt(tabTxt string) TelegramOpt {
	return func(v *telegram) {
		v.tabTxt = tabTxt
	}
}

func DeleteMsg() TelegramOpt {
	return func(v *telegram) {
		v.deleteMsg = true
	}
}

func EnablePreview() TelegramOpt {
	return func(v *telegram) {
		v.enablePreview = true
	}
}

func ResendMsg() TelegramOpt {
	return func(v *telegram) {
		v.resendMsg = true
	}
}
