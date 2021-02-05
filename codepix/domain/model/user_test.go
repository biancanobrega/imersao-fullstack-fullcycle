package model_test

import (
	"testing"

	"github.com/biancanobrega/imersao-fullstack-fullcycle/codepix/domain/model"
	uuid "github.com/satori/go.uuid"
	"github.com/stretchr/testify/require"
)

func TestModel_NewUser(t *testing.T) {
	name := "Bianca"
	email := "teste@teste.com"

	user, err := model.NewUser(name, email)

	require.Nil(t, err)
	require.NotNil(t, uuid.FromStringOrNil(user.ID))
	require.Equal(t, user.Name, name)
	require.Equal(t, user.Email, email)

	_, err = model.NewUser("", "")
	require.NotNil(t, err)
}
