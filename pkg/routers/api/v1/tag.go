package v1

import (
	"github.com/gin-gonic/gin"
	"github.com/wekeeroad/GoWeb/global"
	"github.com/wekeeroad/GoWeb/pkg/app"
	"github.com/wekeeroad/GoWeb/pkg/convert"
	"github.com/wekeeroad/GoWeb/pkg/errcode"
	"github.com/wekeeroad/GoWeb/pkg/service"
)

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

//		@Summary	Get multiple tags
//	 @Tags       tag
//		@Produce	json
//		@Param		name		query		string			false	"Name of tag "	maxlength(100)
//		@Param		state		query		int				false	"states"		Enums(0,1)	default(1)
//		@Param		page		query		int				false	"page"
//		@Param		page_size	query		int				false	"number per page"
//		@Success	200			{object}	model.TagSwagger		"success"
//		@Failure	400			{object}	errcode.Error	"requirment error"
//		@Failure	500			{object}	errcode.Error	"internal error"
//		@Router		/api/v1/tags [get]
func (t Tag) List(c *gin.Context) {
	param := service.TagListRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindandValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	pager := app.Pager{Page: app.GetPage(c), PageSize: app.GetPageSize(c)}
	totalRows, err := svc.CountTag(&service.CountTagRequest{Name: param.Name, State: param.State})
	if err != nil {
		global.Logger.Errorf("svc.CountTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCountTagFail)
		return
	}

	tags, err := svc.GetTagList(&param, &pager)
	if err != nil {
		global.Logger.Errorf("svc.GetTagList err: %v", err)
		response.ToErrorResponse(errcode.ErrorGetTagListFail)
		return
	}

	response.ToResponseList(tags, totalRows)
	return
}

//		@Summary	Create tag
//	 @Tags       tag
//		@Produce	json
//		@Param		name		body		string			true	"name of tag"	minlength(3)	maxlengh(100)
//		@Param		state		body		int				false	"states"		Enums(0,1)		default(1)
//		@Param		created_by	body		string			true	"creator"		minlength(3)	maxlength(100)
//		@Success	200			{object}	model.Tag		"success"
//		@Failure	400			{object}	errcode.Error	"requirment error"
//		@Failure	500			{object}	errcode.Error	"internal error"
//		@Router		/api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {
	param := service.CreateTagRequest{}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindandValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.CreateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.CreateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorCreateTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}

//		@Summary	Update tag
//	 @Tags       tag
//		@Produce	json
//		@Param		id			path		int				true	"id of tag"
//		@Param		name		body		string			false	"name of tag"	minlength(3)	maxlength(100)
//		@Param		state		body		int				false	"states"		Enums(0 ,1)		default(1)
//		@Param		modified_by	body		string			true	"editor"		minlength(3)	maxlength(100)
//		@Success	200			{array}		model.Tag		"success"
//		@Failure	400			{object}	errcode.Error	"requirment error"
//		@Failure	500			{object}	errcode.Error	"internal error"
//		@Router		/api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {
	param := service.UpdateTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindandValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.UpdateTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.UpdateTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorUpdateTagFail)
		return
	}
	response.ToResponse(gin.H{})
	return
}

// @Summary		Delete tag
// @Description	delete tag
// @Tags			tag
// @Produce		json
// @Param			id	path		int				true	"id of tag"
// @Success		200	{string}	string			"success"
// @Failure		400	{object}	errcode.Error	"requirment error"
// @Failure		500	{object}	errcode.Error	"internal error"
// @Router			/api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {
	param := service.DeleteTagRequest{
		ID: convert.StrTo(c.Param("id")).MustUInt32(),
	}
	response := app.NewResponse(c)
	valid, errs := app.BindAndValid(c, &param)
	if !valid {
		global.Logger.Infof("app.BindandValid errs: %v", errs)
		errRsp := errcode.InvalidParams.WithDetails(errs.Errors()...)
		response.ToErrorResponse(errRsp)
		return
	}

	svc := service.New(c.Request.Context())
	err := svc.DeleteTag(&param)
	if err != nil {
		global.Logger.Errorf("svc.DeleteTag err: %v", err)
		response.ToErrorResponse(errcode.ErrorDeleteTagFail)
		return
	}

	response.ToResponse(gin.H{})
	return
}
