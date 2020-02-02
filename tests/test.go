package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/hex"
	"encoding/json"
	"errors"
	"fmt"
	"math/big"

	"github.com/jneubaum/honestvote/core/core-database/database"
)

func main() {
	fuckthis := "election1position1candidate2"
	public_key := database.PublicKey("0437931ff7885e8ba8465ab5a725b31624f208aab3143c8734f8ac38f5f1f0642a922cfac4ecceda4458efd2d634fb2359f687b19ca587196ce3b92edce7b04124")
	private_key := "30450220038c9d8cd5bc75f66ffb735b87a7efb78646c9a44e00e8776eabeba4f6e71c0f022100d1d057acffa94e49c34a8397a7f256b96a362ddcbd6f1ec66d6e90c9666dc897"
	valid, err := Verify([]byte(fuckthis), public_key, private_key)
	if valid {
		fmt.Println(valid)
	} else {
		fmt.Println(err)
	}
}

func Verify(hash []byte, public_key_hex database.PublicKey, signature_hex string) (result bool, err error) {
	// decode public key from hex
	public_key_bytes, err := hex.DecodeString(string(public_key_hex))
	if err != nil {
		return false, err
	}

	// x509 parse public key
	public_key, err := x509.ParsePKIXPublicKey(public_key_bytes)
	if err != nil {
		return false, err
	}

	// check that parse key is ecdsa.PublicKey
	switch public_key := public_key.(type) {
	case *ecdsa.PublicKey:
		// decode signature json from hex
		signature_json, err := hex.DecodeString(signature_hex)
		if err != nil {
			return false, err
		}

		// unmarhsal signature structure to extract signature from
		signature := new(signature)
		err = json.Unmarshal(signature_json, signature)
		if err != nil {
			return false, err
		}

		// verify signature
		return ecdsa.Verify(public_key, hash, signature.R, signature.S), nil

	default:
		// only ECDSA public keys are supported
		return false, errors.New("only ECDSA public keys supported")
	}
}

type signature struct {
	R,
	S *big.Int
}
