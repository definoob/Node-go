package main

import (
	"os"
	"math/big"
	"testing"

	"github.com/vordev/VOR/core/adapters"
	"github.com/vordev/VOR/core/services/vrf"
	"github.com/vordev/VOR/core/store/models"
	"github.com/vordev/VOR/core/utils"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/common/hexutil"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func main() {
	/*adapter := adapters.Random{PublicKey: "0xce3cc486a3aa567a2e15707cd5b874170e746f4db45ec084918e5b2b5ec6c6cf01"}
	hash := utils.MustHash("a random string")
	seed := big.NewInt(0x10)
	blockNum := 10
	jsonInput, err := models.JSON{}.MultiAdd(models.KV{
		"seed":      utils.Uint64ToHex(seed.Uint64()),
		"keyHash":   publicKey.MustHash().Hex(),
		"blockHash": hash.Hex(),
		"blockNum":  blockNum,
	})
	require.NoError(t, err) // Can't fail
	input := models.NewRunInput(&models.ID{}, models.ID{}, jsonInput,
		models.RunStatusUnstarted)
	result := adapter.Perform(*input, store)
	require.NoError(t, result.Error(), "while running random adapter")*/
}
