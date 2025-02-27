package serviceInterfaces

type Reverser interface {
	getAction(actionId int) (string, []string, error)
	Reverse(actionId, usrId, usrRole int) error 
}