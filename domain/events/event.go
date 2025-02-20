package events

type Event interface {
	Name() string
	IsAsynchronous() bool 
}

type AccountBelonging struct {
	doesBelongTo bool
}

func (*AccountBelonging) Name() string {
	return "event.account.belonging"
}

func (event *AccountBelonging) DoesBelongTo() bool {
	return event.doesBelongTo
}