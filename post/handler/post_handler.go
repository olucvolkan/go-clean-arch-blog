package handler

import (
	"github.com/labstack/echo"
	"github.com/olucvolkan/go-clean-arch-blog/domain"
	validator "gopkg.in/go-playground/validator.v9"
	"net/http"

	"strconv"
)

// ResponseError represent the reseponse error struct
type ResponseError struct {
	Message string `json:"message"`
}

// ArticleHandler  represent the httphandler for article
type PostHandler struct {
	PostService domain.PostService
}

// NewArticleHandler will initialize the articles/ resources endpoint
func NewPostHandler(e *echo.Echo, p domain.PostService) {
	handler := &PostHandler{
		PostService: p,
	}
	e.GET("/posts", handler.GetPosts)
	e.POST("/posts", handler.CreatePost)
}

//Return post lists endpoint
func (p *PostHandler) GetPosts(c echo.Context) error {
	limitS := c.QueryParam("limit")
	limit, _ := strconv.Atoi(limitS)

	ctx := c.Request().Context()

	posts := p.PostService.GetPosts(ctx, limit)

	return c.JSON(http.StatusOK, posts)
}

//Create a post
func (p *PostHandler) CreatePost(c echo.Context) error {

	var post domain.Post

	err := c.Bind(&post)
	if err != nil {
		return c.JSON(http.StatusUnprocessableEntity, err.Error())
	}

	var ok bool
	if ok, err = isRequestValid(&post); !ok {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	err = p.PostService.CreatePost(&post)

	if err != nil {
		return c.JSON(http.StatusBadRequest, ResponseError{Message: err.Error()})
	}

	return c.JSON(http.StatusCreated, post)
}

func isRequestValid(m *domain.Post) (bool, error) {
	validate := validator.New()
	err := validate.Struct(m)
	if err != nil {
		return false, err
	}
	return true, nil
}
