package service

import (
	"main/service/repository"
	serviceInterfaces "main/service/service_interfaces"
)

type Reverser struct {
	serviceInterfaces.Reverser
	accountReverser  repository.AccountReverserRepository
	clientReverser   repository.ClientActionsReverserRepository
	operatorReverser repository.OperatorActionsReverserRepository
	infoRepos        repository.ReverserInfoRepository
}

//TODO

func NewReverser(accountReverser repository.AccountReverserRepository, clientReverser repository.ClientActionsReverserRepository,
	operatorReverser repository.OperatorActionsReverserRepository,
	infoRepos repository.ReverserInfoRepository) *Reverser {
	return &Reverser{
		accountReverser:  accountReverser,
		clientReverser:   clientReverser,
		operatorReverser: operatorReverser,
		infoRepos:        infoRepos,
	}
}

func (reverser *Reverser) getAction(actionId int) (string, []string, error) {
	action, err := reverser.infoRepos.GetAction(actionId)
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
