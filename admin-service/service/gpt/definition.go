package gpt

import service "admin-service/service/base"
import model "admin-service/model/gpt"

type GptUserService struct {
	service.BaseService[model.GptUser]
}
