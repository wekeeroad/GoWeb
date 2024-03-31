package v1

import "github.com/gin-gonic/gin"

type Tag struct{}

func NewTag() Tag {
	return Tag{}
}

func (t Tag) Get(c *gin.Context) {}

//	@Summary	Get multiple tags
//  @Tags       tag
//	@Accept		json
//	@Produce	json
//	@Param		name		query		string			false	"Name of tag "	maxlength(100)
//	@Param		state		query		int				false	"states"		Enums(0,1)	default(1)
//	@Param		page		query		int				false	"page"
//	@Param		page_size	query		int				false	"number per page"
//	@Success	200			{object}	model.Tag		"success"
//	@Failure	400			{object}	errcode.Error	"requirment error"
//	@Failure	500			{object}	errcode.Error	"internal error"
//	@Router		/api/v1/tags [get]
func (t Tag) List(c *gin.Context) {}

//	@Summary	Create tag
//  @Tags       tag
//	@Accept		json
//	@Produce	json
//	@Param		name		body		string			true	"name of tag"	minlength(3)	maxlengh(100)
//	@Param		state		body		int				false	"states"		Enums(0,1)		default(1)
//	@Param		created_by	body		string			true	"creator"		minlength(3)	maxlength(100)
//	@Success	200			{object}	model.Tag		"success"
//	@Failure	400			{object}	errcode.Error	"requirment error"
//	@Failure	500			{object}	errcode.Error	"internal error"
//	@Router		/api/v1/tags [post]
func (t Tag) Create(c *gin.Context) {}

//	@Summary	Update tag
//  @Tags       tag
//	@Accept		json
//	@Produce	json
//	@Param		id			path		int				true	"id of tag"
//	@Param		name		body		string			false	"name of tag"	minlength(3)	maxlength(100)
//	@Param		state		body		int				false	"states"		Enums(0 ,1)		default(1)
//	@Param		modified_by	body		string			true	"editor"		minlength(3)	maxlength(100)
//	@Success	200			{array}		model.Tag		"success"
//	@Failure	400			{object}	errcode.Error	"requirment error"
//	@Failure	500			{object}	errcode.Error	"internal error"
//	@Router		/api/v1/tags/{id} [put]
func (t Tag) Update(c *gin.Context) {}

//	@Summary		Delete tag
//	@Description	delete tag
//	@Tags			tag
//	@accept			json
//	@Produce		json
//	@Param			id	path		int				true	"id of tag"
//	@Success		200	{string}	string			"success"
//	@Failure		400	{object}	errcode.Error	"requirment error"
//	@Failure		500	{object}	errcode.Error	"internal error"
//	@Router			/api/v1/tags/{id} [delete]
func (t Tag) Delete(c *gin.Context) {}
