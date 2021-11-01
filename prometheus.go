package prometheus

import (
	"encoding/json"
	"net/http"
	"strings"

	"github.com/portainer/libcrypto"
	"github.com/portainer/libhttp/response"
	portainer "github.com/portainer/portainer/api"
	"github.com/portainer/portainer/api/http/client"
)

type prometheusResponse struct {
	Title         string            `json:"Title"`
	Message       string            `json:"Message"`
	ContentLayout map[string]string `json:"ContentLayout"`
	Style         string            `json:"Style"`
	Hash          []byte            `json:"Hash"`
}

type prometheusData struct {
	Title         string            `json:"title"`
	Message       []string          `json:"message"`
	ContentLayout map[string]string `json:"contentLayout"`
	Style         string            `json:"style"`
}

// @id PROMETHEUS
// @summary fetches the data to be used by PROMETHEUS
// @description **Access policy**: restricted
// @tags prometheus
// @security jwt
// @produce json
// @success 200 {object} JSONResponse
// @router /prometheus [get]
func (handler *Handler) prometheus(w http.ResponseWriter, r *http.Request) {
	prometheus, err := client.Get(portainer.MessageOfTheDayURL, 0)
	if err != nil {
		response.JSON(w, &prometheusResponse{Message: ""})
		return
	}

	var data prometheusData
	err = json.Unmarshal(prometheus, &data)
	if err != nil {
		response.JSON(w, &prometheusResponse{Message: ""})
		return
	}

	message := strings.Join(data.Message, "\n")

	hash := libcrypto.HashFromBytes([]byte(message))
	resp := prometheusResponse{
		Title:         data.Title,
		Message:       message,
		Hash:          hash,
		ContentLayout: data.ContentLayout,
		Style:         data.Style,
	}

	response.JSON(w, &resp)
}
