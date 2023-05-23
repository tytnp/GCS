package base

import (
	"admin-service/entity"
	"admin-service/model/base"
	baseService "admin-service/service/base"
	"github.com/gin-gonic/gin"
	"github.com/gin-gonic/gin/binding"
	"net/http"
	"reflect"
)

type Api[T base.GcsModel] struct {
	RouterGroup *gin.RouterGroup
	Service     baseService.BaseService[T]
}

func (api *Api[T]) Register() {

}

func (api *Api[T]) Default(r *gin.RouterGroup) {
	api.RouterGroup = r.Group((*new(T)).ModelName())
	api.RouterGroup.POST("/add", api.Add)
	api.RouterGroup.POST("/del", api.Del)
	api.RouterGroup.POST("/edit", api.Edit)
	api.RouterGroup.POST("/list", api.List)
	api.RouterGroup.POST("/one", api.One)
}

func (api *Api[T]) Add(c *gin.Context) {
	ret := new(entity.ApiRet)
	model := new(T)
	if ret.Error = c.ShouldBindBodyWith(&model, binding.JSON); ret.Error == nil {
		if ret.Error = api.Service.Insert(model); ret.Error == nil {
			ret.Ok = true
		}
	}
	ret.Data = model
	c.JSON(http.StatusOK, ret)
}

func (api *Api[T]) Del(c *gin.Context) {
	ret := new(entity.ApiRet)
	model := new(T)
	if ret.Error = c.ShouldBindBodyWith(&model, binding.JSON); ret.Error == nil {
		if ret.Error = api.Service.Delete(model); ret.Error == nil {
			ret.Ok = true
		}
	}
	ret.Data = model
	c.JSON(http.StatusOK, ret)
}

func (api *Api[T]) Edit(c *gin.Context) {
	ret := new(entity.ApiRet)
	model := new(T)
	if ret.Error = c.ShouldBindBodyWith(&model, binding.JSON); ret.Error == nil {
		if ret.Error = api.Service.Update(model); ret.Error == nil {
			ret.Ok = true
		}
	}
	ret.Data = model
	c.JSON(http.StatusOK, ret)
}

func (api *Api[T]) List(c *gin.Context) {
	ret := new(entity.ApiRet)
	model := new(T)
	if ret.Error = c.ShouldBindBodyWith(&model, binding.JSON); ret.Error == nil {
		props := reflect.ValueOf(*model).FieldByName("ExpandProps").Interface().(map[string]interface{})
		pageValue, p := props["page"].(float64)
		pageSizeValue, ps := props["pageSize"].(float64)
		if p && ps {
			ret.Page = int(pageValue)
			ret.PageSize = int(pageSizeValue)
		}
		if ret.Data, ret.TotalCount, ret.Error = api.Service.FindList(model); ret.Error == nil {
			ret.Ok = true
		}

	}
	c.JSON(http.StatusOK, ret)
}

func (api *Api[T]) One(c *gin.Context) {
	ret := new(entity.ApiRet)
	model := new(T)
	if ret.Error = c.ShouldBindBodyWith(&model, binding.JSON); ret.Error == nil {
		if ret.Data, ret.Error = api.Service.FindOne(model); ret.Error == nil {
			ret.Ok = true
		}
	}
	c.JSON(http.StatusOK, ret)
}
