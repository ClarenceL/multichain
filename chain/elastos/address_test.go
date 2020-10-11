package elastos

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	. "github.com/elastos/Elastos.ELA/common"
	"github.com/elastos/Elastos.ELA/core/contract"
	"github.com/elastos/Elastos.ELA/crypto"
	"github.com/itchyny/base58-go"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"golang.org/x/crypto/ripemd160"
	"math/big"
)

var _ = fmt.Printf // DEBUG: delete when done

var _ = Describe("Elastos - Address", func() {
	Context("When converting raw addresses to human readable addresses with Elastos PrefixType", func() {

		It("Sanity Check: Full encode of pub key into a human-readable address with all intermediates", func() {

			// this starts with the basic keypair, but we want a known value
			// crypto.GenerateKeyPair()

			// normally this uncompressed pubkey goes straight into CreateStandardContract, but since the pubkey
			// we usually use is compressed we need to decompress it

			// this is compressed!
			// publicKeyHex := "030837acfaecb5ea16fd9fc0dca7dbc7f745c7fa882e0c87332aa025478f1a95e0"
			publicKeyHex := "022c9652d3ad5cc065aa9147dc2ad022f80001e8ed233de20f352950d351d472b7"

			fmt.Printf("Starting Pub Key compressed: %v\n", publicKeyHex)

			// returns byte[] from hex string
			publicKey, err := hex.DecodeString(publicKeyHex)

			Expect(err).ToNot(HaveOccurred())

			// to arrive back at the raw pubkey from crypto.GenerateKeyPair() we need to DecodePoint
			// which is reversing the point compression
			pub, _ := crypto.DecodePoint(publicKey)

			fmt.Printf("Decompressed point X: %v, Y: %v\n", pub.X, pub.Y)

			// recompress for testing
			compressedPubKey, _ := pub.EncodePoint(true)

			fmt.Printf("Recompressed point: %x\n", compressedPubKey)

			Expect(hex.EncodeToString(compressedPubKey)).To(Equal(publicKeyHex))

			/*
			a contract in Elastos sense is
			type Contract struct {
				Code   []byte
				Prefix PrefixType
			}

			here calling `CreateStandardContract` sets Prefix to PrefixType.PrefixStandard

			"Code" is the redeemScript, this implementation is left out, but it's an encoding of the pubkey with an op CHECKSIG at the end
			(the pubkey fully determines the redeemScript, the PrefixType is not included)

			This also does an EncodePoint/compression for the pub key to arrive at "Code", this is a COMPRESSEDLEN of 33 bytes + 1 byte for CHECKSIG
			 */

			// next line does this pubKey.EncodePoint(true) - compression: true
			ct, err := contract.CreateStandardContract(pub)
			// ct, err = contract.CreateCRIDContractByCode(ct.Code) // no effect?
			Expect(err).ToNot(HaveOccurred())

			fmt.Printf("redeemScript + op.CHECKSIG: %x - len: %d\n", ct.Code, len(ct.Code))

			Expect(len(ct.Code) == 34)

			// this line is expanded below
			// programHash := ct.ToProgramHash()

			// NOTE: code is []byte) Uint168

			hash := sha256.Sum256(ct.Code)

			md160 := ripemd160.New()
			md160.Write(hash[:])

			// fmt.Printf("data - orig %v - len: %d\n", md160.Write(), md160.Size())

			programHash := Uint168{}

			// Prefix does make its way into the address
			// copy(programHash[:], md160.Sum(nil))
			copy(programHash[:], md160.Sum([]byte{byte(contract.PrefixStandard)}))

			// this line is expanded below
			// prettyAddress, _ := programHash.ToAddress()

			data := programHash.Bytes()

			fmt.Printf("data: %v - len: %d\n", data, len(data))
			fmt.Printf("data hex: %x - len: %d\n", data, len(hex.EncodeToString(data)))

			checksum := Sha256D(data)

			fmt.Printf("checksum: %v\n", checksum[0:4])

			// ELA like BTC uses a checksum
			data = append(data, checksum[0:4]...)

			fmt.Printf("data new: %v - len: %d\n", data, len(data))

			bi := new(big.Int).SetBytes(data).String()

			fmt.Printf("data int: %v - bytes: %v, len: %d\n", bi, []byte(bi), len(bi))

			encoded, err := base58.BitcoinEncoding.Encode([]byte(bi))

			Expect(err).ToNot(HaveOccurred())

			fmt.Printf("address: %s", encoded)

			// Expect(string(encoded)).To(Equal("EZjQkZs8QswFSYtACHAeSQvKxbBwsnSe1x"))

			Expect(string(encoded)).To(Equal("ENTogr92671PKrMmtWo3RLiYXfBTXUe13Z"))

		})

		It("Should convert a pub key into a ripemd160", func() {

		})

	})
})