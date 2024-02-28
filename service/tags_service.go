package service

import (
	"github.com/aungkoko1234/gin-crud/data/request"
	"github.com/aungkoko1234/gin-crud/data/response"
)

type TagsService interface {
	Create(tags request.CreateTagsRequest)
	Update(tags request.UpdateTagsRequest)
	Delete (tagsId int)
	FindById(tagsId int) response.TagsResponse
	FindAll() []response.TagsResponse
}