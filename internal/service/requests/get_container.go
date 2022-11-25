package requests

import (
	"encoding/json"
	"gitlab.com/MarkCherepovskyi/KeyStorage/resources"
	"net/http"
)

type GetContainerRequest struct {
	Data resources.Container
}

func NewGetContainerRequest(r *http.Request) (GetContainerRequest, error) {
	request := GetContainerRequest{}
	err := json.NewDecoder(r.Body).Decode(&request)
	return request, err

}
