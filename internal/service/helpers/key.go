package helpers

import (
	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"strings"
)

func Gene(requestKey string) *bip32.Key {
	key := new(bip32.Key)
	seed := []byte("")

	if requestKey == "" {
		// Generate a mnemonic for memorization or user-friendly seeds
		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ := bip39.NewMnemonic(entropy)

		// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
		seed = bip39.NewSeed(mnemonic, "Secret Passphrase")

		key, _ = bip32.NewMasterKey(seed)
	}
	if strings.Contains(requestKey, " ") {

		seed = bip39.NewSeed(requestKey, "TREZOR")

		key, _ = bip32.NewMasterKey(seed)

	}
	return key

}
