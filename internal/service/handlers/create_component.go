package handlers

import "C"
import (
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/requests"
	"gitlab.com/MarkCherepovskyi/KeyStorage/resources"
	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
	"net/http"
)

func CreateContainer(w http.ResponseWriter, r *http.Request) {

	request, err := requests.NewCreateContainerRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	encodedContainer, err := data.Encryption([]byte(request.Data.Attributes.Text), []byte(request.Data.Attributes.Key))
	group := false
	if len(request.Data.Attributes.Recipient) > 0 {
		group = true
	}

	container := data.Container{
		Owner:     request.Data.Attributes.Owner,
		Recipient: request.Data.Attributes.Recipient,
		Group:     group,
		Container: string(encodedContainer),
		Tag:       request.Data.Attributes.Tag,
	}

	container.ID, err = ContainerQ(r).Insert(container)
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
