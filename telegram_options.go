package telegram_tree

type TelegramOptions interface {
	GetMessage() string
	GetHumanText() string
	GetHideBar() bool
	GetSwitchInlineQueryCurrentChat() string
	GetEnablePreview() bool
	DeleteMessage() bool
	setDefaultMessage(string)
}

type telegramOptions struct {
	message                      string
	humanText                    string
	switchInlineQueryCurrentChat string
	hideBar                      bool
	enablePreview                bool
	deleteMsg                    bool
}

func (t *telegramOptions) setDefaultMessage(in string) { t.message = in }
func (t *telegramOptions) GetMessage() string          { return t.message }
func (t *telegramOptions) GetHumanText() string        { return t.humanText }
func (t *telegramOptions) GetHideBar() bool            { return t.hideBar }
func (t *telegramOptions) GetEnablePreview() bool      { return t.enablePreview }
func (t *telegramOptions) DeleteMessage() bool         { return t.deleteMsg }
func (t *telegramOptions) GetSwitchInlineQueryCurrentChat() string {
	return t.switchInlineQueryCurrentChat
}

func NewTelegramOptions(message, humanText, switchInlineQueryCurrentChat string, hideBar, enablePreview, deleteMsg bool) TelegramOptions {
	return &telegramOptions{
		message:                      message,
		humanText:                    humanText,
		switchInlineQueryCurrentChat: switchInlineQueryCurrentChat,
		hideBar:                      hideBar,
		enablePreview:                enablePreview,
		deleteMsg:                    deleteMsg,
	}
}
