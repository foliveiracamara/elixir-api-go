package controller

import (
	"net/http"
	"strconv"

	"github.com/foliveiracamara/elixir-api-go/domain/entity"
	"github.com/labstack/echo/v4"
)

func ListPosts(ctx echo.Context) error {
	if entity.Posts == nil {
		return ctx.String(http.StatusNotFound, "No post was created")
	}
	return ctx.JSON(http.StatusOK, entity.Posts)
}

func GetPostById(ctx echo.Context) error {
	id := ctx.Param("id")
	for _, post := range entity.Posts {
		if post.PostId == id {
			return ctx.JSON(http.StatusOK, entity.Posts)
		}
	}
	return ctx.JSON(http.StatusNotFound, nil)
}

func CreatePost(ctx echo.Context) error {
	newPost := entity.Post{}
	err := ctx.Bind(&newPost)
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err)
	}
	newPost.PostId = GenerateUUID()
	newPost.PostDate = GetCurrentDateTime()
	entity.Posts = append(entity.Posts, newPost)
	return ctx.JSON(http.StatusCreated, newPost)
}

func UpdatePostDescription(ctx echo.Context) error {
	id, _ := strconv.Atoi(ctx.Param("id"))
	p := new(entity.Post)
	if err := ctx.Bind(p); err != nil {
		return ctx.JSON(http.StatusNotFound, nil)
	}
	entity.Posts[id].Description = p.Description
	return ctx.JSON(http.StatusOK, entity.Posts[id])
}

func DeletePost(ctx echo.Context) error {
	id := ctx.Param("id")
	for i := range entity.Posts {
		if entity.Posts[i].PostId == id {
			entity.Posts = append(entity.Posts[:i], entity.Posts[i+1:]...)
			return ctx.JSON(http.StatusOK, nil)
		}
	}
	return ctx.JSON(http.StatusNotFound, nil)
}
