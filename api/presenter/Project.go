package presenter

import (
	"porto-project/pkg/model"
)

type ProjectSuccessResponse struct {
    Status string `json:"status"`
	Message string            `json:"message"`
	Project *model.Project `json:"project"`
}

type ProjectsSuccessResponse struct {
    Status string `json:"status"`
	Message string `json:"message"`
	LastIntId int `json:"lastIntId"`
	PageCount int64            `json:"pageCount"`
	Project []model.Project `json:"projects"`
}

type FailedResponse struct {
    Status string `json:"status"`
	Message string `json:"message"`
	Error string `json:"error"`
}
