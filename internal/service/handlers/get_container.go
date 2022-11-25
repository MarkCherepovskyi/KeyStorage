package handlers

import (
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data"
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
		Log(r).WithError(err).Error("failed to create tx")
		ape.Render(w, problems.InternalError())
		return
	}

	if *request.Data.Attributes.Key == "" {
		//generate new kay
	}

	id, err := strconv.Atoi(request.Data.ID)
	if err != nil {
		Log(r).WithError(err).Error("failed to create tx")
		ape.Render(w, problems.InternalError())
		return
	}
	container, err := ContainerQ(r).FilterByID(int64(id)).Get()
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if container == nil {
		ape.Render(w, problems.NotFound())
		return
	}

	bufferContainer, err := data.Decryption([]byte(*request.Data.Attributes.Text), []byte(*request.Data.Attributes.Key))
	if err != nil {
		Log(r).WithError(err).Error("failed to get blob from DB")
		ape.Render(w, problems.InternalError())
		return
	}
	if err != nil {
		Log(r).WithError(err).Error("failed to create tx")
		ape.Render(w, problems.InternalError())
		return
	}

	container.Container = string(bufferContainer)
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
