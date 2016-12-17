package handlers

import (
	"net/http"
	"github.com/sansaralab/buddha/helpers"
)

type ProjectHandler struct {
}

func NewProjectHandler() *ProjectHandler {
	return &ProjectHandler{}
}

func (h *ProjectHandler) ServeHTTP(resp http.ResponseWriter, req *http.Request) {
	helpers.JsonResp(resp, []string{"asdasd", "sddsfsdf"}, 200)
}
