package service

import (
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)


type Reverser struct {
	serviceInterfaces.Reverser 
	repos repository.ReverserRepository
}
//TODO

func NewReverser(repos repository.ReverserRepository) *Reverser {
	return &Reverser{repos: repos}
}

func (reverser *Reverser) getAction(actionId int) (string, []string, error) {
	action, err := reverser.repos.GetAction(actionId)
	if err != nil {
		return "", nil, err
	}
	return action.ActionName(), action.ActionArgs(), nil
}

// func (reverser *Reverser) Reverse(actionId int) error {
// 	actionName, args, err := reverser.getAction(actionId)
// 	if err != nil {
// 		return err
// 	}
// 	switch actionName {
	
// 	}
// 	return nil
// }