package token

import (
	"fmt"
	"testing"
	"time"

	"github.com/stretchr/testify/require"
)

func TestCreatePasetoToken(t *testing.T) {

	pasetoM, err := NewPasetoMaker()
	require.NoError(t, err)
	require.NotEmpty(t, pasetoM)

	encrypted, err := pasetoM.CreateToken("paul", 3*time.Second)
	require.NoError(t, err)

	payload, err := pasetoM.VerifyToken(encrypted)
	fmt.Printf("Current value of payload is: %+v", payload)
	require.NoError(t, err)
	require.NotEmpty(t, payload)

}
