package usecase

import (
	"errors"

	"github.com/biancanobrega/imersao-fullstack-fullcycle/codepix/domain/model"
)

type PixUseCase struct {
	PixKeyRepository model.PixKeyRepositoryInterface
}

func (p *PixUseCase) RegisterKey(key string, kind string, accountId string) (*model.PixKey, error) {
	account, err := p.PixKeyRepository.FindAccount(accountId)
	if err != nil {
		return nil, err
	}

	pixKey, err := model.NewPixKey(account, kind, key)
	if err != nil {
		return nil, err
	}

	p.PixKeyRepository.RegisterKey(pixKey)
	if pixKey.ID == "" {
		return nil, errors.New("unable to create new key at the moment")
	}

	return pixKey, err
}

func (p *PixUseCase) FindKey(key string, kind string) (*model.PixKey, error) {
	pixKey, err := p.PixKeyRepository.FindKeyByKind(key, kind)
	if err != nil {
		return nil, err
	}
	return pixKey, nil
}
