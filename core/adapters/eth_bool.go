package adapters

import (
	"github.com/vordev/VOR/core/store"
	"github.com/vordev/VOR/core/store/models"

	"github.com/tidwall/gjson"
)

var evmFalse = "0x0000000000000000000000000000000000000000000000000000000000000000"
var evmTrue = "0x0000000000000000000000000000000000000000000000000000000000000001"

// EthBool holds no fields
type EthBool struct{}

// TaskType returns the type of Adapter.
func (e *EthBool) TaskType() models.TaskType {
	return TaskTypeEthBool
}

// Perform returns the abi encoding for a boolean
//
// For example, after converting the result false to hex encoded Ethereum
// ABI, it would be:
// "0x0000000000000000000000000000000000000000000000000000000000000000"
func (*EthBool) Perform(input models.RunInput, _ *store.Store) models.RunOutput {
	if boolean(input.Result().Type) {
		return models.NewRunOutputCompleteWithResult(evmTrue)
	}

	return models.NewRunOutputCompleteWithResult(evmFalse)
}

func boolean(t gjson.Type) bool {
	switch t {
	case gjson.False, gjson.Null:
		return false
	default:
		return true
	}
}
