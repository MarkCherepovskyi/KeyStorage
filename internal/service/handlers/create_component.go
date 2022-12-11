package handlers

import "C"
import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/requests"
	"gitlab.com/MarkCherepovskyi/KeyStorage/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"io"

	"github.com/tyler-smith/go-bip32"
	"github.com/tyler-smith/go-bip39"
	"net/http"
	"strings"
)

func CreateContainer(w http.ResponseWriter, r *http.Request) {

	request, err := requests.NewCreateContainerRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	key := new(bip32.Key)
	seed := []byte("")

	if request.Data.Attributes.Key == "" {
		// Generate a mnemonic for memorization or user-friendly seeds
		entropy, _ := bip39.NewEntropy(256)
		mnemonic, _ := bip39.NewMnemonic(entropy)

		// Generate a Bip32 HD wallet for the mnemonic and a user supplied password
		seed = bip39.NewSeed(mnemonic, "Secret Passphrase")

		key, _ = bip32.NewMasterKey(seed)
	}
	if strings.Contains(request.Data.Attributes.Key, " ") {

		seed = bip39.NewSeed(request.Data.Attributes.Key, "TREZOR")

		key, _ = bip32.NewMasterKey(seed)

	}

	encodedContainer := []byte("")

	if key.Key != nil {

		hash := sha256.Sum256([]byte(key.String()))

		encodedContainer, err = data.Encryption([]byte(request.Data.Attributes.Text), hash[:])
		if err != nil {
			Log(r).WithError(err).Error("failed to create blob in DB")
			ape.Render(w, problems.InternalError())
			return
		}
	} else {
		hash := sha256.Sum256([]byte(request.Data.Attributes.Key))
		encodedContainer, err = data.Encryption([]byte(request.Data.Attributes.Text), hash[:])
		if err != nil {
			Log(r).WithError(err).Error("failed to create blob in DB")
			ape.Render(w, problems.InternalError())
			return
		}
	}

	//group := false
	//if len(request.Data.Attributes.Recipient) > 0 {
	//	group = true
	//}

	container := data.Container{
		Owner:     request.Data.Attributes.Owner,
		Recipient: request.Data.Attributes.Recipient,
		Container: encodedContainer,
		Tag:       request.Data.Attributes.Tag,
	}

	err = ContainerQ(r).Insert(container)
	if err != nil {
		Log(r).WithError(err).Error("failed to create blob in DB")
		ape.Render(w, problems.InternalError())
		return
	}

	result := resources.ContainerResponse{
		Data: newContainerModel(container),
	}
	ape.Render(w, result)
}

func md(str string) string {
	h := md5.New()
	io.WriteString(h, str)

	return fmt.Sprintf("%x", h.Sum(nil))
}
