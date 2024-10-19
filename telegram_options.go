package telegram_tree

type TelegramOpt func(v *telegram)

type Telegram interface {
	GetMessage() string
	GetTabTxt() string
	GetHideBar() bool
	GetSwitchInlineQueryCurrentChat() *string
	GetLink() *string
	GetEnablePreview() bool
	DeleteMessage() bool
	GetResendMsg() bool
	setDefaultMessage(string)
	GetColumns() int
}

type telegram struct {
	message                      string
	tabTxt                       string
	hideBar                      bool
	enablePreview                bool
	deleteMsg                    bool
	resendMsg                    bool
	switchInlineQueryCurrentChat *string
	link                         *string
	columns                      int
}

func (t *telegram) GetMessage() string     { return t.message }
func (t *telegram) GetColumns() int        { return t.columns }
func (t *telegram) GetTabTxt() string      { return t.tabTxt }
func (t *telegram) GetHideBar() bool       { return t.hideBar }
func (t *telegram) GetEnablePreview() bool { return t.enablePreview }
func (t *telegram) DeleteMessage() bool    { return t.deleteMsg }
func (t *telegram) GetResendMsg() bool     { return t.resendMsg }
func (t *telegram) GetSwitchInlineQueryCurrentChat() *string {
	return t.switchInlineQueryCurrentChat
}
func (t *telegram) GetLink() *string {
	return t.link
}
func (t *telegram) setDefaultMessage(in string) { t.message = in }

func NewTelegram(options ...TelegramOpt) Telegram {
	tg := telegram{
		columns: 1,
	}
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

func WithLink(link string) TelegramOpt {
	return func(v *telegram) {
		v.link = &link
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

func WithColumns(columnsCount int) TelegramOpt {
	return func(v *telegram) {
		v.columns = columnsCount
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
