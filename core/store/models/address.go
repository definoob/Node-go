package models

import (
	"database/sql/driver"
	"encoding/hex"
	"fmt"
	"math/big"
	"strings"

	"github.com/ethereum/go-ethereum/common"
	"github.com/vordev/VOR/core/logger"
	"github.com/vordev/VOR/core/utils"
)

// EIP55Address is a newtype for string which persists an ethereum address in
// its original string representation which includes a leading 0x, and EIP55
// checksum which is represented by the case of digits A-F.
type EIP55Address string

// NewEIP55Address creates an EIP55Address from a string, an error is returned if:
//
// 1) There is no leading 0x
// 2) The length is wrong
// 3) There are any non hexadecimal characters
// 4) The checksum fails
//
func NewEIP55Address(s string) (EIP55Address, error) {
	address := common.HexToAddress(s)
	if s != address.Hex() {
		return EIP55Address(""), fmt.Errorf(`"%s" is not a valid EIP55 formatted address`, s)
	}
	return EIP55Address(s), nil
}

// Bytes returns the raw bytes
func (a EIP55Address) Bytes() []byte { return a.Address().Bytes() }

// Big returns a big.Int representation
func (a EIP55Address) Big() *big.Int { return a.Address().Hash().Big() }

// Hash returns the Hash
func (a EIP55Address) Hash() common.Hash { return a.Address().Hash() }

// Address returns EIP55Address as a go-ethereum Address type
func (a EIP55Address) Address() common.Address { return common.HexToAddress(a.String()) }

// String implements the stringer interface and is used also by the logger.
func (a EIP55Address) String() string {
	return string(a)
}

// Hex is idential to String but makes the API similar to common.Address
func (a EIP55Address) Hex() string {
	return a.String()
}

// Format implements fmt.Formatter
func (a EIP55Address) Format(s fmt.State, c rune) {
	_, err := fmt.Fprint(s, a.String())
	logger.ErrorIf(err, "failed when format EIP55Address to state")
}

// UnmarshalText parses a hash from plain text
func (a *EIP55Address) UnmarshalText(input []byte) error {
	var err error
	*a, err = NewEIP55Address(string(input))
	return err
}

// UnmarshalJSON parses a hash from a JSON string
func (a *EIP55Address) UnmarshalJSON(input []byte) error {
	input = utils.RemoveQuotes(input)
	return a.UnmarshalText(input)
}

// Value returns this instance serialized for database storage.
func (a EIP55Address) Value() (driver.Value, error) {
	return a.Bytes(), nil

}

// Scan reads the database value and returns an instance.
func (a *EIP55Address) Scan(value interface{}) error {
	switch v := value.(type) {
	case string:
		*a = EIP55Address(v)
	case []byte:
		address := common.HexToAddress("0x" + hex.EncodeToString(v))
		*a = EIP55Address(address.Hex())
	default:
		return fmt.Errorf("unable to convert %v of %T to EIP55Address", value, value)
	}
	return nil
}

// EIP55AddressCollection is an array of EIP55Addresses.
type EIP55AddressCollection []EIP55Address

// Value returns this instance serialized for database storage.
func (c EIP55AddressCollection) Value() (driver.Value, error) {
	// Unable to convert copy-free without unsafe:
	// https://stackoverflow.com/a/48554123/639773
	converted := make([]string, len(c))
	for i, e := range c {
		converted[i] = string(e)
	}
	return strings.Join(converted, ","), nil
}

// Scan reads the database value and returns an instance.
func (c *EIP55AddressCollection) Scan(value interface{}) error {
	temp, ok := value.(string)
	if !ok {
		return fmt.Errorf("unable to convert %v of %T to EIP55AddressCollection", value, value)
	}

	arr := strings.Split(temp, ",")
	collection := make(EIP55AddressCollection, len(arr))
	for i, r := range arr {
		collection[i] = EIP55Address(r)
	}
	*c = collection
	return nil
}
