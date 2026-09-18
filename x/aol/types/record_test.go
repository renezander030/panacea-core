package types_test

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/v2/x/aol/types"
	"github.com/stretchr/testify/require"
)

func TestRecordValidateValueLength(t *testing.T) {
	record := types.Record{
		Key:           []byte("record-key"),
		WriterAddress: sdk.AccAddress(bytes.Repeat([]byte{1}, 20)).String(),
	}

	record.Value = bytes.Repeat([]byte{1}, 5000)
	require.NoError(t, record.Validate())

	record.Value = bytes.Repeat([]byte{1}, 5001)
	require.ErrorIs(t, record.Validate(), types.ErrMessageTooLarge)
}
