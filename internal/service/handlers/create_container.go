package handlers

import (
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/data"
	"gitlab.com/MarkCherepovskyi/KeyStorage/internal/service/requests"
	"gitlab.com/MarkCherepovskyi/KeyStorage/resources"
	"net/http"

	"gitlab.com/distributed_lab/ape"
	"gitlab.com/distributed_lab/ape/problems"
)

func CreateContainer(w http.ResponseWriter, r *http.Request) {

	request, err := requests.NewCreateContainerRequest(r)
	if err != nil {
		Log(r).WithError(err).Info("invalid request")
		ape.RenderErr(w, problems.BadRequest(err)...)
		return
	}

	blob := data.Container{
		Owner:     request.Data.Attributes.Owner,
		Recipient: request.Data.Attributes.Recipient,
	}

	blob.ID, err = helpers.BlobsQ(r).Insert(blob)
	if err != nil {
		helpers.Log(r).WithError(err).Error("failed to create blob in DB")
		ape.Render(w, problems.InternalError())
		return
	}

	result := resources.BlobResponse{
		Data: newBlobModel(blob),
	}
	ape.Render(w, result)
}
