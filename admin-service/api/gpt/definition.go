package gpt

import api "admin-service/api/base"
import model "admin-service/model/gpt"
import service "admin-service/service/gpt"

type GptUserApi struct {
	api.Api[model.GptUser]
	service.GptUserService
}

func init() {
	api.ApiContext = append(api.ApiContext, &GptUserApi{})
}
