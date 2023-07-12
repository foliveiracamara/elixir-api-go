package controller

import (
	"fmt"
	"net/http"

	"github.com/foliveiracamara/elixir-api-go/src/adapter/controller/dto/request"
	"github.com/foliveiracamara/elixir-api-go/src/adapter/controller/dto/response"
	"github.com/foliveiracamara/elixir-api-go/src/adapter/view"
	"github.com/foliveiracamara/elixir-api-go/src/config/logger"
	"github.com/foliveiracamara/elixir-api-go/src/config/validation"
	"github.com/foliveiracamara/elixir-api-go/src/domain/entity"
	"github.com/foliveiracamara/elixir-api-go/src/domain/service"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
)

var (
	PostDomainInterface entity.PostDomainInterface
)

type PostControllerInterface interface {
	ListPosts(c *gin.Context)
	CreatePost(c *gin.Context) 
}

type postControllerInterface struct {
	service service.PostDomainService
}

func NewPostControllerInterface(svc service.PostDomainService) PostControllerInterface {
	return &postControllerInterface{
		service: svc,
	}
}


// func GetPostById(ctx *gin.Context) error {
// 	id := ctx.Param("id")
// 	// for _, post := range entity.Posts {
// 	// 	if post.PostId == id {
// 	// 		return ctx.JSON(http.StatusOK, entity.Posts)
// 	// 	}
// 	// }
// 	return ctx.JSON(http.StatusNotFound, nil)
// }

func (p *postControllerInterface) CreatePost(ctx *gin.Context) {
	var postRequest request.PostRequest
	if err := ctx.ShouldBindJSON(&postRequest); err != nil {
		logger.Error(
			"Error trying to validate post info", err,
			zap.String("journey", "createPostController"),
		)
		appErr := validation.ValidateDonorError(err)
		ctx.JSON(appErr.Code, appErr)
	}

	domain := entity.NewPostDomain(
		postRequest.PostId,
		postRequest.Description,
		postRequest.PostDate,
		postRequest.Receiver,
	)

	if err := p.service.CreatePost(domain); err != nil {
		ctx.JSON(err.Code, err)
	}
	logger.Info("Post created successfully",
		zap.String("journey", "createPostController"),
)

	ctx.JSON(http.StatusCreated, view.PostResponse(domain))
}

func (p *postControllerInterface) ListPosts(ctx *gin.Context)  {
	if entity.Posts == nil {
		return  
	}
	
	ctx.JSON(http.StatusAccepted, response.PostResponses)
	fmt.Println(response.PostResponses) 
}

func UpdatePostDescription(ctx *gin.Context) error {
// 	id, _ := strconv.Atoi(ctx.Param("id"))
// 	p := new(entity.Post)
// 	if err := ctx.Bind(p); err != nil {
// 		return ctx.JSON(http.StatusNotFound, nil)
// 	}
// 	entity.Posts[id].Description = p.Description
// 	return ctx.JSON(http.StatusOK, entity.Posts[id])
return nil
}

// func DeletePost(ctx echo.Context) error {
// 	id := ctx.Param("id")
// 	for i := range entity.Posts {
// 		if entity.Posts[i].PostId == id {
// 			entity.Posts = append(entity.Posts[:i], entity.Posts[i+1:]...)
// 			return ctx.JSON(http.StatusOK, nil)
// 		}
// 	}
// 	return ctx.JSON(http.StatusNotFound, nil)
// }
