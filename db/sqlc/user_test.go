package db

import (
	"context"
	"fmt"
	"testing"

	"github.com/stinekamau/simplebank/utils"
	"github.com/stretchr/testify/require"
)

func createRandomUser(t *testing.T) User {

	hashedPass, err := utils.HashPassword(utils.RandomString(6))
	require.NoError(t, err)
	args := CreateUserParams{
		Username:       utils.RandomOwner(),
		HashedPassword: hashedPass,
		FullName:       utils.RandomOwner() + " " + utils.RandomOwner(),
		Email:          utils.RandomOwner() + "@gmail.com",
	}

	user, err := testQueries.CreateUser(context.Background(), args)
	fmt.Printf("%+v", user)

	require.NotNil(t, user)
	require.Nil(t, err)

	require.Equal(t, user.Email, args.Email)
	require.Equal(t, user.Username, args.Username)
	require.Equal(t, user.HashedPassword, args.HashedPassword)
	require.Equal(t, args.FullName, user.FullName)

	return user

}

func TestCreateUser(t *testing.T) {
	createRandomUser(t)
}

func TestGetUser(t *testing.T) {
	user1 := createRandomUser(t)
	user2, err := testQueries.GetUser(context.Background(), user1.Username)

	require.NoError(t, err)
	require.NotEmpty(t, user2)

	require.Equal(t, user1.Username, user2.Username)
	require.Equal(t, user1.HashedPassword, user2.HashedPassword)
	require.Equal(t, user1.FullName, user2.FullName)
	require.Equal(t, user1.Email, user2.Email)

}
