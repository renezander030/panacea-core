package types_test

import (
	"testing"

	"github.com/medibloc/panacea-core/v2/x/did/types"
	"github.com/stretchr/testify/require"
)

func TestDefaultGenesisState(t *testing.T) {
	defaultState := types.DefaultGenesis()
	require.Empty(t, defaultState.Documents)
}

func TestGenesisStateValidateRejectsNilDocuments(t *testing.T) {
	did := "did:panacea:7Prd74ry1Uct87nZqL3ny7aR7Cg46JamVbJgk8azVgUm"
	validDocument := getValidDIDDocument()
	validDocumentWithSeq := types.NewDIDDocumentWithSeq(&validDocument, types.InitialSequence)

	testCases := []struct {
		name     string
		document *types.DIDDocumentWithSeq
		wantErr  bool
	}{
		{
			name:     "valid document",
			document: &validDocumentWithSeq,
			wantErr:  false,
		},
		{
			name:     "nil map value",
			document: nil,
			wantErr:  true,
		},
		{
			name:     "nil nested document",
			document: &types.DIDDocumentWithSeq{},
			wantErr:  true,
		},
	}

	for _, testCase := range testCases {
		t.Run(testCase.name, func(t *testing.T) {
			state := types.GenesisState{
				Documents: map[string]*types.DIDDocumentWithSeq{
					did: testCase.document,
				},
			}

			err := state.Validate()
			if testCase.wantErr {
				require.ErrorIs(t, err, types.ErrInvalidDIDDocumentWithSeq)
				return
			}
			require.NoError(t, err)
		})
	}
}
