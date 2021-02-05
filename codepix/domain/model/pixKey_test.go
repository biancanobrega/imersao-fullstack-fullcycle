package model_test

import (
	"testing"

	uuid "github.com/satori/go.uuid"

	"github.com/biancanobrega/imersao-fullstack-fullcycle/codepix/domain/model"
	"github.com/stretchr/testify/require"
)

func TestModel_NewPixKey(t *testing.T) {
	code := "001"
	name := "Banco do Brasil"
	bank, err := model.NewBank(code, name)

	accountNumber := "abcnumber"
	ownerName := "Bianca"
	account, err := model.NewAccount(bank, accountNumber, ownerName)

	kind := "email"
	key := "teste@teste.com"
	pixKey, err := model.NewPixKey(account, kind, key)

	require.NotEmpty(t, uuid.FromStringOrNil(pixKey.ID))
	require.Equal(t, pixKey.Kind, kind)
	require.Equal(t, pixKey.Key, key)
	require.Equal(t, pixKey.Status, "active")

	kind = "cpf"
	_, err = model.NewPixKey(account, kind, key)
	require.Nil(t, err)

	_, err = model.NewPixKey(account, "name", key)
	require.NotNil(t, err)

}
