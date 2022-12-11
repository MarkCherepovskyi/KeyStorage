package requests

import (
	"encoding/json"
	"gitlab.com/MarkCherepovskyi/KeyStorage/resources"
	"net/http"
)

type CreateContainerRequest struct {
	Data resources.Container
}

func NewCreateContainerRequest(r *http.Request) (CreateContainerRequest, error) {
	request := CreateContainerRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err

}
