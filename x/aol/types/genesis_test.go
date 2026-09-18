package types_test

import (
	"bytes"
	"testing"

	sdk "github.com/cosmos/cosmos-sdk/types"
	"github.com/medibloc/panacea-core/v2/types/compkey"
	"github.com/medibloc/panacea-core/v2/x/aol/types"
	"github.com/stretchr/testify/require"
)

func TestGenesisStateValidateRejectsNilEntries(t *testing.T) {
	ownerAddress := sdk.AccAddress(bytes.Repeat([]byte{1}, 20))
	writerAddress := sdk.AccAddress(bytes.Repeat([]byte{2}, 20))

	ownerKey := compkey.EncodeToString(&types.OwnerCompositeKey{
		OwnerAddress: ownerAddress,
	}, types.GenesisKeySeparator)
	topicKey := compkey.EncodeToString(&types.TopicCompositeKey{
		OwnerAddress: ownerAddress,
		TopicName:    "topic",
	}, types.GenesisKeySeparator)
	writerKey := compkey.EncodeToString(&types.WriterCompositeKey{
		OwnerAddress:  ownerAddress,
		TopicName:     "topic",
		WriterAddress: writerAddress,
	}, types.GenesisKeySeparator)
	recordKey := compkey.EncodeToString(&types.RecordCompositeKey{
		OwnerAddress: ownerAddress,
		TopicName:    "topic",
		Offset:       1,
	}, types.GenesisKeySeparator)

	testCases := []struct {
		name   string
		insert func(*types.GenesisState)
	}{
		{
			name: "nil owner",
			insert: func(state *types.GenesisState) {
				state.Owners[ownerKey] = nil
			},
		},
		{
			name: "nil topic",
			insert: func(state *types.GenesisState) {
				state.Topics[topicKey] = nil
			},
		},
		{
			name: "nil writer",
			insert: func(state *types.GenesisState) {
				state.Writers[writerKey] = nil
			},
		},
		{
			name: "nil record",
			insert: func(state *types.GenesisState) {
				state.Records[recordKey] = nil
			},
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			state := types.DefaultGenesis()
			testCase.insert(state)

			require.ErrorContains(t, state.Validate(), "must not be nil")
		})
	}
}
