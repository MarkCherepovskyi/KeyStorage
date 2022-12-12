package handlers

import "C"
import (
	"crypto/md5"
	"crypto/sha256"
	"fmt"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/helpers"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/requests"
	"gitlab.com/MarkCherepovskyi/KeyStorage/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"io"

	"net/http"
)

func CreateContainer(w http.ResponseWriter, r *http.Request) {

	request, err := requests.NewCreateContainerRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Info("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	key := helpers.Gene(request.Data.Attributes.Key)

	encodedContainer := []byte("")

	if key.Key != nil {

		hash := sha256.Sum256([]byte(key.String()))

		encodedContainer, err = data.Encryption([]byte(request.Data.Attributes.Text), hash[:])
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to create blob in DB")
			ape.Render(w, problems.InternalError())
			return
		}
	} else {
		hash := sha256.Sum256([]byte(request.Data.Attributes.Key))
		encodedContainer, err = data.Encryption([]byte(request.Data.Attributes.Text), hash[:])
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to create blob in DB")
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

	err = helpers.ContainerQ(r).Insert(container)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create blob in DB")
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
