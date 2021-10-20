package poetryctl

import (
	"context"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"falcon/service/poetrysrv"
	"falcon/web/errorenum"
	"falcon/web/middleware"
)

// url: /apienum/v1/poetry/search
func SearchApi(c *gin.Context) {

	searchType := c.Query("t")
	query := c.Query("q")
	pageStr := c.DefaultQuery("page", "1")
	sizeStr := c.DefaultQuery("size", "20")

	page, err := strconv.Atoi(pageStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorenum.ErrorInvalidArgument)
		return
	}
	size, err := strconv.Atoi(sizeStr)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorenum.ErrorInvalidArgument)
		return
	}

	response := &SearchResponse{
		Total:  0,
		Poetry: make([]*PoetryData, 0),
	}

	var (
		offset = (page - 1) * size
		limit  = size
	)

	ifr := middleware.GetInfra(c)

	poetrySrv := poetrysrv.New(ifr)
	args := &poetrysrv.SearchArgs{Query: query}
	switch searchType {
	case "word":
		args.Type = poetrysrv.TypeContent
	case "author":
		args.Type = poetrysrv.TypeAuthor
	case "title":
		args.Type = poetrysrv.TypeTitle
	case "dynastymapper":
		args.Type = poetrysrv.TypeDynasty
	}
	if len(query) > 0 {
		total, rows, err := poetrySrv.Search(context.Background(), args, offset, limit)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusBadRequest, errorenum.DBError)
			return
		}

		response.Total = total
		for _, c := range rows {
			p := &PoetryData{
				Id:      c.Id,
				Title:   c.Title,
				Content: c.Content,
			}
			response.Poetry = append(response.Poetry, p)
		}
	}

	c.JSON(http.StatusOK, response)
}

// url :/apienum/v1/detail
func DetailApi(c *gin.Context) {
	poetryId := c.Param("id")

	ifr := middleware.GetInfra(c)

	srv := poetrysrv.New(ifr)
	p, err := srv.Detail(context.Background(), poetryId)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, errorenum.DBError)
		return
	}
	if len(p.Content) == 0 {
		c.AbortWithStatusJSON(http.StatusNotFound, errorenum.NotFoundError)
		return
	}

	response := &DetailResponse{
		Title:   p.Title,
		Author:  p.Author,
		Dynasty: p.Dynasty,
		Content: p.Content,
	}
	c.JSON(http.StatusOK, response)
}
