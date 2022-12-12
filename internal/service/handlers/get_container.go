package handlers

import (
	"crypto/sha256"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/helpers"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/requests"
	"gitlab.com/MarkCherepovskyi/KeyStorage/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
	"strconv"
)

func GetContainer(w http.ResponseWriter, r *http.Request) {
	request, err := requests.NewGetContainerRequest(r)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create tx")
		ape.Render(w, problems.InternalError())
		return
	}

	id, err := strconv.Atoi(request.Data.ID)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create tx")
		ape.Render(w, problems.InternalError())
		return
	}

	container, err := helpers.ContainerQ(r).FilterByID(int64(id)).Get()
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to get blob from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if container == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	key := helpers.Gene(request.Data.Attributes.Key)

	var bufferContainer []byte

	if key.Key != nil {

		hash := sha256.Sum256([]byte(key.String()))

		bufferContainer, err = data.Decryption([]byte(request.Data.Attributes.Text), hash[:])
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to create blob in DB")
			ape.Render(w, problems.InternalError())
			return
		}
	} else {
		hash := sha256.Sum256([]byte(request.Data.Attributes.Key))
		bufferContainer, err = data.Decryption(container.Container, hash[:])
		if err != nil {
			helpers.Log(r).WithError(err).Error("failed to create blob in DB")
			ape.Render(w, problems.InternalError())
			return
		}
	}

	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create tx")
		ape.Render(w, problems.InternalError())
		return
	}

	container.Container = bufferContainer
	result := resources.ContainerResponse{
		Data: newContainerModel(*container),
	}
	ape.Render(w, result)
}

func newContainerModel(container data.Container) resources.Container {

	result := resources.Container{
		Key: resources.NewKeyInt64(container.ID, resources.CONTAINER),
		Attributes: resources.ContainerAttributes{
			Container: container.Container,
			Tag:       container.Tag,
		},
	}
	return result

}
