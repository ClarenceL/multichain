package elastos

import (
	"fmt"
	"github.com/itchyny/base58-go"
	"math/big"

	"github.com/renproject/multichain/api/address"
	"github.com/renproject/pack"
)


var _ = fmt.Printf // DEBUG: delete when done

/*
A ELA standard account is a set of private key, public key, redeem script, program hash
and address data.

The pub key is converted into a program hash, which is the sha256 value of redeem script
and converted to ripemd160 format with a (Type) prefix - as given by PrefixType.

Elastos has a series of address prefixes:

const (
	PrefixStandard   PrefixType = 0x21
	PrefixMultiSig   PrefixType = 0x12
	PrefixCrossChain PrefixType = 0x4B
	PrefixDeposit    PrefixType = 0x1F
	PrefixCRDID      PrefixType = 0x67
)
 */

type AddressEncoder struct {

}

type AddressDecoder struct {

}

/*
Converts a raw bytes address into a human-readable address

TODO: I assume there is a prefix?

implements the address.Encoder interface

 */
func (encoder AddressEncoder) EncodeAddress(rawAddr address.RawAddress) (address.Address, error) {

	bi := new(big.Int).SetBytes(rawAddr).String()

	encoded, err := base58.BitcoinEncoding.Encode([]byte(bi))

	return address.Address(string(encoded)), err
}

func (decoder AddressDecoder) DecodeAddress(addr address.Address) (pack.Bytes, error) {


	return nil, nil
}